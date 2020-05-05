package chapter9

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// プロンプトを表示するだけのやつ
func DoPrompt() {
	ch := input(os.Stdin)

	for {
		fmt.Print(">")

		// チャネルからの受信する
		if text := <-ch; text == "exit" {
			fmt.Println("Exited from my prompt.")
			break
		} else {
			fmt.Println(text)
		}
	}
}

func input(r io.Reader) <-chan string {
	ch := make(chan string)

	go func() {
		s := bufio.NewScanner(r)

		for s.Scan() { // os.Stdinの場合はScanするとinputを待ち受ける
			if text := s.Text(); text == "exit" {
				ch <- text
				break
			} else {
				ch <- text
			}
		}
		close(ch)
	}()

	return ch
}
