package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func sortWord(word string) string {
	runes := []rune(word)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func findAnagrams(words *[]string) map[string][]string {
	mapAnagrams := make(map[string][]string)
	for _, word := range *words {
		word = strings.ToLower(word)
		sortedWord := sortWord(word)
		if _, ok := mapAnagrams[sortedWord]; ok {
			mapAnagrams[sortedWord] = append(mapAnagrams[sortedWord], word)
		} else {
			mapAnagrams[sortedWord] = []string{word}
		}
	}

	for key, value := range mapAnagrams {
		if len(value) == 1 {
			delete(mapAnagrams, key)
		} else {
			sort.Strings(value)
			mapAnagrams[key] = value
		}
	}
	return mapAnagrams
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	result := findAnagrams(&words)
	fmt.Println(result)
}
