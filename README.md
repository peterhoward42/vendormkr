# vendormkr

A vehicle for vendoring a set of Go package dependencies that can be exported 
and used by another project..

The internet proxy servers in Go development environments in some locked 
down enterprise environments make the standard Go development iteration model
awkward because ad-hoc ```go get``` requests when you start using new
external packages can be blocked by the proxy.

This repo provides the means to vendor those dependencies in an environment
that is not awkward - (for example your personal computer), and then you can 
copy out the generated ```/vendor``` directory and 
the ```go.mod``` and ```go.sum``` files.

Don't forget that in versions of Go before 1.14, you have to tell the ```go```
tools to use vendored dependencies. E.g.

```go build -mod=vendor ./...```

## How To Use It

Edit ```foo.go``` to import and refer to something in each of the packages 
you need. (Follow the pattern that's alredy in there).

Remove the entire requirements block from```go.mod``` leaving only 
something like:

```
module github.com/storgen/smeapi
go 1.12
```

Remove ```go.sum```

Fetch the dependencies the normal way:

```
go get ./...
```

Now autogenerate the vendored copies:

```
go mod vendor
```

You now have a suitably populated vendor directory, go.mod and go.sum to 
copy and use in the awkward environment.
