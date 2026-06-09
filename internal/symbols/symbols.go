package symbols

import (
	"os/exec"
	"strings"
)

var nerdFontSymbols = map[string]string{
	"command_palette": " ",
	"oci":             "󱋩",
	"boot":            "󰟀",
	"healthy":         "󰄳",
	"website":         "󰌹",
	"issues":          "󰊤",
	"docs":            "󰈙",
	"discuss":         "󰊌",
	"discord":         "󰙯",
	"matrix":          "󰊌",
	"bluesky":         "",
	"mastodon":        "󰫑",
	"donate":          "󱢏",
	"link":            "󰌹",
	"fedora":          "",
	"steam":           "󰓓",
	"kubernetes":      "󱃾",
	"penguin":         "󰻀",
}

var asciiSymbols = map[string]string{
	"command_palette": ">_",
	"oci":             "[Ci]",
	"healthy":         "✓",
}

func getNerdFontSymbols() bool {
	out, err := exec.Command("fc-list").Output()
	if err != nil {
		return false
	}
	lower := strings.ToLower(string(out))
	return strings.Contains(lower, "symbolsnerdfont") ||
		strings.Contains(lower, "nerdfontssymbolsonly")
}

var hasNerdFontSymbols bool = getNerdFontSymbols()

func GetSymbol(symbolName string) string {
	if hasNerdFontSymbols {
		if symbol, ok := nerdFontSymbols[symbolName]; ok {
			return symbol
		}
	}
	if symbol, ok := asciiSymbols[symbolName]; ok {
		return symbol
	}
	return ""
}
