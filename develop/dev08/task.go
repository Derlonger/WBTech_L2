package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

// Основная функция программы
func main() {
	// Создание нового читателя для чтения ввода пользователя
	reader := bufio.NewReader(os.Stdin)

	for {
		// Вывод приглашения для ввода команды
		fmt.Print("> ")
		// Чтение строки ввода от пользователя
		input, err := reader.ReadString('\n')
		if err != nil {
			// Проверка на конец файла (Ctrl+D)
			if err == io.EOF {
				fmt.Println("До свидания!")
				os.Exit(0)
			}
			// В случае другой ошибки выводим сообщение
			fmt.Fprintln(os.Stderr, "Ошибка чтения ввода:", err)
			continue
		}

		// Удаление символа новой строки из ввода
		input = strings.TrimSuffix(input, "\n")

		// Разбивка введенной строки на команды
		commands := strings.Fields(input)

		if len(commands) == 0 {
			continue
		}

		// Проверка на встроенные команды
		switch commands[0] {
		case "cd": // Смена текущей директории
			if len(commands) < 2 {
				fmt.Println("Используйте: cd <директория>")
			} else {
				err := os.Chdir(commands[1])
				if err != nil {
					fmt.Fprintln(os.Stderr, "Ошибка при смене директории:", err)
				}
			}
		case "pwd": // Вывод текущей директории
			dir, err := os.Getwd()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка при получении текущей директории:", err)
			} else {
				fmt.Println(dir)
			}
		case "echo": // Вывод аргументов
			fmt.Println(strings.Join(commands[1:], " "))
		case "kill": // Отправка сигнала завершения процессу
			if len(commands) < 2 {
				fmt.Println("Используйте: kill <PID>")
			} else {
				pid := commands[1]
				err := exec.Command("kill", pid).Run()
				if err != nil {
					fmt.Fprintln(os.Stderr, "Ошибка при отправке сигнала:", err)
				}
			}

		case "ps": // Вывод списка процессов
			cmd := exec.Command("ps", "aux")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка при выполнении команды ps:", err)
			}
		case "fork-exec": // Запуск процесса
			if len(commands) < 2 {
				fmt.Println("Используйте: fork-exec <команда>")
			} else {
				cmd := exec.Command(commands[1], commands[2:]...)
				err := cmd.Start()
				if err != nil {
					fmt.Fprintln(os.Stderr, "Ошибка при запуске процесса:", err)
				}
			}

		default: // Выполнение внешней команды
			cmd := exec.Command(commands[0], commands[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка при выполнении команды:", err)
			}
		}
	}
}
