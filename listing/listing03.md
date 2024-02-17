Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
<nil>
false

Функция Foo() (error) возвращает пустой интерфейс имеющий тип *os.PathError имеющее значение nil,
При сравнении nil и возвращаемый параметр функции мы получим false так как интерфейс является уже не пустым.
```
