# Noire [![GoDoc](https://godoc.org/github.com/teacat/noire?status.svg)](https://godoc.org/github.com/teacat/noire) [![Coverage Status](https://coveralls.io/repos/github/teacat/noire/badge.svg?branch=master#1)](https://coveralls.io/github/teacat/noire?branch=master) [![Build Status](https://travis-ci.com/teacat/noire.svg?branch=master#1)](https://travis-ci.com/teacat/noire) [![Go Report Card](https://goreportcard.com/badge/github.com/teacat/noire#1)](https://goreportcard.com/report/github.com/teacat/noire)

A color library which supports converting between the RGB, HSL, HSV, CMYK, Hex, HTML and some additional functions (tint, saturation).

Requires atleast **Go 1.10** version due to the `math.Round`(https://golang.org/pkg/math/#Round) function call.

## Supported colors

Noire is able to convert the colors between:

-   RGB
-   CMYK
-   HSL
-   HSV
-   Hex
-   HTML

## Benchmark

```
Specification:
4.2 GHz Intel Core i7 (8750H)
32 GB 2666 MHz DDR4

goos: windows
goarch: amd64
pkg: github.com/teacat/noire
BenchmarkCMYKToRGB-12         	100000000	        22.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkRGBToCMYK-12         	50000000	        26.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkRGBToHSL-12          	50000000	        26.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkHSLToRGB-12          	100000000	        17.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkHSVToRGB-12          	100000000	        15.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkRGBToHSV-12          	50000000	        29.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkRGBToHex-12          	20000000	        85.4 ns/op	      32 B/op	       4 allocs/op
BenchmarkHexToRGB-12          	50000000	        36.1 ns/op	       8 B/op	       1 allocs/op
BenchmarkHTMLToRGBName-12     	20000000	       118 ns/op	      40 B/op	       3 allocs/op
BenchmarkHTMLToRGBHex-12      	30000000	        41.0 ns/op	       8 B/op	       1 allocs/op
BenchmarkRGBToHTML-12         	20000000	       103 ns/op	      32 B/op	       4 allocs/op
BenchmarkMix-12               	 5000000	       292 ns/op	     112 B/op	       5 allocs/op
BenchmarkHue-12               	50000000	        33.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkSaturation-12        	50000000	        33.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkLightness-12         	50000000	        33.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkAdjustHue-12         	20000000	        79.0 ns/op	      32 B/op	       1 allocs/op
BenchmarkLighten-12           	20000000	        78.2 ns/op	      32 B/op	       1 allocs/op
BenchmarkDarken-12            	20000000	        77.8 ns/op	      32 B/op	       1 allocs/op
BenchmarkSaturate-12          	20000000	        74.1 ns/op	      32 B/op	       1 allocs/op
BenchmarkDesaturate-12        	20000000	        79.0 ns/op	      32 B/op	       1 allocs/op
BenchmarkGrayscale-12         	20000000	        71.8 ns/op	      32 B/op	       1 allocs/op
BenchmarkComplement-12        	20000000	        79.9 ns/op	      32 B/op	       1 allocs/op
BenchmarkTint-12              	30000000	        40.8 ns/op	      32 B/op	       1 allocs/op
BenchmarkShade-12             	30000000	        40.9 ns/op	      32 B/op	       1 allocs/op
BenchmarkInvert-12            	50000000	        24.9 ns/op	      32 B/op	       1 allocs/op
BenchmarkLuminanaceWCAG-12    	10000000	       224 ns/op	       0 B/op	       0 allocs/op
BenchmarkLuminanace-12        	300000000	         6.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkContrast-12          	 5000000	       257 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsLight-12           	2000000000	         0.26 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsDark-12            	2000000000	         0.26 ns/op	       0 B/op	       0 allocs/op
BenchmarkHSV-12               	50000000	        34.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkHSVA-12              	50000000	        35.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkHSL-12               	50000000	        31.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkHSLA-12              	50000000	        32.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkRGB-12               	2000000000	         0.52 ns/op	       0 B/op	       0 allocs/op
BenchmarkRGBA-12              	2000000000	         0.26 ns/op	       0 B/op	       0 allocs/op
BenchmarkCMYK-12              	50000000	        32.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkHex-12               	20000000	        90.1 ns/op	      32 B/op	       4 allocs/op
BenchmarkHTMLHex-12           	10000000	       146 ns/op	      40 B/op	       5 allocs/op
BenchmarkHTMLName-12          	20000000	       107 ns/op	      32 B/op	       4 allocs/op
BenchmarkHTMLRGBA-12          	 2000000	       839 ns/op	      96 B/op	       5 allocs/op
PASS
ok  	github.com/teacat/noire	67.640s
Success: Benchmarks passed.
```

## Installation

To install Noire by simply typing `go get` in the terminal.

```bash
$ go get github.com/teacat/noire
```

## Usage

Initialize a new color with `noire.NewRGB` (or `NewHex`) to modify the color with `Lighten` or `Tint`, etc.

```go
package main

import (
	"fmt"

	"github.com/teacat/noire"
)

func main() {
	c := noire.NewRGB(255, 255, 255)
	fmt.Println(c.Invert().Hex())       // Output: 000000
	fmt.Println(c.Invert().HTML())      // Output: Black
	fmt.Println(c.Lighten(1).RGB())     // Output: 255, 255, 255
}
```

## Description

There are few functions results cannot be visualized, so make sure to check the [GoDoc](https://godoc.org/github.com/teacat/noire) to see how they work.

-   `Hue`: Get the Hue angle of the current color based on the HSL algorithm.
-   `Saturation`: Get the Saturation of the current color based on the HSL algorithm.
-   `Lightness`： Get the Lightness of the current color based on the HSL algorithm.
-   `LuminanaceWCAG`：Get the Luminance of the current color based on the WCAG 2.0 algorithm.
-   `Luminanace`: Get the Luminance of the current color.
-   `Contrast`: Get the Contrast of the current color based on the WCAG Luminance algorithm.
-   `IsLight`: Returns true if the color is a light scheme, it might not be the same as what human eyes can see.
-   `IsDark`: Returns true if the color is a dark scheme, it might not be the same as what human eyes can see.

### Lighten

![Result preview](./assets/lighten.png)

Lighten a color based on HSL mode, it might makes the color a bit way too bright or washed out.

```go
func main() {
	c := NewRGB(219, 112, 148)
	fmt.Println(c.Lighten(0.15).Hex())   // Output: EAADC2
}
```

### Brighten

![Result preview](./assets/brighten.png)

Increases the brightness of the color based on RGB mode.

```go
func main() {
	c := NewRGB(0, 0, 0)
	fmt.Println(c.Brighten(0.1).Hex())   // Output: 1A1A1A
}
```

### Tint

![Result preview](./assets/tint.png)

Mixing with a white color as base to get the best balance to increase the brightness of a color.

```go
func main() {
	c := NewRGB(0, 0, 0)
	fmt.Println(c.Tint(0.1).Hex())   // Output: 1A1A1A
}
```

### Darken

![Result preview](./assets/darken.png)

Darken a color based on HSL mode, it might makes the color a bit way too dark or dimmed.

```go
func main() {
	c := NewRGB(219, 112, 148)
	fmt.Println(c.Darken(0.15).Hex())   // Output: CB3366
}
```

### Shade

![Result preview](./assets/shade.png)

Mixing with a black color as base to get the best balance to increase the brightness of a color.

```go
func main() {
	c := NewRGB(219, 112, 148)
	fmt.Println(c.Shade(0.15).Hex())   // Output: BA5F7E
}
```

### Saturate

![Result preview](./assets/saturate.png)

Increases the saturation of the color based on HSL mode.

```go
func main() {
	c := NewRGB(219, 112, 148)
	fmt.Println(c.Saturate(0.5).Hex())   // Output: FF4C88
}
```

### Desaturate

![Result preview](./assets/desaturate.png)

Decreases the saturation of the color based on HSL mode.

```go
func main() {
	c := NewRGB(219, 112, 148)
	fmt.Println(c.Desaturate(0.15).Hex())   // Output: AE9DA3
}
```

### AdjustHue

![Result preview](./assets/adjust-hue.png)

Rotates the Hue angle of the color based on HSL mode, it still goes clockwise if the value was set over than 360 degree.

```go
func main() {
	c := NewRGB(219, 112, 148)
	fmt.Println(c.AdjustHue(30).Hex())   // Output: DB8270
}
```

### Mix

![Result preview](./assets/mix.png)

Mixing two colors with a value that specified the weight of the second color.

```go
func main() {
	c1 := NewHex("F00")
	c2 := NewHex("00F")
	fmt.Println(c1.Mix(c2, 0.5).HTML())   // Output: Purple
}
```

### Invert

![Result preview](./assets/invert.png)

Get the opposite color that based on the RGB color map (it's not a complementary color).

```go
func main() {
	c := NewRGB(219, 112, 148)
	fmt.Println(c.Invert().Hex())   // Output: 248F6B
}
```

### Complement

![Result preview](./assets/complement.png)

Get the complementary color of the current color, same as `AdjustHue(180)`.

```go
func main() {
	c := NewRGB(219, 112, 148)
	fmt.Println(c.Complement().Hex())   // Output: 70DBB7
}
```

### Grayscale

![Result preview](./assets/grayscale.png)

Converts the color to grayscale, same as `Desaturate(1)`.

```go
func main() {
	c := NewRGB(219, 112, 148)
	fmt.Println(c.Grayscale().Hex())   // Output: A5A5A5
}
```

### Foreground

![Result preview](./assets/foreground.png)

Get the suggested foreground color by calculating the color luminance, it returns a white color when the color is dark, vise versa.

```go
func main() {
	c := NewHTML("Green")
	fmt.Println(c.Foreground().HTML())   // Output: White
	c = NewHTML("Red")
	fmt.Println(c.Foreground().HTML())   // Output: White
	c = NewHTML("Yellow")
	fmt.Println(c.Foreground().HTML())   // Output: Black
}
```

## References

[RGB 轉 HSV, HSL (線上色碼轉換 HSL, HSV, RGB, HEX)](https://www.ginifab.com.tw/tools/colors/rgb_to_hsv_hsl.html)

[ozdemirburak/iris: PHP library for color manipulation and conversion.](https://github.com/ozdemirburak/iris)

[G17: Ensuring that a contrast ratio of at least 7:1 exists between text (and images of text) and background behind the text | Techniques for WCAG 2.0](https://www.w3.org/TR/WCAG20-TECHS/G17.html#G17-tests)

[Using Sass to automatically pick text colors](https://medium.com/dev-channel/using-sass-to-automatically-pick-text-colors-4ba7645d2796)

[Relative luminance - Wikipedia](https://en.wikipedia.org/wiki/Relative_luminance)

[user interface - Given a background color, how to get a foreground color that makes it readable on that background color? - Stack Overflow](https://stackoverflow.com/questions/3116260/given-a-background-color-how-to-get-a-foreground-color-that-makes-it-readable-o)

[image - Formula to determine brightness of RGB color - Stack Overflow](https://stackoverflow.com/questions/596216/formula-to-determine-brightness-of-rgb-color)

[Ant Design 色板生成算法演进之路 - 知乎](https://zhuanlan.zhihu.com/p/32422584)

[Sass 基础——颜色函数\_Preprocessor, Sass, SCSS 教程\_w3cplus](https://www.w3cplus.com/preprocessor/sass-color-function.html)
