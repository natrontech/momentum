package clients_test

import (
	"fmt"
	"momentum-core/clients"
	"momentum-core/utils"
	"os"
	"testing"
)

const TEST_LOCATION = "testdata"
const TEST_REPO = "https://github.com/Joel-Haeberli/momentum-test-strcuture-2"

func FILESYSTEMTEST_TestTransactionCycle(t *testing.T) {

	clients.CloneRepoTo(TEST_REPO, "", "", TEST_LOCATION)

	transactionId, err := clients.InitGitTransaction(TEST_LOCATION, "test transaction")
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}

	filePath := utils.BuildPath(TEST_LOCATION, "testtransactionfile.txt")
	f, err := utils.FileOpen(filePath, os.O_CREATE|os.O_RDWR)
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}

	_, err = f.WriteString("this is a test")
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}
	f.Close()

	err = clients.GitTransactionWrite(filePath, transactionId)
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}

	err = clients.GitTransactionCommit(transactionId)
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}

	utils.DirDelete(TEST_LOCATION)
}
