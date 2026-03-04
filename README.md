[SikaLabs (sikalabs.com)](https://sikalabs.com) | [Ondrej Sika (sika.io)](https://sika.io)  | [**Skoleni Go**](https://ondrej-sika.cz/skoleni/go/) 🚀💻

# sikalabs-go-tools

`sikalabs-go-tools` is a single binary which contains [slu](https://github.com/sikalabs/slu), [slr](https://github.com/sikalabs/slr), [tergum](https://github.com/sikalabs/tergum), [gobble](https://github.com/sikalabs/gobble) and [mon](https://github.com/sikalabs/mon).

Idea behind it is to save disk space mostly in Docker images. Instead of installing 5 separate binaries, you can install just one and use it for all 5 tools. Instead of 5x 200MB you have just 1x 250MB.

The binary works like `sikalabs-go-tools slu` or `sikalabs-go-tools slr` and so on or you can symlink it to `slu`, `slr` and so on and use it like that.

```
ln -s /usr/local/bin/sikalabs-go-tools /usr/local/bin/slu
ln -s /usr/local/bin/sikalabs-go-tools /usr/local/bin/slr
# and so on
```

## Installation

You can install `sikalabs-go-tools` using `go install`:

```
go install github.com/sikalabs/sikalabs-go-tools@latest
```

## Inspiration

The idea behind `sikalabs-go-tools` is inspired by my example project [ondrejsika/large](https://github.com/ondrejsika/large) for my [Docker Training](https://ondrej-sika.cz/skoleni/docker).
