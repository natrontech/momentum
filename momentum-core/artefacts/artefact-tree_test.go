package artefacts_test

import (
	"fmt"
	"momentum-core/artefacts"
	"momentum-core/config"
	"testing"
)

func TestArtefactTree(t *testing.T) {

	_, err := config.InitializeMomentumCore()
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}

	artfcts, err := artefacts.LoadArtefactTree()
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}

	fmt.Println(artefacts.WriteToString(artfcts))
}
