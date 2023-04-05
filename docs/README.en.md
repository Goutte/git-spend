## Architectural Decision Records

### No dashes in subcommand names

> Dashes in subcommands' names do not play well with `man` generation.

### Some messages are not translated

Those are messages handled by Cobra, like the usage flag for `--help` and some titles.

> I think we should try our best to [fix these upstream in Cobra](https://github.com/spf13/cobra/pull/1944).
> For now we're using our branch `feat-i18n` of Cobra, see `go.mod`.

### Debian package

> I think `git-spend` could be rewritten in less than `2 Mio`, in pure `bash`.
> A wise friend told me that no-one enjoys maintaining huge and complex bash scripts.
> So I'm undecided, for now ; I'll let _you_ decide if _git-spend_ is worthy of debian packages.

### Golang v1.20 and upwards

Golang `1.20` introduces new tools for:
- code coverage _(we use those a lot already)_
- i18n _(same)_


## Manpages

Since `git help sum` will try to fetch a manpage, we're providing one.

To install `man` pages, use:

    sudo git spend man --install

> That API is very experimental and is likely to change in the future, which is why it's hidden for now.
> Perhaps it should/will become `git spend man install` ?

or, if you cloned this source tree:

    make install-man

> … and now a debian package makes sense ;
> can't find yet if `go get` allows hooks to install manpages.

We could also perhaps show a message somewhere if we detect that manpages are missing
and provide there the command hint to install them.
