# Conflicts

We now have the understanding necessary to understand conflicts.

Modify your first change (using either `jj edit` or `jj new`+`jj squash`) in
such a way that the other change will conflict wi  th it. For example, change
"hello" to "goodbye".

```
$ jj edit q
$ my_editor foo  # introduce conflict, and save
```

Now run `jj` to see what happened:

```
$ jj
×  pwnrkwpn evan.martin@gmail.com 08b3e414 conflict
│  make foo say hello, world
@  qlmqnzqo evan.martin@gmail.com d6b14a5d
│  add a foo file that says hello
◆  zzzzzzzz root() 00000000
```

The top commit is marked with an x, in red, and the `conflict` marker to
indicate that this commit is now conflicting: we made a change to commit `q`
that means that `p` may no longer make sense.

In jj conflicts are just allowed to happen, without requiring any immediate
action. You can go on working on other things and the conflict state is stored
in the repository.

When it's time to fix the conflict, you edit that commit as you would any other,
either with `jj edit` or `jj new`+`jj squash`. jj will warn you about the
conflicting state:

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

If you make an edit early in a stack of commits, it's possible each will
conflict. But it's also possible after fixing the first, jj will update the
downstream ones and discover the conflict has been resolved. So to resolve a
series of conflicts, just start at the earliest conflict and work your way
forwards.

## Review

In this chapter, we learned:

- `jj new` can create commits anywhere, producing a non-linear history
- `jj squash`: squashes a commit into its parent
- the "squash" workflow create new commits any time you want to make a change
- edits may produce conflicts, which show up in the history
- fixing conflicting files in a conflicting commit implicitly fixes the conflict
  state that that commit and possibly downstream ones
