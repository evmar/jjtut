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

type PageParams struct {
	Title   string
	WebPath string
	Body    template.HTML
	Root    string
}

func (s *State) render(path string) error {
	md, err := os.ReadFile(filepath.Join(s.src, path))
	if err != nil {
		return err
	}

	firstLine, rest, _ := bytes.Cut(md, []byte("\n"))
	title, found := strings.CutPrefix(string(firstLine), "# ")
	if !found {
		return fmt.Errorf("expected first line to contain title")
	}

	extensions := parser.FencedCode | parser.BackslashLineBreak
	options := html.RendererOptions{
		Flags: html.FlagsNone,
		//RenderNodeHook: syntaxHighlightRenderHook,
	}
	html := markdown.ToHTML(rest, parser.NewWithExtensions(extensions), html.NewRenderer(options))

	var webPath string
	if p, ok := strings.CutSuffix(path, "index.md"); ok {
		webPath = p
	} else {
		webPath = strings.TrimSuffix(path, ".md") + "/"
	}
	dst := webPath + "index.html"
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

	return s.tmpl.ExecuteTemplate(f, "page.gotmpl", &PageParams{
		Title:   title,
		WebPath: webPath,
		Body:    template.HTML(string(html)),
		Root:    root,
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
		path, err = filepath.Rel(s.src, path)
		if err != nil {
			return err
		}
		fmt.Println(path)
		return s.render(path)
	})
}

func loadTemplates() (*template.Template, error) {
	funcs := template.FuncMap{
		"pagelink": func(cur *PageParams, title, path string) template.HTML {
			if path == cur.WebPath {
				return template.HTML(fmt.Sprintf("<b>%s</b>", title))
			} else {
				return template.HTML(fmt.Sprintf("<a href='%s%s'>%s</a>", cur.Root, path, title))
			}
		},
	}
	return template.New("").Funcs(funcs).ParseFiles("book/page.gotmpl", "text/toc.gotmpl")
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
