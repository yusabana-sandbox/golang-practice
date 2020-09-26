package chapter5

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func Do() {
	fmt.Println("chapter5はじまるよーーーー")

	//doGoroutin()
	//
	//doGoroutinChannel()
	//
	//doSelect()
	//
	doBuffer()
}

// https://gihyo.jp/dev/feature/01/go_4beginners/0005?page=3
// バッファなしチャネルは同期制御に使うことができる
// バッファ付きのチャネルはメッセージキューのような挙動になる(バッファの分は非同期にできる)
func doBuffer() {
	// チャネルをバッファ 3 として作る
	ch := make(chan string, 3)

	go func() {
		time.Sleep(time.Second)
		fmt.Println(<-ch) // 1秒後にデータを読み出す
	}()

	// バッファ分はブロックされない
	ch <- "a" // ブロックしない
	ch <- "b" // ブロックしない
	ch <- "c" // ブロックしない
	ch <- "d" // 1秒後にデータが読み出されるまでブロック

	fmt.Println("ccc")
}

func doSelect() {
	urls := []string{
		"http://example.com",
		"http://example.net",
		"http://example.org",
	}

	// 1秒後に値が読み出せるチャネル
	// time.After()関数は時間を指定するとその時間後にデータを書き込むチャネルを返す関数です
	timeout := time.After(time.Second)

	statusChan := getStatus(urls)

	// LOOPというラベルをつける
	// つけないでcaseの中でbreakを呼ぶとselect文は抜けられるが、その外側のforが抜けられないので、
	// ループのラベルを付けてそのラベルの場所から抜けられるようにしている
LOOP:
	for {
		select {
		case status := <-statusChan:
			fmt.Println(status)
		case <-timeout: // timeoutのチャネルが書き込まれたらループを抜ける
			break LOOP // このfor/selectを抜ける
		}
	}
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

// getStatus改良版、バッファを使って非同期にする方法
// いずれのゴルーチンもstatusChanに値を書き込むことで終了するけど、main側の読み取り処理が遅かった場合、
// ゴルーチンはステータスの処理が終わっているのに書き込みでブロックして閉じることが出来ない。
// この場合はstatusChanに必要な分だけのバッファをつけることでmain()の処理が遅くても、
// チャネルに値を書き込んでゴルーチンを終了できメモリ負荷を下げられる
func getStatusAdvanced(urls []string) <-chan string {
	// バッファをURLの数（3）に
	statusChan := make(chan string, len(urls))
	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			// main()の読み出しが遅くても
			// 3つのゴルーチンはすぐに終わる
			statusChan <- res.Status
		}(url)
	}
	return statusChan
}

// limitというバッファ付きのチャネルを用いて、このチャネルに値が書き込める場合はゴルーチンを起動
// ゴルーチンが終わったらlimitから値を読み出すことでゴルーチンの同時期同数を制御する
// limitの数を超えた場合はlimitへの書き込みが同期処理になるため、そこで一旦止まる感じになる
// この場合limitチャネルに入れるデータには意味がないためサイズゼロの構造体を使う
var empty struct{} // サイズがゼロの構造体
func getStatusAdvanced2(urls []string) <-chan string {
	statusChan := make(chan string, len(urls))
	// バッファを5に指定して生成
	limit := make(chan struct{}, 5)
	go func() {
		for _, url := range urls {
			select {
			case limit <- empty:
				// limitに書き込みが可能な場合は取得処理を実施
				go func(url string) {
					// このゴルーチンは同時に5つしか起動しない
					res, err := http.Get(url)
					if err != nil {
						log.Fatal(err)
					}
					statusChan <- res.Status
					// 終わったら1つ読み出して空きを作る
					<-limit
				}(url)
			}
		}
	}()
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
