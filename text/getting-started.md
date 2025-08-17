# Getting Started

## jj setup

Install by following the
[installation instructions](https://jj-vcs.github.io/jj/latest/install-and-setup/).

Tell jj how to author your commits. (It will complain when you run commands
otherwise, so it's worth doing now.)

```
$ jj config set --user user.name "My Name"
$ jj config set --user user.emil "my@email.address"
```

If you run `jj` with no arguments, it prints a help message that suggests a
default configuration that you should adopt.

```
$ jj
Hint: Use `jj -h` for a list of available commands.
Run `jj config set --user ui.default-command log` to disable this message.
```

So run that too:

```
$ jj config set --user ui.default-command log
```

## Repo setup

To learn, you'll need a repository to run commands on. You can create a new
repository or pick an existing Git repository so you have some real files to
work with.

jj atop Git can work in two modes: the default, where the Git parts are kept
hidden within the `.jj` directory, or _colocated_, where your directory is both
a Git and jj repository. The latter is better for now because it means Git
commands will continue to work.

To create a new colocated repository, create a new directory (or go to an
existing directory already managed by Git) and:

```
$ jj git init --colocate
```

If you wanted to clone a Git repository, e.g. one from GitHub:

```
$ jj git clone --colocate <someurl>
```

This tutorial assumes you're starting with a fresh repository, but it will be
almost the same if you clone an existing repository instead.
