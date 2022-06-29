package page

type HTMLPageBuilder struct {
	Page
}

func (builder *HTMLPageBuilder) AddTitle(title string) {
	builder.Title = title
}

func (builder *HTMLPageBuilder) AddSection(section string) {
	if section == "" {
		return
	}
	builder.ContentSections = append(builder.ContentSections, section)
}

func (builder *HTMLPageBuilder) SetType() {
	builder.Type = "html"
}

func (builder *HTMLPageBuilder) SetRenderer() {
	builder.renderer = &HTMLRenderer{}
}

func (builder *HTMLPageBuilder) GetPage() Page {
	return builder.Page
}
