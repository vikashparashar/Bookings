package config

import (
	"html/template"
	"log"
)

// AppConfig hold's application config

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLogger    *log.Logger
}
