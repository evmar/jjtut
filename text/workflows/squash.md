# jj workflow: squash instead of edit

When you want to make a quick fix in a commit, using `jj edit` to jump directly
to it and editing the file directly is the quickest way. But any edits you make
will directly modify the commit; there isn't a quick way (outside of the op log)
to diff or undo your new edits as distinct from the existing commit you're
fixing up.

An alternative workflow is instead of using `jj edit the-commit` to jump to the
commit you want to modify, you instead `jj new the-commit` to create a new
commit starting from at that point. Any edits you make will be kept distinct
from the commit you're trying to fix up, so commands `jj diff` and `jj restore`
will only apply to the new edits you're making.

Once you're done, you can squash the fixups, similar to how the
[staging workflow](staging) goes.

After running `jj squash`, you're left pointed at an empty commit; if you go
elsewhere with it will be automatically abandoned.

## Visualized

If that text was hard to follow, here's an example of using it to repair a
conflict.

To start with, suppose your project is in some conflicted state:

```
$ jj log
×  vlzuomou my@email dd8df68c (conflict)
│  some other change
×  xnrzmulp my@email 3ca17944 (conflict)
│  make foo say hello, world
@  kzynprlx my@email 112ed9e4
│  add the foo file
◆  zzzzzzzz root() 00000000
```

I first make a new change at the first point I want to fix (the commit named
`x`) and repair it. I can run `jj diff` as I go to see what I changed.

```
$ jj new x
$ [fix the conflict]
$ jj log
@  yszmpkop my@email 451d4230
│  (no description set)
│ ×  vlzuomou my@email dd8df68c (conflict)
├─╯  some other change
×  xnrzmulp my@email 3ca17944 (conflict)
│  make foo say hello, world
○  kzynprlx my@email 112ed9e4
│  add the foo file
◆  zzzzzzzz root() 00000000
```

Once I'm ready, I can `jj squash` my fixup commit, and abandon this now empty
change `y`.

## The squash workflow

Depending on how you think about things, way to use jj is to _never_ use
`jj edit` to edit history, and instead _always_ use `jj new` and squash. That
is, if you want to make a change to some commit X, you instead say `jj new X` to
start a new commit based on X, and then `jj squash` when you're done with your
edits.

In the jj community this is called "the squash workflow". martinvonz, the jj
project leader works this way. In my experience I use `jj edit` for simple
fixups and the squash workflow when I think a fix might be more complex.

[Back to the workflow list](../).
