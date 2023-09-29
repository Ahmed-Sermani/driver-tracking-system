package users

//go:generate buf generate

//go:generate swagger generate client -q -f ./internal/rest/api.swagger.json -c usersclient -m usersclient/models --with-flatten=remove-unused
