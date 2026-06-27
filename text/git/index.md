# Working with Git

Your jj repository can push and pull Git branches, and jj has additional
functionality for working with Git. Try `jj git --help` to see a list of
subcommands.

## First steps

For learning purposes, you could clone this tutorial's Git repository to look
around:

```
$ jj git clone https://github.com/evmar/jjtut.git
$ jj log
@  kpmzwpnw my@email fc5fe481
│  (empty) (no description set)
◆  mzrrnrom evan.martin@gmail.com main@origin ead50b5a
│  broken link
~
```

Here are some things to notice:

1. The commit log ends with a `~` to show it was truncated.
1. The commit `mzrrnrom` is from the upstream repository. It is marked with a
   diamond to show it is _immutable_, a new jj concept.
1. The commit is tagged with `main@origin`.

We'll go into these in order.

## Truncated log

The `jj log` command by default will attempt to omit some commits, including
commits from remote repositories. This is configurable, but for now if you'd
like to see all commits you can use:

```
$ jj log -r ::
```

The `::` is the jj revset expression for "all commits".

## Immutable commits

If you attempt to edit the immutable commit, jj will refuse:

```
$ jj edit mzrrnrom
Error: Commit ead50b5a160a is immutable
Hint: Could not modify commit: mzrrnrom main@origin | broken link
Hint: Immutable commits are used to protect shared history.
Hint: For more information, see:
      - https://docs.jj-vcs.dev/latest/config/#set-of-immutable-commits
      - `jj help -k config`, "Set of immutable commits"
Hint: This operation would rewrite 1 immutable commits.
```

jj considers commits from remote branches immutable so you don't accidentally
edit them.

This immutability is more of a guardrail than a strict law. You can not only
reconfigure which commits are consided immutable, you can also often pass
`--ignore-immutable` to commands that complain in cases where you know better.
It's there to help prevent accidents.

But it's probably not a good idea to bypass the immutability here, because
editing this commit is likely to make things conflict if you attempt to sync
with upstream again.

## Bookmarks

Finally, the `main@origin` bit in the log on the upstream commit is a jj
_bookmark_, a jj feature seen here for the first time.

A jj bookmark is a name attached to a commit, similar to a Git tag. Here, the
bookmark named `main@origin` points at the current `main` branch in the upstream
repository. We'll go into them next.

## Review

- commits coming from Git remotes are immutable and not shown in the default log
- attempting to edit an immutable commit is an error
- Git branches show up as jj bookmarks

## Next step

To push Git changes we need to [learn about bookmarks](bookmarks).
