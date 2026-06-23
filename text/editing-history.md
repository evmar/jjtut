# Editing history

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

If you look at the file `foo` now, you'll see it's back to the state of the
world when we made that commit, with only the one line added.

Similarly if you now run a command like `jj st` or `jj diff`, the output is as
if you were back at that old commit, showing that you are adding a new file. And
running `jj
desc` will modify the current commit's description.

If you run `jj log` now, you will notice two things.

```
$ jj log
○  pwnrkwpn my@email 08b3e414
│  make foo say hello, world
@  qlmqnzqo my@email d6b14a5d
│  add a foo file that says hello
◆  zzzzzzzz root() 00000000
```

First, note that `@` points at the current commit, which is no longer the top.

Second, notice that the top empty commit disappeared! This is because jj
abandons empty commits when you move away from them.

You can start a new commit at the top with `jj new p`, where `p` is the name of
the commit to start from. Alternatively, if you had made any changes (or given a
description) to your new commit, it would not have been abandoned.

## Modifying files

Let's make a change to the file `foo` in our initial commit -- a change that
**doesn't** introduce a conflict, which we'll get to later. Open up `foo` in
your editor and insert a line at the top.

```
$ jj edit q       # q, the first commit
$ my_editor foo   # edit the file, insert a line at the top
```

Now run `jj status`. Similarly to after editing the description, jj rebases the
downstream commits on top of this change.

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

- history is mutable
- editing history causes downstream changes to update
- `jj edit`: jump to a specific change and begin editing it
- `jj new`: accepts an argument for which commit to start from
- moving away from empty commits causes them to be automatically abandoned
- `jj restore`: copy file contents from a change
- `jj abandon`: abandon (delete) a change

## Next step

[more](more-history.html)
