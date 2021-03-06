package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/fatih/color"
	"github.com/gomarkdown/markdown"
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
	startTime := time.Now()
	endTime := generateSite()
	timeDiff := endTime.Sub(startTime).Seconds()

	printSuccessMessage(timeDiff)
}

func generateSite() time.Time {
	filename := flag.String("file", "", "Name of a text file in the current directory.")
	dirpath := flag.String("dir", "./", "A path to a directory containing text files. Defaults to ./")
	templateDir := flag.String("templateDir", "./templates/", "A path to a directory containing template files. Defaults to ./templates/")
	outputDir := flag.String("outputDir", "./output/", "A path to the desired output directory. Defaults to ./output/")
	fileExtension := flag.String("ext", ".md", "The extension of files to be parsed. defaults to .md")
	flag.Parse()

	createDirectoryIfDoesNotExist(*outputDir)

	if *filename != "" {
		generatePageFromFile("./", *filename, *fileExtension, *templateDir, *outputDir)
		return time.Now()
	}

	generateSiteFromDir(*dirpath, *templateDir, *outputDir, *fileExtension)
	return time.Now()
}

func printSuccessMessage(timeDiff float64) {
	white := color.New(color.FgWhite)
	green := color.New(color.FgGreen)
	boldGreen := green.Add(color.Bold)

	outputSize := fmt.Sprintf("%.1fkB", float64(bytesGenerated)/1000)
	timeString := fmt.Sprintf("%.2fs.", float64(timeDiff)/1000)

	// generate size for banner by generatedFilesCount char length
	countTextSize := len(strconv.Itoa(generatedFilesCount))
	sizeTextSize := len(outputSize)
	timeDiffSize := len(timeString)

	banner := strings.Repeat("-", 40+countTextSize+sizeTextSize+timeDiffSize)

	white.Print(banner + " \n")

	boldGreen.Print(" Success! ")
	fmt.Print("Generated ")
	boldGreen.Printf("%d ", generatedFilesCount)
	fmt.Printf("pages (")
	boldGreen.Printf(outputSize)
	fmt.Printf(" total) in ")
	boldGreen.Printf(timeString + " \n")
	white.Print(banner + " \n")
}

func createDirectoryIfDoesNotExist(path string) {
	_, err := os.Stat(path)
	if err == nil {
		return
	}

	create_err := os.Mkdir(path, 0755)
	if create_err != nil {
		panic(create_err)
	}
}

func generateSiteFromDir(dirpath string, templatePath string, outputDir string, ext string) {
	textFilesInDir := getFilesInDirectory(dirpath, ext)
	var pages []Page

	for _, file := range textFilesInDir {
		pages = append(
			pages,
			generatePageFromFile(file.DirPath, file.FileName, "", templatePath, outputDir),
		)
	}
	generateHomePage(pages, "Home", templatePath, outputDir)
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
			filesInSubDir := getFilesInDirectory(dirpath+file.Name()+"/", extension)
			textFiles = append(textFiles, filesInSubDir...)
		}
	}

	return textFiles
}

// Take a file path and save that file's contents as a new html post
func generatePageFromFile(
	dirpath string,
	filename string,
	extension string,
	templateDir string,
	outputDir string,
) Page {
	// Check if parent directory exists at output dirpath yet
	fileContents, err := ioutil.ReadFile(dirpath + filename + extension)
	if err != nil {
		panic(err)
	}

	output := markdown.ToHTML(fileContents, nil, nil)
	outputFilePath := outputDir + filename + ".html"

	page := Page{
		TextFilePath: "./first-post",
		TextFileName: filename,
		HTMLPagePath: outputFilePath,
		Content:      string(output),
	}

	// Create a new template in memory named "template.tmpl".
	// When the template is executed, it will parse template.tmpl,
	// looking for {{ }} where we can inject content.
	t := template.Must(template.New("template.tmpl").ParseFiles(templateDir + "template.tmpl"))

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

func generateHomePage(pages []Page, title string, templateDir string, outputDir string) {
	homePage := PageList{
		HTMLPagePath: (outputDir + "index.html"),
		Title:        title,
		Pages:        pages,
	}

	// Create a new, blank HTML file.
	newFile, err := os.Create(homePage.HTMLPagePath)
	if err != nil {
		panic(err)
	}

	t := template.Must(template.New("home.tmpl").ParseFiles(templateDir + "home.tmpl"))
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
