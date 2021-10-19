package interp

import (
	"fmt"
	"go/constant"
	"go/token"
	"go/types"
	"reflect"
	"runtime"
	"strings"
)

var (
	rtyp          = NewRtyp()
	instPkgs      = make(map[string]*Package)
	instTypesPkgs = make(map[string]*types.Package)
)

func loadTypesPackage(path string) (*types.Package, bool) {
	if p, ok := instTypesPkgs[path]; ok {
		return p, true
	}
	if pkg, ok := instPkgs[path]; ok {
		err := rtyp.InsertPackage(pkg)
		if err != nil {
			panic(fmt.Errorf("insert package %v failed: %v", path, err))
		}
		if p, ok := rtyp.Packages[path]; ok {
			p.MarkComplete()
			instTypesPkgs[path] = p
			return p, true
		}
	}
	return nil, false
}

func InstallPackage(pkg *Package) {
	if p, ok := instPkgs[pkg.Path]; ok {
		p.Types = append(p.Types, pkg.Types...)
		for k, v := range pkg.Vars {
			p.Vars[k] = v
		}
		for k, v := range pkg.Funcs {
			p.Funcs[k] = v
		}
		for k, v := range pkg.Methods {
			p.Methods[k] = v
		}
		for k, v := range pkg.UntypedConsts {
			p.UntypedConsts[k] = v
		}
		return
	}
	instPkgs[pkg.Path] = pkg
	externPackages[pkg.Path] = true
}

type TypedConst struct {
	Typ   reflect.Type
	Value constant.Value
}

type UntypedConst struct {
	Typ   string
	Value constant.Value
}

type Package struct {
	Name          string
	Path          string
	Types         []reflect.Type
	AliasTypes    map[string]reflect.Type
	Vars          map[string]reflect.Value
	Funcs         map[string]reflect.Value
	Methods       map[string]reflect.Value
	TypedConsts   map[string]TypedConst
	UntypedConsts map[string]UntypedConst
	Deps          map[string]string
}

type Rtyp struct {
	Packages map[string]*types.Package
	Rcache   map[reflect.Type]types.Type
	Tcache   map[types.Type]reflect.Type
	pkgs     map[string]*Package
	curpkg   *Package
}

func NewRtyp() *Rtyp {
	r := &Rtyp{
		Packages: make(map[string]*types.Package),
		Rcache:   make(map[reflect.Type]types.Type),
		Tcache:   make(map[types.Type]reflect.Type),
		pkgs:     make(map[string]*Package),
	}
	r.Rcache[reflect.TypeOf((*error)(nil)).Elem()] = types.Universe.Lookup("error").Type()
	r.Rcache[reflect.TypeOf((*interface{})(nil)).Elem()] = types.NewInterfaceType(nil, nil)
	return r
}

var (
	rinit = reflect.ValueOf(func() bool { return true })
)

func (r *Rtyp) InsertPackage(pkg *Package) (err error) {
	if _, ok := r.pkgs[pkg.Name]; ok {
		return nil
	}
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
		r.curpkg = nil
	}()
	r.pkgs[pkg.Path] = pkg
	r.curpkg = pkg
	finit := pkg.Path + ".init"
	if _, ok := pkg.Funcs[finit]; !ok {
		pkg.Funcs[finit] = rinit
	}
	for _, typ := range pkg.Types {
		r.InsertType(typ)
	}
	for name, typ := range pkg.AliasTypes {
		r.InsertAlias(name, typ)
	}
	for name, fn := range pkg.Funcs {
		r.InsertFunc(name, fn)
	}
	for name, v := range pkg.Vars {
		r.InsertVar(name, v.Elem())
	}
	for name, c := range pkg.TypedConsts {
		r.InsertTypedConst(name, c)
	}
	for name, c := range pkg.UntypedConsts {
		r.InsertUntypedConst(name, c)
	}
	return
}

func (r *Rtyp) InsertType(rt reflect.Type) {
	r.ToType(rt)
}

func (r *Rtyp) InsertAlias(path string, rt reflect.Type) {
	p, name := r.parserNamed(path)
	typ := r.ToType(rt)
	p.Scope().Insert(types.NewTypeName(token.NoPos, p, name, typ))
}

func (r *Rtyp) InsertFunc(path string, v reflect.Value) {
	p, name := r.parserNamed(path)
	typ := r.ToType(v.Type())
	p.Scope().Insert(types.NewFunc(token.NoPos, p, name, typ.(*types.Signature)))
}

func (r *Rtyp) InsertVar(path string, v reflect.Value) {
	p, name := r.parserNamed(path)
	typ := r.ToType(v.Type())
	p.Scope().Insert(types.NewVar(token.NoPos, p, name, typ))
}

func (r *Rtyp) InsertConstEx(path string, typ types.Type, c constant.Value) {
	p, name := r.parserNamed(path)
	p.Scope().Insert(types.NewConst(token.NoPos, p, name, typ, c))
}

func splitPath(path string) (pkg string, name string, ok bool) {
	pos := strings.LastIndex(path, ".")
	if pos == -1 {
		return path, "", false
	}
	return path[:pos], path[pos+1:], true
}

func (r *Rtyp) parserNamed(path string) (*types.Package, string) {
	if pkg, name, ok := splitPath(path); ok {
		if p := r.GetPackage(pkg); p != nil {
			return p, name
		}
	}
	panic(fmt.Errorf("parse path failed: %v", path))
}

func (r *Rtyp) LookupType(typ string) types.Type {
	if t, ok := basicTypes[typ]; ok {
		return t
	}
	p, name := r.parserNamed(typ)
	return p.Scope().Lookup(name).Type()
}

func (r *Rtyp) InsertTypedConst(path string, v TypedConst) {
	typ := r.ToType(v.Typ)
	r.InsertConstEx(path, typ, v.Value)
}

func (r *Rtyp) InsertUntypedConst(path string, v UntypedConst) {
	var typ types.Type
	if t, ok := basicTypes[v.Typ]; ok {
		typ = t
	} else {
		typ = r.LookupType(v.Typ)
	}
	r.InsertConstEx(path, typ, v.Value)
}

var (
	basicTypes = make(map[string]*types.Basic)
)

func init() {
	for i := types.Invalid; i < types.UntypedNil; i++ {
		typ := types.Typ[i]
		basicTypes[typ.String()] = typ
	}
}

func (r *Rtyp) GetPackage(pkg string) *types.Package {
	if p, ok := r.Packages[pkg]; ok {
		return p
	}
	var name string
	if r.curpkg != nil {
		name = r.curpkg.Deps[pkg]
	}
	if name == "" {
		pkgs := strings.Split(pkg, "/")
		name = pkgs[len(pkgs)-1]
	}
	p := types.NewPackage(pkg, name)
	r.Packages[pkg] = p
	return p
}

func toDir(dir reflect.ChanDir) types.ChanDir {
	switch dir {
	case reflect.RecvDir:
		return types.RecvOnly
	case reflect.SendDir:
		return types.SendOnly
	case reflect.BothDir:
		return types.SendRecv
	}
	panic("unreachable")
}

func (r *Rtyp) Insert(v reflect.Value) {
	typ := r.ToType(v.Type())
	if v.Kind() == reflect.Func {
		name := runtime.FuncForPC(v.Pointer()).Name()
		names := strings.Split(name, ".")
		pkg := r.GetPackage(names[0])
		pkg.Scope().Insert(types.NewFunc(token.NoPos, pkg, names[1], typ.(*types.Signature)))
	}
}

func (r *Rtyp) toFunc(recv *types.Var, inoff int, rt reflect.Type) *types.Signature {
	numIn := rt.NumIn()
	numOut := rt.NumOut()
	in := make([]*types.Var, numIn-inoff, numIn-inoff)
	out := make([]*types.Var, numOut, numOut)
	for i := inoff; i < numIn; i++ {
		it := r.ToType(rt.In(i))
		in[i-inoff] = types.NewVar(token.NoPos, nil, "", it)
	}
	for i := 0; i < numOut; i++ {
		it := r.ToType(rt.Out(i))
		out[i] = types.NewVar(token.NoPos, nil, "", it)
	}
	return types.NewSignature(recv, types.NewTuple(in...), types.NewTuple(out...), rt.IsVariadic())
}

var (
	typDummy = types.NewStruct(nil, nil)
	sigDummy = types.NewSignature(nil, nil, nil, false)
)

func (r *Rtyp) ToType(rt reflect.Type) types.Type {
	if t, ok := r.Rcache[rt]; ok {
		return t
	}
	var typ types.Type
	var fields []*types.Var
	var imethods []*types.Func
	kind := rt.Kind()
	switch kind {
	case reflect.Invalid:
		typ = types.Typ[types.Invalid]
	case reflect.Bool:
		typ = types.Typ[types.Bool]
	case reflect.Int:
		typ = types.Typ[types.Int]
	case reflect.Int8:
		typ = types.Typ[types.Int8]
	case reflect.Int16:
		typ = types.Typ[types.Int16]
	case reflect.Int32:
		typ = types.Typ[types.Int32]
	case reflect.Int64:
		typ = types.Typ[types.Int64]
	case reflect.Uint:
		typ = types.Typ[types.Uint]
	case reflect.Uint8:
		typ = types.Typ[types.Uint8]
	case reflect.Uint16:
		typ = types.Typ[types.Uint16]
	case reflect.Uint32:
		typ = types.Typ[types.Uint32]
	case reflect.Uint64:
		typ = types.Typ[types.Uint64]
	case reflect.Uintptr:
		typ = types.Typ[types.Uintptr]
	case reflect.Float32:
		typ = types.Typ[types.Float32]
	case reflect.Float64:
		typ = types.Typ[types.Float64]
	case reflect.Complex64:
		typ = types.Typ[types.Complex64]
	case reflect.Complex128:
		typ = types.Typ[types.Complex128]
	case reflect.Array:
		elem := r.ToType(rt.Elem())
		typ = types.NewArray(elem, int64(rt.Len()))
	case reflect.Chan:
		elem := r.ToType(rt.Elem())
		dir := toDir(rt.ChanDir())
		typ = types.NewChan(dir, elem)
	case reflect.Func:
		typ = r.toFunc(nil, 0, rt)
	case reflect.Interface:
		n := rt.NumMethod()
		imethods = make([]*types.Func, n, n)
		for i := 0; i < n; i++ {
			im := rt.Method(i)
			sig := sigDummy
			imethods[i] = types.NewFunc(token.NoPos, nil, im.Name, sig)
		}
		typ = types.NewInterfaceType(imethods, nil)
	case reflect.Map:
		key := r.ToType(rt.Key())
		elem := r.ToType(rt.Elem())
		typ = types.NewMap(key, elem)
	case reflect.Ptr:
		elem := r.ToType(rt.Elem())
		typ = types.NewPointer(elem)
	case reflect.Slice:
		elem := r.ToType(rt.Elem())
		typ = types.NewSlice(elem)
	case reflect.String:
		typ = types.Typ[types.String]
	case reflect.Struct:
		n := rt.NumField()
		fields = make([]*types.Var, n, n)
		tags := make([]string, n, n)
		for i := 0; i < n; i++ {
			f := rt.Field(i)
			ft := types.Typ[types.UnsafePointer] //r.ToType(f.Type)
			var pkg *types.Package
			if f.PkgPath != "" {
				pkg = r.GetPackage(f.PkgPath)
			}
			fields[i] = types.NewVar(token.NoPos, pkg, f.Name, ft)
			tags[i] = string(f.Tag)
		}
		typ = types.NewStruct(fields, tags)
	case reflect.UnsafePointer:
		typ = types.Typ[types.UnsafePointer]
	default:
		panic("unreachable")
	}
	var named *types.Named
	if path := rt.PkgPath(); path != "" {
		pkg := r.GetPackage(path)
		obj := types.NewTypeName(token.NoPos, pkg, rt.Name(), nil)
		named = types.NewNamed(obj, typ, nil)
		typ = named
		pkg.Scope().Insert(obj)
	}
	r.Rcache[rt] = typ
	r.Tcache[typ] = rt
	if kind == reflect.Struct {
		n := rt.NumField()
		for i := 0; i < n; i++ {
			f := rt.Field(i)
			ft := r.ToType(f.Type)
			var pkg *types.Package
			if f.PkgPath != "" {
				pkg = r.GetPackage(f.PkgPath)
			}
			fields[i] = types.NewField(token.NoPos, pkg, f.Name, ft, f.Anonymous)
		}
	} else if kind == reflect.Interface {
		n := rt.NumMethod()
		recv := types.NewVar(token.NoPos, nil, "", typ)
		for i := 0; i < n; i++ {
			im := rt.Method(i)
			sig := r.toFunc(recv, 0, im.Type)
			imethods[i] = types.NewFunc(token.NoPos, nil, im.Name, sig)
		}
		typ.Underlying().(*types.Interface).Complete()
	}
	if named != nil {
		if kind != reflect.Interface {
			pkg := named.Obj().Pkg()
			recv := types.NewVar(token.NoPos, nil, "", typ)
			n := rt.NumMethod()
			for i := 0; i < n; i++ {
				im := rt.Method(i)
				sig := r.toFunc(recv, 1, im.Type)
				named.AddMethod(types.NewFunc(token.NoPos, pkg, im.Name, sig))
			}
			prt := reflect.PtrTo(rt)
			ptyp := r.ToType(prt)
			precv := types.NewVar(token.NoPos, nil, "", ptyp)
			n = prt.NumMethod()
			for i := 0; i < n; i++ {
				im := prt.Method(i)
				sig := r.toFunc(precv, 1, im.Type)
				named.AddMethod(types.NewFunc(token.NoPos, pkg, im.Name, sig))
			}
		}
	}
	return typ
}
