package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// Page holds all the information we need to generate a new
// HTML page from a text file on the filesystem.
type Page struct {
	TextFilePath string
	TextFileName string
	HTMLPagePath string
	Content      string
}

type File struct {
	FileName string
	FilePath string
}

func main() {
	filename := flag.String("file", "", "Name of a text file in the current directory.")
	dirpath := flag.String("dir", "./", "A path to a directory containing text files.")
	flag.Parse()

	if *filename != "" {
		generatePageFromFile("./"+*filename, *filename, "txt")
		return
	}

	generateSiteFromDir(*dirpath)
}

func generateSiteFromDir(dirpath string) {
	textFilesInDir := getTextFilesInDirectory(dirpath, "txt")

	for _, file := range textFilesInDir {
		generatePageFromFile(file.FilePath, file.FileName, "")
	}
}

func getTextFilesInDirectory(dirpath string, extension string) []File {
	directory := dirpath
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	var textFiles []File
	for _, file := range files {
		if filepath.Ext(file.Name()) == "."+extension {
			filePath := dirpath + file.Name()
			fileInfo := File{
				FileName: file.Name(),
				FilePath: filePath,
			}
			textFiles = append(textFiles, fileInfo)
		}

		// get files recursively
		// if file.IsDir() {
		// 	filesInSubDir := getTextFilesInDirectory(dirpath + "/" + file.Name())
		// 	textFiles = append(textFiles, filesInSubDir...)
		// }
	}

	return textFiles
}

// Take a file path and save that file's contents as a new html post
func generatePageFromFile(dirpath string, filename string, extension string) {
	fileContents, err := ioutil.ReadFile(dirpath + extension)
	if err != nil {
		// A common use of `panic` is to abort if a function returns an error
		// value that we don’t know how to (or want to) handle. This example
		// panics if we get an unexpected error when creating a new file.
		panic(err)
	}
	outputFilePath := "./output/" + filename + ".html"

	page := Page{
		TextFilePath: "./first-post",
		TextFileName: filename,
		HTMLPagePath: outputFilePath,
		Content:      string(fileContents),
	}

	// Create a new template in memory named "template.tmpl".
	// When the template is executed, it will parse template.tmpl,
	// looking for {{ }} where we can inject content.
	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	// Create a new, blank HTML file.
	newFile, err := os.Create(page.HTMLPagePath)
	if err != nil {
		panic(err)
	}

	// Executing the template injects the Page instance's data,
	// allowing us to render the content of our text file.
	// Furthermore, upon execution, the rendered template will be
	// saved inside the new file we created earlier.
	t.Execute(newFile, page)
}
