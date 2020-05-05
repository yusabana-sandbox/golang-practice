package chapter9

import (
	"fmt"
	"sync"
	"time"
)

func DoLock() {
	var m sync.Mutex // ゼロ値でそのまま扱える
	m.Lock()
	go func() {
		time.Sleep(2 * time.Second)
		m.Unlock()
		fmt.Println("unlock 1")
	}()
	m.Lock() // ここでブロック
	m.Unlock()
	fmt.Println("unlock 2")
}
