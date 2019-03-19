package noire

import (
	"encoding/hex"
	"fmt"
	"math"
	"strings"
)

// Color 呈現了一個可供操作與轉換的顏色資料。
type Color struct {
	Red   float64
	Green float64
	Blue  float64
	Alpha float64
}

func newColor(r float64, g float64, b float64, a float64) *Color {
	if r > 255 {
		r = 255
	} else if r < 0 {
		r = 0
	}
	if g > 255 {
		g = 255
	} else if g < 0 {
		g = 0
	}
	if b > 255 {
		b = 255
	} else if b < 0 {
		b = 0
	}
	if a > 1 {
		a = 1
	} else if a < 0 {
		a = 0
	}
	return &Color{
		Red:   r,
		Green: g,
		Blue:  b,
		Alpha: a,
	}
}

// CMYKToRGB 能夠將 CMYK 的顏色以有損的方式轉換成 RGB。
func CMYKToRGB(c float64, m float64, y float64, k float64) (r float64, g float64, b float64) {
	c = c / 100
	m = m / 100
	y = y / 100
	k = k / 100

	r = 1 - math.Min(1, c*(1-k)+k)
	g = 1 - math.Min(1, m*(1-k)+k)
	b = 1 - math.Min(1, y*(1-k)+k)

	r = math.Round(r * 255)
	g = math.Round(g * 255)
	b = math.Round(b * 255)

	return
}

// RGBToCMYK 能夠將 RGB 的顏色以有損的方式轉換成 CMYK。
func RGBToCMYK(r float64, g float64, b float64) (c float64, m float64, y float64, k float64) {
	r = r / 255
	g = g / 255
	b = b / 255

	k = math.Min(1-r, 1-g)
	k = math.Min(k, 1-b)
	if (1 - k) == 0 {
		c = 0
		m = 0
		y = 0
	} else {
		c = (1 - r - k) / (1 - k)
		m = (1 - g - k) / (1 - k)
		y = (1 - b - k) / (1 - k)
	}
	c = math.Round(c * 100)
	m = math.Round(m * 100)
	y = math.Round(y * 100)
	k = math.Round(k * 100)
	return
}

// HueToRGB 能夠以明亮色相顏色轉換成 RGB。
func HueToRGB(p float64, q float64, t float64) float64 {
	if t < 0 {
		t++
	}
	if t > 1 {
		t--
	}
	if t < 1/6 {
		return p + (q-p)*6*t
	}
	if t < 1/2 {
		return q
	}
	if t < 2/3 {
		return p + (q-p)*(2/3-t)*6
	}
	return p
}

// RGBToHSL 能夠將 RGB 的顏色以有損的方式轉換成 HSL。
func RGBToHSL(r float64, g float64, b float64) (h float64, s float64, l float64) {
	r = r / 255
	g = g / 255
	b = b / 255
	max := math.Max(r, g)
	max = math.Max(max, b)
	min := math.Min(r, g)
	min = math.Min(min, b)
	l = (max + min) / 2
	if max == min {
		h = 0
		s = 0
	} else {
		d := max - min
		if l > 0.5 {
			s = d / (2 - max - min)
		} else {
			s = d / (max + min)
		}
		switch max {
		case r:
			if g < b {
				h = (g-b)/d + 6
			} else {
				h = (g-b)/d + 0
			}
			break
		case g:
			h = (b-r)/d + 2
			break
		case b:
			h = (r-g)/d + 4
			break
		}
	}
	h = math.Round(h * 60)
	s = math.Round(s*1000) / 10
	l = math.Round(l*1000) / 10
	return
}

// HSLToRGB 能夠將 HSL 的顏色以有損的方式轉換成 RGB。
func HSLToRGB(h float64, s float64, l float64) (r float64, g float64, b float64) {
	h = h / 360
	s = s / 100
	l = l / 100

	if s == 0 {
		r = l
		g = l
		b = l
	} else {
		var q float64
		if l < 0.5 {
			q = l * (1 + s)
		} else {
			q = l + s - l*s
		}
		p := 2*l - q

		r = HueToRGB(p, q, h+1/3)
		g = HueToRGB(p, q, h)
		b = HueToRGB(p, q, h-1/3)
	}
	r = math.Round(r * 255)
	g = math.Round(g * 255)
	b = math.Round(b * 255)
	return
}

// HSVToRGB 能夠將 HSV 的顏色以有損的方式轉換成 RGB。
func HSVToRGB(h float64, s float64, v float64) (r float64, g float64, b float64) {
	h = h / 360
	s = s / 100
	v = v / 100

	if s == 0 {
		r = v * 255
		g = v * 255
		b = v * 255
	} else {
		localH := h * 6
		localI := math.Floor(localH)
		local1 := v * (1 - s)
		local2 := v * (1 - s*(localH-localI))
		local3 := v * (1 - s*(1-(localH-localI)))

		var localR float64
		var localG float64
		var localB float64
		switch localI {
		case 0:
			localR = v
			localG = local3
			localB = local1
			break
		case 1:
			localR = local2
			localG = v
			localB = local1
			break
		case 2:
			localR = local1
			localG = v
			localB = local3
			break

		case 3:
			localR = local1
			localG = local2
			localB = v
			break

		case 4:
			localR = local3
			localG = local1
			localB = v
			break
		default:
			localR = v
			localG = local1
			localB = local2
			break
		}
		r = localR * 255
		g = localG * 255
		b = localB * 255
		r = math.Round(r * 255)
		g = math.Round(g * 255)
		b = math.Round(b * 255)
	}
	return
}

// RGBToHSV 能夠將 RGB 的顏色以有損的方式轉換成 HSV。
func RGBToHSV(r float64, g float64, b float64) (h float64, s float64, v float64) {
	r = r / 255
	g = g / 255
	b = b / 255

	minValue := math.Min(r, g)
	minValue = math.Min(minValue, b)
	maxValue := math.Max(r, g)
	maxValue = math.Max(maxValue, b)

	delta := maxValue - minValue
	v = maxValue

	if delta == 0 {
		h = 0
		s = 0
	} else {
		s = delta / maxValue
		deltaR := (((maxValue - r) / 6) + (delta / 2)) / delta
		deltaG := (((maxValue - g) / 6) + (delta / 2)) / delta
		deltaB := (((maxValue - b) / 6) + (delta / 2)) / delta
		switch maxValue {
		case r:
			h = deltaB - deltaG
			break
		case g:
			h = (1 / 3) + deltaR - deltaB
			break
		case b:
			h = (2 / 3) + deltaG - deltaR
			break
		}
		if h < 0 {
			h++
		}
		if h > 1 {
			h--
		}

	}
	h = math.Round(h * 360)
	s = math.Round(s*1000) / 10
	v = math.Round(v*1000) / 10
	return
}

// RGBToHex 能夠將 RGB 的顏色轉換成 Hex 十六進制字串（不包含 `#` 井字符號）。
func RGBToHex(r float64, g float64, b float64) string {
	h := []byte{uint8(math.Round(r)), uint8(math.Round(g)), uint8(math.Round(b))}
	return strings.ToUpper(hex.EncodeToString(h))
}

// HexToRGB 能夠將 Hex 十六進制字串（任意包含 `#` 井字符號）轉換成 RGB。
func HexToRGB(h string) (r float64, g float64, b float64) {
	if string(h[0]) == "#" {
		h = h[1:]
	}
	byteArray, _ := hex.DecodeString(h)
	r = float64(byteArray[0])
	g = float64(byteArray[1])
	b = float64(byteArray[2])
	return
}

// HTMLToRGB 能夠將 HTML 的網頁顏色名稱或帶有井字號的十六進制色彩轉換成 RGB。
func HTMLToRGB(h string) (r float64, g float64, b float64) {
	if string(h[0]) == "#" {
		h = h[1:]
		r, g, b = HexToRGB(h)
	} else {
		v, ok := colorNames[strings.ToUpper(h)]
		if !ok {
			r = 0
			g = 0
			b = 0
			return
		}
		r, g, b = HexToRGB(v)
	}
	return
}

// RGBToHTML 能夠將 RGB 的顏色轉換成網頁的顏色名稱（如：`Red`、`White`）或帶有井字號的十六進制色彩。
func RGBToHTML(r float64, g float64, b float64) (h string) {
	h = RGBToHex(r, g, b)
	v, ok := hexNames[h]
	if !ok {
		h = "#" + h
		return
	}
	h = v
	return
}

// NewHTML 會初始化一個基於 HTML 的顏色資訊。
func NewHTML(color string) *Color {
	r, g, b := HTMLToRGB(color)
	return newColor(r, g, b, 1)
}

// NewHTMLA 會初始化一個基於 HTML 且帶有 Alpha 的顏色資訊。
func NewHTMLA(color string, a float64) *Color {
	r, g, b := HTMLToRGB(color)
	return newColor(r, g, b, a)
}

// NewHex 會初始化一個基於 Hex 的顏色資訊。
func NewHex(color string) *Color {
	r, g, b := HexToRGB(color)
	return newColor(r, g, b, 1)
}

// NewHexA 會初始化一個基於 Hex 且帶有 Alpha 的顏色資訊。
func NewHexA(color string, a float64) *Color {
	r, g, b := HexToRGB(color)
	return newColor(r, g, b, a)
}

// NewHSL 會初始化一個基於 HSL 的顏色資訊。
func NewHSL(h float64, s float64, l float64) *Color {
	r, g, b := HSLToRGB(h, s, l)
	return newColor(r, g, b, 1)
}

// NewHSLA 會初始化一個基於 HSL 且帶有 Alpha 的顏色資訊。
func NewHSLA(h float64, s float64, l float64, a float64) *Color {
	r, g, b := HSLToRGB(h, s, l)
	return newColor(r, g, b, a)
}

// NewHSV 會初始化一個基於 HSV 的顏色資訊。
func NewHSV(h float64, s float64, v float64) *Color {
	r, g, b := HSVToRGB(h, s, v)
	return newColor(r, g, b, 1)
}

// NewHSVA 會初始化一個基於 HSV 且帶有 Alpha 的顏色資訊。
func NewHSVA(h float64, s float64, v float64, a float64) *Color {
	r, g, b := HSVToRGB(h, s, v)
	return newColor(r, g, b, a)
}

// NewRGB 會初始化一個基於 RGB 的顏色資訊。
func NewRGB(r float64, g float64, b float64) *Color {
	return newColor(r, g, b, 1)
}

// NewRGBA 會初始化一個基於 RGB 且帶有 Alpha 的顏色資訊。
func NewRGBA(r float64, g float64, b float64, a float64) *Color {
	return newColor(r, g, b, a)
}

// NewCMYK 會初始化一個基於 CMYK 的顏色資訊。
func NewCMYK(c float64, m float64, y float64, k float64) *Color {
	r, g, b := CMYKToRGB(c, m, y, k)
	return newColor(r, g, b, 1)
}

// NewCMYKA 會初始化一個基於 CMYK 且帶有 Alpha 的顏色資訊。
func NewCMYKA(c float64, m float64, y float64, k float64, a float64) *Color {
	r, g, b := CMYKToRGB(c, m, y, k)
	return newColor(r, g, b, a)
}

// Mix 會將傳入的兩個顏色混合在一起。
// func (c *Color) Mix(color *Color, percent float64) *Color {
// 	r := math.Round((1 - (1-c.Red/255)*(1-color.Red/255)) * 255)
// 	g := math.Round((1 - (1-c.Green/255)*(1-color.Green/255)) * 255)
// 	b := math.Round((1 - (1-c.Blue/255)*(1-color.Blue/255)) * 255)
// 	a := c.Alpha + color.Alpha*(1-c.Alpha)
// 	return newColor(r, g, b, a)
// }

// Mix 會將傳入的兩個顏色混合在一起，並且透過權重表示後者混入顏色的佔比（預設為 `0.5`，即為 `50%`）。
func (c *Color) Mix(color *Color, weight float64) *Color {
	oWeight := 1 - weight
	r := math.Round(oWeight*c.Red + weight*color.Red)
	g := math.Round(oWeight*c.Green + weight*color.Green)
	b := math.Round(oWeight*c.Blue + weight*color.Blue)
	a := math.Round(oWeight*c.Alpha + weight*color.Alpha)
	return newColor(r, g, b, a)
}

// Hue 會取得顏色的色相角度值。
func (c *Color) Hue() float64 {
	h, _, _ := c.HSL()
	return h
}

// Saturation 會取得顏色的飽和百分比。
func (c *Color) Saturation() float64 {
	_, s, _ := c.HSL()
	return s
}

// Lightness 會取得顏色的明亮百分比。
func (c *Color) Lightness() float64 {
	_, _, l := c.HSL()
	return l
}

// AdjustHue 會以角度旋轉顏色的色相角度。
func (c *Color) AdjustHue(degrees float64) *Color {
	h, s, v := c.HSV()
	h += degrees
	for {
		if h >= 0 && h <= 360 {
			break
		}
		if h < 0 {
			h += 360
		} else {
			h += -360
		}
	}
	r, g, b := HSVToRGB(h, s, v)
	return newColor(r, g, b, c.Alpha)
}

// Lighten 會以百分比率將顏色加亮。
func (c *Color) Lighten(percent float64) *Color {
	h, s, l := c.HSL()
	l += percent
	if l > 100 {
		l = 100
	}
	r, g, b := HSLToRGB(h, s, l)
	return newColor(r, g, b, c.Alpha)
}

// Darken 會以百分比率將顏色變暗。
func (c *Color) Darken(percent float64) *Color {
	h, s, l := c.HSL()
	l -= percent
	if l < 0 {
		l = 0
	}
	r, g, b := HSLToRGB(h, s, l)
	return newColor(r, g, b, c.Alpha)
}

// Saturate 會以百分比率提高顏色的飽和度。
func (c *Color) Saturate(percent float64) *Color {
	h, s, l := c.HSL()
	s += percent
	if s > 100 {
		s = 100
	}
	r, g, b := HSLToRGB(h, s, l)
	return newColor(r, g, b, c.Alpha)
}

// Desaturate 會以百分比率降低顏色的飽和度。
func (c *Color) Desaturate(percent float64) *Color {
	h, s, l := c.HSL()
	s -= percent
	if s < 0 {
		s = 0
	}
	r, g, b := HSLToRGB(h, s, l)
	return newColor(r, g, b, c.Alpha)
}

// Grayscale 會將顏色轉換成灰階色調，相當於 `Desaturate(100)`。
func (c *Color) Grayscale() *Color {
	return c.Desaturate(100)
}

// Complement 會取得顏色的互補色，相當於 `AdjustHue(180)`。
func (c *Color) Complement() *Color {
	return c.AdjustHue(180)
}

// Tint 會以百分比率將顏色盡可能地保持原色調的方式加亮，相當於基於白色的 `Mix`。
func (c *Color) Tint(percent float64) *Color {
	return c.Mix(newColor(255, 255, 255, c.Alpha), percent)
}

// Shade 會以百分比率將顏色盡可能地保持原色調的方式變暗，相當於基於黑色的 `Mix`。
func (c *Color) Shade(percent float64) *Color {
	return c.Mix(newColor(0, 0, 0, c.Alpha), percent)
}

// Invert 會將整個顏色反轉。
func (c *Color) Invert() *Color {
	r, g, b := c.RGB()
	h := int(r * g * b)
	r, g, b = HexToRGB(fmt.Sprintf("%06x\n", 0xFFFFFF-h))
	return newColor(r, g, b, c.Alpha)
}

// HSV 會將目前的顏色以 HSV 格式回傳。
func (c *Color) HSV() (float64, float64, float64) {
	return RGBToHSV(c.Red, c.Green, c.Blue)
}

// HSVA 會將目前的顏色以 HSVA 格式回傳。
func (c *Color) HSVA() (float64, float64, float64, float64) {
	h, s, v := RGBToHSV(c.Red, c.Green, c.Blue)
	return h, s, v, c.Alpha
}

// HSL 會將目前的顏色以 HSL 格式回傳。
func (c *Color) HSL() (float64, float64, float64) {
	return RGBToHSL(c.Red, c.Green, c.Blue)
}

// HSLA 會將目前的顏色以 HSLA 格式回傳。
func (c *Color) HSLA() (float64, float64, float64, float64) {
	h, s, l := RGBToHSL(c.Red, c.Green, c.Blue)
	return h, s, l, c.Alpha
}

// RGB 會將目前的顏色以 RGB 格式回傳。
func (c *Color) RGB() (float64, float64, float64) {
	return c.Red, c.Green, c.Blue
}

// RGBA 會將目前的顏色以 RGBA 格式回傳。
func (c *Color) RGBA() (float64, float64, float64, float64) {
	return c.Red, c.Green, c.Blue, c.Alpha
}

// CMYK 會將目前的顏色以 CMYK 格式回傳。
func (c *Color) CMYK() (float64, float64, float64, float64) {
	return RGBToCMYK(c.Red, c.Green, c.Blue)
}

// Hex 會將目前的顏色以 Hex 十六進制字串回傳（不帶有 `#` 井字符號）。
func (c *Color) Hex() string {
	return RGBToHex(c.Red, c.Green, c.Blue)
}

// HTML 會將目前的顏色轉換成帶 `#` 井字符號的色彩代碼，如果該色彩與某個網頁色彩名稱相等，則會轉換成網頁色彩名稱（如：`red`、`yellow`）。
// 如果該顏色帶有 Alpha 透明通道，那麼將會轉譯成 `rgba(x, x, x, x)` 的字串格式。
func (c *Color) HTML() string {
	if c.Alpha != 1 {
		return fmt.Sprintf("rgba(%f, %f, %f, %f)", c.Red, c.Green, c.Blue, c.Alpha)
	}
	return RGBToHTML(c.Red, c.Green, c.Blue)
}
