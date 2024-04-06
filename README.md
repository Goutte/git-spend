git-spend : time tracker using git commit message commands
=======================================================

[![MIT](https://img.shields.io/github/license/Goutte/git-spend?style=for-the-badge)](LICENSE)
[![Release](https://img.shields.io/github/v/release/Goutte/git-spend?include_prereleases&style=for-the-badge)](https://github.com/Goutte/git-spend/releases)
[![Build Status](https://img.shields.io/github/actions/workflow/status/Goutte/git-spend/go.yml?style=for-the-badge)](https://github.com/Goutte/git-spend/actions/workflows/go.yml)
[![Coverage](https://img.shields.io/codecov/c/github/Goutte/git-spend?style=for-the-badge)](https://app.codecov.io/gh/Goutte/git-spend/)
[![A+](https://img.shields.io/badge/go%20report-A+-brightgreen.svg?style=for-the-badge)](https://goreportcard.com/report/github.com/Goutte/git-spend)
[![Code Quality](https://img.shields.io/codefactor/grade/github/Goutte/git-spend?style=for-the-badge)](https://www.codefactor.io/repository/github/Goutte/git-spend)
[![Download](https://img.shields.io/github/downloads/Goutte/git-spend/total?style=for-the-badge)](https://github.com/Goutte/git-spend/releases/latest/download/git-spend)


Purpose
-------

Collect, addition and return all the `/spend` and `/spent` time-tracking directives in git commit messages.

> This looks at the `git log` of the currently checked out branch of the working directory,
> and therefore requires `git` to be installed on your system.

**TLDR; JUST [DOWNLOAD LINUX/MAC] — [DOWNLOAD WINDOWS]**

[DOWNLOAD LINUX/MAC]: https://github.com/Goutte/git-spend/releases/latest/download/git-spend
[DOWNLOAD WINDOWS]: https://github.com/Goutte/git-spend/releases/latest/download/git-spend.exe


### By Example

Say you are in the directory of a project with one commit like so :

```
feat(crunch): implement a nice feature

Careful, it's still sharp.
/spend 10h30
```

Running:
```
$ git spend sum
```
would yield:
> `1 day 2 hours 30 minutes`

Of course, _git-spend_ really shines when you have multiple commits with `/spend` commands that you want to tally and sum.

> 💡 You can use `git-spend sum` or `git spend sum`, they are equivalent. 

### Specifications

We assume `8` hours per day, `5` days per week, `4` weeks per month. _(like Gitlab does)_
These can be configured at runtime if needed, using environment variables.

The **complete specification** can be found in [the rules](./gitime/gitime_test_data.yaml) of the test data,
and in excruciating detail in [the grammar](./gitime/grammar.go).

The [acceptance testing suite](./test/features.bats) also holds many usage examples.


Usage
-----

Go into your git-versioned project's directory:

```
cd <some git versioned project with commits using /spend directives>
```

and run:

```
git spend sum
```
> `2 days 1 hour 42 minutes`

Or run `git-spend` from anywhere, but specify the `--target` directory (which defaults to `.`):

```
git spend sum --target <some git versioned project dir>
```
> `2 days 1 hour 42 minutes`


> ⛑ Use `git spend sum --help` or `man git-spend-sum` to see all the options.
> Meanwhile, let's look at some available options, below.


### Format the output

You can get the spent time in a specific unit :

```
git spend sum --minutes
git spend sum --hours
git spend sum --days
```
> These values will always be rounded to integers, for convenience,
> although _git-spend_ does understand floating point numbers in `/spend` directives.


### Filter by commit authors

You can track the time of specified authors only, by `name` or `email` :

```
git spend sum --author Alice --author bob@email.net
```


### Exclude merge commits

You can also exclude merge commits :

```
git spend sum --no-merges
```


### Restrict to a range of commits

You can restrict to a range of commits, using a commit hash, a tag, or even `HEAD~N`.

```
git spend sum --since <ref> --until <ref>
```

For example, to get the time spent on the last `15` commits :

```
git spend sum --since HEAD~15
```

Or the time spent on a tag since previous tag :

```
git spend sum --since 0.1.0 --until 0.1.1
```

You can also use _dates_ and _datetimes_, but remember to quote them if you specify the time:

```
git spend sum --since 2023-03-21
git spend sum --since "2023-03-21 13:37:00"
```

> 📅 Other supported time formats: [`RFC3339`], [`RFC822`], [`RFC850`].
> If you need a specific timezone, try setting the `TZ` environment variable:
> `TZ="Europe/Paris" git-spend sum --since 2023-03-21`

[`RFC3339`]: https://www.rfc-editor.org/rfc/rfc3339
[`RFC822`]: https://www.w3.org/Protocols/rfc822/
[`RFC850`]: https://www.rfc-editor.org/rfc/rfc850


Download
--------

### Direct download

You can [⮋ download the binary](https://github.com/Goutte/git-spend/releases/latest/download/git-spend) straight from the [latest build in the releases](https://github.com/Goutte/git-spend/releases),
and move it anywhere in your `$PATH`, such as `/usr/local/bin/git-spend` for example.

> ⚠ Remember to enable the execution bit with `chmod u+x ./git-spend`, for example.

There is an _experimental_ install script that does exactly this, plus `man` pages generation:

    curl https://raw.githubusercontent.com/Goutte/git-spend/main/install.sh | sh

> 🐧 This script only works for `linux/amd64`, for now.   _Stigmergy?_

### Via `go get`

You can also install via `go get` (hopefully) :

```
go get -u github.com/goutte/git-spend
```

or `go install`:

```
go install github.com/goutte/git-spend
```

> If that fails, you can install by cloning and running `make install`.


Advanced Usage
--------------

### Read from standard input

You can also directly parse messages from `stdin`
instead of attempting to read the git log:

```
git log > git.log
cat git.log | git-spend sum --stdin
```

> `git spend` ignores standard input otherwise.


### Configure the time modulo

If you live somewhere where work hours per week are limited (to 35 for example)
in order to mitigate labor oppression tactics from monopoly hoarders,
you can use environment variables to control how time is "rolled over" between units :

```
GIT_SPEND_HOURS_PER_DAY=7 git-spend sum
```

Here are the available environment variables :

- `GIT_SPEND_MINUTES_PER_HOUR` (default: `60`)
- `GIT_SPEND_HOURS_PER_DAY` (default: `8`)
- `GIT_SPEND_DAYS_PER_WEEK` (default: `5`)
- `GIT_SPEND_WEEKS_PER_MONTH` (default: `4`)


### Install the man pages

If you installed via direct download, you might want to install the `man` pages:

```
sudo git spend man --install
```

> `git help spend` will then work as expected.



Develop
-------

First, you'll need to [install Golang](https://go.dev/dl/).

```
git clone https://github.com/Goutte/git-spend.git
cd git-spend
go get
go run main.go
```


Build & Run & Install
---------------------

The binaries in the releases are built by our [Continuous Integration](./.github/workflows/release.yml).

Nevertheless, if you want to build your own `git-spend`, you can clone this project and run: 

```
make
make install
```

> [`upx`] is used to reduce the binary size in `make install-release`.

[`upx`]: https://upx.github.io/

---

You can compare the checksums, and they should be the same unless microsoft is being naughty.


### Build for other platforms

You may use the `GOOS` and `GOARCH` environment variables to control the build targets:

```
GOOS=<target-OS> GOARCH=<target-architecture> go build -o build/git-spend .
```

To list available targets (`os`/`arch`), you can run:

```
go tool dist list
```

> There's an example in the [`Makefile`], with the recipe `make build-windows-amd64`.

[`Makefile`]: ./Makefile

Contribute
----------

Merge requests are welcome.  Make sure you record the time you `/spend` in your commit messages.  :)

### Translations

Translations files are in `locale/*.toml`.
To add another language, add a new file, some sugar, some water, and … _voilà !_

### Ideas Stash

> You can pick and start any, or do something else entirely.
> If you don't like any of these, please voice your concerns as early as possible.

- [x] `curl install.sh | sudo sh` _(ongoing, wider support needed)_
- [x] i18n _(ongoing, [cobra forked](https://github.com/Goutte/cobra/tree/feat-i18n))_
- [ ] `git-spend sum --format <custom format>`
- [ ] `git-spend sum --short` → `1d3h27m`
- [ ] `git-spend chrono start` → start an internal chronometer
- [ ] `git-spend chrono add` → add time to the chronometer
- [ ] Rewriting of `/spend chrono [± <duration>]` by commit hook
- [ ] Rewriting of `/spend <ref> [± <duration>]` by commit hook 🌟
- [ ] `git spend hook --install` to install git hooks for rewriting
- [ ] `git spend hook --remove` to remove installed git hooks
- [ ] `git-spend amend <duration>` → amend previous commit with `/spend <duration>`
- [ ] `git-spend amend --add <duration>` → same but adds
- [ ] `git-spend amend --subtract <duration>` → same but subtracts  (alias: `--sub` ?)
- [ ] docker _(`docker run git-spend` -- awkward? ; would need a volume)_
- [ ] flatpak perhaps (road blocked, see [`packaging/`](./packaging))
- [ ] Right-To-Left _(ساعد)_
