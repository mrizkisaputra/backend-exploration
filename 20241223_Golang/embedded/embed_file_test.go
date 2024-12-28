package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

//go:embed file/config.txt
var config string

func TestEmbedFile(t *testing.T) {
	fmt.Println(config)
}

//go:embed file/images.png
var image []byte

func TestEmbedImage(t *testing.T) {
	if err := os.WriteFile("logo.png", image, fs.ModePerm); err != nil {
		panic(err)
	}
}

//go:embed templates/index.html
//go:embed templates/header.html
var files embed.FS

func TestEmbedMultipleFile(t *testing.T) {
	bytes, _ := files.ReadFile("templates/index.html")
	fmt.Println(string(bytes))
}


//go:embed templates/*.html
var templates embed.FS

func TestEmbedMultipleWithPathMather(t *testing.T) {
	
}
