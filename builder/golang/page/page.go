package page

type Page struct {
	Title           string
	ContentSections []string
	Type            string
	renderer        Renderer
}

func (p Page) Render() string {
	return p.renderer.Render(p)
}

type PageBuilder interface {
	AddTitle(title string)
	AddSection(section string)
	SetType()
	SetRenderer()
	GetPage() Page
}
