package chapter5

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func Do() {
	fmt.Println("chapter5はじまるよーーーー")

	doGoroutin()
}

func doGoroutin() {
	wait := new(sync.WaitGroup)
	urls := []string{
		"http://example.com",
		"http://example.net",
		"http://example.org",
	}

	for _, url := range urls {
		// waitGroupに追加
		wait.Add(1)

		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}

			defer res.Body.Close()
			fmt.Println(url, res.Status)

			// waitGroupから削除
			wait.Done()
		}(url)
	}

	// 全部が完了するまで待ち合わせる
	wait.Wait()
}
