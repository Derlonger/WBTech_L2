Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
Вывод будет таким 
0
1
2
3
4
5
6
7
8
9
deadlock
Происходит это потому что цикл for range пытается читать канал пока он не закроется.
```
