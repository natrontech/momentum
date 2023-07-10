package momentumservices_test

import (
	"fmt"
	services "momentum/momentum-core/momentum-services"
	utils "momentum/momentum-core/momentum-utils"
	"os"
	"testing"
)

func FILESYSTEMTEST_TestApplyDeploymentReleaseTemplate(t *testing.T) {

	var wd, _ = os.Getwd()
	var TEMPLATE_TEST_FILE_PATH = utils.BuildPath(wd, "template_test.yaml")
	var templateService = services.NewTemplateService()

	template := "Hello my name is {{ .ApplicationName }}"
	writeTestFile(t, TEMPLATE_TEST_FILE_PATH, []string{template})

	appName := "{CoolApp}"

	err := templateService.ApplyDeploymentReleaseTemplate(TEMPLATE_TEST_FILE_PATH, appName)
	if err != nil {
		fmt.Println("applying template failed:", err.Error())
		t.FailNow()
	}

	f, err := utils.FileOpen(TEMPLATE_TEST_FILE_PATH, utils.FILE_ALLOW_READ_WRITE_ALL)
	if err != nil {
		fmt.Println("failed opening file:", err.Error())
	}

	lines := utils.FileAsLines(f)

	check := lines[0]
	fmt.Println(check)
	if "Hello my name is {CoolApp}" != check {
		fmt.Println(check, "does not satisfies assumptions")
	}

	cleanup(t, TEMPLATE_TEST_FILE_PATH)
}

func FILESYSTEMTEST_TestApplyDeploymentStageDeploymentDescriptionTemplate(t *testing.T) {

	var wd, _ = os.Getwd()
	var TEMPLATE_TEST_FILE_PATH = utils.BuildPath(wd, "template_test.yaml")
	var templateService = services.NewTemplateService()

	template := "Hello my name is '{{ .DeploymentName }}' and im living in {{ .RepositoryName }}"
	writeTestFile(t, TEMPLATE_TEST_FILE_PATH, []string{template})

	deploymentName := "CoolDeployment"
	repositoryName := "CoolRepository"

	err := templateService.ApplyDeploymentStageDeploymentDescriptionTemplate(TEMPLATE_TEST_FILE_PATH, deploymentName, repositoryName)
	if err != nil {
		fmt.Println("applying template failed:", err.Error())
		t.FailNow()
	}

	cleanup(t, TEMPLATE_TEST_FILE_PATH)
}

func FILESYSTEMTEST_TestApplyDeploymentKustomizationTemplate(t *testing.T) {

	var wd, _ = os.Getwd()
	var TEMPLATE_TEST_FILE_PATH = utils.BuildPath(wd, "template_test.yaml")
	var templateService = services.NewTemplateService()

	template := "Hello my name is {{ .DeploymentName }}"
	writeTestFile(t, TEMPLATE_TEST_FILE_PATH, []string{template})

	deploymentName := "CoolDeployment"

	err := templateService.ApplyDeploymentKustomizationTemplate(TEMPLATE_TEST_FILE_PATH, deploymentName)
	if err != nil {
		fmt.Println("applying template failed:", err.Error())
		t.FailNow()
	}

	cleanup(t, TEMPLATE_TEST_FILE_PATH)
}

func writeTestFile(t *testing.T, p string, lines []string) {
	if !utils.FileWriteLines(p, lines) {
		fmt.Println("failed writing test file")
		t.FailNow()
	}
}

func cleanup(t *testing.T, p string) {
	err := os.Remove(p)

	if err != nil {
		fmt.Println("unable to clean up after test")
		t.FailNow()
	}
}
