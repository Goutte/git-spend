## Architectural Decision Records

### No dashes in subcommand names

> Dashes in subcommands' names do not play well with `man` generation.

### Some messages are not translated

Those are messages handled by Cobra, like the usage flag for `--help` and some titles.

> I think we should try our best to fix these upstream in Cobra.

### Debian package

> I think `git-spend` could be rewritten in less than `2 Mio`, in pure `bash`.
> A wise friend told me that no-one enjoys maintaining huge and complex bash scripts.
> So I'm undecided, for now ; I'll let _you_ be the judge.


## Manpages

Since `git help sum` will try to fetch a manpage, we're providing one.

To install `man` pages, use:

    sudo git spend man --install

> That API is very experimental and is likely to change in the future, which is why it's hidden for now.

or, if you cloned this source tree:

    make install-man

> … and now a debian package makes sense ;
> can't find yet if `go get` allows hooks to install manpages.

