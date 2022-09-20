package main

import (
	_ "embed"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

//go:embed 仿宋_GB2312.ttf
var font []byte

var myfont = &fyne.StaticResource{
	StaticName:    "LXGWWenKaiMono",
	StaticContent: font,
}

type myTheme struct{}

var _ fyne.Theme = (*myTheme)(nil)

// return bundled font resource
func (*myTheme) Font(s fyne.TextStyle) fyne.Resource {
	if s.Monospace {
		return theme.LightTheme().Font(s)
		// return theme.DefaultTheme().Font(s)
	}
	if s.Bold {
		if s.Italic {
			return theme.LightTheme().Font(s)
			// return theme.DefaultTheme().Font(s)
		}
		return myfont
	}
	if s.Italic {
		return theme.LightTheme().Font(s)
		// return theme.DefaultTheme().Font(s)
	}
	return myfont
}

func (*myTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	return theme.LightTheme().Color(n, v)
}

func (*myTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.LightTheme().Icon(n)
	// return theme.DefaultTheme().Icon(n)
}

func (*myTheme) Size(n fyne.ThemeSizeName) float32 {
	return theme.LightTheme().Size(n)
	// return theme.DefaultTheme().Size(n)
}
