//go:build go1.17 && !go1.18
// +build go1.17,!go1.18

package pkg

import (
	_ "github.com/goplus/gossa/pkg/archive/tar"
	_ "github.com/goplus/gossa/pkg/archive/zip"
	_ "github.com/goplus/gossa/pkg/bufio"
	_ "github.com/goplus/gossa/pkg/bytes"
	_ "github.com/goplus/gossa/pkg/compress/bzip2"
	_ "github.com/goplus/gossa/pkg/compress/flate"
	_ "github.com/goplus/gossa/pkg/compress/gzip"
	_ "github.com/goplus/gossa/pkg/compress/lzw"
	_ "github.com/goplus/gossa/pkg/compress/zlib"
	_ "github.com/goplus/gossa/pkg/container/heap"
	_ "github.com/goplus/gossa/pkg/container/list"
	_ "github.com/goplus/gossa/pkg/container/ring"
	_ "github.com/goplus/gossa/pkg/context"
	_ "github.com/goplus/gossa/pkg/crypto"
	_ "github.com/goplus/gossa/pkg/crypto/aes"
	_ "github.com/goplus/gossa/pkg/crypto/cipher"
	_ "github.com/goplus/gossa/pkg/crypto/des"
	_ "github.com/goplus/gossa/pkg/crypto/dsa"
	_ "github.com/goplus/gossa/pkg/crypto/ecdsa"
	_ "github.com/goplus/gossa/pkg/crypto/ed25519"
	_ "github.com/goplus/gossa/pkg/crypto/elliptic"
	_ "github.com/goplus/gossa/pkg/crypto/hmac"
	_ "github.com/goplus/gossa/pkg/crypto/md5"
	_ "github.com/goplus/gossa/pkg/crypto/rand"
	_ "github.com/goplus/gossa/pkg/crypto/rc4"
	_ "github.com/goplus/gossa/pkg/crypto/rsa"
	_ "github.com/goplus/gossa/pkg/crypto/sha1"
	_ "github.com/goplus/gossa/pkg/crypto/sha256"
	_ "github.com/goplus/gossa/pkg/crypto/sha512"
	_ "github.com/goplus/gossa/pkg/crypto/subtle"
	_ "github.com/goplus/gossa/pkg/crypto/tls"
	_ "github.com/goplus/gossa/pkg/crypto/x509"
	_ "github.com/goplus/gossa/pkg/crypto/x509/pkix"
	_ "github.com/goplus/gossa/pkg/database/sql"
	_ "github.com/goplus/gossa/pkg/database/sql/driver"
	_ "github.com/goplus/gossa/pkg/debug/dwarf"
	_ "github.com/goplus/gossa/pkg/debug/elf"
	_ "github.com/goplus/gossa/pkg/debug/gosym"
	_ "github.com/goplus/gossa/pkg/debug/macho"
	_ "github.com/goplus/gossa/pkg/debug/pe"
	_ "github.com/goplus/gossa/pkg/debug/plan9obj"
	_ "github.com/goplus/gossa/pkg/embed"
	_ "github.com/goplus/gossa/pkg/encoding"
	_ "github.com/goplus/gossa/pkg/encoding/ascii85"
	_ "github.com/goplus/gossa/pkg/encoding/asn1"
	_ "github.com/goplus/gossa/pkg/encoding/base32"
	_ "github.com/goplus/gossa/pkg/encoding/base64"
	_ "github.com/goplus/gossa/pkg/encoding/binary"
	_ "github.com/goplus/gossa/pkg/encoding/csv"
	_ "github.com/goplus/gossa/pkg/encoding/gob"
	_ "github.com/goplus/gossa/pkg/encoding/hex"
	_ "github.com/goplus/gossa/pkg/encoding/json"
	_ "github.com/goplus/gossa/pkg/encoding/pem"
	_ "github.com/goplus/gossa/pkg/encoding/xml"
	_ "github.com/goplus/gossa/pkg/errors"
	_ "github.com/goplus/gossa/pkg/expvar"
	_ "github.com/goplus/gossa/pkg/flag"
	_ "github.com/goplus/gossa/pkg/fmt"
	_ "github.com/goplus/gossa/pkg/go/ast"
	_ "github.com/goplus/gossa/pkg/go/build"
	_ "github.com/goplus/gossa/pkg/go/build/constraint"
	_ "github.com/goplus/gossa/pkg/go/constant"
	_ "github.com/goplus/gossa/pkg/go/doc"
	_ "github.com/goplus/gossa/pkg/go/format"
	_ "github.com/goplus/gossa/pkg/go/importer"
	_ "github.com/goplus/gossa/pkg/go/parser"
	_ "github.com/goplus/gossa/pkg/go/printer"
	_ "github.com/goplus/gossa/pkg/go/scanner"
	_ "github.com/goplus/gossa/pkg/go/token"
	_ "github.com/goplus/gossa/pkg/go/types"
	_ "github.com/goplus/gossa/pkg/hash"
	_ "github.com/goplus/gossa/pkg/hash/adler32"
	_ "github.com/goplus/gossa/pkg/hash/crc32"
	_ "github.com/goplus/gossa/pkg/hash/crc64"
	_ "github.com/goplus/gossa/pkg/hash/fnv"
	_ "github.com/goplus/gossa/pkg/hash/maphash"
	_ "github.com/goplus/gossa/pkg/html"
	_ "github.com/goplus/gossa/pkg/html/template"
	_ "github.com/goplus/gossa/pkg/image"
	_ "github.com/goplus/gossa/pkg/image/color"
	_ "github.com/goplus/gossa/pkg/image/color/palette"
	_ "github.com/goplus/gossa/pkg/image/draw"
	_ "github.com/goplus/gossa/pkg/image/gif"
	_ "github.com/goplus/gossa/pkg/image/jpeg"
	_ "github.com/goplus/gossa/pkg/image/png"
	_ "github.com/goplus/gossa/pkg/index/suffixarray"
	_ "github.com/goplus/gossa/pkg/io"
	_ "github.com/goplus/gossa/pkg/io/fs"
	_ "github.com/goplus/gossa/pkg/io/ioutil"
	_ "github.com/goplus/gossa/pkg/log"
	_ "github.com/goplus/gossa/pkg/math"
	_ "github.com/goplus/gossa/pkg/math/big"
	_ "github.com/goplus/gossa/pkg/math/bits"
	_ "github.com/goplus/gossa/pkg/math/cmplx"
	_ "github.com/goplus/gossa/pkg/math/rand"
	_ "github.com/goplus/gossa/pkg/mime"
	_ "github.com/goplus/gossa/pkg/mime/multipart"
	_ "github.com/goplus/gossa/pkg/mime/quotedprintable"
	_ "github.com/goplus/gossa/pkg/net"
	_ "github.com/goplus/gossa/pkg/net/http"
	_ "github.com/goplus/gossa/pkg/net/http/cgi"
	_ "github.com/goplus/gossa/pkg/net/http/cookiejar"
	_ "github.com/goplus/gossa/pkg/net/http/fcgi"
	_ "github.com/goplus/gossa/pkg/net/http/httptest"
	_ "github.com/goplus/gossa/pkg/net/http/httptrace"
	_ "github.com/goplus/gossa/pkg/net/http/httputil"
	_ "github.com/goplus/gossa/pkg/net/http/pprof"
	_ "github.com/goplus/gossa/pkg/net/mail"
	_ "github.com/goplus/gossa/pkg/net/rpc"
	_ "github.com/goplus/gossa/pkg/net/rpc/jsonrpc"
	_ "github.com/goplus/gossa/pkg/net/smtp"
	_ "github.com/goplus/gossa/pkg/net/textproto"
	_ "github.com/goplus/gossa/pkg/net/url"
	_ "github.com/goplus/gossa/pkg/os"
	_ "github.com/goplus/gossa/pkg/os/exec"
	_ "github.com/goplus/gossa/pkg/os/signal"
	_ "github.com/goplus/gossa/pkg/os/user"
	_ "github.com/goplus/gossa/pkg/path"
	_ "github.com/goplus/gossa/pkg/path/filepath"
	_ "github.com/goplus/gossa/pkg/plugin"
	_ "github.com/goplus/gossa/pkg/reflect"
	_ "github.com/goplus/gossa/pkg/regexp"
	_ "github.com/goplus/gossa/pkg/regexp/syntax"
	_ "github.com/goplus/gossa/pkg/runtime"
	_ "github.com/goplus/gossa/pkg/runtime/debug"
	_ "github.com/goplus/gossa/pkg/runtime/metrics"
	_ "github.com/goplus/gossa/pkg/runtime/pprof"
	_ "github.com/goplus/gossa/pkg/runtime/trace"
	_ "github.com/goplus/gossa/pkg/sort"
	_ "github.com/goplus/gossa/pkg/strconv"
	_ "github.com/goplus/gossa/pkg/strings"
	_ "github.com/goplus/gossa/pkg/sync"
	_ "github.com/goplus/gossa/pkg/sync/atomic"
	_ "github.com/goplus/gossa/pkg/syscall"
	_ "github.com/goplus/gossa/pkg/text/scanner"
	_ "github.com/goplus/gossa/pkg/text/tabwriter"
	_ "github.com/goplus/gossa/pkg/text/template"
	_ "github.com/goplus/gossa/pkg/text/template/parse"
	_ "github.com/goplus/gossa/pkg/time"
	_ "github.com/goplus/gossa/pkg/unicode"
	_ "github.com/goplus/gossa/pkg/unicode/utf16"
	_ "github.com/goplus/gossa/pkg/unicode/utf8"
)
