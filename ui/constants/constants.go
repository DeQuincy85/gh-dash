package constants

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/lipgloss"
	"github.com/dlvhdr/gh-dash/ui/styles"
)

type KeyMap struct {
	Up            key.Binding
	Down          key.Binding
	FirstItem     key.Binding
	LastItem      key.Binding
	TogglePreview key.Binding
	OpenGithub    key.Binding
	Refresh       key.Binding
	PageDown      key.Binding
	PageUp        key.Binding
	NextSection   key.Binding
	PrevSection   key.Binding
	Help          key.Binding
	Quit          key.Binding
}

type Dimensions struct {
	Width  int
	Height int
}

var (
	WaitingGlyph = lipgloss.NewStyle().Foreground(styles.DefaultTheme.FaintText).Render("")
	FailureGlyph = lipgloss.NewStyle().Foreground(styles.DefaultTheme.WarningText).Render("")
	SuccessGlyph = lipgloss.NewStyle().Foreground(styles.DefaultTheme.SuccessText).Render("")
)
