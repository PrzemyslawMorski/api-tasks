package handlers

import (
	"os"
	"testing"
)

const testDbFileName = "__test__.db"

func DeleteTestDb(t *testing.T) {
	if _, err := os.Stat(testDbFileName); os.IsNotExist(err) {
		return
	}

	err := os.Remove(testDbFileName)
	if err != nil {
		t.Log(err)
	}
}
