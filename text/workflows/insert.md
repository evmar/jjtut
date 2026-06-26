# jj workflow: inserting commits

Suppose you are halfway through a change and you notice some smaller change that
you had ought to have done first.

One way to do this is `jj new` at the earlier point, with the plan to eventually
move your in-progress change after this new change. But this means you must do
the work to transplant that change.

An easier way is to know that `jj new` can insert a commit in between two
others. It takes a flag of where to insert, either `-B` (insert Before) or `-A`
(insert After) another commit.

So you can just run:

```
$ jj new -B @
```

and you will now be editing a new commit, inserted just before your current
work. Your in-progress change will be rebased on top of any edits you make.

[Back to the workflow list](../).
