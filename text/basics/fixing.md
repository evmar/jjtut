# Fixing mistakes

Sometimes you start on a change and realize it's a bad idea. Sometimes only
after you've moved on do you realize your mistake. Both are easy to fix.

## jj restore

`jj restore` copies file contents from the previous commit. With no arguments,
this empties the current commit of file changes, effectively deleting all
changes you've made. This preserves the current commit's description and change
ID.

```
$ rm foo
[... whoops, didn't mean to do that!]
$ jj restore
Working copy  (@) now at: vwvywzku (empty) (no description set)
Parent commit (@-)      : [...]
Added 1 files, modified 0 files, removed 0 files
```

You also can specify which paths to restore to only undo specific files.

## jj abandon

The `jj abandon` command throws away a commit. If you abandon the current
commit, jj will recreate a new empty commit in its place, with a new change ID
and empty description.

```
$ echo "rewrite from scratch" > foo
$ jj desc -m "a bold new plan"
[... wait a sec, that was a bad idea!]
$ jj abandon
Abandoned 1 commits:
[...]
```

The distinction with `jj restore` may seem small, but it will become more
important in more complex workflows.

## jj squash

Even after you've finished a commit and started a new one, you might notice you
made a typo in some file that you really ought to have included.

To fix this, you can fix the typo in the current commit, then run `jj squash` It
takes any changes in the current commit and merges them into the previous one.
(Git users might compare this to `git commit --amend`, and also might notice the
`squash` terminology from `git rebase -i`.)

`jj squash` becomes more important in more complex workflows, which we'll get to
later.

## The operation log

These commands that throw away code may seem risky, but it's also easy to fix
mistakes when using them.

jj records any changes you make in a log called the
"[operation log](https://docs.jj-vcs.dev/operation-log/)". This includes not
only the state after you run a jj command, but also any time it updates the
current commit due to noticing a changed file. This log is private and distinct
from your history of commits.

The `jj op` command has a collection of subcommands to view this log, see diffs,
and restore state from it, but I've almost never needed to use it because
there's a simpler alternative:

## jj undo

The `jj undo` command undoes the most recent entry in the op log, and `jj redo`
goes the other way.

```
$ jj restore
[... whoops, I didn't mean to delete everything!]
$ jj undo
[... files are back.]
```

A Git note: this feature is comparable to the Git reflog. But in my experience,
sometimes I lost things with Git that weren't contained in the reflog. In
contrast, because jj commands always snapshot your current commit, the jj op log
tends to be more fine-grained and pervasive.

## Review

In this chapter, we learned:

- `jj restore`: restore files
- `jj abandon`: throw away a commit
- `jj squash`: merge a commit into its parent, fixing up previous commit
- `jj op`: entry point to op log subcommands
- `jj undo`: undo any previous `jj` command

## Next step

This ends the "basics" part of the tutorial. So far all the commands we've
discussed have worked on the top of a single linear history, but they all also
can work across different branches as well as earlier points in history.

Read on to [part 2, history](../../history), to learn about this.
