package formatter

import (
	"github.com/areon546/go-files/files"
	"github.com/areon546/go-files/table"
)

type Formatter interface {
	Link(displayText, path string) string
	Embed(path string) string
	Heading(tier int, heading string) string
	Table(t table.Table) string
	Bold(s string) string
	Italic(s string) string
}

// ~~~~~~~~~~~~~~~~~~~~ FormattedFile
type FormattedFile struct {
	files.TextFile
	Fmt Formatter
}

// func NewHTMLFile(path, filename, IWANTERRORS string) *FormattedFile {
// 	filepath := files.AddFileType(path+filename, "html")
// 	return newFormattedFile(NewMarkdownFormatter(), filepath)
// }

func NewMarkdownFile(path, filename, IWantErrorsWrong string) *FormattedFile {
	filepath := files.AddFileType(path+filename, "md")
	return newFormattedFile(NewMarkdownFormatter(), filepath)
}

func newFormattedFile(fmt Formatter, filePath string) *FormattedFile {
	return &FormattedFile{TextFile: *files.NewTextFile(filePath), Fmt: fmt}
}

func (m *FormattedFile) AppendLink(displayText, path string) {
	m.Append(m.Fmt.Link(displayText, path), false)
}

func (m *FormattedFile) AppendEmbed(path string) {
	m.Append(m.Fmt.Embed(path), false)
}

func (m *FormattedFile) AppendHeading(tier int, heading string) {
	m.Append(m.Fmt.Heading(tier, heading), false)
}

func (m *FormattedFile) AppendItalics(heading string) {
	m.Append(m.Fmt.Italic(heading), false)
}

func (m *FormattedFile) AppendBold(heading string) {
	m.Append(m.Fmt.Bold(heading), false)
}
