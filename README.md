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

##

![Noire 可用色相參考表](https://user-images.githubusercontent.com/7308718/54680503-fdd58900-4b44-11e9-952d-c464fbc5252e.png)

## 可參考文件

[RGB 轉 HSV, HSL (線上色碼轉換 HSL, HSV, RGB, HEX)](https://www.ginifab.com.tw/tools/colors/rgb_to_hsv_hsl.html)

[ozdemirburak/iris: PHP library for color manipulation and conversion.](https://github.com/ozdemirburak/iris)

[G17: Ensuring that a contrast ratio of at least 7:1 exists between text (and images of text) and background behind the text | Techniques for WCAG 2.0](https://www.w3.org/TR/WCAG20-TECHS/G17.html#G17-tests)

[Using Sass to automatically pick text colors](https://medium.com/dev-channel/using-sass-to-automatically-pick-text-colors-4ba7645d2796)

[Relative luminance - Wikipedia](https://en.wikipedia.org/wiki/Relative_luminance)

[user interface - Given a background color, how to get a foreground color that makes it readable on that background color? - Stack Overflow](https://stackoverflow.com/questions/3116260/given-a-background-color-how-to-get-a-foreground-color-that-makes-it-readable-o)

[image - Formula to determine brightness of RGB color - Stack Overflow](https://stackoverflow.com/questions/596216/formula-to-determine-brightness-of-rgb-color)

[Ant Design 色板生成算法演进之路 - 知乎](https://zhuanlan.zhihu.com/p/32422584)

[Sass基础——颜色函数_Preprocessor, Sass, SCSS 教程_w3cplus](https://www.w3cplus.com/preprocessor/sass-color-function.html)