

## jj abandon

The `jj abandon` command modifies history to throw away a commit. If you abandon the
current commit, jj will recreate a new empty commit in its place, with a new
change ID.

Like other jj commands, `jj abandon` accepts `-r` to abandon an arbitrary
historical commit, in which case it will rebase.

## jj abandon

Suppose you started a new change, but then changed your mind about it:

```
$ jj new
$ rm foo   # or edited it in some bad way, etc.
```

## jj restore

One way to fix things is `jj restore`, which copies file contents from a different commit. With no
arguments it copies all files from the previous commit, emptying the current
commit of any file changes, effectively clobbering any changes you've made. This
preserves the current commit's description and change ID.
