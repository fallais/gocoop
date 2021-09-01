package main

import (
	"embed"

	"github.com/fallais/gocoop/cmd"
	"github.com/fallais/gocoop/internal"
	"github.com/fallais/gocoop/internal/routes"
)

//go:embed static
var staticFS embed.FS

//go:embed templates
var templatesFS embed.FS

func main() {
	internal.StaticFS = staticFS
	routes.TemplatesFS = templatesFS

	cmd.Execute()
}
