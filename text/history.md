# History

Run `jj log` to see the repository history. (I've removed dates from my output
just to reduce visual noise.)

```
@  umnvtwlo my@email 20de4517
│  (empty) (no description set)
○  pwnrkwpn my@email 08b3e414
│  make foo say hello, world
○  qlmqnzqo my@email git_head() d6b14a5d
│  add the foo file
◆  zzzzzzzz root() 00000000
```

From the top:

1. umnvtwlo is an empty commit
1. pwnrkwpn and qlmqnzqo were our two edits
1. zzzzzzzz is a special "root" commit that starts the repository and is is
   always empty

In a terminal, these commits will have some prefix (often the initial letter)
highlighted or in bold. This indicates the unique prefix of the commit that can
be used to refer to it in commands. jj accepts either jj or Git commit IDs.

## Diffs and revsets

We've used `jj diff` to see the diff of the current change. `diff` (and many
other jj commands) can also be told which change to work with by using the `-r`
flag.

Here, I use `q` as the unique prefix of the commit above:

```
$ jj diff -r q
[... same diff output as earlier ...]
```

The argument passed to `-r` is called a _revset_, and in jj it is a miniature
[programming language for specifying commits](https://jj-vcs.github.io/jj/latest/revsets/).

In practice, almost all you need to know is that the alias "`@`" refers to the
current commit, and the operator "`-`" means "the commit before". So the revset
"`@-`" means "the commit before the current one". You can see both of these in
the `jj status` output.

Putting it all together, you can now understand that `jj diff` is a short way of
saying `jj diff -r @`. (This is quite different from Git, with its flags for
diffs involving the working copy or the index!)

## Modifying history

Another command that accepts a revision by `-r` is `desc`. You can change the
description of our first commit to be more descriptive:

```
$ jj desc -r q -m "add a foo file that says hello"
```

Like `diff`, you can now see that `jj desc` without flags edits `@`.

Users coming to jj from another version control system might be surprised here:
in jj, your commit history is generally freely editable.

In case of making mistakes, jj has powerful undo functionality. And when working
with Git, jj has additional functionality related to not accidentally modifying
commits you shouldn't. We'll get to both of these later.

## Rebasing

After making the above modification to history, you'll notice a new line in the
output:

```
$ jj desc [...as above...]
Rebased 2 descendant commits
[...]
```

What happened? Whenever you modify history, jj updates downstream commits atop
that change. In Git terms this is a "rebase", but in jj these happen implicitly
and frequently.

## Review

In this chapter, we learned:

- specify revisions using the `-r` flag to `diff` and `desc`
- there exists a 'revset' language for specifying revisions
- history is mutable
- editing history causes downstream changes to update

## Next step

[Learn about editing history](editing-history.html).
