# The basics

On a fresh respository, run `jj status` (or `jj st` for short) to show the
current status.

```
$ jj st
The working copy has no changes.
Working copy  (@) : qlmqnzqo 7437703d (empty) (no description set)
Parent commit (@-): zzzzzzzz 00000000 (empty) (no description set)
```

This shows the current working copy commit is named `qlmqnzqo`. (It will be
different for you.) jj change IDs are stable: if you change their description,
edit files, or move them around, the change ID will not change.

Currently the commit is `(empty)` (with no file changes in it) and it has
`(no description set)`.

## Changes versus commits

The hex `7437703d` is the Git commit ID. (It will be different for you.) This
will change as the commit changes. You can distinguish jj change IDs from Git
commit IDs because they don't use the same letters; Git commit IDs are hex,
while change IDs don't use those digits.

A terminology note. To be precise, `qlmqnzqo` identifies a "change": the thing
you can edit over time. `7437703d` identifies a "commit" (or synonymously
"revision"): an immutable snapshot. In practice the distinction is effectively
an implementation detail, and you will almost never need to think about Git
commits or commit IDs.

In this tutorial I try to call everything "commits" for simplicity, in
particular because I use the word "change" in its informal sense of "the code
change you're working on".

## First change

Create or edit a file:

```
$ echo hello > foo
```

Per the rules of jj, this edit to a file is considered part of your current
commit. There is no "I'm done, commit this" command to run.

You can see this by running `jj st` again:

```
$ jj st
Working copy changes:
A foo
Working copy : qlmqnzqo 2a5c0d7e (no description set)
Parent commit: zzzzzzzz 00000000 (empty) (no description set)
```

This is not saying "there is a new file named `foo` that is ready to add". It is
saying "the current commit contains a file add named `foo`".

You can also see this by running `jj diff`, which shows the diff of a commit:

```
$ jj diff
Added regular file foo:
        1: hello
```

You can edit the description of the commit with `jj desc`, which with no flags
will pop up your editor. For tutorial purposes I'll provide a description on the
command line:

```
$ jj desc -m 'add the foo file'
```

Coming from other version control, this might feel like you're "committing"
something. But note that this command only adds information to the current
commit, which existed from the start. We could have just as well done it before
making any file edits.

## Second change

To start a new change, use `jj new`, which creates a new commit as a child of
the current commit.

```
$ jj new
Working copy now at: pwnrkwpn ac9121f7 (empty) (no description set)
Parent commit      : qlmqnzqo b2fa5372 add the foo file
```

We now are editing a new commit and can do the same edits with it as the first
commit.

> [!WARNING]
> As you get started with jj, forgetting to start new commits before editing
> files will likely be the most common mistake you make. It will feel unnatural
> at first but will get easier.

## jj commit

As a helpful alias, the command `jj commit` combines running `jj desc` followed
by `jj new`. It's useful if you haven't given your commit a description already.

Add another line to foo and use `jj commit` to describe the current commit and
start a new one:

```
$ echo world >> foo
$ jj commit -m 'make foo say hello, world'
Working copy now at: nnlkypwz f58d0c2c (empty) (no description set)
Parent commit      : pwnrkwpn 3d263ba5 make foo say hello, world
```

## Review

In this chapter, we learned:

- `jj status` (`jj st`): show status (files included, etc.) of a change
- `jj diff`: show diff of a change
- `jj desc`: edit description of a change
- `jj new` and `jj commit`: create a new change

## Next step

If you stopped reading here, you have all you need to start using jj for adding
commits to a project. The next most important thing to learn is
[how to fix mistakes](../fixing).
