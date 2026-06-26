# Workflow: pausing work in progress

Suppose you're working on a new change and notice you made some typo in an
earlier commit. In Git, you might use `git stash` to save your current work to
go make that fix.

In jj, because your work is already saved in the current commit, there is no
need to explicitly stash. (If you want to put a note on the current commit
before moving away from it, you can use `jj desc`.)

[Back to the workflow list](./).
