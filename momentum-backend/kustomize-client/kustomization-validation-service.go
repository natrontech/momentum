package kustomizeclient

import (
	"errors"
	"fmt"
	config "momentum/momentum-core/momentum-config"
	services "momentum/momentum-core/momentum-services"
	utils "momentum/momentum-core/momentum-utils"

	"github.com/pocketbase/pocketbase/models"
	"github.com/spf13/cobra"
	kustomize "sigs.k8s.io/kustomize/kustomize/v5/commands"
)

const KUSTOMIZE_BUILD_COMMAND_NAME = "build"

type KustomizationValidationService struct {
	config            *config.MomentumConfig
	repositoryService *services.RepositoryService
}

func (kustomizationService *KustomizationValidationService) Validate(repo *models.Record) bool {

	if repo.Collection().Name != config.TABLE_REPOSITORIES_NAME {
		fmt.Println("cannot validate non repositories collection record")
		return false
	}

	repoName := repo.GetString(config.TABLE_REPOSITORIES_FIELD_NAME)
	path := utils.BuildPath(kustomizationService.config.ValidationTmpDir(), repoName)
	src := utils.BuildPath(kustomizationService.config.DataDir(), repoName)

	err := kustomizationService.prepareValidation(path, src)
	if err != nil {
		fmt.Println("error while validating kustomize structure:", err.Error())
		kustomizationService.validationCleanup(path)
		return false
	}
	err = kustomizationService.check(path)
	if err != nil {
		fmt.Println("error while validating kustomize structure:", err.Error())
		kustomizationService.validationCleanup(path)
		return false
	}

	success := kustomizationService.checkSuccessful(path)

	err = kustomizationService.validationCleanup(path)
	if err != nil {
		fmt.Println("error while validating kustomize structure:", err.Error())
	}

	return success
}

func (kustomizationService *KustomizationValidationService) prepareValidation(path string, src string) error {

	_, err := utils.DirCopy(src, path)
	return err
}

func (kustomizationService *KustomizationValidationService) check(path string) error {

	var buildCmd *cobra.Command = nil
	for _, cmd := range kustomize.NewDefaultCommand().Commands() {
		if cmd.Name() == KUSTOMIZE_BUILD_COMMAND_NAME {
			buildCmd = cmd
			break
		}
	}

	if buildCmd == nil {
		return errors.New("build command was not found")
	}

	// TODO: FIND OUT HOW TO SET PATH HERE???
	buildCmd.Args = kustomizationService.kustomizeBuildArgs

	err := buildCmd.Execute()

	return err
}

func (kustomizationService *KustomizationValidationService) checkSuccessful(path string) bool {

	return true
}

func (kustomizationService *KustomizationValidationService) validationCleanup(path string) error {

	err := utils.DirDelete(path)
	if err != nil {
		return err
	}
	return nil
}

func (kustomizationService *KustomizationValidationService) kustomizeBuildArgs(cmd *cobra.Command, args []string) error {

	return nil
}
