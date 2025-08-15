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
It's an implementation detail, but in practice whenever you run a jj
command it updates its state to make the above stay true.
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
- When an operation like a rebase or a merge causes conflicts, jj just stores
  this conflicted state in the relevant commits. You resolve the conflicts by
  editing those commits like any other. (This subsumes the Git "currently in a
  rebase" state and associated `rebase --continue` etc. commands.)

jj has other nice features beyond this one, including an "operation log" that
lets you undo any previous operation you've done (similar to `git reflog` but
broader in scope). We'll get into those later.
