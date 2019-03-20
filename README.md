# Noire [![GoDoc](https://godoc.org/github.com/teacat/noire?status.svg)](https://godoc.org/github.com/teacat/noire) [![Coverage Status](https://coveralls.io/repos/github/teacat/noire/badge.svg?branch=master)](https://coveralls.io/github/teacat/noire?branch=master) [![Build Status](https://travis-ci.org/teacat/noire.svg?branch=master)](https://travis-ci.org/teacat/noire) [![Go Report Card](https://goreportcard.com/badge/github.com/teacat/noire)](https://goreportcard.com/report/github.com/teacat/noire)

支援 RGB、HSL、HSV、CMYK、Hex、HTML 顏色代碼進行轉換與顏色（亮度、飽和度…等）編輯的套件。

## 這是什麼？

Trillium 是一個基於 TeaCat 所需而提出的分散式不重複唯一編號演算規則，其方式與 Twitter 所設計的 [Snowflake 雪花編號](https://developer.twitter.com/en/docs/basics/twitter-ids.html)類似。

* 不須要中心化服務，任何一個服務都能獨立產生唯一編號。
* 基於時間順序而排定的編號，能夠更方便排序。
* 無流水號問題而能被得知總體數量或是資料被電腦程式爬取。
* 單個服務每秒可以產生 100,000 個唯一編號；若額度耗盡，將會暫停動作並延遲到下一秒才繼續配發唯一編號。
* 最高可高達 99,999 個服務同時使用，或是 9,999 個（隨機碰撞較安全的範圍）。
* 可用編號時間配發時間高達 292 年。

簡單來說：

> Trillium 最可以同時執行在至少 9,999 個服務中；每個服務每秒最高可以處理 100,000 個請求（略估為每毫秒 100 個請求）；照這個方式下去，編號將能持續提供到 292 年後。

## 效能比較

下列效能測試會因為受到每秒僅能產生 100,000 個唯一編號導致延遲推後而有所影響。

```
測試規格：
4.2 GHz Intel Core i7 (8750H)
32 GB 2666 MHz DDR4

goos: linux
goarch: amd64
pkg: github.com/teacat/trillium
BenchmarkString-12        300000              7512 ns/op             112 B/op          9 allocs/op
BenchmarkInt-12           300000              9853 ns/op             160 B/op         10 allocs/op
PASS
ok      github.com/teacat/trillium      5.350s
```

## 安裝方式

打開終端機並且透過 `go get` 安裝此套件即可。

```bash
$ go get github.com/teacat/trillium
```

## 使用方式

透過 `trillium.New` 建立一個新的唯一編號產生器，並且以 `Generate` 來產生。

```go
package main

import (
	"fmt"

	"github.com/teacat/trillium"
)

func main() {
	t := trillium.New(trillium.DefaultConfig())  // 傳入 `0` 會採用預設的起始日期，亦能自訂。
	fmt.Println(t.Generate().Int())              // 輸出：647219540794334229
}
```

## 構造

Trillium 只能執行在 64 位元的電腦中，因為其編號長度高達 20 字元寬度。

```txt
   已過時間      機器隨機編號    流水編號
+------------+-------------+---------+
| 1547491194 |    61835    |  01824  |  = "15474911946183501824"
+------------+-------------+---------+
    10 字元        5 字元      5 字元
```

## 問與答

**問：如何確保不同服務不會產生重複的編號？**

答：當 Trillium 被建立時，會自動產生長度為 5 的隨機數字來作為服務的唯一辨識號碼並避免與其他正在執行的服務產生出相同的編號。

**問：看起來編號中有隨機要素存在，那麼為什麼這個編號還可以按照時間排序？**

答：因為 Trillium 以時間（秒數）為主並將其擺放在開頭，所以就算後面的要素導致亂序，但因為最前面是以時間開頭為基準，所以這個唯一編號仍能保持有序。Trillium 是以秒數為基準單位，因此排序僅能精準到「秒」，無法算出到「毫秒」的排序（_不過也夠大部分場合使用了_）。

**問：以數字編號作為唯一編號，會很容易被流水號攻擊、得知資料總筆數量，或是被機器人爬取資料嗎？**

答：不會。Trillium 每秒就會更改唯一編號的規則（開頭秒數的異動），加上不同服務有著不同的隨機編號，因此無法透過遞增的方式來爬取、或是得知資料的總筆數量為何。

**問：唯一編號中如果有時間要素，是不是就能猜測到資料建立的日期與時間？**

答：這點是肯定的。但以自動遞增的編號索引也能夠透過增加的「速率」來推算出資料建立的日期，因此無論是何種方式都能推算出大略時間（_除非是完全隨機字串，但如此一來就無法排序_）。

## 可參考文件

[分布式系统中 Unique ID 的生成方法](https://darktea.github.io/notes/2013/12/08/Unique-ID)

[扯扯ID - 掘金](https://juejin.im/post/593d0821128fe1006ae47e3c)