package page

type pageDirector struct {
	builder PageBuilder
}

func (d *pageDirector) BuildPage(title string, contentSections []string) Page {
	d.builder.SetType()
	d.builder.SetRenderer()
	d.builder.AddTitle(title)
	for _, s := range contentSections {
		d.builder.AddSection(s)
	}
	return d.builder.GetPage()
}

func NewPageDirector(builder PageBuilder) *pageDirector {
	return &pageDirector{
		builder: builder,
	}
}
