# History

Run `jj log` to see the repository history.

```
@  umnvtwlo my@email 2025-08-17 09:00:50 20de4517
│  (empty) (no description set)
○  pwnrkwpn my@email 2025-03-18 12:53:50 08b3e414
│  make foo say hello, world
○  qlmqnzqo my@email 2025-03-18 12:53:50 git_head() d6b14a5d
│  add a foo file that says hello
◆  zzzzzzzz root() 00000000
```

From the top:

1. umnvtwlo is an empty commit
1. pwnrkwpn and qlmqnzqo were our two edits
1. zzzzzzzz is a special "root" commit that starts the repository and is is
   always empty

In a terminal, these commits will have some prefix (often the initial letter)
highlighted or in bold. This indicates the unique prefix of the commit that can
be used to refer to it in commands.

## Diffs and revsets

We've used `jj diff` to see the diff of the current change. `diff` (and many
other jj commands) can also be told which change to show by using the `-r` flag.

Here, I use `q` as the unique prefix of the commit above:

```
$ jj diff -r q
[... same diff output as earlier ...]
```

The argument passed to `-r` is called a _revset_, and in jj it is a miniature
programming language for specifying commits. For tutorial purposes we can just
continue to pass the explicit names we've been using, but there is one alias
worth knowing about:

The alias `@` refers to the current change. Putting it together, you can now
understand that `jj diff` is a short way of saying `jj diff -r @`.

## Modifying history

Another command that accepts which revision to edit by `-r` is `desc`. You can
change the description of our first commit to be more descriptive:

```
$ jj desc -r q -m "add a foo file that says hello"
```

Like `diff`, you can now see that `jj desc` without flags edits `@`.

Users coming to jj from another version control system might raise an eyebrow
here: in jj, your commit history is generally freely editable. In case of making
mistakes, jj has powerful undo functionality that we'll get to later. (When
working with Git, jj has additional functionailty related to not accidentally
rewriting commits from remote branches; we'll get to it later.)

## Jumping around in history

Alternatively, you can change which commit you're currently editing with
`jj edit`. Jump back to the first commit we created:

```
$ jj edit q
```

If you look at the file `foo` now, you'll see it's back to the state of the
world when we made that commit, with only the one line added. Similarly if you
now run a command like `jj st` or `jj diff`, the output is as if you were back
at that first commit, showing that you are adding a new file. And `jj desc` will
edit the description of this first commit.

If you run `jj log` now, you might notice that our topmost commit disappeared!
This is because the commit was empty, and `jj` abandons empty commits when you
move away from them. You can create a new one in its place with `jj new p`,
where `p` is the name of the commit to start from. Alternatively, if you had
made any changes (or given a description) to your new commit, it would not have
been abandoned.

### Editing history

Let's make a change to the `foo` in our initial commit -- a change that
**doesn't** introduce a conflict, which we'll get to later. Open up `foo` in
your editor and insert a line at the top.

```
$ jj edit q       # q, the first commit
$ my_editor foo   # edit the file, insert a line at the top
```

Save the edit, and run a command like `jj st`. You'll notice a new line in the
output:

<pre>
$ jj st
<b>Rebased 1 descendant commits onto updated working copy</b>
Working copy changes:
A foo
Working copy : <b>q</b>lmqnzqo 763dc940 add a foo file that says hello
Parent commit: <b>z</b>zzzzzzz 00000000 (empty) (no description set)
</pre>

Note the line that says "Rebased 1 descendant commits onto updated working
copy". What happened? When you ran the jj command, jj implicitly integrated the
file edits you have made into the current commit, and then rebased any
downstream commits to integrate that change. Because your edit didn't conflict
with anything that came later, everything went fine.

## Review

In this chapter, we learned:

- specify revisions using the `-r` flag to `diff` and `desc`
- there exists a 'revset' language for specifying revisions
- history is mutable
- `jj edit`: jump to a specific change and begin editing it
- moving away from empty commits causes them to be automatically abandoned
- editing history causes downstream changes to update
