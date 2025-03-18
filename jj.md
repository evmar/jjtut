# Jujutsu tutorial

Jujutsu ("jj") is a version control system that combines some simple but
powerful ideas into a useful and sensible tool.

Underneath, your data is in a Git repository, which means using jj you can still
collaborate with all the existing Git tooling. Much like Git, jj has a unifying
underlying model of how your history is tracked that will feel familiar coming
from Git. Unlike Git, jj's model further subsumes many of the bells and whistles
of Git's clunky user interface, providing the same functionality of many Git
workflows (including the different modes of reset, stashes, rebase conflicts,
amending, the index, and so on) in a pleasant, simple, and coherent set of
commands.

> If this tutorial doesn't work out for you, try another! There is
> [one included in the Jujutsu docs](https://jj-vcs.github.io/jj/latest/tutorial/)
> and also
> [Steve's Jujutsu tutorial](https://steveklabnik.github.io/jujutsu-tutorial/).

## The central idea

Most of jj's functionality comes from one big idea: **the working copy is a
commit**. Understanding this concept will explain much of how jj works.

What does it mean? Version control systems store commits. In most, you can
"check out" some version of the code, edit some files, and run commands to push
those edits back in. In jj, you instead are always (conceptually) **directly
editing** a current commit; the commit and the files you are working with are
kept in sync. Saving a file means updating the current commit with the new
file's contents. The command most analagous to a "check out" operation is
literally `jj edit <commit>`.

<details>
<summary>How can that actually work?</summary>
It's an implementation detail, but in practice
whenever you run a jj command it updates its state to make the above stay
true.
</details>

This idea seems small at first but as we'll see it has far-reaching
consequences. For a sales pitch, though, here are some:

- You can write the commit description of a new change before you're done with
  it. (You can still do it at the end too, if you like.)
- If you're partway through writing a change and realize you need to do
  something else, you can jump to that other thing without explicitly saving
  your current work, because your current work is already held in a commit.
  (Subsumes `git stash`.)
- If you want to fix a typo or change the description of an old commit, you can
  jump directly to that commit and just make your edits, just as you would in a
  new change. (Subsumes `git commit --amend` and various `rebase -i` workflows.)
- It's inevitable that some operations, like rebases and merges, can create
  conflicts. In jj rebases and merges immediately succeed and this conflict
  state is just stored in the relevant commits. You resolve the conflicts by
  editing those commits like any other. (This subsumes the Git "currently in a
  rebase" state and associated `rebase --continue` etc. commands.)

jj has other nice features beyond this one, including an "operation log" that
lets you undo any previous operation you've done (similar to `git reflog` but
broader in scope). We'll get into those later.

## Setup

Install by following the
[installation instructions](https://jj-vcs.github.io/jj/latest/install-and-setup/).

Tell jj how to author your commits. (It will complain later otherwise, so it's
worth doing first.)

```
$ jj config set --user user.name "My Name"
$ jj config set --user user.emil "my@email.address"
```

## First commands

### Repo init

To learn, you'll need a repository to run commands on. You can create a new
repository or (maybe better for learning) pick an existing Git repository so you
have some real files to work with.

jj atop Git can work in two modes: the default, where the Git parts are kept
hidden within the `.jj` directory, or _colocated_, where your directory is both
a Git and jj repository. The latter is better for now because it means Git
commands will continue to work.

To create a new colocated repository, create a new directory (or go to an
existing directory already managed by Git) and:

```
$ jj git init --colocate
```

If you wanted to clone a Git repository, e.g. one from GitHub:

```
$ jj git clone --colocate <someurl>
```

This tutorial assumes you're starting with a fresh repository, but it will be
almost the same if you clone an existing repository instead.

### Initial status

On a fresh repository, jj creates an initial empty commit for you.

Run `jj` with no arguments, which is an alias for `jj log`. You'll see a history
of recent commits, with output like this:

<pre>
@  <b>q</b>lmqnzqo my@email 2025-03-18 10:30:32 6e075e98
│  (empty) (no description set)
◆  <b>z</b>zzzzzzz root() 00000000
</pre>

The line marked with `@` is the current commit, the empty one jj created. `@` is
also an alias to refer to the current commit in jj commands. (In Git, this is
analagous to maybe `HEAD` or the index.)

The first string of letters (<code><b>q</b>lmqnzqo</code> above) is the _change
ID_. This is jj's name for the commit, and how you refer to it in commands. In your
terminal, a prefix of the letters will be highlighted; in this tutorial, we mark
them with an underline.  That prefix is
sufficient to uniquely identify the commit, so you can write the change ID as just `q`
if you need to refer to this commit in a command.

> [!IMPORTANT]
> In this tutorial the commit was named `q`, but in your checkout it will likely
> have a different name. Substitute your commit's name in place of `q` in all the commands that follow.

On the far right is the Git ID for this commit. This may change as the commit
changes, and is mostly only useful for Git commands. You can distinguish jj
change IDs from Git IDs because they don't use the same letters; Git IDs are
hex, jj IDs are other letters.

In the next line, `(empty)` means this commit contains no file changes, and
`(no description set)` is where the description would be if we had one. But this
commit is brand new so it's both empty and descriptionless.

(The subsequent commit labelled `root()` in the above output just indicates the
beginning of history, since my repository is new. In a repository with more
history you'd instead see those commits.)

### Review

In this chapter, we learned about:

- `jj git init`, `jj git clone`: commands to set up a repo
- `jj log`: show commit history
- commits have IDs with short names

## Basic workflow

Next, let's create some changes.

### My first change

Create or edit a file:

```
$ echo hello > foo
```

Per the rules of jj, this edit to a file is integrated into your current commit.
There is no "I'm done, commit this" step; it's already done.

You can see this by running `jj status` (or `jj st`, for short), which
summarizes the current commit and mentions that it adds a file:

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

### My second change

To start a new change, use `jj new`.

<pre>
$ jj new
Working copy now at: <b>p</b>wnrkwpn ac9121f7 (empty) (no description set)
Parent commit      : <b>q</b>lmqnzqo b2fa5372 add the foo file
</pre>

We now are editing a new commit `p`, and can do the same commands as before.

> [!WARNING]
> As you get started with jj, forgetting to start new commits before editing
> files will likely be the most common mistake you make. It will feel unnatural
> at first but will get easier.

As a helpful alias, the command `jj commit` combines running `jj desc` followed
by `jj new`. It's especially useful if you haven't given your commit a
description already.

Add another line to foo and use `jj commit` to describe the current commit and start 
a new one:

<pre>
$ echo world &gt;&gt; foo
$ jj commit -m 'make foo say hello, world'
Working copy now at: <b>n</b>nlkypwz f58d0c2c (empty) (no description set)
Parent commit      : <b>p</b>wnrkwpn 3d263ba5 make foo say hello, world
</pre>

### Abandoning a change

Suppose next you accidentally clobbered your important work:

```
$ rm foo   # or edited it, etc.
```

There are two ways to undo this, with two different ways of thinking about it.

One is `jj restore`, which copies file contents from a different commit. With no
arguments it copies all files from the previous commit, emptying the commit of
any file changes, effectively clobbering any changes you've made. This preserves
the current commit's description and change ID.

The other option is `jj abandon`, which throws away the current commit. jj will
recreate a new empty commit in its place, with a new change ID.

If you did try making a change here, undo it using one of the above commands to prepare for
the next chapter.

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

## A taste of history

Run `jj log`. Our repository should have three commits. From the top: an empty
commit, the one that added "world", and the one that added the "hello" file
(which is named `q` in this tutorial).

We've used `jj diff` to see the diff of the current change. `diff` (and many
other jj commands) can also be told which change to show by using the `-r` flag.

To see the diff of a specific change:

```
$ jj diff -r q
[... same diff output as earlier ...]
```

The argument passed to `-r` is called a _revset_, and in jj it is a miniature
programming language for specifying commits. For tutorial purposes we can just
continue to pass the explicit names we've been using, but there is one alias
worth knowing about:

The alias `@` refers to the current change. Putting it together, you can now
understand that `jj diff` is a short way of saying `jj diff -r @`. (Coming from
Git, with its various modes for comparing files, the index, and trees, this was
such breath of fresh air!)

### Modifying history

Another command that accepts which revision to edit by `-r` is `desc`. You can
change the description of our first commit to be more descriptive:

```
$ jj desc -r q -m "add a foo file that says hello"
```

Like `diff`, you can now see that `jj desc` without flags edits `@`.

Users coming to jj from another version control system might raise an eyebrow
here: in jj, your commit history is generally freely editable. In case of making
mistakes, jj has powerful undo functionality that we'll get to later. When
working with Git, jj (by default) won't let you edit commits that (to a first
approximation) exist in remote branches already.

### Jumping around in history

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
made any changes (or given a description to) your new commit, it would not have
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

What happened? When you ran the jj command, jj implicitly integrated the file
edits you have made into the current commit, and then rebased any downstream
commits to integrate that change. Because your edit didn't conflict with
anything that came later, nothing else complained.

### Review

In this chapter, we learned:

- specifying revisions using the `-r` flag to `diff` and `desc`
- there exists a 'revset' language for specifying revisions
- history is mutable
- `jj edit`: jump to a specific change and begin editing it
- moving away from empty commits causes them to be automatically abandoned
- editing history causes downstream changes to update

## More history

So far we've created a linear sequence of commits. But we are free to start new
commits from anywhere. Suppose we want to try out greeting someone else. Let's
start a new change starting from our initial "hello" file and add a different
greeting:

<pre>
$ jj new q
Working copy now at: <b>u</b>kuqsrnl 247f710f (empty) (no description set)
Parent commit      : <b>q</b>lmqnzqo 763dc940 add a foo file that says hello
</pre>

Now run `jj log`:

<pre>
@  <b>v</b>ukpmpko my@email 2025-03-18 11:49:59 2d92d549
│  (empty) (no description set)
│ ○  <b>p</b>wnrkwpn my@email 2025-03-18 11:25:43 9759c1c7
├─╯  make foo say hello, world
○  <b>q</b>lmqnzqo my@email 2025-03-18 11:25:43 git_head() 763dc940
│  add a foo file that says hello
◆  <b>z</b>zzzzzzz root() 00000000
</pre>

The log now shows a branched tree. Unlike Git, jj does not have a notion of a
"current branch" that has a single head commit. Instead, any commits you create
are saved, and you can jump between them with `jj edit` as you desire.

### squash and the squash workflow

When you want to fix a typo in a commit, using `jj edit` to jump directly to it
and editing the file directly is the quickest way. But if you make a mistake
(cat on the keyboard?) there isn't a way to see your new edits as distinct from
the existing edits.

An alternative workflow is instead to create a new commit before editing, like
we just did in the previous section. After editing files, you run `jj diff` to see just the
new changes you've made. When you're happy with them, you can run `jj squash` to
"squash" the current commit into the previous one, merging their changes
together. (Git users might recognize the "squash" terminology from
`git rebase -i`.) After running `jj squash`, you're left pointed at an empty
commit; if you go elsewhere with `jj edit`, it will be automatically abandoned.

Because it's free to make new commits anywhere, one way to look at the squash
workflow is to *always* use `jj new` when editing history.  That is, if you want
to make a change to some commit X, you instead say `jj new X` to start a new commit
based on X, and then `jj squash` when you're done with your edits.  When you
move elsewhere (via `jj new` or `jj edit`) the empty temporary commit you've been using
on will be automatically cleaned up.

### squash and the Git index

Advanced Git users might be accustomed to building up a single
change by incrementally adding parts of it to the index, which in Git can be
thought of as a staging area for a pending change. jj models this using regular
commits:

1. Start a new commit (`jj new`), possibly give it an initial description
2. Start a new commit on top of that (`jj new` again).
3. Edit files again; whenever you're happy with them, `jj squash` them into the
   first commit.
4. Repeat step 3 as many times as necessary.

### Conflicts

We now have the understanding necessary to understand conflicts.

Modify your first
change (using either `jj edit` or `jj new`+`jj squash`) in such a way that the other
change will conflict with it.  For example, change "hello" to "goodbye".

```
$ jj edit q
$ my_editor foo  # introduce conflict, and save
```

Now run `jj` to see what happened:

<pre>
$ jj
×  <b>p</b>wnrkwpn evan.martin@gmail.com 2025-03-18 12:53:50 08b3e414 conflict
│  make foo say hello, world
@  <b>q</b>lmqnzqo evan.martin@gmail.com 2025-03-18 12:53:50 d6b14a5d
│  add a foo file that says hello
◆  zzzzzzzz root() 00000000
</pre>

The top commit is marked with an x, in red, and the `conflict` marker to indicate that
this commit is now conflicting: we made a change to commit `q` that means that `p` may
no longer make sense.

Unlike Git (or any other VCS I've used!), in jj conflicts are just allowed to happen,
without requiring any immediate action.  You can go on working on other things and the
conflict state is just stored in the repository.

When it's time to fix the conflict, you edit that commit as you would any other, either
with `jj edit` or `jj new`+`jj squash`.  jj will warn you about the conflicting state:

<pre>
$ jj edit p
Working copy now at: <b>p</b>wnrkwpn 08b3e414 (conflict) make foo say hello, world
Parent commit      : <b>q</b>lmqnzqo d6b14a5d add a foo file that says hello
Added 0 files, modified 1 files, removed 0 files
There are unresolved conflicts at these paths:
foo    2-sided conflict
</pre>

When you open the conflicting file in your editor, you will see conflict markers in
the file.  If you fix these conflicts (by editing them out), jj will notice and mark
the commit as no longer conflicting.

If you make an edit early in a tall stack of commits, it's possible each will conflict.
But it's also possible after fixing the first, jj will update the downstream ones and
discover the conflict has been resolved.  So to resolve a series of conflicts, just
start at the earliest commit and work your way forwards.

### Review

In this chapter, we learned:

- `jj new` can create commits anywhere, producing a non-linear history
- `jj squash`: squashes a commit into its parent
- the "squash" workflow involves creating new commits any time you want to make a change
- edits may produce conflicts, which show up in the history
- fixing conflicting files in a conflicting commit implicitly fixes the conflict state
  that that commit and possibly downstream ones

## Remaining outline/notes

> [!WARNING]
> Everything below this point is TODO

## Working with Git

## Workflow recipes

- jj split
- stash-like changes
- create commits before/after point

## where to fit in

- mention the "don't auto-add files" feature