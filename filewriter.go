package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type FileWriter struct {
	templateFile *os.File
}

func (w FileWriter) WriteLine(str string) {
	w.templateFile.WriteString(str + "\n")
}

func (w FileWriter) WriteHeading(str string, headingType int) {
	var LOWER_HEADING_LIMIT = 1;
	var UPPER_HEADING_LIMIT = 6;

	if (headingType < LOWER_HEADING_LIMIT  || headingType > UPPER_HEADING_LIMIT) {
		headingError := fmt.Sprintf("unsupported heading type : %d", headingType);
		panic(errors.New(headingError));
	}

	var heading = strings.Repeat("#", headingType);
	w.templateFile.WriteString(heading + " " + str + "\n");
}

func (w FileWriter) WriteCodeBlock(codeBlock []string) {
	w.WriteLine("```");
	for _, code := range codeBlock {
		w.WriteLine(code);
	}
	w.WriteLine("```");
}

