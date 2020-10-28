# Installation on Linux

Please use the following steps to build and install Delve on Linux.

There are two ways to install on Linux. First is the standard `go get` method:

```
go get github.com/gadelkareem/delve/cmd/dlv
```

Note: if you are using Go in modules mode you must execute this command outside of a module directory or Delve will be added to your project as a dependency.

Alternatively make sure $GOPATH is set (e.g. as `~/.go`) and:

```
$ git clone https://github.com/gadelkareem/delve.git $GOPATH/src/github.com/gadelkareem/delve
$ cd $GOPATH/src/github.com/gadelkareem/delve
$ make install
```
