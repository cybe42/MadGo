package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"

	"mime"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const (
	defaultFileType = "text/plain; charset=utf-8" // The default file type if we fail to parse from extension.
)

func getFileExtension(filePath string) string {
	var fileSplitted []string = strings.Split(filePath, ".")
	var splittedLength int = len(fileSplitted)
	if splittedLength < 2 {
		return ""
	} else {
		return "." + fileSplitted[splittedLength-1] // get last element from array
	}
}

func getFileType(filePath string) string {
	var fileExtension string = getFileExtension(filePath)
	var fileType string = mime.TypeByExtension(fileExtension)
	if fileType == "" {
		return defaultFileType // return file type if we could not find file extension
	} else {
		return fileType
	}
}

func cropFileExtension(filePath string) string {
	return filePath[:len(filePath)-len(getFileExtension(filePath))]
}

func splitPathSeperator(path string) []string {
	return strings.Split(path, string(os.PathSeparator))
}

func pathSeperatorJoiner(splitted []string) string {
	return strings.Join(splitted, string(os.PathSeparator))
}

func decideURLPath(urlPath string) string {
	var filebase string = filepath.Base(urlPath)
	var dirbase string = filepath.Dir(urlPath)
	switch filebase {
	case "index.html":
		return "/"
	default:
		if getFileExtension(filebase) == ".html" { // remove .html from path
			return "/" + filepath.Join(dirbase, cropFileExtension(filebase))
		} else { // else just serve file regularly
			return "/" + urlPath
		}
	}
}

//go:embed frontend/dist/*
var dist embed.FS

type fileHandler struct {
	app *fiber.App
}

func (handler *fileHandler) addFileHandle(urlPath string, filePath string) {
	fmt.Println("Adding handle --> ", urlPath)
	file, _ := dist.ReadFile(filePath)
	var fileContent string = string(file)

	handler.app.Get(urlPath, func(c *fiber.Ctx) error {
		c.Set("Content-Type", getFileType(filePath))
		return c.SendString(fileContent)
	}) // serve static file
}

func (handler *fileHandler) addFileHandleRecursively(basePath string, path string, entry fs.DirEntry) {
	if entry.IsDir() {
		distDir, _ := dist.ReadDir(filepath.Join(path, entry.Name()))
		for _, direntry := range distDir {
			handler.addFileHandleRecursively(basePath, filepath.Join(path, entry.Name()), direntry)
		}
	} else {
		filePath := filepath.Join(path, entry.Name())
		oldURLPath := pathSeperatorJoiner(splitPathSeperator(filePath)[len(splitPathSeperator(basePath)):])
		newURLPath := decideURLPath(oldURLPath)
		handler.addFileHandle(newURLPath, filePath)
	}
}

func Static() *fiber.App {
	app := fiber.New()
	var handler fileHandler = fileHandler{app: app}
	path := "frontend/dist"
	distDir, _ := dist.ReadDir(path)
	for _, entry := range distDir {
		handler.addFileHandleRecursively(path, path, entry)
	}

	return app
}
