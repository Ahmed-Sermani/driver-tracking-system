package emulation

//go:generate buf generate
//go:generate swagger generate client -q -f ./internal/rest/api.swagger.json -c emulationclient -m emulationclient/models --with-flatten=remove-unused
