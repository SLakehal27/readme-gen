package main

import (
	"fmt"
)

func main() {
	var templateFile = CreateTemplate();
	fmt.Printf("Created README file : %s", templateFile.Name());
	var writer = Writer{templateFile: templateFile};
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