package config

import "html/template"

// AppConfig hold's application config

type AppConfig struct {
	TemplateCache map[string]*template.Template
}
