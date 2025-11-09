package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Writer struct {
	templateFile *os.File
}

func (w Writer) WriteLine(str string) {
	w.templateFile.WriteString(str + "\n")
}

func (w Writer) WriteHeading(str string, headingType int) {
	var LOWER_HEADING_LIMIT = 1;
	var UPPER_HEADING_LIMIT = 6;

	if (headingType < LOWER_HEADING_LIMIT  || headingType > UPPER_HEADING_LIMIT) {
		headingError := fmt.Sprintf("unsupported heading type : %d", headingType);
		panic(errors.New(headingError));
	}

	var heading = strings.Repeat("#", headingType);
	w.templateFile.WriteString(heading + " " + str + "\n");
}

func (w Writer) WriteCodeBlock(codeBlock []string) {
	w.WriteLine("```");
	for _, code := range codeBlock {
		w.WriteLine(code);
	}
	w.WriteLine("```");
}

// TODO :  find a good spot for this method.
func CreateTemplate() *os.File {
	filePath := "README.md";

	templateFile, err := os.Create(filePath);
	if err != nil { panic(err); }

	return templateFile;
}