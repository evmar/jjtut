# Conflicts

We now have the tools necessary to understand conflicts.

Let's modify your first change in such a way that a later change will conflict
with it. For example, change "hello" to "goodbye".

```
$ jj edit q
$ my_editor foo  # introduce conflict, and save
```

Now run `jj` to see what happened:

```
$ jj
×  pwnrkwpn my@email 08b3e414 (conflict)
│  make foo say hello, world
@  qlmqnzqo my@email d6b14a5d
│  add a foo file that says hello
◆  zzzzzzzz root() 00000000
```

The top commit is marked with an x, in red, and the `(conflict)` marker to
indicate that this commit is now conflicting: we made a change to commit `q`
that means that the change `p` no longer makes sense.

In jj, conflicts are just allowed to happen, without requiring any immediate
action. You can go on working on other things and the conflict state is stored
in the repository.

## Fixing conflicts

When it's time to fix the conflict, you edit that commit as you would any other.
Switch to it using `jj edit`. jj will warn you about the conflicting state:

```
$ jj edit p
Working copy now at: pwnrkwpn 08b3e414 (conflict) make foo say hello, world
Parent commit      : qlmqnzqo d6b14a5d add a foo file that says hello
Added 0 files, modified 1 files, removed 0 files
There are unresolved conflicts at these paths:
foo    2-sided conflict
```

When you open the conflicting file in your editor, you will see conflict markers
in the file. If you fix these conflicts (by editing them out), jj will notice
and mark the commit as no longer conflicting.

If you make an edit early in a stack of commits, it's possible you then have a
stack of conflicts. But it's also possible after fixing the first, jj will
update the downstream ones and discover the conflict has been resolved. So to
resolve a series of conflicts, just start at the earliest conflict and work your
way forwards.

## Review

In this chapter, we learned:

- edits may produce conflicts, which show up in the history
- fixing conflicting files in a conflicting commit implicitly fixes the conflict
  state that that commit and possibly downstream ones

## Next step

So far we've worked on a single linear history. Next learn about
[nonlinear history and branches](nonlinear.html).
