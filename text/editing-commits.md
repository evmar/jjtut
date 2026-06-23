# Editing commits

Another command that accepts a revision by `-r` is `desc`. You can change the
description of our first commit to be more descriptive:

```
$ jj desc -r q -m "add a foo file that says hello"
```

Like with `jj diff`, you can now see that running `jj desc` without flags edits
`@`, but both can be pointed at any commit.

Users coming to jj from another version control system might be surprised here
by how making new changes and modifying history are the same commands just
pointed at different places. In jj, your commit history is generally freely
editable.

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
if you were back at that old commit, showing that you are adding a new file.

## In-place changes

Once you've switched to an old commit, you edit it using the same tools as before.

`jj desc` will still modify the current commit's description. This will
cause a rebase just as above.

And if you edit a file, it will update the file in the current commit. (If you try
this, be careful to only insert a line at the top, to not cause a conflict --
we'll get to those later.) The next time you run a `jj` command you will also
see it rebase.

## Empty commits and jumping back

A second thing to notice about the `jj log` output is that the top empty commit
disappeared! This is because jj abandons empty commits when you move away from
them.

You can start a new commit at the top with `jj new p`, where `p` is the name of
the commit to start from. Alternatively, if you had made any changes (or given a
description) to your new commit, it would not have been abandoned.

## Review

- history is mutable
- editing history causes downstream changes to update
- `jj edit`: jump to a specific change and begin editing it
- `jj new`: accepts an argument for which commit to start from
- moving away from empty commits causes them to be automatically abandoned

## Next step

We've modified commits in place, but in jj it's also just as easy to modify
the sequence of commits.  XXX [editing history](editing-history.html).
