package proxy

//go:generate buf generate
//go:generate swagger generate client -q -f ./internal/rest/api.swagger.json -c proxyclient -m proxyclient/models --with-flatten=remove-unused
