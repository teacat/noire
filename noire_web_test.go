package noire

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestColorPlate(t *testing.T) {
	var c string
	colors := []*Color{NewHTML("Red"), NewHTML("Orange"), NewHTML("Yellow"), NewHTML("Green"), NewHTML("Blue"), NewHTML("White")}

	before := func(fn string) {
		c += fmt.Sprintf(`<div class="header"></div><div class="section">`)
	}
	after := func() {
		c += fmt.Sprintf(`</div>`)
	}
	do := func(fn string, a string, v interface{}) {
		var d string
		switch b := v.(type) {
		case string:
			d = b
			break
		case float64:
			d = fmt.Sprintf("%.1f", b)
			break
		}
		c += fmt.Sprintf(`<div class="group">
						<div class="color" style="background-color: %s"></div>
						<div class="label">%s(%s)</div>
					</div>`, a, fn, d)
	}

	before("Lighten")
	for i := 0.0; i <= 1; i += 0.2 {
		h := NewHex("00ADEA").Lighten(i).HTML()
		do("Lighten", h, i)
	}
	after()

	before("Brighten")
	for i := 0.0; i <= 1; i += 0.2 {
		h := NewHex("00ADEA").Brighten(i).HTML()
		do("Brighten", h, i)
	}
	after()

	before("Tint")
	for i := 0.0; i <= 1; i += 0.2 {
		h := NewHex("00ADEA").Tint(i).HTML()
		do("Tint", h, i)
	}
	after()

	before("Darken")
	for i := 0.0; i <= 1; i += 0.2 {
		h := NewHex("00ADEA").Darken(i).HTML()
		do("Darken", h, i)
	}
	after()

	before("Shade")
	for i := 0.0; i <= 1; i += 0.2 {
		h := NewHex("00ADEA").Shade(i).HTML()
		do("Shade", h, i)
	}
	after()

	before("Saturate")
	for i := 0.0; i <= 1; i += 0.2 {
		h := NewHex("4a5f45").Saturate(i).HTML()
		do("Saturate", h, i)
	}
	after()

	before("Desaturate")
	for i := 0.0; i <= 1; i += 0.2 {
		h := NewHex("00ADEA").Desaturate(i).HTML()
		do("Desaturate", h, i)
	}
	after()

	before("AdjustHue")
	for i := float64(0); i <= float64(360); i += float64(72) {
		h := NewHex("00ADEA").AdjustHue(i).HTML()
		do("AdjustHue", h, i)
	}
	after()

	before("Mix")
	for i := 0.0; i <= 1; i += 0.2 {
		h := NewHex("00ADEA").Mix(NewRGB(255, 0, 0), i).HTML()
		do("Mix", h, i)
	}
	after()

	before("Invert")
	for _, v := range colors {
		h := v.Invert().HTML()
		do("Invert", h, v.HTML())
	}
	after()

	before("Complement")
	for _, v := range colors {
		h := v.Complement().HTML()
		do("Complement", h, v.HTML())
	}
	after()

	before("Grayscale")
	for _, v := range colors {
		h := v.Grayscale().HTML()
		do("Grayscale", h, v.HTML())
	}
	after()

	before("Foreground")
	for _, v := range colors {
		h := v.Foreground().HTML()
		do("Foreground", h, v.HTML())
	}
	after()

	ioutil.WriteFile("./test/index.html", []byte(fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Test</title>
    <style type="text/css">
    html, body {
        padding: 30px;
        text-align: center;
        font-family: "Noto Sans CJK TC";
    }
    .header {
        margin-top: 32px;
        margin-bottom: 16px;
        font-size: 24px;
        font-weight: bold;
        color: #555;
        /*text-transform: uppercase;*/
    }
    .group {
        display: inline-block;
    }
    .color {
        height: 70px;
        width: 150px;
    }
    .section .group:first-child .color {
        border-top-left-radius: 6px;
        border-bottom-left-radius: 6px;
    }
    .section .group:last-child .color {
        border-top-right-radius: 6px;
        border-bottom-right-radius: 6px;
    }
    .label {
        text-align: center;
        color:#BBB;
        font-size: 13px;
        font-weight: bold;
        padding-top: 16px;
    }
    </style>
</head>
<body>
%s
</body>
</html>`, c)), os.ModePerm)
}
