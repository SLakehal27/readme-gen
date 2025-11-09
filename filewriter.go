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
	w.templateFile.WriteString(str)
}

func formatHeading(str string, headingType int) string {
	var LOWER_HEADING_LIMIT = 1;
	var UPPER_HEADING_LIMIT = 6;

	if (headingType < LOWER_HEADING_LIMIT  || headingType > UPPER_HEADING_LIMIT) {
		headingError := fmt.Sprintf("unsupported heading type : %d", headingType);
		panic(errors.New(headingError));
	}

	var headings = strings.Repeat("#", headingType);
	return headings + " " + str;
}

func (w FileWriter) WritePartialHeading(str string, headingType int) {
	w.WriteLine(formatHeading(str, headingType));
}

func (w FileWriter) WriteHeading(str string, headingType int) {
	w.WriteLine(formatHeading(str, headingType) + "\n\n");
}

func formatURL(name string, url string) string {
	return "[" + name + "]" + "(" + url + ")";
}

func (w FileWriter) WriteImage(name string, url string) {
	w.WriteLine("!" + formatURL(name, url) + "\n\n");
}

func (w FileWriter) WriteCodeBlock(codeBlock []string, language string) {
	w.WriteLine("```" + language + "\n");
	for _, code := range codeBlock {
		w.WriteLine(code + "\n");
	}
	w.WriteLine("```" + "\n\n");
}

