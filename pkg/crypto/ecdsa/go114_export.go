// export by github.com/goplus/gossa/cmd/qexp

//+build go1.14,!go1.15

package ecdsa

import (
	q "crypto/ecdsa"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "ecdsa",
		Path: "crypto/ecdsa",
		Deps: map[string]string{
			"crypto":                   "crypto",
			"crypto/aes":               "aes",
			"crypto/cipher":            "cipher",
			"crypto/elliptic":          "elliptic",
			"crypto/internal/randutil": "randutil",
			"crypto/sha512":            "sha512",
			"encoding/asn1":            "asn1",
			"errors":                   "errors",
			"io":                       "io",
			"math/big":                 "big",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"PrivateKey": {reflect.TypeOf((*q.PrivateKey)(nil)).Elem(), "", "Public,Sign"},
			"PublicKey":  {reflect.TypeOf((*q.PublicKey)(nil)).Elem(), "", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"GenerateKey": reflect.ValueOf(q.GenerateKey),
			"Sign":        reflect.ValueOf(q.Sign),
			"Verify":      reflect.ValueOf(q.Verify),
		},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
