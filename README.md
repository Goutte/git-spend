gitime : time tracker using git commit message commands
=======================================================

Purpose
-------

Collect, addition and return all the `/spend` and `/spent` time-tracking directives in git commit messages.

> This only looks at the `git log` of the currently checked out branch.


Download
--------

You can download the latest build from the releases.


Usage
-----

Go into your git-versioned project's directory, and run:

```
    gitime
```

You can also get the spent time in a specific unit :

```
    gitime --minutes
```


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
    build/gitime
```


Dependencies
------------

	go get -u github.com/tsuyoshiwada/go-gitlog

