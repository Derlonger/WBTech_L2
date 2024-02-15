package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

/*
=== Утилита wget ===

# Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
func download(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Ошибка при запросе", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Ошибка при запросе. Код статуса: ", resp.StatusCode)
	}

	fileName := path.Base(url)

	err = os.Mkdir(fileName, os.ModePerm)
	if err != nil {
		log.Println("Ошибка при создании директории:", err)
	}

	filePath := path.Join(fileName, "index.html")
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Ошибка при копирование данных:", err)
		return
	}

	fmt.Printf("Сайт успешно скачан в директорию: %s\n", fileName)

}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Использование: go run main.go <URL>")
		os.Exit(1)
	}

	url := os.Args[1]
	download(url)

}
