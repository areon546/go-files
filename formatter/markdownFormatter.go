package formatter

import "github.com/areon546/go-ds/table"

type markdownFormatter struct{}

func NewMarkdownFormatter() markdownFormatter {
	return markdownFormatter{}
}

func (m markdownFormatter) Link(displayText, link string) string {
	return markdownLink(false, displayText, link)
}

func (m markdownFormatter) Embed(embed, alt string) string {
	return markdownLink(true, alt, embed)
}

func (m markdownFormatter) Heading(tier int, heading string) string {
	s := ""

	for range tier {
		s += "#"
	}

	return s + " " + heading
}

func (m markdownFormatter) Table(t table.Table) string {
	s := ""
	headers := ""

	headersRow, _ := t.Headers()
	if t.HasHeaders() {
		headers = constructRow(headersRow)
		s += headers + "\n"
	}
	headerDecleration := markdownHeaderDeclarationRow(t.Width())

	s += headerDecleration + "\n"

	// Add records
	for _, rec := range t.Iter() {
		s += constructRow(rec) + "\n"
	}

	return s
}

func (m markdownFormatter) Bold(s string) string {
	return format("**%s**", s)
}

func (m markdownFormatter) Italic(s string) string {
	return format("*%s*", s)
}

// Helper functions.
func markdownLink(embed bool, displayText, path string) (s string) {
	if embed {
		s += "!"
	}
	s += format("[%s](%s)", displayText, path)
	return
}

func constructRow(r table.Row) string {
	return r.Join("|", "", "|", "", "|")
}

func markdownHeaderDeclarationRow(length int) string {
	headerDecleration := table.NewRow(length)

	for i := range length {
		headerDecleration.Set(i, "---")
	}

	return constructRow(*headerDecleration)
}
