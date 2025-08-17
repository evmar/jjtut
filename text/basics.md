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
different for you.) jj commit IDs are stable, in that if you edit files, change
their description, or move them around, the commit ID will not change. Currently
the commit is `(empty)` (with no file changes in it) and it has
`(no description set)`.

The hex `7437703d` is the Git ID for the commit. (It will be different for you.)
This will change as the commit changes. You can distinguish jj change IDs from
Git IDs because they don't use the same letters; Git IDs are hex, while jj IDs
don't use those digits.

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

You can also see this by running `jj diff`, which shows a diff of the current
commit:

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

## My second change

To start a new change, use `jj new`.

<pre>
$ jj new
Working copy now at: <b>p</b>wnrkwpn ac9121f7 (empty) (no description set)
Parent commit      : <b>q</b>lmqnzqo b2fa5372 add the foo file
</pre>

We now are editing a new commit `p`, and can do the same commands with it as the
first commit.

> [!WARNING]
> As you get started with jj, forgetting to start new commits before editing
> files will likely be the most common mistake you make. It will feel unnatural
> at first but will get easier.

As a helpful alias, the command `jj commit` combines running `jj desc` followed
by `jj new`. It's useful if you haven't given your commit a description already.

Add another line to foo and use `jj commit` to describe the current commit and
start a new one:

<pre>
$ echo world &gt;&gt; foo
$ jj commit -m 'make foo say hello, world'
Working copy now at: <b>n</b>nlkypwz f58d0c2c (empty) (no description set)
Parent commit      : <b>p</b>wnrkwpn 3d263ba5 make foo say hello, world
</pre>

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

The other option is `jj abandon`, which throws away the current commit. jj will
recreate a new empty commit in its place, with a new change ID.

If you did try making a change here, undo it using one of the above commands to
prepare for the next chapter.

### Review

In this chapter, we learned about:

- `jj status` (`jj st`): show status (files included, etc.) of a change
- `jj diff`: show diff of a change
- `jj desc`: edit description of a change
- `jj new`, `jj commit`: create a new change
- `jj restore`: copy file contents from a change
- `jj abandon`: abandon (delete) a change

If you stopped reading here, you have all you need to productively start using
jj for adding commits to a project. On the other hand, there's little point in
using version control unless you make use of the history you've been
accumulating, so in the next chapter we'll go into the history.
