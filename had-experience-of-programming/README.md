# README


## サンプルコード

- 各chapter2 - chapter5 については以下の記事を動作させている
  - [はじめてのGo―シンプルな言語仕様，型システム，並行処理：特集｜gihyo.jp … 技術評論社](https://gihyo.jp/dev/feature/01/go_4beginners)
- 以下の `hoge` のmain関数の実行によって動作確認できる


## Usage build

- 参考
    - https://github.com/Songmu/horenso/blob/master/Makefile
    - https://qiita.com/kitsuyui/items/d03a9de90330d8c275c8

go build に `-w -s` などのオプションを付けてもいいかも


## hogeのビルド

```
$ go build ./cmd/hoge
```

- hogeは以下のサイトのサンプルを実行したもの
    - https://gihyo.jp/dev/feature/01/go_4beginners
    - わかりやすい（5章のゴルーチンとチャネルは特に）

## fugaのビルド

```
$ go build ./cmd/fuga
```

- fugaは以下のサンプルを実行したもの
    - [プログラミング経験者がGo言語を本格的に勉強する前に読むための本](https://www.amazon.co.jp/dp/B06XJ86BFZ)

## piyoのビルド

```
$ go build ./cmd/piyo
```

- piyoは以下のサンプルを実行したもの
    - [実践Go言語 - golang.jp](http://golang.jp/effective_go)
