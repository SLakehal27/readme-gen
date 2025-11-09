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

func Test_WriteHeadingMin(t *testing.T){
	template, err := CreateTestFile(t.TempDir());
	if err != nil {
		t.Errorf("Unable to create file at directory %s", template.Name());
	}
	defer template.Close();

	var testFileWriter = FileWriter{template};
	var testLine = "Text with heading under 1";

	defer func() {
		if r := recover(); r == nil {
			t.Error("Invalid heading : under 1");
		}
	}()

	testFileWriter.WriteHeading(testLine, -1);
}

func Test_WriteHeadingMax(t *testing.T){
	template, err := CreateTestFile(t.TempDir());
	if err != nil {
		t.Errorf("Unable to create file at directory %s", template.Name());
	}
	defer template.Close();

	var testFileWriter = FileWriter{template};
	var testLine = "Text with heading over 6";

	defer func() {
		if r := recover(); r == nil {
			t.Error("Invalid heading : over 6");
		}
	}()

	testFileWriter.WriteHeading(testLine, 7);
}

func Test_WriteHeading(t *testing.T){
	template, err := CreateTestFile(t.TempDir());
	if err != nil {
		t.Errorf("Unable to create file at directory %s", template.Name());
	}
	defer template.Close();

	var testFileWriter = FileWriter{template};
	var testLine = "Test line";
	testFileWriter.WriteHeading(testLine, 1);
	
	fileContents, err := os.ReadFile(template.Name());

	if err != nil {
		t.Errorf("Unable to read file at directory %s", template.Name());
	}

	if (!strings.Contains(string(fileContents), testLine)){
		t.Errorf("%s does not contain the line %s", template.Name(), "# " + testLine)
	}
}

// TODO: Add case for WriteCodeBlock
