// export by github.com/goplus/gossa/cmd/qexp

//+build go1.14,!go1.15

package hex

import (
	q "encoding/hex"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "hex",
		Path: "encoding/hex",
		Deps: map[string]string{
			"errors":  "errors",
			"fmt":     "fmt",
			"io":      "io",
			"strings": "strings",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"InvalidByteError": {reflect.TypeOf((*q.InvalidByteError)(nil)).Elem(), "Error", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrLength": reflect.ValueOf(&q.ErrLength),
		},
		Funcs: map[string]reflect.Value{
			"Decode":         reflect.ValueOf(q.Decode),
			"DecodeString":   reflect.ValueOf(q.DecodeString),
			"DecodedLen":     reflect.ValueOf(q.DecodedLen),
			"Dump":           reflect.ValueOf(q.Dump),
			"Dumper":         reflect.ValueOf(q.Dumper),
			"Encode":         reflect.ValueOf(q.Encode),
			"EncodeToString": reflect.ValueOf(q.EncodeToString),
			"EncodedLen":     reflect.ValueOf(q.EncodedLen),
			"NewDecoder":     reflect.ValueOf(q.NewDecoder),
			"NewEncoder":     reflect.ValueOf(q.NewEncoder),
		},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
