package main

import (
	"errors"
	"flag"
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
func moverArchivoAlaCarpeta(file fs.FileInfo, format string, dir string) {
	direccion := dir + "/" + format + "/"
	crearCarpetaSiNoExiste(direccion)
	fmt.Println(dir + file.Name())
	fmt.Println(direccion + file.Name())
	err := os.Rename(dir+file.Name(), direccion+file.Name())
	if err != nil {
		panic(err)
	}
}

func main() {
	var direction string
	directory := flag.String("dir", "directory", "Directory to organize")
	flag.Parse()
	if *directory == "." {
		direction, _ = os.Getwd()
	} else {
		direction = *directory
	}
	fmt.Println("Organizando los elementos del directorio: ", direction)
	files := archivosDelDIrectoriO(direction)
	i := len(files)
	fmt.Printf("Organizando %d archivos", i)
	for _, fileInfo := range files {
		f, err := formato(fileInfo)
		i--
		if err != nil {
			continue
		}
		fmt.Printf("El archivo nro %d, %s fue organizado en %s\n", i, fileInfo.Name(), f)
		moverArchivoAlaCarpeta(fileInfo, f, direction)
	}
}
