package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/fatih/color"
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
	DirPath  string
}

type PageList struct {
	HTMLPagePath string
	Title        string
	Pages        []Page
}

var generatedFilesCount int = 0
var bytesGenerated int64 = 0

func main() {
	filename := flag.String("file", "", "Name of a text file in the current directory.")
	dirpath := flag.String("dir", "./", "A path to a directory containing text files.")
	flag.Parse()

	if *filename != "" {
		generatePageFromFile("./", *filename, ".txt")
		return
	}

	generateSiteFromDir(*dirpath)
	printSuccessMessage()
}

func printSuccessMessage() {
	white := color.New(color.FgWhite)
	green := color.New(color.FgGreen)
	boldGreen := green.Add(color.Bold)

	outputSize := fmt.Sprintf("(%.1fkB) total.", float64(bytesGenerated)/1000)

	// generate size for banner by generatedFilesCount char length
	banner := strings.Repeat("-", 29+len(strconv.Itoa(generatedFilesCount))+len(outputSize))

	white.Print(banner + " \n")
	boldGreen.Print(" Success! ")
	fmt.Print("Generated ")
	boldGreen.Printf("%d ", generatedFilesCount)
	fmt.Printf("pages " + outputSize + "\n")
	white.Print(banner + " \n")
}

func generateSiteFromDir(dirpath string) {
	textFilesInDir := getFilesInDirectory(dirpath, ".txt")
	var pages []Page

	for _, file := range textFilesInDir {
		pages = append(
			pages,
			generatePageFromFile(file.DirPath, file.FileName, ""),
		)
	}
	generateHomePage(pages, "Home")
}

func getFilesInDirectory(dirpath string, extension string) []File {
	directory := dirpath
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	var textFiles []File
	for _, file := range files {
		if filepath.Ext(file.Name()) == extension {

			filePath := dirpath + file.Name()

			fileInfo := File{
				FileName: file.Name(),
				FilePath: filePath,
				DirPath:  dirpath,
			}

			textFiles = append(textFiles, fileInfo)
		}

		// get files recursively
		if file.IsDir() {
			filesInSubDir := getFilesInDirectory(dirpath+file.Name()+"/", ".txt")
			textFiles = append(textFiles, filesInSubDir...)
		}
	}

	return textFiles
}

// Take a file path and save that file's contents as a new html post
func generatePageFromFile(dirpath string, filename string, extension string) Page {
	// Check if parent directory exists at output dirpath yet
	fileContents, err := ioutil.ReadFile(dirpath + filename + extension)
	if err != nil {
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
	incrementStatCounter(*newFile)

	return page
}

func generateHomePage(pages []Page, title string) {
	homePage := PageList{
		HTMLPagePath: "./output/index.html",
		Title:        title,
		Pages:        pages,
	}

	// Create a new, blank HTML file.
	newFile, err := os.Create(homePage.HTMLPagePath)
	if err != nil {
		panic(err)
	}

	t := template.Must(template.New("home.tmpl").ParseFiles("home.tmpl"))
	t.Execute(newFile, homePage)

	incrementStatCounter(*newFile)
}

func incrementStatCounter(file os.File) {
	// Get size of new file
	fileStat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	bytesGenerated += fileStat.Size()
	generatedFilesCount += 1
}
