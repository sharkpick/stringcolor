package stringcolor

import (
	"fmt"
	"testing"
)

var TheColors = []StringColor{
	Reset, Bold, Light, Italic, Underscore, Highlight,
	Score, DoubleUnderscore, DarkGray, Red, Green,
	Yellow, Blue, Purple, Teal, White, WhiteOnBlack,
	WhiteOnRed, BlackOnGreen, BlackOnYellow, WhiteOnBlue,
	WhiteOnPurple, BlackOnTeal, BlackOnWhite,
	WhiteOnGray, BlackOnRed, GrayOnGreen,
	GrayOnYellow, BlackOnBlue, BlackOnPurple,
	GrayOnTeal, GrayOnWhite,
}

func TestColors(t *testing.T) {
	for _, color := range TheColors {
		fmt.Println(ColorWrapString(color, "hello"))
	}
}

func TestComboColors(t *testing.T) {
	s := Reset + Italic + Red + "hello world" + Reset
	fmt.Println(s)
}
