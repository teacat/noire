package noire

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tColorRGB = []float64{219, 112, 147}
var tColorCMYK = []float64{0, 49, 33, 14}
var tColorHSV = []float64{340, 48.9, 85.9}
var tColorHSL = []float64{340, 59.8, 64.9}
var tColorHex = "DB7093"
var tColorHTML = "PaleVioletRed"
var tColorHexInvert = "248F6C"
var tColorHexGrayscale = "A6A6A6"

func TestCMYKToRGB(t *testing.T) {
	assert := assert.New(t)
	r, g, b := CMYKToRGB(tColorCMYK[0], tColorCMYK[1], tColorCMYK[2], tColorCMYK[3])
	assert.Equal(tColorRGB, []float64{r, g, b})
}
func TestRGBToCMYK(t *testing.T) {
	assert := assert.New(t)
	c, m, y, k := RGBToCMYK(tColorRGB[0], tColorRGB[1], tColorRGB[2])
	assert.Equal(tColorCMYK, []float64{c, m, y, k})
}

func TestHueToRGB(t *testing.T) {
	//assert := assert.New(t)
}

func TestRGBToHSL(t *testing.T) {
	assert := assert.New(t)
	h, s, l := RGBToHSL(tColorRGB[0], tColorRGB[1], tColorRGB[2])
	assert.Equal(tColorHSL, []float64{h, s, l})
}

func TestHSLToRGB(t *testing.T) {
	assert := assert.New(t)
	r, g, b := HSLToRGB(tColorHSL[0], tColorHSL[1], tColorHSL[2])
	assert.Equal(tColorRGB, []float64{r, g, b})
}

func TestHSVToRGB(t *testing.T) {
	assert := assert.New(t)
	h, s, v := HSVToRGB(tColorHSV[0], tColorHSV[1], tColorHSV[2])
	assert.Equal(tColorRGB, []float64{h, s, v})
}

func TestRGBToHSV(t *testing.T) {
	assert := assert.New(t)
	r, g, b := RGBToHSV(tColorRGB[0], tColorRGB[1], tColorRGB[2])
	assert.Equal(tColorHSV, []float64{r, g, b})
}

func TestRGBToHex(t *testing.T) {
	assert := assert.New(t)
	h := RGBToHex(tColorRGB[0], tColorRGB[1], tColorRGB[2])
	assert.Equal(tColorHex, h)
}

func TestHexToRGB(t *testing.T) {
	assert := assert.New(t)
	r, g, b := HexToRGB(tColorHex)
	assert.Equal(tColorRGB, []float64{r, g, b})
}

func TestHTMLToRGB(t *testing.T) {
	assert := assert.New(t)
	r, g, b := HTMLToRGB(tColorHTML)
	assert.Equal(tColorRGB, []float64{r, g, b})
}

func TestRGBToHTML(t *testing.T) {
	assert := assert.New(t)
	h := RGBToHTML(tColorRGB[0], tColorRGB[1], tColorRGB[2])
	assert.Equal(tColorHTML, h)
}

func TestMix(t *testing.T) {

}

func TestHue(t *testing.T) {

}

func TestSaturation(t *testing.T) {

}

func TestLightness(t *testing.T) {

}

func TestAdjustHue(t *testing.T) {

}

func TestLighten(t *testing.T) {

}

func TestDarken(t *testing.T) {

}

func TestSaturate(t *testing.T) {

}

func TestDesaturate(t *testing.T) {

}

func TestGrayscale(t *testing.T) {

}

func TestComplement(t *testing.T) {

}

func TestTint(t *testing.T) {

}

func TestShade(t *testing.T) {

}

func TestInvert(t *testing.T) {
	assert := assert.New(t)
	c := NewRGB(tColorRGB[0], tColorRGB[1], tColorRGB[2])
	c = c.Invert()
	assert.Equal(tColorHexInvert, c.Hex())
}

func TestHSV(t *testing.T) {

}

func TestHSVA(t *testing.T) {

}

func TestHSL(t *testing.T) {

}

func TestHSLA(t *testing.T) {

}

func TestRGB(t *testing.T) {

}

func TestRGBA(t *testing.T) {

}

func TestCMYK(t *testing.T) {

}

func TestHex(t *testing.T) {

}

func TestHTML(t *testing.T) {

}

//func BenchmarkString(b *testing.B) {
//	t := New(DefaultConfig())
//	b.ReportAllocs()
//	b.ResetTimer()
//	for n := 0; n < b.N; n++ {
//		_ = t.Generate().String()
//	}
//}
