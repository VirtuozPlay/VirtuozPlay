package actions

import (
	"encoding/json"
	"fmt"
	"github.com/gobuffalo/helpers/tags"
	"github.com/gobuffalo/plush/v4"
	tags2 "github.com/gobuffalo/tags/v3"
	"io/fs"
	"virtuozplay/templates"

	"github.com/gobuffalo/buffalo/render"
	"html/template"
)

var r *render.Engine

// The global manifest instance, created by loadManifest
var m *manifest

// JSON to Go translation of Vite's manifest.json file
type manifest = map[string]mapping

type mapping struct {
	File    string   `json:"file"`
	Src     string   `json:"src"`
	CSS     []string `json:"css"`
	IsEntry bool     `json:"isEntry"`
}

func init() {
	r = render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: "application.plush.html",

		// fs.FS containing templates
		TemplatesFS: templates.FS(),

		// Add template helpers here:
		Helpers: render.Helpers{
			// for non-bootstrap form helpers uncomment the lines
			// below and import "github.com/gobuffalo/helpers/forms"
			// forms.FormKey:     forms.Form,
			// forms.FormForKey:  forms.FormFor,
			"viteClientTag":      viteClientTag(),
			"viteEntryPointTags": viteEntryPointTags(DistFS()),
		},
	})
}

// Generates the HTML tags needed to load the given entrypoint in both dev and prod mode.
func viteEntryPointTags(assetsFS fs.FS) func(entrypoint string, c plush.Context) (template.HTML, error) {
	return func(entrypoint string, c plush.Context) (template.HTML, error) {
		// When not in production, load entrypoint from Vite server
		if ENV != "production" {
			return jsm("/" + entrypoint), nil
		}

		// When in production use the manifest to get the location of the entrypoint
		if m == nil {
			if err := loadManifest(assetsFS); err != nil {
				return "", err
			}
		}

		entryMapping, ok := (*m)[entrypoint]

		if !ok {
			return "", fmt.Errorf("no mapping found for entrypoint %s", entrypoint)
		}

		container := tags2.New("div", tags2.Options{})
		container.Append(jsm(entryMapping.File))

		for _, style := range entryMapping.CSS {
			container.Append(tags.CSS(style, tags2.Options{}))
		}

		return container.HTML(), nil
	}
}

// Generates the tag needed to enable Vite hot reloading functionality in development mode
func viteClientTag() func() template.HTML {
	return func() template.HTML {
		if ENV != "production" {
			return jsm("/@vite/client")
		}
		return ""
	}
}

// Generates a JavaScript module (ESM) HTML tag
func jsm(src string) template.HTML {
	return tags.JS(src, tags2.Options{
		"type": "module",
	})
}

func loadManifest(assetsFS fs.FS) error {
	mFile, err := assetsFS.Open("manifest.json")
	if err != nil {
		return err
	}

	m2 := map[string]mapping{}
	err = json.NewDecoder(mFile).Decode(&m)

	if closeErr := mFile.Close(); closeErr != nil {
		fmt.Println("Failed to close manifest file", closeErr)
	}
	if err != nil {
		*m = m2
	}

	return err
}
