## First commands

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
ID_. This is jj's name for the commit, and how you refer to it in commands. In
your terminal, a prefix of the letters will be highlighted; in this tutorial, we
mark them with an underline. That prefix is sufficient to uniquely identify the
commit, so you can write the change ID as just `q` if you need to refer to this
commit in a command.

> [!IMPORTANT]
> In this tutorial the commit was named `q`, but in your checkout it will likely
> have a different name. Substitute your commit's name in place of `q` in all
> the commands that follow.

On the far right is the Git ID for this commit. This may change as the commit
changes, and is mostly only useful for Git commands. You can distinguish jj
change IDs from Git IDs because they don't use the same letters; Git IDs are
hex, while jj IDs don't use those digits.

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
