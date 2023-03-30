In this directory are the manifest configurations for various packaging platforms.


## Go get

    go get -u github.com/goutte/git-spend


## Github Releases

Things I built using `make release`:

- Linux: https://github.com/Goutte/git-spend/releases/latest/download/git-spend
- Windows: https://github.com/Goutte/git-spend/releases/latest/download/git-spend.exe

> Wanted: CI releases with artifacts


## Flatpak

Packaging with flatpak is not yet supported.

The manifest is half-there, but we need to solve these roadblocks first. 

### Roadblocks

- Git Access (how?)
- Filesystem Access:
    - global (not recommended)
    - https://docs.flatpak.org/en/latest/sandbox-permissions.html#portals
    - overrides (meh, but quick)


## Debian

The final binary weighs almost `2 Mio`.
I believe that with enough patience, skill and work, one could make
a shell (or bash) version of `git-spend` for a fraction of that.

I'd reserve the debian package for such a rewrite.


