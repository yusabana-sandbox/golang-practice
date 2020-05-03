package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var optN = flag.Bool("n", false, "行番号")

func main() {
	println("CATコマンドです")

	flag.Parse()
	for _, fn := range flag.Args() {
		err := readFile(fn, *optN)
		if err != nil { log.Fatal(err) }
	}
}

func readFile(fn string, optN bool) (error) {
	sf, err := os.Open(fn)
	if err != nil { return err }
	defer sf.Close() // fileハンドラを閉じるのを予約

	num := 1
	scanner := bufio.NewScanner(sf)
	for scanner.Scan() {
		text := scanner.Text()

		if optN {
			fmt.Fprintln(os.Stdout, num, ":", text)
		} else {
			fmt.Fprintln(os.Stdout, text)
		}

		num++
	}
	if err := scanner.Err(); err != nil { return err }


	return nil
}
