package main

import (
  "image/color"
  "fyne.io/fyne/v2"
  "fyne.io/fyne/v2/theme"
)

type mikesTheme struct {}

var _ fyne.Theme = (*mikesTheme)(nil)

func (m *mikesTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
  if name == theme.ColorNameBackground {
    if variant == theme.VariantLight {
      return color.White
    }
    return color.Gray{0x7F}
  }
  return theme.DefaultTheme().Color(name, variant)
}

func (m *mikesTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
  if name == theme.IconNameHome {
    return fyne.NewStaticResource("myHome", []byte{0})
  }

  return theme.DefaultTheme().Icon(name)
}

func (m *mikesTheme) Font(style fyne.TextStyle) fyne.Resource {
  return theme.DefaultTheme().Font(style)
}

func (m *mikesTheme) Size(name fyne.ThemeSizeName) float32 {
  return theme.DefaultTheme().Size(name)
}
