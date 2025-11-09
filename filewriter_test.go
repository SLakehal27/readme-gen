package main

import (
	"os"
	"strings"
	"testing"
)

func Test_WriteLine(t *testing.T){
	template, err := CreateTestFile(t.TempDir());
	if err != nil {
		t.Errorf("Unable to create file at directory %s", template.Name());
	}
	defer template.Close();

	var testFileWriter = FileWriter{template};
	var testLine = "Test line";
	testFileWriter.WriteLine(testLine);
	
	fileContents, err := os.ReadFile(template.Name());

	if err != nil {
		t.Errorf("Unable to read file at directory %s", template.Name());
	}

	if (!strings.Contains(string(fileContents), testLine)){
		t.Errorf("%s does not contain the line %s", template.Name(), testLine)
	}

}

// TODO: Add 3 cases for WriteHeading that take into account it panicking.
// TODO: Add case for WriteCodeBlock