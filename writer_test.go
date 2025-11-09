package main

import (
	"os"
	"path"
	"strings"
	"syscall"
	"testing"
)

func CreateTestFile(tempDir string) (*os.File, error) {
	var filePath = path.Join(tempDir, "test.md");
	template, err := os.Create(filePath);
	return template, err;
}

func Test_CreateTemplate(t *testing.T) {
	template, err := CreateTestFile(t.TempDir());
	e, ok := err.(*os.PathError);
	defer template.Close(); 
	
	if ok && e.Err == syscall.ENOSPC {
		t.Errorf("Unable to create file at directory %s", template.Name()); 
	}
}

func Test_WriteLine(t *testing.T){
	template, err := CreateTestFile(t.TempDir());
	if err != nil {
		t.Errorf("Unable to create file at directory %s", template.Name());
	}
	defer template.Close();

	var testWriter = Writer{template};
	var testLine = "Test line";
	testWriter.WriteLine(testLine);
	
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