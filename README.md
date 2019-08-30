# GOES
Inspired by go-billy, "goes" wraps "syscall/js" as an interface to make it easier
to test the Go side of interactions with the Javascript runtime without needing
to run the project in a browser or with Node.

## Contributing

Please cake sure to run `go generate ./...` when contributing to the interface declarations.

## Some random `syscall/js` tidbits

I use the following alias.
```sh
alias goes='GOOS=js GOARCH=wasm go'
```

I really like Agniva De Sarker's [wasmbrowsertest](https://github.com/agnivade/wasmbrowsertest).
