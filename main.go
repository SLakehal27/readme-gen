package main

func main() {
	var fileCreator = FileCreator{filePath: "README.md"};
	var templateFile = fileCreator.CreateTemplate();
	defer templateFile.Close();
	fileCreator.WriteDefaultReadme(templateFile);
}