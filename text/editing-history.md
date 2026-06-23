## Abandoning a change

Suppose next you accidentally clobbered your important work:

```
$ rm foo   # or edited it, etc.
```

There are two ways to undo this, with two different ways of thinking about it.

One is `jj restore`, which copies file contents from a different commit. With no
arguments it copies all files from the previous commit, emptying the current
commit of any file changes, effectively clobbering any changes you've made. This
preserves the current commit's description and change ID.

The other option is `jj abandon`, which throws away a commit. If you abandon the
current commit, jj will recreate a new empty commit in its place, with a new
change ID.

Like other jj commands, `jj abandon` accepts `-r` to abandon an arbitrary
historical commit.

## Review

- `jj restore`: copy file contents from a change
- `jj abandon`: abandon (delete) a change
