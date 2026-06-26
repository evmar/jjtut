# The Sales Pitch

Is it really worth your time to learn a new version control system? I think the
answer is yes. Certainly it was for me and many other fans of this tool! I adopt
new tools rarely and evangelize them even less often, but jj has been my
favorite of the last few years.

In this chapter I try to give the high-level sales pitch for why jj is worth
your time.

## Interoperation and risk

To start with, it's easy to try out and you have little to lose.

jj stores its data in Git and can be thought of as a different user interface
for Git, so you still can use Git tools like GitHub as before. If you try jj and
don't like it, abandoning it is almost as simple as `rm -rf .jj` and switching
back to Git commands in your same checkout.

## Simple but powerful

jj improves upon Git by achieving two opposing goals at the same time: it has a
_simpler_ mental model and command set, but simultaneously provides _more
powerful_ functionality. jj simplifies Git not by hiding complexity, but by
using a better conceptual model.

Git has a beautiful conceptual core of how commits build atop one another --
also used by jj -- but from there Git made a mess of its user interface
including the index, stashes, different modes of reset, in-progress rebases,
inconsistent undo, and
[zillions of flags for everything](https://git-man-page-generator.lokaltog.net/).

In contrast, with jj you can still do the same kinds of workflows as what those
Git tools do, but with just a few orthogonal commands that compose well
together.

## My favorites

Here are a few of my favorite features of jj:

- Making a fix-up to some earlier change is an easy and trivial workflow.

- If you're halfway through some change and realize there's something else you
  needed to do first, it's easy to start a new change before the current one and
  then pick up where you left off.

- You can write the description of a change before you start working on it. (You
  can also describe it after, like in traditional version control.)

- Rebases (which are more common in jj, for reasons that will become apparent)
  in jj always succeed. There is no "rebase in progress" state like in Git. If a
  rebase causes a conflict, that is just recorded in the branch history and
  resolving it is similar to making fixups.

- jj has pervasive "undo" support, including operations outside of Git's reflog,
  making all of the above kinds of things easy to try out and undo.

The best part about these -- hard to convey in a feature checklist -- is that
they are not just separate subcommands, but rather different ways the same core
features compose.

## Coming from Git

As a Git expert I was already comfortable using tools like `git stash` and
`git rebase -i` to move code around. What I found after adopting jj is that
while I always _could_ make Git do what I wanted it to do, with jj I can do the
same things more easily and quickly.

Common tasks "while writing this change I realized there's a change I ought to
do first" are possible with Git but trivial with jj. As a developer I've found
it's made me more confident in how I experiment, because it's so easy to move
code around.

## Beyond Git

Finally, jj also exposes some workflows that will be new even to Git experts.
For example, the "megamerge" is a way of working on multiple branches at the
same time, letting you test how multiple in-progress changes work together.

## Next step

[Read about the big idea](../basics).
