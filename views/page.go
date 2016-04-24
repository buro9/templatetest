package views

type Page struct {
	// Name is used to provide a non-localised page name that can be used to
	// determine contextual features, i.e. on a tabbed navigation to highlight
	// the current page
	Name string

	// Title appears in the <title></title> tag
	Title string

	// Description appears in the <meta name="Description" /> tag
	Description string
}
