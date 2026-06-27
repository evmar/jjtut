# History

Run `jj log` (or just `jj` by itself, if you followed the default configuration
in the [setup](../../basics/setup) section) to see the repository history. (In
this tutorial, I've removed dates from my output just to reduce visual noise.)

```
@  umnvtwlo my@email 20de4517
│  (empty) (no description set)
○  pwnrkwpn my@email 08b3e414
│  make foo say hello, world
○  qlmqnzqo my@email d6b14a5d
│  add the foo file
◆  zzzzzzzz root() 00000000
```

From the top:

1. umnvtwlo is the empty current commit
1. pwnrkwpn and qlmqnzqo were two edits I made
1. zzzzzzzz is a special "root" commit that starts the repository and is is
   always empty

In a terminal, these commits will have some prefix (often the initial letter)
highlighted or in bold. This indicates the unique prefix of the commit that can
be used to refer to it in commands. jj commands accept either jj or Git commit
IDs.

## Diffs and revsets

We've used `jj diff` to see the diff of the current change. `diff` (and many
other jj commands) can also be told which change to work with by using the `-r`
flag.

Here, I use `q` as the unique prefix of my first commit to specify it. (You'll
need to use your own in place of `q`):

```
$ jj diff -r q
[... same diff output as earlier ...]
```

The argument passed to `-r` is called a _revset_, and in jj it is a miniature
[programming language for specifying commits](https://jj-vcs.github.io/jj/latest/revsets/).
The `root()` shown in the above log output is a function in that language that
gets the repository's root commit, and it's shown in the log output as another
way of specifying that commit.

In practice, almost all you need to know is that the alias "`@`" refers to the
current commit, and the operator "`-`" means "the commit before". So the revset
"`@-`" means "the commit before the current one", while `p-` means "the commit
before commit `p`". You have seen `@` and `@-` in the `jj status` output.

Putting it all together, you can now understand that `jj diff` is a short way of
saying `jj diff -r @`. (This is quite different from Git, with its flags for
diffs involving the working copy or the index!)

## Review

In this chapter, we learned:

- specify revisions using the `-r` flag to `diff`
- there exists a 'revset' language for specifying revisions
- `@` refers to the current commit, `@-` refers to the previous one

## Next step

[Learn about editing commits](editing).
