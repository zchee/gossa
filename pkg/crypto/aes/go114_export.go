// export by github.com/goplus/gossa/cmd/qexp

//+build go1.14,!go1.15

package aes

import (
	q "crypto/aes"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "aes",
		Path: "crypto/aes",
		Deps: map[string]string{
			"crypto/cipher":          "cipher",
			"crypto/internal/subtle": "subtle",
			"crypto/subtle":          "subtle",
			"encoding/binary":        "binary",
			"errors":                 "errors",
			"internal/cpu":           "cpu",
			"strconv":                "strconv",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"KeySizeError": {reflect.TypeOf((*q.KeySizeError)(nil)).Elem(), "Error", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewCipher": reflect.ValueOf(q.NewCipher),
		},
		TypedConsts: map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{
			"BlockSize": {"untyped int", constant.MakeInt64(int64(q.BlockSize))},
		},
	})
}
