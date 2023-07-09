# stringcolor
package stringcolor provides Linux-only text color and appearance manipulation using ANSI Escape Codes. Not all terminals provide the same support - in general you could use the `ColorCodeString(int) StringColor` function for any 8-bit integer

## structure
useful colors are predefined as variables in the package to be used with the ColorWrapString function. They can also be combined however you feel is useful. The `Reset` string must be used to reset, or odd results may occur.