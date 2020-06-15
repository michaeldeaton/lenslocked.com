package views

import "html/template"

func NewView(layout string, files ...string) *View {
	files = append(files,
		"views/layouts/footer.gohtml",
		"views/layouts/bootstrap.gohtml")
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	test := "g"
	return &View{
		Template: t,
		Layout:   layout,
		Test:     test,
	}
}

type View struct {
	Template *template.Template
	Layout   string
	Test     string
}
