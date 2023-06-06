package utils_test

import (
	"strings"
	"testing"

	utils "momentum/momentum-core/momentum-utils"
)

func TestBuildPath(t *testing.T) {

	absoluteExpectedResult := "/from/root/append/suffix"
	relativeExpectedResult := "./from/root/append/suffix"
	relativeExpectedResult2 := "from/root/append/suffix"
	fileSuffixExpectedResult := "/from/root/append/config.yml"

	absoluteBase := "/from/root/"
	absoluteBase2 := "/from/root"
	relativeBase := "from/root/"
	relativeBase2 := "./from/root/"
	relativeBase3 := "from/root"

	suffix := "append/suffix"
	suffix2 := "append/suffix/"
	suffix3 := "/append/suffix"  // This is convenience and might be misleading for user -> supplying absolute path as suffix and accepting
	suffix4 := "/append/suffix/" // This is convenience and might be misleading for user -> supplying absolute path as suffix and accepting

	fileSuffix := "append/config.yml"

	absolutePaths := []string{absoluteBase, absoluteBase2}
	relativePaths := []string{relativeBase, relativeBase2, relativeBase3}
	suffixes := []string{suffix, suffix2, suffix3, suffix4}
	fileSuffixes := []string{fileSuffix}

	for _, ap := range absolutePaths {
		for _, s := range suffixes {
			p := utils.BuildPath(ap, s)
			if absoluteExpectedResult != p {
				t.Error("Test failed with", ap, s, p)
			}
		}
	}

	for _, rp := range relativePaths {

		check := relativeExpectedResult2
		if strings.Index(strings.Trim(rp, " "), ".") == 0 {
			check = relativeExpectedResult
		}

		for _, s := range suffixes {
			p := utils.BuildPath(rp, s)
			if check != p {
				t.Error("Test failed with", rp, s, p)
			}
		}
	}

	for _, ap := range absolutePaths {
		for _, s := range fileSuffixes {
			p := utils.BuildPath(ap, s)
			if fileSuffixExpectedResult != p {
				t.Error("Test failed with", ap, s, p)
			}
		}
	}
}

func TestFileEmpty(t *testing.T) {
	testFile := "/home/joel/projects/natrium-cli/natrium-demo/natrium/mywebserver/_templates/release.yaml"

	if utils.FileExists(testFile) {

		isEmpty := utils.FileEmpty(testFile)
		if !isEmpty {
			t.Errorf("Test file is empty but not recongized as such.")
		}
		return
	}

	t.Errorf("Test File %s is not existent", testFile)
}
