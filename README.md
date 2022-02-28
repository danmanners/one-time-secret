# Project for learning go - One Time Secret

This project is loosely based on my [Non-Disclosure-Agreement](https://github.com/danmanners/non-disclosure-agreement) Python project, but written in Go.

## What it can do

- Store text
- Store URLs and redirect
- Store whatever you want

## What it cannot do

- Be highly-available
- Persistently store anything
- Encrypt Secrets

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

## Generating your own AES-256-CFB8 keys

```bash
# Generate the AES Keys
openssl enc -aes-256-cfb8 -k secret -P -md sha512 -pbkdf2 -iter 100000

# stdout
salt=E25EB815671DC132
key=4230E0F4626FBC1071D3818909D7953D02A77F2A89B9D0A72E50B2322F3334A8
iv =9919B3F5064249DE3BC767916B0417A6
```

## To-Do

- [ ] Make it highly scalable
- [ ] Add in Multiple User-Selectable Backends
  - [x] Built-in Standalone
  - [ ] Redis
  - [ ] etcd
- [ ] Encrypt _literally_ anything; shit isn't even encoded.
