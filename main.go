package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func formato(file fs.FileInfo) (string, error) {
	//var format []rune
	if file.IsDir() {
		return "", errors.New("No tiene formato ES UN DIRECTORIO!")
	}
	nombre := file.Name()
	indicePunto := strings.LastIndex(nombre, ".")
	formato := nombre[indicePunto+1:]
	if formato == "" || formato == " " {
		return "", errors.New("No tiene formato")
	}
	return formato, nil
}
func archivosDelDIrectoriO(directory string) []fs.FileInfo {
	dir, _ := os.Open(directory)
	fs, _ := dir.Readdir(0)

	return fs
}
func crearCarpetaSiNoExiste(dir string) {
	_, err := os.Stat(dir)

	if err != nil {
		os.Mkdir(dir, os.FileMode(0755))
	}
}
func moverArchivoAlaCarpeta(file fs.FileInfo, format string) {
	direccion := "/home/ezequiel-k/Descargas/" + format + "/"
	crearCarpetaSiNoExiste(direccion)
	fmt.Println("/home/ezequiel-k/Descargas/" + file.Name())
	fmt.Println(direccion + file.Name())
	err := os.Rename("/home/ezequiel-k/Descargas/"+file.Name(), direccion+file.Name())
	if err != nil {
		panic(err)
	}
}

func main() {
	//args := os.Args[1:]
	files := archivosDelDIrectoriO("/home/ezequiel-k/Descargas")
	for _, fileInfo := range files {
		f, err := formato(fileInfo)
		if err != nil {
			continue
		}
		fmt.Printf("El archivo %f fue organizado en %d\n", fileInfo.Name(), f)
		moverArchivoAlaCarpeta(fileInfo, f)
	}
}
