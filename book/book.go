package main

import (
	"fmt"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type State struct {
	src, dst string

	toc  template.HTML
	tmpl *template.Template
}

func (s *State) render(path string) error {
	md, err := os.ReadFile(filepath.Join(s.src, path))
	if err != nil {
		return err
	}
	extensions := parser.FencedCode | parser.BackslashLineBreak
	options := html.RendererOptions{
		Flags: html.FlagsNone,
		//RenderNodeHook: syntaxHighlightRenderHook,
	}
	html := markdown.ToHTML(md, parser.NewWithExtensions(extensions), html.NewRenderer(options))

	root := strings.Repeat("../", strings.Count(path, "/"))

	dst := filepath.Join(s.dst, strings.TrimSuffix(path, ".md")+".html")
	f, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer f.Close()

	return s.tmpl.Execute(f, struct {
		Nav, Body template.HTML
		Root      string
	}{
		Nav:  template.HTML(s.toc),
		Body: template.HTML(string(html)),
		Root: root,
	})
}

func (s *State) renderAll() error {
	os.Mkdir(s.dst, 0777)
	return filepath.WalkDir(s.src, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !strings.HasSuffix(path, ".md") {
			return nil
		}
		if filepath.Base(path) == "contents.md" {
			return nil
		}
		path, err = filepath.Rel(s.src, path)
		if err != nil {
			return err
		}
		fmt.Println(path)
		return s.render(path)
	})
}

func (s *State) loadTOC() (template.HTML, error) {
	md, err := os.ReadFile(filepath.Join(s.src, "contents.md"))
	if err != nil {
		return template.HTML(""), err
	}
	html := markdown.ToHTML(md, parser.New(), html.NewRenderer(html.RendererOptions{}))
	return template.HTML(string(html)), nil
}

func run() error {
	if err := os.MkdirAll("html", 0777); err != nil {
		return err
	}

	page, err := os.ReadFile("book/page.gotmpl")
	if err != nil {
		return err
	}
	tmpl, err := template.New("page").Parse(string(page))
	if err != nil {
		return err
	}

	css, err := os.ReadFile("book/style.css")
	if err != nil {
		return err
	}
	if err := os.WriteFile("html/style.css", css, 0666); err != nil {
		return err
	}

	state := State{
		src:  "text",
		dst:  "html",
		tmpl: tmpl,
	}
	toc, err := state.loadTOC()
	if err != nil {
		return err
	}
	state.toc = toc

	return state.renderAll()
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
	}
}
