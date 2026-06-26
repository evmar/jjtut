# Editing commits

Another command that accepts a revision by `-r` is `desc`. You can change the
description of our first commit to be more descriptive:

```
$ jj desc -r q -m "add a foo file that says hello"
```

Like with `jj diff`, you can now see that running `jj desc` without flags edits
`@`, but both can be pointed at any commit.

Users coming to jj from another version control system might be surprised (or
alarmed!) here by how making new changes and modifying history are the same
commands, just pointed at different places. In jj, your commit history is
generally freely editable and feels pretty similar to editing the newest change.

Recall that `jj undo` helps if you make any mistakes. And when working with Git,
jj has additional functionality related to not accidentally modifying commits
you shouldn't. We'll get that later.

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

## Jumping around

You can change which commit you're currently editing with `jj edit`. Switch back
to the first commit we created:

```
$ jj edit q
```

If you run `jj log`, note that `@` now points at the current commit, which is no
longer the top:

```
$ jj log
○  pwnrkwpn my@email 08b3e414
│  make foo say hello, world
@  qlmqnzqo my@email d6b14a5d
│  add a foo file that says hello
◆  zzzzzzzz root() 00000000
```

If you look at the file `foo` now, you'll see it's back to the state of the
world when we made that commit, with only the one line added.

Similarly if you now run a command like `jj st` or `jj diff`, the output is as
if you were back at that old commit, showing that you were adding a new file.

## In-place changes

Once you've switched to an old commit, you edit it using the same tools you've
already used.

`jj desc` will still modify the current commit's description. This will cause a
rebase just as above.

If you edit a file on disk, it will update the file in the current commit. (If
you try this, be careful to only insert a line at the top, to not cause a
conflict -- we'll get to those later.) The next time you run any `jj` command
you will also see it rebase as it automatically resynchronizes the commit with
your working copy.

## Recipe: pausing work in progress

Suppose you're working on a new change and notice you made some typo in an
earlier commit. In Git, you might use `git stash` to save your current work to
go make that fix.

In jj, because your work is already saved in the current commit, there is no
need to explicitly stash. (If you want to put a note on the current commit
before moving away from it, you can use `jj desc`.)

Instead, you can immediately jump to that commit, make the fix, and jump back:

```
$ jj log
[... spot the name of the commit that had the typo ...]
$ jj edit that-commit
$ edit the file
$ jj edit commit-i-was-on-before
```

## Empty commits and jumping back

A second thing to notice about the `jj log` output is that the top empty commit
disappeared. This is because jj abandons empty commits when you move away from
them.

You can start a new commit at the top with `jj new p`, where `p` is the name of
the commit to start from. (You can now see that `jj new` both creates the new
commit and switches to it as `jj edit` does.) Alternatively, if you had made any
changes (or given a description) to your new commit, it would not have been
abandoned.

## Review

In this chapter, we learned:

- history is mutable
- editing history causes downstream changes to update
- `jj edit`: jump to a specific change and begin editing it
- there is no `git stash` equivalent because work is always already saved
- `jj new`: accepts an argument for which commit to start from
- moving away from empty commits causes them to be automatically abandoned

## Next step

We've modified history, but so far have been careful to avoid making a changes
that causes a conflict. What happens if we make an edit that conflicts?
[Read next about conflicts](conflicts.html).
