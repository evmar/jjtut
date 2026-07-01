# jj workflow: pausing work in progress

Suppose you're working on a new change and notice you made some typo in an
earlier commit. In Git, you might use `git stash` to save your current work to
go make that fix.

In jj, because your work is already saved in the current commit, there is no
need to explicitly stash. So this "workflow" is kind of empty: if you want to
edit a commit somewhere else or make a new one, just run `jj edit` or `jj new`,
as your current state has already been saved.

One small tip: to make it easier to pick up where you left off before moving
away from it, give your current commit a note using `jj desc`. I sometimes use
'wip' for "work in progress" as a reminder.

```
$ jj desc -m 'wip sketching out new api for ...'
$ jj edit ...
```

[Back to the workflow list](../).
