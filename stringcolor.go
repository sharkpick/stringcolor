package stringcolor

import "fmt"

type StringColor string

var (
	Reset            = ColorCodeString(0)
	Bold             = ColorCodeString(1)
	Light            = ColorCodeString(2)
	Italic           = ColorCodeString(3)
	Underscore       = ColorCodeString(4)
	Highlight        = ColorCodeString(7)
	Score            = ColorCodeString(9)
	DoubleUnderscore = ColorCodeString(21)
	DarkGray         = ColorCodeString(30)
	Red              = ColorCodeString(31)
	Green            = ColorCodeString(32)
	Yellow           = ColorCodeString(33)
	Blue             = ColorCodeString(34)
	Purple           = ColorCodeString(35)
	Teal             = ColorCodeString(36)
	White            = ColorCodeString(37)
	WhiteOnBlack     = ColorCodeString(40)
	WhiteOnRed       = ColorCodeString(41)
	BlackOnGreen     = ColorCodeString(42)
	BlackOnYellow    = ColorCodeString(43)
	WhiteOnBlue      = ColorCodeString(44)
	WhiteOnPurple    = ColorCodeString(45)
	BlackOnTeal      = ColorCodeString(46)
	BlackOnWhite     = ColorCodeString(47)
	WhiteOnGray      = ColorCodeString(100)
	BlackOnRed       = ColorCodeString(101)
	GrayOnGreen      = ColorCodeString(102)
	GrayOnYellow     = ColorCodeString(103)
	BlackOnBlue      = ColorCodeString(104)
	BlackOnPurple    = ColorCodeString(105)
	GrayOnTeal       = ColorCodeString(106)
	GrayOnWhite      = ColorCodeString(107)
)

const (
	ANSIEscape       = "\033"
	ColorCodeFString = ANSIEscape + "[%dm"
)

func ColorCodeString(code int) StringColor {
	return StringColor(fmt.Sprintf(ColorCodeFString, code))
}

func ColorWrapString(color StringColor, s string) string {
	return fmt.Sprintf("%s%s%s%s", Reset, string(color), s, Reset)
}
