// export by github.com/goplus/gossa/cmd/qexp

//+build go1.16,!go1.17

package png

import (
	q "image/png"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "png",
		Path: "image/png",
		Deps: map[string]string{
			"bufio":           "bufio",
			"compress/zlib":   "zlib",
			"encoding/binary": "binary",
			"fmt":             "fmt",
			"hash":            "hash",
			"hash/crc32":      "crc32",
			"image":           "image",
			"image/color":     "color",
			"io":              "io",
			"strconv":         "strconv",
		},
		Interfaces: map[string]reflect.Type{
			"EncoderBufferPool": reflect.TypeOf((*q.EncoderBufferPool)(nil)).Elem(),
		},
		NamedTypes: map[string]gossa.NamedType{
			"CompressionLevel": {reflect.TypeOf((*q.CompressionLevel)(nil)).Elem(), "", ""},
			"Encoder":          {reflect.TypeOf((*q.Encoder)(nil)).Elem(), "", "Encode"},
			"EncoderBuffer":    {reflect.TypeOf((*q.EncoderBuffer)(nil)).Elem(), "", ""},
			"FormatError":      {reflect.TypeOf((*q.FormatError)(nil)).Elem(), "Error", ""},
			"UnsupportedError": {reflect.TypeOf((*q.UnsupportedError)(nil)).Elem(), "Error", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Decode":       reflect.ValueOf(q.Decode),
			"DecodeConfig": reflect.ValueOf(q.DecodeConfig),
			"Encode":       reflect.ValueOf(q.Encode),
		},
		TypedConsts: map[string]gossa.TypedConst{
			"BestCompression":    {reflect.TypeOf(q.BestCompression), constant.MakeInt64(int64(q.BestCompression))},
			"BestSpeed":          {reflect.TypeOf(q.BestSpeed), constant.MakeInt64(int64(q.BestSpeed))},
			"DefaultCompression": {reflect.TypeOf(q.DefaultCompression), constant.MakeInt64(int64(q.DefaultCompression))},
			"NoCompression":      {reflect.TypeOf(q.NoCompression), constant.MakeInt64(int64(q.NoCompression))},
		},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}