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
go build -o ots main.go
```

## To-Do

- [ ] Add in Multiple User-Selectable Backends
  - [ ] Redis
  - [ ] etcd
- [ ] Encrypt _literally_ anything; shit isn't even encoded.
