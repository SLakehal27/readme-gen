package main

import (
	"fmt"
	"os"
)

type FileCreator struct {
	filePath string
}

func (f FileCreator) CreateTemplate() *os.File {
	templateFile, err := os.Create(f.filePath);
	if err != nil { panic(err); }
	fmt.Printf("Created README file : %s", templateFile.Name());
	return templateFile;
}

func(f FileCreator) WriteDefaultReadme(templateFile *os.File) {
	var writer = FileWriter{templateFile: templateFile};
	writer.WriteHeading("README Template", 1);
	writer.WriteLine("Replace this with the purpose of your project.");

	writer.WriteHeading("Usage", 2);

	writer.WriteHeading("Execution", 3);
	var runCmds = []string{"go run ."};
	writer.WriteCodeBlock(runCmds);

	writer.WriteHeading("Testing", 3);
	var testCmds = []string{"go test -cover -v"};
	writer.WriteCodeBlock(testCmds);
}