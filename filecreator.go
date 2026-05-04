package main

import (
	"fmt"
	"os"
)

type FileCreator struct {
	fileName string
}

func (f FileCreator) CreateTemplateFile() *os.File {
	templateFile, err := os.Create(f.fileName);
	if err != nil { panic(err); }
	dir, dirErr := os.Getwd();
	if dirErr != nil { panic(dirErr);}
	fmt.Printf("Created %s at : %s\n", templateFile.Name(), dir);
	return templateFile;
}

func(f FileCreator) WriteDefaultReadme(templateFile *os.File) {
	var writer = FileWriter{templateFile: templateFile};

	writer.WriteHeading("Project Title", 1);
	writer.WriteLine("Simple overview of use/purpose.\n\n");

	writer.WriteHeading("Description", 2);
	writer.WriteLine("An in-depth paragraph about your project and overview of use.\n\n");

	writer.WriteHeading("Getting Started", 2);

	writer.WriteHeading("Dependencies", 3);
	writer.WriteLine("- Describe any prerequisites, libraries, OS version, etc., needed before installing the program\n");
	writer.WriteLine("- ex: Windows 10\n\n");

	writer.WriteHeading("Installing", 3);
	writer.WriteLine("- How/where to download your program\n")
	writer.WriteLine("- Any modifications needed to be made to files/folders\n\n");

	writer.WriteHeading("Execution", 3);
	writer.WriteLine("- How to run the program\n");
	writer.WriteLine("- Step-by-step bullets\n");
	writer.WriteCodeBlock([]string{"code blocks for commands"}, "");
	

	writer.WriteHeading("Contribution", 2);
	writer.WriteLine("Any advice for common problems or issues. \n\n");
	writer.WriteCodeBlock([]string{"command to run if program contains helper info"}, "");

	writer.WriteHeading("Authors", 2);
	writer.WriteLine("Contributors names and contact info\n\n");
	writer.WriteLine("ex. John Doe\n");
	
	writer.WriteHeading("License", 2);
	writer.WriteLine("Contributors names and contact info.\n\n");

	writer.WriteHeading("Acknowledgments", 2);
	writer.WriteLine("Inspiration, code snippets, etc.\n\n");
}