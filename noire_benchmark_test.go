package noire

import (
	"testing"
)

func BenchmarkCMYKToRGB(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		CMYKToRGB(0, 49, 33, 14)
	}
}
func BenchmarkRGBToCMYK(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		RGBToCMYK(219, 112, 148)
	}
}

func BenchmarkRGBToHSL(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		RGBToHSL(219, 112, 148)
	}
}

func BenchmarkHSLToRGB(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		HSLToRGB(340, 59.8, 64.9)
	}
}

func BenchmarkHSVToRGB(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		HSVToRGB(340, 48.9, 85.9)
	}
}

func BenchmarkRGBToHSV(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		RGBToHSV(219, 112, 148)
	}
}

func BenchmarkRGBToHex(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		RGBToHex(219, 112, 148)
	}
}

func BenchmarkHexToRGB(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		HexToRGB("DB7093")
	}
}

func BenchmarkHTMLToRGBName(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		HTMLToRGB("PaleVioletRed")
	}
}

func BenchmarkHTMLToRGBHex(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		HTMLToRGB("#F0F0F0")
	}
}

func BenchmarkRGBToHTML(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		RGBToHTML(219, 112, 147)
	}
}

func BenchmarkMix(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c1 := NewHex("F00")
		c2 := NewHex("00F")
		c1.Mix(c2, 0.5)
	}
}

func BenchmarkHue(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		NewRGB(219, 112, 148).Hue()
	}
}

func BenchmarkSaturation(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		NewRGB(219, 112, 148).Saturation()
	}
}

func BenchmarkLightness(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		NewRGB(219, 112, 148).Lightness()
	}
}

func BenchmarkAdjustHue(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		NewRGB(219, 112, 148).AdjustHue(30)
	}
}

func BenchmarkLighten(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		NewRGB(219, 112, 148).Lighten(0.15)
	}
}

func BenchmarkDarken(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		NewRGB(219, 112, 148).Darken(0.15)
	}
}

func BenchmarkSaturate(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		NewRGB(219, 112, 148).Saturate(0.5)
	}
}

func BenchmarkDesaturate(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		NewRGB(219, 112, 148).Desaturate(0.5)
	}
}

func BenchmarkGrayscale(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		NewRGB(219, 112, 148).Grayscale()
	}
}

func BenchmarkComplement(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		NewRGB(219, 112, 148).Complement()
	}
}

func BenchmarkTint(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		NewRGB(219, 112, 148).Tint(0.15)
	}
}

func BenchmarkShade(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		NewRGB(219, 112, 148).Shade(0.15)
	}
}

func BenchmarkInvert(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		NewRGB(219, 112, 148).Invert()
	}
}

func BenchmarkLuminanaceWCAG(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		NewRGB(219, 112, 148).LuminanaceWCAG()
	}
}

func BenchmarkLuminanace(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		NewRGB(219, 112, 148).Luminanace()
	}
}

func BenchmarkContrast(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c1 := NewRGB(219, 112, 148)
		c2 := NewRGB(0, 0, 0)
		c1.Contrast(c2)
	}
}

func BenchmarkIsLight(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		NewRGB(0, 0, 0).IsLight()
	}
}

func BenchmarkIsDark(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		NewRGB(0, 0, 0).IsDark()
	}
}

func BenchmarkHSV(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		NewRGB(219, 112, 148).HSV()
	}
}

func BenchmarkHSVA(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		NewRGB(219, 112, 148).HSVA()
	}
}

func BenchmarkHSL(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		NewRGB(219, 112, 148).HSL()
	}
}

func BenchmarkHSLA(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		NewRGB(219, 112, 148).HSLA()
	}
}

func BenchmarkRGB(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		NewRGB(219, 112, 148).RGB()
	}
}

func BenchmarkRGBA(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		NewRGB(219, 112, 148).RGBA()
	}
}

func BenchmarkCMYK(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		NewRGB(219, 112, 148).CMYK()
	}
}

func BenchmarkHex(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		NewRGB(219, 112, 148).Hex()
	}
}

func BenchmarkHTMLHex(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		NewRGB(219, 112, 148).HTML()
	}
}

func BenchmarkHTMLName(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		NewRGB(219, 112, 147).HTML()
	}
}

func BenchmarkHTMLRGBA(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c := NewRGB(219, 112, 147)
		c.Alpha = 0.5
		c.HTML()
	}
}
