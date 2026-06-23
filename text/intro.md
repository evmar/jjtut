# Introduction

jj stores your data in a Git repository. This means using jj you can still
collaborate using existing Git tooling. jj's underlying model of how your
history is tracked is more or less the same as Git, with branches and commits.

## The central idea

Most of jj's functionality comes from one big idea: **the working copy is a
commit**. Understanding this concept will explain much of how jj works.

What does it mean? Version control systems store commits. In most, you "check
out" some version of the code, edit some files, then run commands to push those
edits back in.

In jj, you instead are always (conceptually) **directly editing** a current
commit; the commit and the files you are working with are kept in sync. Editing
a file means implicitly updating the current commit with the new file's
contents.

<details>
<summary>How can that actually work?</summary>
It's an implementation detail, but in practice whenever you run a jj
command it updates its state to make the above stay true.
</details>

This idea might seem small but it has far-reaching consequences.

To start with, when you make a new change you start with a new empty commit and
then edit files, modifying the commit. This means you can also edit the
description on your new commit before finishing it. If you need to switch to
working on something else, your work is already saved. (You no longer need Git's
"stash" tool.)

For another example, if you need to fix a typo in an old commit, you can jump
directly to it and just make the edits. (This replaces many `git rebase -i`
workflows.) Changing the description of an old commit is the same command as
changing the description of the newest one.

This approach has further implications on how merge/rebase conflicts work. There
is no "rebase in progress" state or commands like `git rebase --continue`.
Instead you use the same core commands to resolve conflicts.

Finally, jj records a log of every operation on the respository to provide
pervasive undo. Because working copy state is always mirrored into the
repository, this covers all the above operations. (Like an enhanced
`git reflog`.)
