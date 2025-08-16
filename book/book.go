package main

import (
	"fmt"
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

	body := fmt.Sprintf(`<!doctype html>
<meta name="viewport" content="width=device-width, initial-scale=1">
<link href="%sstyle.css" rel="stylesheet">
<link href="https://fonts.googleapis.com/css2?family=Open+Sans:ital,wght@0,300..800;1,300..800&display=swap" rel="stylesheet">
<main>%s</main>`, root, html)

	dst := filepath.Join(s.dst, strings.TrimSuffix(path, ".md")+".html")
	return os.WriteFile(dst, []byte(body), 0666)
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

func run() error {
	state := State{
		src: "text",
		dst: "html",
	}

	css, err := os.ReadFile("book/style.css")
	if err != nil {
		return err
	}
	if err := os.WriteFile("html/style.css", css, 0666); err != nil {
		return err
	}

	return state.renderAll()
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
	}
}
