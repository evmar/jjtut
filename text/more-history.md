# More History

So far we've created a linear sequence of commits. But we are free to start new
commits from anywhere. Suppose we want to try out greeting someone else. Let's
start a new change starting from our initial "hello" file and add a different
greeting:

```
$ jj new q
Working copy now at: vukpmpko 2d92d549 (empty) (no description set)
Parent commit      : qlmqnzqo 763dc940 add a foo file that says hello
```

Now run `jj log`:

```
@  vukpmpko my@email 2025-03-18 11:49:59 2d92d549
│  (empty) (no description set)
│ ○  pwnrkwpn my@email 2025-03-18 11:25:43 9759c1c7
├─╯  make foo say hello, world
○  qlmqnzqo my@email 2025-03-18 11:25:43 git_head() 763dc940
│  add a foo file that says hello
◆  zzzzzzzz root() 00000000
```

The log now shows a branched tree. Unlike Git, jj does not have a notion of a
"current branch" that has a single head commit. Instead, any commits you create
are saved, and you can jump between them with `jj edit` as you desire.

## squash and the squash workflow

When you want to fix a typo in a commit, using `jj edit` to jump directly to it
and editing the file directly is the quickest way. But if you make a mistake
(cat on the keyboard?) there isn't a way to diff or undo your new edits as
distinct from the existing edits.

An alternative workflow is instead to create a new commit before editing, like
we just did in the previous section. After editing files, you run `jj diff` to
see just the new changes you've made.

When you're happy with the change, the `jj squash` command will "squash" the
current commit into the previous one, merging their changes together. (Git users
might recognize the "squash" terminology from `git rebase -i`.) After running
`jj squash`, you're left pointed at an empty commit; if you go elsewhere with
`jj edit`, it will be automatically abandoned.

Because it's free to make new commits anywhere, one way to look at the squash
workflow is to _always_ use `jj new` when editing history. That is, if you want
to make a change to some commit X, you instead say `jj new X` to start a new
commit based on X, and then `jj squash` when you're done with your edits. When
you move elsewhere (via `jj new` or `jj edit`) the empty temporary commit you've
been using on will be automatically cleaned up.

## squash and the Git index

Advanced Git users might be accustomed to building up a single change by
incrementally adding parts of it to the index, which in Git can be thought of as
a staging area for a pending change. jj models this using regular commits:

1. Start a new commit (`jj new`), possibly give it an initial description
2. Start a new commit on top of that (`jj new` again).
3. Edit files again; whenever you're happy with them, `jj squash` them into the
   first commit.
4. Repeat step 3 as many times as necessary.

## Conflicts

We now have the understanding necessary to understand conflicts.

Modify your first change (using either `jj edit` or `jj new`+`jj squash`) in
such a way that the other change will conflict with it. For example, change
"hello" to "goodbye".

```
$ jj edit q
$ my_editor foo  # introduce conflict, and save
```

Now run `jj` to see what happened:

```
$ jj
×  pwnrkwpn evan.martin@gmail.com 2025-03-18 12:53:50 08b3e414 conflict
│  make foo say hello, world
@  qlmqnzqo evan.martin@gmail.com 2025-03-18 12:53:50 d6b14a5d
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
