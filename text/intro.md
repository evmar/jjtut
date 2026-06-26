# Introduction

jj stores your data in a Git repository. jj's underlying model of how your
history is tracked is more or less the same as Git, with branches and commits,
and when using it you still interoperate with Git hosting.

jj differs from Git in its improved user interface and the way it represents
history.

## The central idea

Most of jj's functionality comes from one big idea: **the working copy is a
commit**. Understanding this concept will explain much of how jj works.

What does it mean? Version control systems store commits. In most, you "check
out" some version of the code, edit some files, then run commands to push those
edits back in.

In jj, you instead are always (conceptually) _directly editing the current
commit_; the commit and the files you are working with are kept in sync. Editing
a file means implicitly updating the current commit with the new file's
contents.

(An aside: how can that actually work? It's an implementation detail, but in
practice whenever you run a jj command it updates its state to make the above
stay true.)

## Consequences

This idea might seem small but it has far-reaching consequences.

To start with, when you make a new change you start with a new empty commit and
then edit files, modifying the commit. This means you can also edit the
description on your new commit before finishing it. If you need to switch to
working on something else, your work is already saved. (This replaces Git's
"stash" tool.)

For another example, if you need to fix a typo in an old commit, you can switch
directly to it and just make the edits. (This replaces many `git rebase -i`
workflows.) Changing the description or contents of an old commit is the same as
changing the description of the newest one. (This replaces `commit --amend`.)

This approach has further implications on how merge/rebase conflicts work. There
is no "rebase in progress" state or commands like `git rebase --continue`.
Instead rebases create conflicting commits, and you resolve conflicts by
directly editing those commits using the same core commands as elsewhere.

Finally, jj records a log of every operation on the respository to provide
pervasive undo. Because working copy state is always mirrored into the
repository, this covers all the above operations. (Like an enhanced
`git reflog`.)

## Next step

[Start by setting things up](setup.html).
