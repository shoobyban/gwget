# Wget based on file list

[![Documentation](https://godoc.org/github.com/shoobyban/gwget?status.svg)](http://godoc.org/github.com/shoobyban/gwget)
[![Go Report Card](https://goreportcard.com/badge/github.com/shoobyban/gwget)](https://goreportcard.com/report/github.com/shoobyban/gwget)
[![Build Status](https://travis-ci.org/shoobyban/gwget.svg?branch=master)](https://travis-ci.org/ShoobyBan/gwget)

The tool can (for example) mirror a media folder to a staging or developer server from the current live, over standard http requests, when `scp` is way too slow, creating tarball is not an option because of space requirements, VPN makes it hard to run `tar | tar` or `rsync`.

## Usage

```sh
source$ cd ~/web/media
source$ find . -type f | grep -v '.thumb' | grep -v product/cache | sed 's#^./##' > ~/media.txt
```

`scp media.txt` onto the target server with a gwget binary

```sh
target$ cd ~/web/media
target$ ~/gwget ~/media.txt https://www.somesite.com/media/
```
