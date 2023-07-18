package clients

import (
	"fmt"
	"momentum-core/config"
	"momentum-core/tree"
	"momentum-core/utils"

	"sigs.k8s.io/kustomize/api/krusty"
	"sigs.k8s.io/kustomize/kyaml/filesys"
)

type KustomizationValidationClient struct {
	config *config.MomentumConfig
}

func NewKustomizationValidationClient(config *config.MomentumConfig) *KustomizationValidationClient {

	validator := new(KustomizationValidationClient)

	validator.config = config

	return validator
}

// returns nil if repo path is valid kustomizable
func (kustomizationService *KustomizationValidationClient) Validate(repoName string) error {

	path := utils.BuildPath(kustomizationService.config.ValidationTmpDir(), repoName)
	src := utils.BuildPath(kustomizationService.config.DataDir(), repoName)

	err := kustomizationService.prepareValidation(path, src)
	if err != nil {
		fmt.Println("error while validating kustomize structure (prepare):", err.Error())
		kustomizationService.validationCleanup(path)
		return err
	}

	repoTree, err := tree.Parse(path)
	if err != nil {
		fmt.Println("failed parsing validation directory")
		return err
	}

	err = kustomizationService.check(repoTree.MomentumRoot().FullPath())
	if err != nil {
		fmt.Println("error while validating kustomize structure (check):", err.Error())
		kustomizationService.validationCleanup(path)
		return err
	}

	err = kustomizationService.validationCleanup(path)
	if err != nil {
		fmt.Println("error while validating kustomize structure (cleanup):", err.Error())
		return err
	}

	return nil
}

func (kustomizationService *KustomizationValidationClient) prepareValidation(path string, src string) error {

	_, err := utils.DirCopy(src, path)
	return err
}

func (kustomizationService *KustomizationValidationClient) check(path string) error {

	fs := filesys.MakeFsOnDisk()

	// TODO ->  OpenAI ApiSchemes for FluxCD -> Kubeconform
	kustomizer := krusty.MakeKustomizer(krusty.MakeDefaultOptions())

	_, err := kustomizer.Run(fs, path)
	if err != nil {
		return err
	}
	return nil
}

func (kustomizationService *KustomizationValidationClient) validationCleanup(path string) error {

	err := utils.DirDelete(path)
	if err != nil {
		return err
	}
	return nil
}
