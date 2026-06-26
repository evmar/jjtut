# Fixing mistakes

Sometimes you start on a change and realize it's a bad idea.  Sometimes
after you've moved on you realize your mistake.

## jj restore

`jj restore` copies file contents from the previous commit. You can specify
which paths to restore, and with no arguments it copies all files.

This empties the current commit of file changes, effectively deleting the changes you've
made. This preserves the current commit's description and change ID.

## jj abandon

The `jj abandon` command modifies history to throw away a commit. If you abandon
the current commit, jj will recreate a new empty commit in its place, with a new
change ID and empty description.

The distinction with `jj restore` may seem small, but it will become more
important in more complex workflows.

## jj squash

Even after you've finished a commit and started a new one, you might notice
you made a typo in some file that you really ought to have included.

To fix this, you can fix the typo in the current commit, then run `jj squash`
It takes any changes in the current foo commit and merges them into the
previous one.
