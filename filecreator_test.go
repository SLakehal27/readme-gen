package main

import (
	"os"
	"path"
	"syscall"
	"testing"
)

func CreateTestFile(tempDir string) (*os.File, error) {
	var filePath = path.Join(tempDir, "test.md");
	var fileCreator = FileCreator{filePath: filePath};
	template, err := os.Create(fileCreator.filePath);
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

// TODO: Add test case for WriteDefaultReadme()