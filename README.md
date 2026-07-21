# Evan's Jujutsu Tutorial

(This README is for the underlying text of the tutorial. For the tutorial
itself, see <https://evmar.github.io/jjtut/>.)

This text was all written by hand by [Evan Martin](https://neugierig.org/), a
human. I used no AI, and disclaim this in particular so that you can expect that
the choices made in what to mention, what to leave out, and in what order were
carefully considered. (After writing most of it I did use AI to check for
mistakes.)

I welcome feedback and suggestions for the tutorial, please file issues!

## Text

The tutorial text is in `text/`. Run `dprint fmt` to reformat Markdown after
edits.

## Code

The site generation code is in `book/`. (I considered tools like mkdocs but it
seemed easier to just generate it myself.)

## Hacking

Run `./watch.sh` to automatically updated HTML output in `html/` while you edit.

I use `python3 -m http.server` to then serve the content locally in my browser.
(You can't point a browser directly at the directory and instead need a web
server because content is rendered to index.html files.)
