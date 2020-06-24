package main

import (
	"fmt"
	"runtime"
	"sync"
)

// https://github.com/golang/go/wiki/LearnConcurrency
func main() {
	fmt.Println("start")
	go func() {
		fmt.Println("inside")
	}()
	fmt.Println("finish")
	// просим планировщик подождать окончания выполнения горутин
	// но если горутина сама доходит до своей блокировки, ее он уже не будет ждать
	runtime.Gosched()

	// Пример ожидания завершения с WaitGorup
	var wg sync.WaitGroup
	numbers := [3]int{1, 2, 3}
	for _, n := range numbers {
		// +1 в счетчик ожидания
		wg.Add(1)
		// не забываем динамически прокинуть параметры
		go func(n int) {
			fmt.Println("inside ", n)
			// -1 в счетчик ожидания
			wg.Done()
		}(n)
	}
	// остановка, когда счетчик ожидания сойдется
	wg.Wait()
}
