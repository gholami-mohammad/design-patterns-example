package main

import (
	"builder/page"
	"fmt"
)

func main() {
	htmlpage := page.
		NewPageDirector(&page.HTMLPageBuilder{}).
		BuildPage("First Page Header", []string{
			"<p>this the section one</p>",
			"<small>this the second section</small>",
		})
	fmt.Println(htmlpage.Render())

	mdpage := page.
		NewPageDirector(&page.MarkdownPageBuilder{}).
		BuildPage("Readme Header", []string{
			"Section 1",
			"Section 2",
		})
	fmt.Println(mdpage.Render())

}
