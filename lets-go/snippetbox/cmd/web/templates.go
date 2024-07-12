package main

import "github.com/ashidiqidimas/snippetbox/internal/models"

type TemplateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
