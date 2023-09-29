package auth

import (
	"context"

	"github.com/golang-jwt/jwt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/selector"
	"google.golang.org/grpc"
)

type matcher struct {
	excluded map[string]bool
}

func (m *matcher) Match(ctx context.Context, callMeta interceptors.CallMeta) bool {
	return !m.excluded[callMeta.FullMethod()]
}

// validateToken validates the JWT token from the context and returns a new context with the claim values.
// secret is the secret key used to sign the tokens.
// claimKeys is a slice of keys of the claims that should be put in the context, e.g. ["user_id", "role"].
func validateToken(ctx context.Context, secret string, claimKeys []string) (context.Context, error) {
	token, err := auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}
	claims := &jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	for _, k := range claimKeys {
		if claimValue, ok := (*claims)[k]; ok {
			ctx = context.WithValue(ctx, k, claimValue)
		}
	}
	return ctx, nil
}

// NewUnaryAuthInterceptor returns a new unary server interceptor that validates JWT tokens and allows some methods to be excluded from the validation.
// secret is the secret key used to sign the tokens.
// excludedMethods is a slice of full method names that should be excluded from the validation, e.g. "/helloworld.Greeter/SayHello".
// claimKeys is a slice of keys of the claims that should be put in the context, e.g. ["user_id", "role"].
func NewUnaryAuthInterceptor(secret string, excludedMethods []string, claimKeys []string) grpc.UnaryServerInterceptor {
	excludedMap := make(map[string]bool)
	for _, m := range excludedMethods {
		excludedMap[m] = true
	}
	return selector.UnaryServerInterceptor(
		auth.UnaryServerInterceptor(func(ctx context.Context) (context.Context, error) {
			return validateToken(ctx, secret, claimKeys)
		}),
		&matcher{excludedMap},
	)
}

// NewStreamAuthInterceptor returns a new stream server interceptor that validates JWT tokens and allows some methods to be excluded from the validation.
// secret is the secret key used to sign the tokens.
// excludedMethods is a slice of full method names that should be excluded from the validation, e.g. "/helloworld.Greeter/SayHello".
// claimKeys is a slice of keys of the claims that should be put in the context, e.g. ["user_id", "role"].
func NewStreamAuthInterceptor(secret string, excludedMethods []string, claimKeys []string) grpc.StreamServerInterceptor {
	excludedMap := make(map[string]bool)
	for _, m := range excludedMethods {
		excludedMap[m] = true
	}
	return selector.StreamServerInterceptor(
		auth.StreamServerInterceptor(func(ctx context.Context) (context.Context, error) {
			return validateToken(ctx, secret, claimKeys)
		}),
		&matcher{excludedMap},
	)
}
