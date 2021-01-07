package noire

import (
	"encoding/hex"
	"fmt"
	"math"
	"strings"
)

// Color represents a manipulable color.
type Color struct {
	Red   float64
	Green float64
	Blue  float64
	Alpha float64
}

// newColor returns a new color.
func newColor(r float64, g float64, b float64, a float64) Color {
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
	return Color{
		Red:   r,
		Green: g,
		Blue:  b,
		Alpha: a,
	}
}

// CMYKToRGB converts the color from CMYK to RGB with a lossy algorithm.
//
// reference: https://www.ginifab.com.tw/tools/colors/js/colorconverter.js
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

// RGBToCMYK converts the color from RGB to CMYK with a lossy algorithm.
//
// reference: https://www.ginifab.com.tw/tools/colors/js/colorconverter.js
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

// HueToRGB converts the color from Hue to RGB.
//
// reference: https://www.ginifab.com.tw/tools/colors/js/colorconverter.js
func HueToRGB(p float64, q float64, t float64) float64 {

	if t < 0 {
		t++
	}
	if t > 1 {
		t--
	}
	if t < float64(1)/float64(6) {
		return p + (q-p)*6*t
	}
	if t < float64(1)/float64(2) {
		return q
	}
	if t < float64(2)/float64(3) {
		return p + (q-p)*(float64(2)/float64(3)-t)*6
	}
	return p
}

// RGBToHSL converts the color from RGB to HSL with a lossy algorithm.
//
// reference: https://www.ginifab.com.tw/tools/colors/js/colorconverter.js
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

// HSLToRGB converts the color from HSL to RGB with a lossy algorithm.
//
// reference: https://www.ginifab.com.tw/tools/colors/js/colorconverter.js
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
		r = HueToRGB(p, q, h+float64(float64(1)/float64(3)))
		g = HueToRGB(p, q, h)
		b = HueToRGB(p, q, h-float64(float64(1)/float64(3)))
	}
	r = math.Round(r * 255)
	g = math.Round(g * 255)
	b = math.Round(b * 255)
	return
}

// HSVToRGB converts the color from HSV to RGB with a lossy algorithm.
//
// reference: https://www.rapidtables.com/convert/color/hsv-to-rgb.html
func HSVToRGB(h float64, s float64, v float64) (r float64, g float64, b float64) {
	s = s / 100
	v = v / 100
	c := v * s
	hh := h / 60
	x := c * (1 - math.Abs(math.Mod(hh, 2)-1))
	if hh >= 0 && hh < 1 {
		r = c
		g = x
	} else if hh >= 1 && hh < 2 {
		r = x
		g = c
	} else if hh >= 2 && hh < 3 {
		g = c
		b = x
	} else if hh >= 3 && hh < 4 {
		g = x
		b = c
	} else if hh >= 4 && hh < 5 {
		r = x
		b = c
	} else {
		r = c
		b = x
	}
	m := v - c
	r += m
	g += m
	b += m
	r = math.Round(r * 255)
	g = math.Round(g * 255)
	b = math.Round(b * 255)
	return
}

// RGBToHSV converts the color from RGB to HSV with a lossy algorithm.
//
// reference: https://www.ginifab.com.tw/tools/colors/js/colorconverter.js
func RGBToHSV(r float64, g float64, b float64) (h float64, s float64, v float64) {
	r = r / 255
	g = g / 255
	b = b / 255

	minValue := math.Min(math.Min(r, g), b)
	maxValue := math.Max(math.Max(r, g), b)

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
			h = (float64(1) / float64(3)) + deltaR - deltaB
			break
		case b:
			h = (float64(2) / float64(3)) + deltaG - deltaR
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

// RGBToHex converts the color from RGB to a uppercased Hex string (without the `#` prefix).
func RGBToHex(r float64, g float64, b float64) string {
	h := []byte{uint8(math.Round(r)), uint8(math.Round(g)), uint8(math.Round(b))}
	return strings.ToUpper(hex.EncodeToString(h))
}

// HexToRGB converts the Hex string (can be `#` prefixed or either a 3 characters shorthand) to RGB.
func HexToRGB(h string) (r float64, g float64, b float64) {
	if string(h[0]) == "#" {
		h = h[1:]
	}
	if len(h) == 3 {
		h = string(h[0]) + string(h[0]) + string(h[1]) + string(h[1]) + string(h[2]) + string(h[2])
	}
	byteArray, _ := hex.DecodeString(h)
	r = float64(byteArray[0])
	g = float64(byteArray[1])
	b = float64(byteArray[2])
	return
}

// HTMLToRGB converts the color from HTML color name or a Hex string (can be `#` prefixed or either a 3 characters shorthand) to RGB.
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

// RGBToHTML converts the color from RGB to a `#` prefixed Hex string if it doesn't have a HTML color name.
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

// NewHTML initializes a color based on the HTML color name.
func NewHTML(color string) Color {
	r, g, b := HTMLToRGB(color)
	return newColor(r, g, b, 1)
}

// NewHTMLA initializes a color based on the HTML color name with an alpha channel.
func NewHTMLA(color string, a float64) Color {
	r, g, b := HTMLToRGB(color)
	return newColor(r, g, b, a)
}

// NewHex initializes a color based on a Hex string.
func NewHex(color string) Color {
	r, g, b := HexToRGB(color)
	return newColor(r, g, b, 1)
}

// NewHexA initializes a color based on a Hex string with an alpha channel.
func NewHexA(color string, a float64) Color {
	r, g, b := HexToRGB(color)
	return newColor(r, g, b, a)
}

// NewHSL initializes a color based on HSL.
func NewHSL(h float64, s float64, l float64) Color {
	r, g, b := HSLToRGB(h, s, l)
	return newColor(r, g, b, 1)
}

// NewHSLA initializes a color based on HSL with an alpha channel.
func NewHSLA(h float64, s float64, l float64, a float64) Color {
	r, g, b := HSLToRGB(h, s, l)
	return newColor(r, g, b, a)
}

// NewHSV initializes a color based on HSV.
func NewHSV(h float64, s float64, v float64) Color {
	r, g, b := HSVToRGB(h, s, v)
	return newColor(r, g, b, 1)
}

// NewHSVA initializes a color based on HSV with an alpha channel.
func NewHSVA(h float64, s float64, v float64, a float64) Color {
	r, g, b := HSVToRGB(h, s, v)
	return newColor(r, g, b, a)
}

// NewRGB initializes a color based on RGB.
func NewRGB(r float64, g float64, b float64) Color {
	return newColor(r, g, b, 1)
}

// NewRGBA initializes a color based on RGB with an alpha channel.
func NewRGBA(r float64, g float64, b float64, a float64) Color {
	return newColor(r, g, b, a)
}

// NewCMYK initializes a color based on CMYK.
func NewCMYK(c float64, m float64, y float64, k float64) Color {
	r, g, b := CMYKToRGB(c, m, y, k)
	return newColor(r, g, b, 1)
}

// NewCMYKA initializes a color based on CMYK with an alpha channel.
func NewCMYKA(c float64, m float64, y float64, k float64, a float64) Color {
	r, g, b := CMYKToRGB(c, m, y, k)
	return newColor(r, g, b, a)
}

// Mix mixs both color with the specified weight of the second color. (`0.5` as `50%`)
func (c Color) Mix(color Color, weight float64) Color {
	oWeight := 1 - weight
	r := math.Round(oWeight*c.Red + weight*color.Red)
	g := math.Round(oWeight*c.Green + weight*color.Green)
	b := math.Round(oWeight*c.Blue + weight*color.Blue)
	a := math.Round(oWeight*c.Alpha + weight*color.Alpha)
	return newColor(r, g, b, a)
}

// Hue returns the Hue angle of the current color based on the HSL algorithm.
func (c Color) Hue() float64 {
	h, _, _ := c.HSL()
	return h
}

// Saturation returns the percentage of the current color saturation based on the HSL algorithm.
func (c Color) Saturation() float64 {
	_, s, _ := c.HSL()
	return s
}

// Lightness returns the Lightness of the current color based on the HSL algorithm.
func (c Color) Lightness() float64 {
	_, _, l := c.HSL()
	return l
}

// AdjustHue rotates the Hue angle of the color based on HSL mode, it still goes clockwise if the value was set over than 360 degree.
func (c Color) AdjustHue(degrees float64) Color {
	h, s, l := c.HSL()
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
	r, g, b := HSLToRGB(h, s, l)
	return newColor(r, g, b, c.Alpha)
}

// Lighten increases the brightness of the color based on HSL mode. (`0.5` as `50%`)
func (c Color) Lighten(percent float64) Color {
	percent = percent * 100
	h, s, l := c.HSL()
	l += percent
	if l > 100 {
		l = 100
	}
	r, g, b := HSLToRGB(h, s, l)
	return newColor(r, g, b, c.Alpha)
}

// Darken decreases the brightness of the color based on HSL mode. (`0.5` as `50%`)
func (c Color) Darken(percent float64) Color {
	percent = percent * 100
	h, s, l := c.HSL()
	l -= percent
	if l < 0 {
		l = 0
	}
	r, g, b := HSLToRGB(h, s, l)
	return newColor(r, g, b, c.Alpha)
}

// Saturate increases the saturation of the color based on HSL mode. (`0.5` as `50%`)
func (c Color) Saturate(percent float64) Color {
	percent = percent * 100
	h, s, l := c.HSL()
	s += percent
	if s > 100 {
		s = 100
	}
	r, g, b := HSLToRGB(h, s, l)
	return newColor(r, g, b, c.Alpha)
}

// Desaturate decreases the saturation of the color based on HSL mode. (`0.5` as `50%`)
func (c Color) Desaturate(percent float64) Color {
	percent = percent * 100
	h, s, l := c.HSL()
	s -= percent
	if s < 0 {
		s = 0
	}
	r, g, b := HSLToRGB(h, s, l)
	return newColor(r, g, b, c.Alpha)
}

// Grayscale converts the color to grayscale, same as `Desaturate(1)`.
func (c Color) Grayscale() Color {
	return c.Desaturate(1)
}

// Complement returns the complementary color of the current color, same as `AdjustHue(180)`.
func (c Color) Complement() Color {
	return c.AdjustHue(180)
}

// Tint increases the brightness of the color while keeping the color tone, same as `Mix` with a white color. (`0.5` as `50%`)
func (c Color) Tint(percent float64) Color {
	return c.Mix(newColor(255, 255, 255, c.Alpha), percent)
}

// Shade decreases the brightness of the color while keeping the color tone, same as `Mix` with a black color. (`0.5` as `50%`)
func (c Color) Shade(percent float64) Color {
	return c.Mix(newColor(0, 0, 0, c.Alpha), percent)
}

// Invert returns the opposite color that based on the RGB color map (it's not a complementary color).
func (c Color) Invert() Color {
	r, g, b := c.RGB()
	r = 255 - r
	g = 255 - g
	b = 255 - b
	return newColor(r, g, b, c.Alpha)
}

// LuminanaceWCAG returns the Luminance of the the current color based on the WCAG 2.0 algorithm.
//
// reference: https://www.w3.org/TR/WCAG20-TECHS/G17.html#G17-tests
//
// reference: https://medium.com/dev-channel/using-sass-to-automatically-pick-text-colors-4ba7645d2796
func (c Color) LuminanaceWCAG() float64 {
	rgb := []float64{c.Red, c.Green, c.Blue}
	for k, v := range rgb {
		v /= 255
		if v <= 0.03928 {
			v = v / 12.92
		} else {
			v = math.Pow((v+0.055)/1.055, 2.4)
		}
		rgb[k] = v
	}
	v := rgb[0]*0.2126 + rgb[1]*0.7152 + rgb[2]*0.0722
	return math.Round(v*100) / 100
}

// Luminanace returns the Luminance of the current color.
//
// reference: https://en.wikipedia.org/wiki/Relative_luminance
//
// reference: https://stackoverflow.com/questions/596216/formula-to-determine-brightness-of-rgb-color
//
// reference: https://stackoverflow.com/questions/3116260/given-a-background-color-how-to-get-a-foreground-color-that-makes-it-readable-o
func (c Color) Luminanace() float64 {
	v := 0.2126*c.Red + 0.7152*c.Green + 0.0722*c.Blue
	return math.Round(v*100) / 100
}

// Foreground returns suggested foreground color by calculating the color luminance, it returns a white color when the color is dark, vise versa.
func (c Color) Foreground() Color {
	white := NewRGB(255, 255, 255)
	black := NewRGB(0, 0, 0)
	if c.Luminanace() < 140 {
		return white
	}
	return black
}

// Brighten increases the brightness of the color. (`0.5` as `50%`)
//
// reference: https://github.com/ozdemirburak/iris
func (c Color) Brighten(percent float64) Color {
	percent *= -100
	r := math.Max(0, math.Min(255, c.Red-math.Round(255*(percent/100))))
	g := math.Max(0, math.Min(255, c.Green-math.Round(255*(percent/100))))
	b := math.Max(0, math.Min(255, c.Blue-math.Round(255*(percent/100))))
	return newColor(r, g, b, c.Alpha)
}

// Contrast returns the Contrast of the current color based on the WCAG Luminance algorithm.
func (c Color) Contrast(color Color) float64 {
	c1 := c.LuminanaceWCAG() + 0.05
	c2 := color.LuminanaceWCAG() + 0.05
	v := math.Max(c1, c2) / math.Min(c1, c2)
	return math.Round(v*100) / 100
}

// IsLight returns true if the color is a light scheme, it might not be the same as what human eyes can see.
func (c Color) IsLight() bool {
	darkness := 1 - (0.299*c.Red+0.587*c.Green+0.114*c.Blue)/255
	return darkness < 0.5
}

// IsDark returns true if the color is a dark scheme, it might not be the same as what human eyes can see.
func (c Color) IsDark() bool {
	return !c.IsLight()
}

// HSV returns the HSV value of the current color.
func (c Color) HSV() (float64, float64, float64) {
	return RGBToHSV(c.Red, c.Green, c.Blue)
}

// HSVA returns the HSVA value of the current color.
func (c Color) HSVA() (float64, float64, float64, float64) {
	h, s, v := RGBToHSV(c.Red, c.Green, c.Blue)
	return h, s, v, c.Alpha
}

// HSL returns the HSL value of the current color.
func (c Color) HSL() (float64, float64, float64) {
	return RGBToHSL(c.Red, c.Green, c.Blue)
}

// HSLA returns the HSLA value of the current color.
func (c Color) HSLA() (float64, float64, float64, float64) {
	h, s, l := RGBToHSL(c.Red, c.Green, c.Blue)
	return h, s, l, c.Alpha
}

// RGB returns the RGB value of the current color.
func (c Color) RGB() (float64, float64, float64) {
	return c.Red, c.Green, c.Blue
}

// RGBA returns the RGBA value of the current color.
func (c Color) RGBA() (float64, float64, float64, float64) {
	return c.Red, c.Green, c.Blue, c.Alpha
}

// CMYK returns the CMYK value of the current color.
func (c Color) CMYK() (float64, float64, float64, float64) {
	return RGBToCMYK(c.Red, c.Green, c.Blue)
}

// Hex returns a Hex string of the current color. (Without the `#` prefix)
func (c Color) Hex() string {
	return RGBToHex(c.Red, c.Green, c.Blue)
}

// HTML returns a `#` prefixed Hex string or the HTML color name (like: `red`, `yellow`) if it had one.
// It returns a format with `rgba()` if the color comes with a alpha channel.
func (c Color) HTML() string {
	if c.Alpha != 1 {
		return fmt.Sprintf("rgba(%f, %f, %f, %f)", c.Red, c.Green, c.Blue, c.Alpha)
	}
	return RGBToHTML(c.Red, c.Green, c.Blue)
}
