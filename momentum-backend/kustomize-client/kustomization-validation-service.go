package kustomizeclient

import (
	"fmt"
	config "momentum/momentum-core/momentum-config"
	services "momentum/momentum-core/momentum-services"
	tree "momentum/momentum-core/momentum-tree"
	utils "momentum/momentum-core/momentum-utils"

	"sigs.k8s.io/kustomize/api/krusty"
	"sigs.k8s.io/kustomize/kyaml/filesys"
)

type KustomizationValidationService struct {
	config            *config.MomentumConfig
	repositoryService *services.RepositoryService
}

func NewKustomizationValidationService(config *config.MomentumConfig, repositoryService *services.RepositoryService) *KustomizationValidationService {

	validator := new(KustomizationValidationService)

	validator.config = config
	validator.repositoryService = repositoryService

	return validator
}

func (kustomizationService *KustomizationValidationService) Validate(repoName string) (bool, error) {

	path := utils.BuildPath(kustomizationService.config.ValidationTmpDir(), repoName)
	src := utils.BuildPath(kustomizationService.config.DataDir(), repoName)

	err := kustomizationService.prepareValidation(path, src)
	if err != nil {
		fmt.Println("error while validating kustomize structure (prepare):", err.Error())
		kustomizationService.validationCleanup(path)
		return false, err
	}

	repoTree, err := tree.Parse(path, []string{".git"})
	if err != nil {
		fmt.Println("failed parsing validation directory")
		return false, err
	}

	for _, app := range repoTree.Apps() {
		err = kustomizationService.check(app.FullPath())
		if err != nil {
			fmt.Println("error while validating kustomize structure (check):", err.Error())
			kustomizationService.validationCleanup(path)
			return false, err
		}
	}

	err = kustomizationService.validationCleanup(path)
	if err != nil {
		fmt.Println("error while validating kustomize structure (cleanup):", err.Error())
		return false, err
	}

	return true, nil
}

func (kustomizationService *KustomizationValidationService) prepareValidation(path string, src string) error {

	_, err := utils.DirCopy(src, path)
	return err
}

func (kustomizationService *KustomizationValidationService) check(path string) error {

	fs := filesys.MakeFsOnDisk()

	kustomizer := krusty.MakeKustomizer(krusty.MakeDefaultOptions())

	_, err := kustomizer.Run(fs, path)
	if err != nil {
		return err
	}
	return nil
}

func (kustomizationService *KustomizationValidationService) validationCleanup(path string) error {

	err := utils.DirDelete(path)
	if err != nil {
		return err
	}
	return nil
}
