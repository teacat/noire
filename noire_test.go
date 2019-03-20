package noire

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCMYKToRGB(t *testing.T) {
	assert := assert.New(t)
	r, g, b := CMYKToRGB(0, 49, 33, 14)
	assert.Equal([]float64{219, 112, 147}, []float64{r, g, b})
}
func TestRGBToCMYK(t *testing.T) {
	assert := assert.New(t)
	c, m, y, k := RGBToCMYK(219, 112, 148)
	assert.Equal([]float64{0, 49, 32, 14}, []float64{c, m, y, k})
}

func TestHueToRGB(t *testing.T) {
	//assert := assert.New(t)
}

func TestRGBToHSL(t *testing.T) {
	assert := assert.New(t)
	h, s, l := RGBToHSL(219, 112, 148)
	assert.Equal([]float64{340, 59.8, 64.9}, []float64{h, s, l})
}

func TestHSLToRGB(t *testing.T) {
	assert := assert.New(t)
	r, g, b := HSLToRGB(340, 59.8, 64.9)
	assert.Equal([]float64{219, 112, 148}, []float64{r, g, b})
}

func TestHSVToRGB(t *testing.T) {
	assert := assert.New(t)
	h, s, v := HSVToRGB(340, 48.9, 85.9)
	assert.Equal([]float64{219, 112, 148}, []float64{h, s, v})
	h, s, v = HSVToRGB(340, 0, 85.9)
	assert.Equal([]float64{219, 219, 0}, []float64{h, s, v})
}

func TestRGBToHSV(t *testing.T) {
	assert := assert.New(t)
	r, g, b := RGBToHSV(219, 112, 148)
	assert.Equal([]float64{340, 48.9, 85.9}, []float64{r, g, b})
}

func TestRGBToHex(t *testing.T) {
	assert := assert.New(t)
	h := RGBToHex(219, 112, 148)
	assert.Equal("DB7094", h)
}

func TestHexToRGB(t *testing.T) {
	assert := assert.New(t)
	r, g, b := HexToRGB("DB7093")
	assert.Equal([]float64{219, 112, 147}, []float64{r, g, b})
}

func TestHTMLToRGB(t *testing.T) {
	assert := assert.New(t)
	r, g, b := HTMLToRGB("PaleVioletRed")
	assert.Equal([]float64{219, 112, 147}, []float64{r, g, b})
	r, g, b = HTMLToRGB("#F0F0F0")
	assert.Equal([]float64{240, 240, 240}, []float64{r, g, b})
	r, g, b = HTMLToRGB("NinjaTurtle")
	assert.Equal([]float64{0, 0, 0}, []float64{r, g, b})
}

func TestRGBToHTML(t *testing.T) {
	assert := assert.New(t)
	h := RGBToHTML(219, 112, 147)
	assert.Equal("PaleVioletRed", h)
}

func TestNewAlpha(t *testing.T) {
	assert := assert.New(t)
	c := NewCMYKA(20, 20, 20, 20, 0.5)
	assert.Equal(0.5, c.Alpha)
	c = NewHSLA(20, 20, 20, 0.5)
	assert.Equal(0.5, c.Alpha)
	c = NewHSVA(20, 20, 20, 0.5)
	assert.Equal(0.5, c.Alpha)
	c = NewHTMLA("Red", 0.5)
	assert.Equal(0.5, c.Alpha)
	c = NewHexA("000", 0.5)
	assert.Equal(0.5, c.Alpha)
	c = NewRGBA(20, 20, 20, 0.5)
	assert.Equal(0.5, c.Alpha)
}

func TestMix(t *testing.T) {
	assert := assert.New(t)
	c1 := NewHex("F00")
	c2 := NewHex("00F")
	assert.Equal("Purple", c1.Mix(c2, 0.5).HTML())
}

func TestHue(t *testing.T) {
	assert := assert.New(t)
	c := NewRGB(219, 112, 148)
	assert.Equal(float64(340), c.Hue())
}

func TestSaturation(t *testing.T) {
	assert := assert.New(t)
	c := NewRGB(219, 112, 148)
	assert.Equal(59.8, c.Saturation())
}

func TestLightness(t *testing.T) {
	assert := assert.New(t)
	c := NewRGB(219, 112, 148)
	assert.Equal(64.9, c.Lightness())
}

func TestAdjustHue(t *testing.T) {
	assert := assert.New(t)
	c := NewRGB(219, 112, 148).AdjustHue(30)
	assert.Equal("DB8270", c.Hex())
	c = NewRGB(219, 112, 148).AdjustHue(360)
	assert.Equal("DB7094", c.Hex())
	c = NewRGB(219, 112, 148).AdjustHue(480)
	assert.Equal("94DB70", c.Hex())
	c = NewRGB(219, 112, 148).AdjustHue(720)
	assert.Equal("DB7094", c.Hex())
	c = NewRGB(219, 112, 148).AdjustHue(-720)
	assert.Equal("DB7094", c.Hex())
}

func TestLighten(t *testing.T) {
	assert := assert.New(t)
	c := NewRGB(219, 112, 148)
	c = c.Lighten(0.15)
	assert.Equal("EAADC2", c.Hex())
}

func TestDarken(t *testing.T) {
	assert := assert.New(t)
	c := NewRGB(219, 112, 148)
	c = c.Darken(0.15)
	assert.Equal("CB3366", c.Hex())
}

func TestSaturate(t *testing.T) {
	assert := assert.New(t)
	c := NewRGB(219, 112, 148)
	c = c.Saturate(0.5)
	assert.Equal("FF4C88", c.Hex())
}

func TestDesaturate(t *testing.T) {
	assert := assert.New(t)
	c := NewRGB(219, 112, 148)
	c = c.Desaturate(0.5)
	assert.Equal("AE9DA3", c.Hex())
}

func TestGrayscale(t *testing.T) {
	assert := assert.New(t)
	c := NewRGB(219, 112, 148)
	c = c.Grayscale()
	assert.Equal("A5A5A5", c.Hex())
}

func TestComplement(t *testing.T) {
	assert := assert.New(t)
	c := NewRGB(219, 112, 148)
	c = c.Complement()
	assert.Equal("70DBB7", c.Hex())
}

func TestTint(t *testing.T) {
	assert := assert.New(t)
	c := NewRGB(219, 112, 148)
	c = c.Tint(0.15)
	assert.Equal("E085A4", c.Hex())
}

func TestShade(t *testing.T) {
	assert := assert.New(t)
	c := NewRGB(219, 112, 148)
	c = c.Shade(0.15)
	assert.Equal("BA5F7E", c.Hex())
}

func TestInvert(t *testing.T) {
	assert := assert.New(t)
	c := NewRGB(219, 112, 148)
	c = c.Invert()
	assert.Equal("248F6B", c.Hex())
}

func TestLuminanaceWCAG(t *testing.T) {
	assert := assert.New(t)
	c := NewRGB(219, 112, 148)
	assert.Equal(0.29, c.LuminanaceWCAG())
}

func TestLuminanace(t *testing.T) {
	assert := assert.New(t)
	c := NewRGB(219, 112, 148)
	assert.Equal(137.35, c.Luminanace())
}

func TestForeground(t *testing.T) {
	assert := assert.New(t)
	c := NewHTML("Green")
	assert.Equal("White", c.Foreground().HTML())
	c = NewHTML("Red")
	assert.Equal("White", c.Foreground().HTML())
	c = NewHTML("Yellow")
	assert.Equal("Black", c.Foreground().HTML())
}

func TestBrighten(t *testing.T) {
	assert := assert.New(t)
	c := NewRGB(0, 0, 0)
	assert.Equal("1A1A1A", c.Brighten(0.1).Hex())
}

func TestContrast(t *testing.T) {
	assert := assert.New(t)
	c1 := NewRGB(219, 112, 148)
	c2 := NewRGB(0, 0, 0)
	assert.Equal(6.8, c1.Contrast(c2))
}

func TestIsLight(t *testing.T) {
	assert := assert.New(t)
	c := NewRGB(0, 0, 0)
	assert.Equal(false, c.IsLight())
	c = NewRGB(255, 255, 255)
	assert.Equal(true, c.IsLight())
	c = NewHTML("Red")
	assert.Equal(false, c.IsLight())
	c = NewHTML("Blue")
	assert.Equal(false, c.IsLight())
	c = NewHTML("Green")
	assert.Equal(false, c.IsLight())
	c = NewHTML("Purple")
	assert.Equal(false, c.IsLight())
}

func TestIsDark(t *testing.T) {
	assert := assert.New(t)
	c := NewRGB(0, 0, 0)
	assert.Equal(true, c.IsDark())
	c = NewRGB(255, 255, 255)
	assert.Equal(false, c.IsDark())
	c = NewHTML("Red")
	assert.Equal(true, c.IsDark())
	c = NewHTML("Blue")
	assert.Equal(true, c.IsDark())
	c = NewHTML("Green")
	assert.Equal(true, c.IsDark())
	c = NewHTML("Purple")
	assert.Equal(true, c.IsDark())
}

func TestHSV(t *testing.T) {
	assert := assert.New(t)
	c := NewRGB(219, 112, 148)
	h, s, v := c.HSV()
	assert.Equal([]float64{340, 48.9, 85.9}, []float64{h, s, v})
}

func TestHSVA(t *testing.T) {
	assert := assert.New(t)
	c := NewRGB(219, 112, 148)
	h, s, v, a := c.HSVA()
	assert.Equal([]float64{340, 48.9, 85.9, 1}, []float64{h, s, v, a})
}

func TestHSL(t *testing.T) {
	assert := assert.New(t)
	c := NewRGB(219, 112, 148)
	h, s, l := c.HSL()
	assert.Equal([]float64{340, 59.8, 64.9}, []float64{h, s, l})
}

func TestHSLA(t *testing.T) {
	assert := assert.New(t)
	c := NewRGB(219, 112, 148)
	h, s, l, a := c.HSLA()
	assert.Equal([]float64{340, 59.8, 64.9, 1}, []float64{h, s, l, a})
}

func TestRGB(t *testing.T) {
	assert := assert.New(t)
	c := NewRGB(219, 112, 148)
	r, g, b := c.RGB()
	assert.Equal([]float64{219, 112, 148}, []float64{r, g, b})
}

func TestRGBA(t *testing.T) {
	assert := assert.New(t)
	c := NewRGB(219, 112, 148)
	r, g, b, a := c.RGBA()
	assert.Equal([]float64{219, 112, 148, 1}, []float64{r, g, b, a})
}

func TestCMYK(t *testing.T) {
	assert := assert.New(t)
	c1 := NewRGB(219, 112, 148)
	c, m, y, k := c1.CMYK()
	assert.Equal([]float64{0, 49, 32, 14}, []float64{c, m, y, k})
}

func TestHex(t *testing.T) {
	assert := assert.New(t)
	c := NewRGB(219, 112, 148)
	assert.Equal("DB7094", c.Hex())
}

func TestHTML(t *testing.T) {
	assert := assert.New(t)
	c := NewRGB(219, 112, 148)
	assert.Equal("#DB7094", c.HTML())
	c = NewRGB(219, 112, 147)
	assert.Equal("PaleVioletRed", c.HTML())
	c = NewRGB(219, 112, 147)
	c.Alpha = 0.5
	assert.Equal("rgba(219.000000, 112.000000, 147.000000, 0.500000)", c.HTML())
}

func BenchmarkString(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {

	}
}
