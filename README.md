# Project for learning go - One Time Secret

This project is loosely based on my [Non-Disclosure-Agreement](https://github.com/danmanners/non-disclosure-agreement) Python project, but written in Go.

## What it can do

- Store text
- Store URLs and redirect
- Store whatever you want

## What it cannot do

- Be highly-available
- Persistently store anything

## Building

You can build this shitty software with:

```bash
# Build the binary
go build -o ots main.go
```

Or, you can build it with `docker` or `podman`:

```bash
# Build Container
sudo podman build -t ghcr.io/danmanners/ots:$(git rev-parse --short HEAD) .
```

## To-Do

- [ ] Add in Multiple User-Selectable Backends
  - [ ] Redis
  - [ ] etcd
- [ ] Encrypt _literally_ anything; shit isn't even encoded.
