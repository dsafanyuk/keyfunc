package keyfunc

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

var (

	// ErrUnsupportedKeyType indicates the JWT key type is an unsupported type.
	ErrUnsupportedKeyType = errors.New("the JWT key type is unsupported")
)

// KeyFunc is a compatibility function that matches the signature of github.com/dgrijalva/jwt-go's KeyFunc function.
func (j *JWKS) KeyFunc() (keyFunc jwt.Keyfunc) {
	return func(token *jwt.Token) (interface{}, error) {

		// Determine the key's algorithm and return the appropriate public key.
		switch keyAlg := token.Header["alg"]; keyAlg {
		case es256, es384, es512:
			return j.ECDSA(token.Header["kid"].(string))
		case ps256, ps384, ps512, rs256, rs384, rs512:
			return j.RSA(token.Header["kid"].(string))
		default:
			return nil, fmt.Errorf("%w: %s: feel free to add a feature request or contribute to https://github.com/MicahParks/keyfunc", ErrUnsupportedKeyType, keyAlg)
		}
	}
}
