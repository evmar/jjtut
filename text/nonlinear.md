# Nonlinear history

So far we've created a linear sequence of commits. But we are free to start new
commits from anywhere. Suppose we want to try out greeting someone else. Let's
start a new change starting from our initial "hello" file and add a different
greeting, by passing which commit to start from to `jj new`.

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
commits you create are saved, and you can jump between them with `jj edit` as
you desire.

When you work with external Git repositories that contain Git branches, they
show up in jj
