package main

import (
	"fmt"
	"os"
	"time"
)

// каналы в go - чтото вроде типизированных сокетов для передачи данных
// между корутинами
func main() {
	// полезно в случаях динамического создания корутин и каналов
	// для избегания утечек, закрывать каналы отправки, используя специальный канал (done)
	done := make(chan bool)
	until := time.After(3 * time.Second)
	// важно - буферизованные каналы зачастую спасают от дедлоков!
	// буферезированный канал автоматически ставит блокировки, аналогично Mutex,
	// если в нем нет свободного места
	message := make(chan []byte, 1)

	go readStdin(message, done)

	for {
		select {
		case buf := <-message:
			os.Stdout.Write(buf)
		case <-until:
			done <- true
			time.Sleep(500 * time.Millisecond)
			return
		}
	}
}

func readStdin(message chan<- []byte, done <-chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("Done")
			close(message)
			return
		default:
			data := make([]byte, 1)
			l, _ := os.Stdin.Read(data)
			if l > 0 {
				message <- data
			}
		}
	}
}
