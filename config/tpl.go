package config

import "html/template"

// TPL : The main template variable - contains all of the templates
var TPL *template.Template

//
// init : Initialize the template
func init() {
	TPL = template.Must(template.ParseGlob("ui/html/*.ghtml"))
}
