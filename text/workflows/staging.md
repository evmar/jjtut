# jj workflow: staging and the Git index

Advanced Git users might be accustomed to building up a single change by
incrementally adding parts of it to the index, which in Git can be thought of as
a staging area for a pending change.

jj models the workflow of building up a change by using regular commits.

1. Start a new staging commit on top of your current one (run `jj new`).
1. Edit files as needed; whenever you're happy with them, `jj squash` them into
   the previous commit, or use `jj restore` to undo them.
1. Repeat that step as many times as necessary.
1. Use `jj abandon`, if needed, to abandon the staging commit.

What if you've already made the commit and want to split it up? Use the
[insert](../insert) workflow to insert an empty commit before the current one, then
follow the same steps.

[Back to the workflow list](../).
