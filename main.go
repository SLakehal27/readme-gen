package main

func main() {
	var fileCreator = FileCreator{fileName: "README.md"};
	var templateFile = fileCreator.CreateTemplate();
	defer templateFile.Close();
	fileCreator.WriteDefaultReadme(templateFile);
}