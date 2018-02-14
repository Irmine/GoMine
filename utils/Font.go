package utils

import "strings"

var colorConvert = map[string]string{
	Black: AnsiBlack,
	Blue: AnsiBlue,
	Green: AnsiGreen,
	Cyan: AnsiCyan,
	Red: AnsiRed,
	Magenta: AnsiMagenta,
	Orange: AnsiYellow,
	BrightGray: AnsiWhite,
	Gray: AnsiGray,
	BrightBlue: AnsiBrightBlue,
	BrightGreen: AnsiBrightGreen,
	BrightCyan: AnsiBrightCyan,
	BrightRed: AnsiBrightRed,
	BrightMagenta: AnsiBrightMagenta,
	Yellow: AnsiBrightYellow,
	White: AnsiBrightWhite,
	Bold: AnsiBold,
	Underlined: AnsiUnderlined,
	Italic: AnsiItalic,
	Reset: AnsiReset,
	StrikeThrough: AnsiUnderlined,
	Obfuscated: AnsiUnderlined,
}

const (
	AnsiPre = "\u001b["
	AnsiReset = AnsiPre + "0m"

	AnsiBold = AnsiPre + "1m"
	AnsiItalic = AnsiPre + "3m"
	AnsiUnderlined = AnsiPre + "4m"

	AnsiBlack = AnsiPre + "30m"
	AnsiRed = AnsiPre + "31m"
	AnsiGreen = AnsiPre + "32m"
	AnsiYellow = AnsiPre + "33m"
	AnsiBlue = AnsiPre + "34m"
	AnsiMagenta = AnsiPre + "35m"
	AnsiCyan = AnsiPre + "36m"
	AnsiWhite = AnsiPre + "37m"
	AnsiGray = AnsiPre + "30;1m"

	AnsiBrightRed = AnsiPre + "31;1m"
	AnsiBrightGreen = AnsiPre + "32;1m"
	AnsiBrightYellow = AnsiPre + "33;1m"
	AnsiBrightBlue = AnsiPre + "34;1m"
	AnsiBrightMagenta = AnsiPre + "35;1m"
	AnsiBrightCyan = AnsiPre + "36;1m"
	AnsiBrightWhite = AnsiPre + "37;1m"
)

const (
	Pre = "§"

	Black = Pre + "0"
	Blue = Pre + "1"
	Green = Pre + "2"
	Cyan = Pre + "3"
	Red = Pre + "4"
	Magenta = Pre + "5"
	Orange = Pre + "6"
	BrightGray = Pre + "7"
	Gray = Pre + "8"
	BrightBlue = Pre + "9"

	BrightGreen = Pre + "a"
	BrightCyan = Pre + "b"
	BrightRed = Pre + "c"
	BrightMagenta = Pre + "d"
	Yellow = Pre + "e"
	White = Pre + "f"

	Obfuscated = Pre + "k"
	Bold = Pre + "l"
	StrikeThrough = Pre + "m"
	Underlined = Pre + "n"
	Italic = Pre + "o"

	Reset = Pre + "r"
)

/**
 * Converts all MCPE Color codes to ANSI for display in the terminal.
 */
func ConvertMcpeColorsToAnsi(text string) string {
	for toConvert, convertValue := range colorConvert {
		text = strings.Replace(text, toConvert, convertValue, -1)
	}
	return text
}

/**
 * Converts all ANSI Color codes to MCPE for display in-game.
 */
func ConvertAnsiColorsToMcpe(text string) string {
	for convertValue, toConvert := range colorConvert {
		text = strings.Replace(text, toConvert, convertValue, -1)
	}
	return text
}

/**
 * Strips all MCPE colors from the given string.
 */
func StripMcpeColors(text string) string {
	for toConvert := range colorConvert {
		text = strings.Replace(text, toConvert, "", -1)
	}
	return text
}

/**
 * Strips all colors (both ANSI and MCPE) from the given string.
 */
func StripAllColors(text string) string {
	for mcpeColor, ansiColor := range colorConvert {
		text = strings.Replace(text, mcpeColor, "", -1)
		text = strings.Replace(text, ansiColor, "", -1)
	}
	return text
}