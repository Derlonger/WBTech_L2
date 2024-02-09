package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	tests := []string{"a4bc2d5e", "abcd", "45", "", "qwe\\4\\5", "qwe\\45", "qwe\\\\5"}
	for _, test := range tests {
		unpacked, err := UnpackString(test)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
		} else {
			fmt.Printf("Исходная строка: %s, Распакованная строка: %s\n", test, unpacked)
		}
	}
}

// UnpackString выполняет примитивную распаковку строки
func UnpackString(str string) (string, error) {
	var result strings.Builder
	var prev rune
	esc := false

	// Проходим по каждому символу в строке
	for _, char := range str {
		// Если символ является цифрой и не находится в экранированном состоянии,
		// то распаковываем предыдущий символ нужное количество раз
		if unicode.IsDigit(char) && !esc {
			// Если предыдущий символ не был установлен, это некорректная строка
			if prev == 0 {
				return "", errors.New("некорректная строка")
			}
			// Преобразуем символ цифры в число
			n, _ := strconv.Atoi(string(char))
			// Добавляем предыдущий символ нужное количество раз в результат
			for i := 0; i < n-1; i++ {
				result.WriteRune(prev)
			}
			continue
		}

		// Если символ - обратный слеш и не находится в экранированном состоянии,
		// переключаем экранированное состояние
		if char == '\\' && !esc {
			esc = true
			continue
		}

		// Если предыдущий символ был установлен, добавляем его в результат
		if prev != 0 {
			result.WriteRune(prev)
		}

		// Устанавливаем текущий символ как предыдущий для следующей итерации
		prev = char
		// Сбрасываем экранированное состояние
		esc = false
	}

	// Добавляем последний символ в результат, если он был установлен
	if prev != 0 {
		result.WriteRune(prev)
	}

	// Возвращаем результат в виде строки и ошибку (если есть)
	return result.String(), nil
}
