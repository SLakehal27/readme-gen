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
	fmt.Printf("Created README file : %s\n", templateFile.Name());
	return templateFile;
}

func(f FileCreator) WriteDefaultReadme(templateFile *os.File) {
	var writer = FileWriter{templateFile: templateFile};
	writer.WritePartialHeading("README Generator ", 1);
	writer.WriteImage("readme-gen-pipeline", "https://github.com/SLakehal27/readme-gen/actions/workflows/pipeline.yml/badge.svg?branch=main");

	writer.WriteLine("A simple README.md generator made in GoLang to generate README.md files for future GitHub projects.\n\n");

	writer.WriteHeading("Why a README generator?", 2)
	writer.WriteLine("Sure, there are other tools to help generate README files that are already implemented and others that use AI for it.");
	writer.WriteLine("I see this project as a learning experience and an opportunity to make my own tooling with what I've learned so far.\n\n");

	writer.WriteHeading("Usage", 2);

	writer.WriteHeading("Execution", 3);
	var runCmds = []string{"go run ."};
	writer.WriteCodeBlock(runCmds, "");

	writer.WriteHeading("Testing", 3);
	var testCmds = []string{"go test -cover -v"};
	writer.WriteCodeBlock(testCmds, "");
}