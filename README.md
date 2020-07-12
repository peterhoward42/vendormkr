# vendormkr

Developing Go projects in a corporate enterprise environment can be hampered
by the enterprises' http proxy, blocking ```go get``` requests to some URLs.

A workaround is to vendor the dependent packages and include them in your own
code repository.

But of course you can't *make* the vendor directory in the corporate 
environment because of the same proxy obstacles.

This repo helps you make the vendor directory externally. So you can then
copy it in.


A vehicle for vendoring a set of Go package dependencies that can be exported 
and used by another project..

The internet proxy servers in Go development environments in some locked 
down enterprise environments make the standard Go development iteration model
awkward because ad-hoc ```go get``` requests when you start using new
external packages can be blocked by the proxy.


## How To Use It

Edit ```foo.go``` to import and refer to something in each of the packages 
you need. (Follow the pattern that's already in there).

Remove the entire requirements block from```go.mod``` leaving only 
something like:

```
module github.com/storgen/smeapi
go 1.12
```

Remove ```go.sum```.

Fetch the dependencies the normal way:

```
go get ./...
```

Now autogenerate the vendored copy of the source code fetched by the 
```go get```:

```
go mod vendor
```

You now have a suitably populated vendor directory, go.mod and go.sum to 
copy and use in your enterprise project.

Don't forget that for Go versions prior to 1.14 you have to tell the Go tool
to use vendoring when it builds etc. like this:

```
go build -mod=vendor ./...
```
