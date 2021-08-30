package main

import (
	"github.com/d3z41k/snippetbox/pkg/models"
)

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
