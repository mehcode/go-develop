# go-develop
> Wrapper for `go get` to setup a project for development.

## Features
 - Uses `git clone` (therefore supports ssh, private repositories, etc.)
 - Operates like `git clone` by creating a symlink to the project directory from GOPATH to the current directory

## Usage
Each of the following commands result in an identical project layout in GOPATH.

```
$ go-develop https://github.com/mehcode/go-develop.git
$ go-develop https://github.com/mehcode/go-develop
$ go-develop git@github.com:mehcode/go-develop
$ go-develop git@github.com:mehcode/go-develop.git
$ go-develop ssh://git@github.com/mehcode/go-develop.git
```
