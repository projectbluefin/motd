package main

import (
	"os/exec"
	"strings"

	"charm.land/glamour/v2/ansi"
)

const defaultMargin uint = 2

func stringPtr(s string) *string { return &s }
func boolPtr(b bool) *bool       { return &b }
func uintPtr(u uint) *uint       { return &u }

// The colors are ANSI color codes (ANSI 256 - https://en.wikipedia.org/wiki/ANSI_escape_code#Colors)

var colorTheme = map[string]map[string]string{
	"blue":   {"accent": "33", "link": "69"},
	"green":  {"accent": "34", "link": "28"},
	"orange": {"accent": "208", "link": "130"},
	"pink":   {"accent": "212", "link": "163"},
	"purple": {"accent": "165", "link": "164"},
	"red":    {"accent": "203", "link": "124"},
	"slate":  {"accent": "104", "link": "104"},
	"teal":   {"accent": "44", "link": "38"},
	"yellow": {"accent": "220", "link": "178"},
}

func getColorTheme() map[string]string {
	defaultColorTheme := colorTheme["blue"]
	cmd := exec.Command("gsettings", "get", "org.gnome.desktop.interface", "accent-color")
	output, err := cmd.Output()
	if err != nil {
		return defaultColorTheme
	}
	accent := strings.Trim(strings.TrimSpace(string(output)), "'")
	if color, ok := colorTheme[accent]; ok {
		return color
	}
	return defaultColorTheme
}

func getAccentStyle() ansi.StyleConfig {
	theme := getColorTheme()
	accent := theme["accent"]
	link := theme["link"]

	return ansi.StyleConfig{
		Document: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				BlockPrefix: "\n",
				BlockSuffix: "\n",
			},
			Margin: uintPtr(defaultMargin),
		},
		BlockQuote: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Italic: boolPtr(true),
			},
			Indent:      uintPtr(1),
			IndentToken: stringPtr("│ "),
		},
		List: ansi.StyleList{
			LevelIndent: defaultMargin,
		},
		Heading: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				BlockSuffix: "\n",
				Color:       stringPtr(accent),
				Bold:        boolPtr(true),
			},
		},
		H1: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				BlockPrefix: "\n",
				BlockSuffix: "\n",
			},
		},
		H2: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{Prefix: "▌ "},
		},
		H3: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{Prefix: "┃ "},
		},
		H4: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{Prefix: "│ "},
		},
		H5: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{Prefix: "┆ "},
		},
		H6: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Prefix: "┊ ",
				Bold:   boolPtr(false),
			},
		},
		Strikethrough: ansi.StylePrimitive{CrossedOut: boolPtr(true)},
		Emph:          ansi.StylePrimitive{Italic: boolPtr(true)},
		Strong: ansi.StylePrimitive{
			Color: stringPtr(accent),
			Bold:  boolPtr(true),
		},
		HorizontalRule: ansi.StylePrimitive{
			Color:  stringPtr(accent),
			Format: "\n──────\n",
		},
		Item: ansi.StylePrimitive{
			BlockPrefix: "• ",
		},
		Enumeration: ansi.StylePrimitive{
			BlockPrefix: ". ",
		},
		Task: ansi.StyleTask{
			Ticked:   "[✓] ",
			Unticked: "[ ] ",
		},
		Link: ansi.StylePrimitive{
			Color:     stringPtr(link),
			Underline: boolPtr(true),
		},
		LinkText: ansi.StylePrimitive{Bold: boolPtr(true)},
		Image:    ansi.StylePrimitive{Underline: boolPtr(true)},
		ImageText: ansi.StylePrimitive{
			Format: "Image: {{.text}}",
		},
		Code: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Prefix: " ",
				Suffix: " ",
				Color:  stringPtr(accent),
				Bold:   boolPtr(true),
			},
		},
		CodeBlock:             ansi.StyleCodeBlock{},
		Table:                 ansi.StyleTable{},
		DefinitionDescription: ansi.StylePrimitive{BlockPrefix: "\n🠶 "},
	}
}
