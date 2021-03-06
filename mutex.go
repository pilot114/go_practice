package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

// пример с контролем целостности при конкурентном доступе
// проверить наличие гонки:
// go run --race .\mutex.go .\files\text.txt .\files\text2.txt

func main() {
	var wg sync.WaitGroup
	w := newWords()
	for _, f := range os.Args[1:] {
		wg.Add(1)
		go func(file string) {
			if err := tallyWords(file, w); err != nil {
				fmt.Println(err.Error())
			}
			wg.Done()
		}(f)
	}
	wg.Wait()
	fmt.Println("Words that appear more than once:")
	// Т.к. в структуре есть мьютекс, её можно блочить (тут - на чтение)
	w.Lock()
	for word, count := range w.found {
		if count > 1 {
			fmt.Printf("%s: %d\n", word, count)
		}
	}
	w.Unlock()
}

type words struct {
	sync.Mutex // общие данные содержат мьютекс
	found      map[string]int
}

func newWords() *words {
	return &words{found: map[string]int{}}
}

func (w *words) add(word string, n int) {
	// "забронировать" данные можно в самом начале функции,
	// где данные используются, для наглядности (тут - на запись)
	w.Lock()
	defer w.Unlock()
	count, ok := w.found[word]
	if !ok {
		w.found[word] = n
		return
	}
	w.found[word] = count + n
}

func tallyWords(filename string, dict *words) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		dict.add(word, 1)
	}
	return scanner.Err()
}
