# The Sales Pitch

Is it really worth your time to learn a new version control system? I think the
answer is yes; certainly it was for me and many other fans of this tool! In this
document I try to give the high-level sales pitch for why jj is worth your time.

To start with, you have little to lose. jj stores its data in Git and can be
thought of as a different user interface for Git, so you still can use Git tools
like GitHub as before. If you try jj and don't like it, abandoning it is about
as simple as switching back to Git commands in your same checkout.

jj improves upon Git by managing to achieve two opposing goals at the same time:
it has a _simpler_ mental model and command set, but simultaneously provides
_more powerful_ functionality.

Git started with a beautiful conceptual core of how commits build atop one
another -- also used by jj -- but from there Git made a mess of its user
interface including the index, stashes, different modes of reset, in-progress
rebases, and inconsistent undo. In contrast, with jj you can still do the same
kinds of workflows as what those Git tools do, but with just a few orthogonal
commands that compose well together. For one example of the consequences of
this, jj has pervasive "undo" support, even for operations outside of Git's
reflog.

As a Git expert I was already comfortable using tools like `git rebase -i` to
move fixups around. What I found adopting jj is that while it was always
possible to make Git do what I wanted it to do, with jj I can do the same things
more easily and quickly. Common tasks like "while writing this change I realized
there's a change I ought to do first" are possible with Git but trivial with jj.
As a developer I've found it's made me more confident in how I experiment,
because it's so easy to move code around.

Finally, jj also exposes some workflows that will be new even to Git experts.
For example, the "megamerge" is a way of working on multiple branches at the
same time, letting you test how multiple in-progress changes work together.
