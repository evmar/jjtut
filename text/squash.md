# jj squash

When you want to make a quick fix in a commit, using `jj edit` to jump directly
to it and editing the file directly is the quickest way. But if you make a
mistake (cat on the keyboard?) there isn't a way to diff or undo your new edits
as distinct from the existing edits.

An alternative workflow is instead to create a new commit before editing, like
we just did in the previous section. After editing files, you run `jj diff` to
see just the new changes you've made.

When you're happy with the change, the `jj squash` command will "squash" the
current commit into the previous one, merging their changes together. (Git users
might recognize the "squash" terminology from `git rebase -i`.) After running
`jj squash`, you're left pointed at an empty commit; if you go elsewhere with
`jj edit`, it will be automatically abandoned.

## Recipe: jj squash and the Git index

Advanced Git users might be accustomed to building up a single change by
incrementally adding parts of it to the index, which in Git can be thought of as
a staging area for a pending change.

jj models the workflow of building up a change by using regular commits.

1. Start a new staging commit on top of your current one (run `jj new`).
1. Edit files as needed; whenever you're happy with them, `jj squash` them into
   the previous commit, or use `jj restore` to undo them.
1. Repeat that step as many times as necessary.
1. Use `jj abandon`, if needed, to abandon the staging commit.

## The squash workflow

Because it's free to make new commits anywhere, one way to use jj is to _always_
use `jj new` when editing history. That is, if you want to make a change to some
commit X, you instead say `jj new X` to start a new commit based on X, and then
`jj squash` when you're done with your edits. This allows you to leave X in its
original statecontained and what new changes you've made atop X.

When you move elsewhere (via `jj new` or `jj edit`) the empty temporary commit
you've been using on will be automatically cleaned up.

## Review

In this chapter, we learned:

- `jj abandon`: abandon (delete) a commit
