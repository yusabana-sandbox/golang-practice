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

	doGoroutinChannel()
}

func doGoroutinChannel() {
	urls := []string{
		"http://example.com",
		"http://example.net",
		"http://example.org",
	}

	// 直接チャネルを処理する書き方
	//// channelを作る
	//statusChan := make(chan string)
	//
	//for _, url := range urls {
	//	go func(url string) {
	//		res, err := http.Get(url)
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//		defer res.Body.Close()
	//
	//		statusChan <- fmt.Sprintf("%s %s", res.Status, url)
	//	}(url)
	//}
	// 関数に切り出して返り値として該当のチャネルを受け取る書き方
	statusChan := getStatus(urls)

	// <-statusChanとチャネルを読み込みする際処理がブロックされるのでwithGroupのような待ち合わせは不要
	// 3個のリクエストを投げているから3回チャネルからのデータを待ち合わせておけばOK
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-statusChan)
	}
}

// チャネルを返す関数
// <-chan string のような戻り値の型を指定し、読み出し専用のチャネルにすることで不用意にchannelへの書き込みを出来ないようにする
func getStatus(urls []string) <-chan string {
	// channelを作る
	statusChan := make(chan string)

	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			statusChan <- fmt.Sprintf("%s %s", res.Status, url)
		}(url)
	}
	return statusChan
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
