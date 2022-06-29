package page

type MarkdownPageBuilder struct {
	Page
}

func (builder *MarkdownPageBuilder) AddTitle(title string) {
	builder.Title = title
}

func (builder *MarkdownPageBuilder) AddSection(section string) {
	if section == "" {
		return
	}
	builder.ContentSections = append(builder.ContentSections, section)
}

func (builder *MarkdownPageBuilder) SetType() {
	builder.Type = "md"
}

func (builder *MarkdownPageBuilder) GetPage() Page {
	return builder.Page
}

func (builder *MarkdownPageBuilder) SetRenderer() {
	builder.renderer = &MarkdownRenderer{}
}
