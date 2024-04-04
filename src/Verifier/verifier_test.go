package Verifier

import (
	"log"
	"os"
	"strings"
	"testing"
)

var csvfile1 string = `name,ring,quadrant,isNew,moved,description
Go,Adopt,Language,true,0,Its a programming Language
Visual Studio Code,Trial,Tool,false,2,An IDE
Dagger IO,Assess,Tool,true,1,Its a workflow thing`

func createCsvFiles() {
	err := os.WriteFile("testFile1.csv", []byte(csvfile1), 0644)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("testFile2.csv", []byte(csvfile1), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func cleanUp() {
	os.Remove("testFile1.csv")
	os.Remove("testFile2.csv")
}

func TestVerifierFunctionDuplicateDeletion(t *testing.T) {
	createCsvFiles()
	defer cleanUp()

	Verifier("./testFile1.csv", "./testFile2.csv")

	csv1, err := os.ReadFile("./testFile1.csv")
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(csv1), csvfile1) {
		t.Fatalf("csvFile1 does not match expected output.\nExpected: %s \n Actual: %s", csvfile1, csv1)
	}

	csv2, err := os.ReadFile("./testFile2.csv")
	if err != nil {
		t.Fatal(err)
	}

	if string(csv2) != "name,ring,quadrant,isNew,moved,description\n" {
		t.Fatalf("csvFile2 does not match expected output.\nExpected: name,ring,quadrant,isNew,moved,description \n Actual: %s",csv2)
	}
}