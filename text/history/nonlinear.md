# Nonlinear history

So far we've created a linear sequence of commits. But we are free to start new
commits from anywhere. Suppose we want to try out greeting someone else. Let's
start a new change starting from our initial "hello" file, by passing which
commit to start from to `jj new`.

```
$ jj new q
Working copy now at: vukpmpko 2d92d549 (empty) (no description set)
Parent commit      : qlmqnzqo 763dc940 add a foo file that says hello
```

Now run `jj log`:

```
@  vukpmpko my@email 2d92d549
│  (empty) (no description set)
│ ○  pwnrkwpn my@email 9759c1c7
├─╯  make foo say hello, world
○  qlmqnzqo my@email 763dc940
│  add a foo file that says hello
◆  zzzzzzzz root() 00000000
```

The log now shows a branched tree.

## Branches

Unlike Git, jj does not have a notion of a "current branch". Instead, any
commits you create are saved, and you jump between them with `jj edit` as you
desire. This also means that there are no names for branches outside of the
names you give to your commits.

When you work with external Git repositories that contain Git branches, they
still show up in jj as branched code. And jj has features for pushing and
pulling to Git branches, which we'll go into later.

In my experience, I expected to miss having having named branches, but actually
it ended up fine.

## Merges

To merge branches, pass multiple arguments to `jj new`:

```
$ jj new commit-a commit-b
```

Exactly like rebases, if there is a merge conflict, the new commit results in a
conflicted state. It's resolved by editing the files.

## Review

- `jj new` can start from anywhere, possibly creating a branch
- there is no explicit branch-like concept, just commits
- `jj new` accepts multiple arguments, creating a merge

## Next step

You now have seen most of the concepts you'll encounter. To apply them in
practice, read [part 3: workflows](../../workflows).
