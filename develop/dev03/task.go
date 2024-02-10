package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type SortConfig struct {
	keyIndex     int
	sortByNumber bool
	reverse      bool
	unique       bool
}

func main() {
	// Парсинг флагов командной строки
	config := parseFlags()

	// Чтение строк из файла
	lines, err := readLines("develop/dev03/input.txt")
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		os.Exit(1)
	}

	// Сортировка строк
	sortLines(lines, config)

	// Запись отсортированных строк в файл
	err = writeLines("develop/dev03/output.txt", lines)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
		os.Exit(1)
	}

	fmt.Println("Успешно отсортировано.")
}

// parseFlags парсит флаги командной строки и возвращает соответствующую конфигурацию сортировки
func parseFlags() SortConfig {
	keyIndex := flag.Int("k", 0, "Индекс колонки для сортировки (начиная с 0)")
	sortByNumber := flag.Bool("n", false, "Сортировать по числовому значению")
	reverse := flag.Bool("r", false, "Сортировать в обратном порядке")
	unique := flag.Bool("u", false, "Не выводить повторяющиеся строки")
	flag.Parse()

	return SortConfig{
		keyIndex:     *keyIndex,
		sortByNumber: *sortByNumber,
		reverse:      *reverse,
		unique:       *unique,
	}
}

// readLines читает строки из файла и возвращает их в виде среза
func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines записывает строки в файл
func writeLines(filename string, lines []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

// sortLines сортирует строки в соответствии с заданной конфигурацией
func sortLines(lines []string, config SortConfig) {
	sort.SliceStable(lines, func(i, j int) bool {
		fieldI := getField(lines[i], config.keyIndex)
		fieldJ := getField(lines[j], config.keyIndex)

		if config.sortByNumber {
			numI, errI := strconv.Atoi(fieldI)
			numJ, errJ := strconv.Atoi(fieldJ)

			if errI == nil && errJ == nil {
				if config.reverse {
					return numI > numJ
				}
				return numI < numJ
			}
		}

		if config.reverse {
			return fieldI > fieldJ
		}
		return fieldI < fieldJ
	})

	if config.unique {
		uniqueLines := make([]string, 0, len(lines))
		prevLine := ""
		for _, line := range lines {
			if line != prevLine {
				uniqueLines = append(uniqueLines, line)
			}
			prevLine = line
		}
		copy(lines, uniqueLines)
		lines = lines[:len(uniqueLines)]
	}
}

// getField возвращает значение поля (колонки) по указанному индексу
func getField(line string, index int) string {
	fields := strings.Fields(line)
	if index < len(fields) {
		return fields[index]
	}
	return ""
}
