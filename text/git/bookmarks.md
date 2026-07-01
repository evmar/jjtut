# Bookmarks and branches

Our Git clone came with a jj bookmark named `main@origin`, as seen in the log in
the previous chapter and in the the output of this command:

```
$ jj bookmark list --all
main@origin: mzrrnrom broken link
```

## Bookmark basics

jj bookmarks are names that point at commits and can be used independently of
working with Git. They can be used in place of IDs in commands, and created,
modified, and destroyed with `jj bookmark` subcommands.

In my experience they aren't too useful on their own. If you modify a commit the
bookmark remains pointing at it, but unlike Git branches, they don't
automatically move around when you create new commits. They only serve to mark a
specific commit, like a Git tag.

## Remote bookmarks

Bookmarks with `@` in their name, like `main@origin`, are called remote
bookmarks. They represent our current knowledge of the state of the remote
repository, comparable to `refs/remotes/origin/main` in Git.

Attempts to change a remote bookmark directly with commands like
`jj bookmark set` will fail.

## Tracking bookmarks

So far, bookmarks resemble Git refs, but the concept of "tracking" in jj is
different. Create a `main` bookmark that is set up as tracking `main@origin`:

```
$ jj bookmark track main --remote origin
Started tracking 1 remote bookmarks.
```

What "tracked" means here is that jj will attempt to keep `main` in sync with
`main@origin`, in both pulls and pushes.

## Pushing tracked bookmarks

When you commit a new change locally and want to push it, you first need to
update `main` to point at the new commit:

```
$ jj bookmark move main -t @
Moved 1 bookmarks to tststvny main* | [...]
```

You'll now see the bookmark marked with a star, `main*`. What this indicates is
that the local bookmark is out of sync with the `main@origin` it tracks. If you
`jj git push` it will update the origin with the new main and bring things back
into sync.

```
$ jj git push
Changes to push to origin:
  Move forward bookmark main from 062b7af99b0d to ad9ce1693435
git: Enumerating objects: 1, done.
git: Counting objects: 100% (1/1), done.
git: Writing objects: 100% (1/1), 208 bytes | 208.00 KiB/s, done.
git: Total 1 (delta 0), reused 0 (delta 0), pack-reused 0 (from 0)
Warning: The working-copy commit in workspace 'default' became immutable, so a new commit has been created on top of it.
Working copy  (@) now at: ntxvxxzz (empty) (no description set)
Parent commit (@-)      : tststvny main | [...]
```

The need to manually update the bookmark before pushing is known to be a bit
clunky. jj is currently designed around the workflow where you don't often need
to move a bookmark around after you push it; see below.

## Pulling tracked bookmarks

If you `jj git fetch` to pull changes from the remote, jj will gather new
commits, update `main@origin`, and also update `main`.

What if you've already changed your local `main`?

```
$ jj log
@  trzlszwl my@email main?? main@git 491fa5d3
│  [my local change]
│ ◆  xysoxspt evan.martin@gmail.com main?? main@origin 4fda16de
├─╯  [a new change from the origin]
◆  xwlqyxry evan.martin@gmail.com 062b7af9
│  touchups
~
```

In the log, the bookmark shows as `main??`, pointing at two different commits
commits: your local commit and whatever commit it pulled from the origin. It is
now _conflicted_. You cannot use it until you fix things by manually pointing it
at a single commit using `jj bookmark move`.

## Creating bookmarks automatically

A common way for Git projects to integrate changes is via pull requests: you
push your changes as a new branch, and after code review it gets integrated into
the main branch.

The `jj git push -c` (`--change`) command generates a new bookmark name and
pushes it in one command. On GitHub this is an easy way to create a pull
request. If you need to make changes in response to review feedback, you can
modify your commits and `jj git push` again to update the pushed branch.

## Review

- `jj bookmark` is the entry point to bookmark commands
- local bookmarks can be set to track remote bookmarks
- `jj git push` pushes tracked bookmarks
- `jj git fetch` updates tracked bookmarks
- use `jj bookmark move` to move around local bookmarks
