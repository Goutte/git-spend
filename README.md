gitime : time tracker using git commit message commands
=======================================================

[![MIT](https://img.shields.io/github/license/Goutte/gitime?style=for-the-badge)](LICENSE)
[![Release](https://img.shields.io/github/v/release/Goutte/gitime?include_prereleases&style=for-the-badge)](https://github.com/Goutte/gitime/releases)
[![Build Status](https://img.shields.io/github/actions/workflow/status/Goutte/gitime/go.yml?style=for-the-badge)](https://github.com/Goutte/gitime/actions/workflows/go.yml)
[![Coverage](https://img.shields.io/codecov/c/github/Goutte/gitime?style=for-the-badge)](https://app.codecov.io/gh/Goutte/gitime/)
[![A+](https://img.shields.io/badge/go%20report-A+-brightgreen.svg?style=for-the-badge)](https://goreportcard.com/report/github.com/Goutte/gitime)
[![Code Quality](https://img.shields.io/codefactor/grade/github/Goutte/gitime?style=for-the-badge)](https://www.codefactor.io/repository/github/Goutte/gitime)
[![Download](https://img.shields.io/github/downloads/Goutte/gitime/total?style=for-the-badge)](https://github.com/Goutte/gitime/releases/latest/download/gitime)


Purpose
-------

Collect, addition and return all the `/spend` and `/spent` time-tracking directives in git commit messages.

> This only looks at the `git log` of the currently checked out branch.

**TLDR; [JUST DOWNLOAD](https://github.com/Goutte/gitime/releases/latest/download/gitime)**


### Example of a parsed commit

```
feat(crunch): implement a nice feature

Careful, it's still sharp.
/spend 10h30
```

Running:
```
$ gitime sum
```
would yield:
> `1 day 2 hours 30 minutes`

Of course, _gitime_ really shines when you have multiple commits with `/spend` commands that you want to sum.


### Specifications

We assume `8` hours per day, `5` days per week, `4` weeks per month. _(like Gitlab does)_

The **complete specification** can be found in [the rules](./gitime/gitime_test_data.yaml) of the test data,
and in excruciating detail in [the grammar](./gitime/grammar.go).


Usage
-----

Go into your git-versioned project's directory, and run:

```
cd <some git versioned project with commits using /spend directives>
gitime sum
```
> `> 2 days 1 hour 42 minutes`


### Format the output

You can also get the spent time in a specific unit :

```
gitime sum --minutes
gitime sum --hours
gitime sum --days
```
> These values will always be rounded to integers, for convenience,
> although _gitime_ does understand floating point numbers in `/spend` directives.


### Filter by commit authors

You can track the time of specified authors only, by `name` or `email` :

```
gitime sum --author Alice --author bob@email.net
```


### Exclude merge commits

You can also exclude merge commits :

```
gitime sum --exclude-merge
```

### Read from stdin

You can also parse messages from `stdin` instead of the git log:

```
git log > git.log
cat git.log | gitime sum
```


Download
--------

You can [download the binary](https://github.com/Goutte/gitime/releases/latest/download/gitime) straight from the [latest build in the releases](https://github.com/Goutte/gitime/releases).

You can also install via `go get` (hopefully) :

```
go get -u github.com/goutte/gitime
```

or `go install`:

```
go install github.com/goutte/gitime
```

> If that fails, you can install by cloning and running `make install`.


Develop
-------

```
git clone https://github.com/Goutte/gitime.git
cd gitime
go get
go run main.go
```


Build & Run & Install
---------------------

```
make
make sum
make install
make install-optimized
```

> `upx` is used to reduce the binary size in `make install-optimized`.


### Build for other platforms

You may use the `GOAS` and `GOARCH` environment variables to control the build targets:

```
GOOS=<target-OS> GOARCH=<target-architecture> go build -o build/gitime .
```

To list available targets (`os`/`arch`), you can run:

```
go tool dist list
```

> There's an example in the `Makefile`, with the recipe `make build-windows-amd64`.


Contribute
----------

Merge requests are welcome.  Make sure you recorded the time you `/spent` in your commit messages.  :)


### Ideas Stash

> You can pick and start any, or do something else entirely.

- `gitime sum --since <commit>`
- `gitime sum --since <tag>`
- `gitime sum --since <datetime>`
- Configure `DaysInOneWeek` and so forth using `ENV`, or config file
- `curl install.sh | bash`
- flatpak
- git extension
- docker
- i18n _(godspeed)_
- Right-To-Left _(help)_
