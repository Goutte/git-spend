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

> This looks at the `git log` of the currently checked out branch of the working directory.

**TLDR; JUST [DOWNLOAD LINUX/MAC](https://github.com/Goutte/gitime/releases/latest/download/gitime) — [DOWNLOAD WINDOWS](https://github.com/Goutte/gitime/releases/latest/download/gitime.exe)**


### Example of a parsed commit

Say you are in the directory of a project with one commit like so :

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
gitime sum
```
> `2 days 1 hour 42 minutes`

Or run `gitime` from anywhere, but specify the `--target` directory (which defaults to `.`):

```
gitime sum --target <some git versioned project dir>
```
> `2 days 1 hour 42 minutes`


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
gitime sum --no-merges
```


### Restrict to a range of commits

You can restrict to a range of commits, using a commit hash, a tag, or even `HEAD~N`.

```
gitime sum --since <ref> --until <ref>
```

For example, to get the time spent on the last `15` commits :

```
gitime sum --since HEAD~15
```

Or the time spent on a tag since previous tag :

```
gitime sum --since 0.1.0 --until 0.1.1
```

You can also use _dates_ and _datetimes_, but remember to quote them if you specify the time:

```
gitime sum --since 2023-03-21
gitime sum --since "2023-03-21 13:37:00"
```

> Other supported time formats: [`RFC3339`], `RFC822`, `RFC850`.
> If you need a specific timezone, try setting the `TZ` environment variable:
> `TZ="Europe/Paris" gitime sum --since 2023-03-21`

[`RFC3339`]: https://www.rfc-editor.org/rfc/rfc3339


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


Advanced Usage
--------------

### Read from stdin

You can also parse messages from `stdin` instead of the git log:

```
git log > git.log
cat git.log | gitime sum
```


### Configure the time modulo

If you live somewhere where work hours per week are limited (to 35 for example)
in order to mitigate labor oppression tactics from monopoly hoarders,
you can use environment variables to control how time is "rolled over" between units :

```
GITIME_HOURS_IN_ONE_DAY=7 gitime sum
```

Here are the available environment variables :

- `GITIME_MINUTES_IN_ONE_HOUR` (default: `60`)
- `GITIME_HOURS_IN_ONE_DAY` (default: `8`)
- `GITIME_DAYS_IN_ONE_WEEK` (default: `5`)
- `GITIME_WEEKS_IN_ONE_MONTH` (default: `4`)




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
```

> `upx` is used to reduce the binary size in `make install-release`.


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

Merge requests are welcome.  Make sure you record the time you `/spend` in your commit messages.  :)


### Ideas Stash

> You can pick and start any, or do something else entirely.

- [ ] `curl install.sh | bash`
- [ ] flatpak
- [ ] git extension
- [ ] docker
- [ ] i18n _(godspeed)_
- [ ] Right-To-Left _(help)_
