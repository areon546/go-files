package formatter

import (
	fileIO "github.com/areon546/go-files/files"
	"github.com/areon546/go-files/table"
)

type Formatter interface {
	FormatLink(displayText, path string) string
	FormatEmbed(path string) string
	FormatHeading(tier int, heading string) string
	FormatTable(t table.Table, headers bool) string
	FormatBold(s string) string
	FormatItalic(s string) string
}

// ~~~~~~~~~~~~~~~~~~~~ FormattedFile
type FormattedFile struct {
	fileIO.TextFile
	Fmt Formatter
}

func NewHTMLFile(path, filename, IWANTERRORS string) *FormattedFile {
	return newFormattedFile(NewMarkdownFormatter(), fileIO.ConstructFilePath(path, filename, "html"))
}

func NewMarkdownFile(path, filename, IWantErrorsWrong string) *FormattedFile {
	return newFormattedFile(NewMarkdownFormatter(), fileIO.ConstructFilePath(path, filename, "md"))
}

func newFormattedFile(fmt Formatter, filePath string) *FormattedFile {
	return &FormattedFile{TextFile: *fileIO.NewTextFile(filePath), Fmt: fmt}
}

func (m *FormattedFile) AppendLink(displayText, path string) {
	m.Append(m.Fmt.FormatLink(displayText, path), false)
}

func (m *FormattedFile) AppendEmbed(path string) {
	m.Append(m.Fmt.FormatEmbed(path), false)
}

func (m *FormattedFile) AppendHeading(tier int, heading string) {
	m.Append(m.Fmt.FormatHeading(tier, heading), false)
}

func (m *FormattedFile) AppendItalics(heading string) {
	m.Append(m.Fmt.FormatItalic(heading), false)
}

func (m *FormattedFile) AppendBold(heading string) {
	m.Append(m.Fmt.FormatBold(heading), false)
}
