Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
test() вернет 2, так как именованные значения могут быть изменены отложенными вызовами (defer).
anotherTest() вернет 1, так как x возвращается как неименованный результат, его значение не изменяется после возврата.
```
