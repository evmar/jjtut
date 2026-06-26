package main

import (
	"bytes"
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

	tmpl *template.Template
}

func (s *State) render(path string) error {
	md, err := os.ReadFile(filepath.Join(s.src, path))
	if err != nil {
		return err
	}

	firstLine, _, _ := bytes.Cut(md, []byte("\n"))
	title, found := strings.CutPrefix(string(firstLine), "# ")
	if !found {
		return fmt.Errorf("expected first line to contain title")
	}
	subtitle := "Evan's Jujutsu Tutorial"
	if title != subtitle {
		title = title + " — " + subtitle
	}

	extensions := parser.FencedCode | parser.BackslashLineBreak
	options := html.RendererOptions{
		Flags: html.FlagsNone,
		//RenderNodeHook: syntaxHighlightRenderHook,
	}
	html := markdown.ToHTML(md, parser.NewWithExtensions(extensions), html.NewRenderer(options))

	var dst string
	if strings.HasSuffix(path, "index.md") {
		dst = strings.TrimSuffix(path, ".md") + ".html"
	} else {
		dst = strings.TrimSuffix(path, ".md") + "/index.html"
	}
	root := strings.Repeat("../", strings.Count(dst, "/"))

	dst = filepath.Join(s.dst, dst)
	dstDir := filepath.Dir(dst)
	if err := os.MkdirAll(dstDir, 0777); err != nil {
		return err
	}
	f, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer f.Close()

	return s.tmpl.Execute(f, struct {
		Title string
		Body  template.HTML
		Root  string
	}{
		Title: title,
		Body:  template.HTML(string(html)),
		Root:  root,
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
		if filepath.Base(path) == "toc.md" {
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

func loadTemplates() (*template.Template, error) {
	tmpl, err := template.ParseFiles("book/page.gotmpl")
	if err != nil {
		return nil, err
	}

	md, err := os.ReadFile("text/toc.md")
	if err != nil {
		return nil, err
	}
	html := markdown.ToHTML(md, parser.New(), html.NewRenderer(html.RendererOptions{}))
	toc, err := template.New("toc").Parse(string(html))
	if err != nil {
		return nil, err
	}
	if _, err := tmpl.AddParseTree("toc", toc.Tree); err != nil {
		return nil, err
	}

	return tmpl, nil
}

func run() error {
	tmpl, err := loadTemplates()
	if err != nil {
		return err
	}

	if err := os.MkdirAll("html", 0777); err != nil {
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

	return state.renderAll()
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
	}
}
