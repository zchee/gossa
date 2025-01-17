// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.18
// +build go1.18

package types

import (
	q "go/types"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "types",
		Path: "go/types",
		Deps: map[string]string{
			"bytes":                  "bytes",
			"container/heap":         "heap",
			"errors":                 "errors",
			"fmt":                    "fmt",
			"go/ast":                 "ast",
			"go/constant":            "constant",
			"go/internal/typeparams": "typeparams",
			"go/parser":              "parser",
			"go/token":               "token",
			"io":                     "io",
			"math":                   "math",
			"regexp":                 "regexp",
			"sort":                   "sort",
			"strconv":                "strconv",
			"strings":                "strings",
			"sync":                   "sync",
			"sync/atomic":            "atomic",
			"unicode":                "unicode",
			"unicode/utf8":           "utf8",
		},
		Interfaces: map[string]reflect.Type{
			"Importer":     reflect.TypeOf((*q.Importer)(nil)).Elem(),
			"ImporterFrom": reflect.TypeOf((*q.ImporterFrom)(nil)).Elem(),
			"Object":       reflect.TypeOf((*q.Object)(nil)).Elem(),
			"Sizes":        reflect.TypeOf((*q.Sizes)(nil)).Elem(),
			"Type":         reflect.TypeOf((*q.Type)(nil)).Elem(),
		},
		NamedTypes: map[string]gossa.NamedType{
			"ArgumentError": {reflect.TypeOf((*q.ArgumentError)(nil)).Elem(), "", "Error,Unwrap"},
			"Array":         {reflect.TypeOf((*q.Array)(nil)).Elem(), "", "Elem,Len,String,Underlying"},
			"Basic":         {reflect.TypeOf((*q.Basic)(nil)).Elem(), "", "Info,Kind,Name,String,Underlying"},
			"BasicInfo":     {reflect.TypeOf((*q.BasicInfo)(nil)).Elem(), "", ""},
			"BasicKind":     {reflect.TypeOf((*q.BasicKind)(nil)).Elem(), "", ""},
			"Builtin":       {reflect.TypeOf((*q.Builtin)(nil)).Elem(), "", "String"},
			"Chan":          {reflect.TypeOf((*q.Chan)(nil)).Elem(), "", "Dir,Elem,String,Underlying"},
			"ChanDir":       {reflect.TypeOf((*q.ChanDir)(nil)).Elem(), "", ""},
			"Checker":       {reflect.TypeOf((*q.Checker)(nil)).Elem(), "", "Files,addDeclDep,allowVersion,applyTypeFunc,arguments,arityMatch,arrayLength,assertableTo,assignError,assignVar,assignVars,assignment,bestContext,binary,blockBranches,bound,brokenAlias,builtin,callExpr,caseTypes,caseValues,checkFiles,cleanup,closeScope,collectMethods,collectObjects,collectParams,collectTypeParams,comparison,constDecl,conversion,convertUntyped,cycleError,declStmt,declare,declareInSet,declarePkgObj,declareTypeParams,definedType,dump,err,error,errorUnusedPkg,errorf,exclude,expr,exprInternal,exprList,exprOrType,exprWithHint,filename,funcBody,funcDecl,funcInst,funcString,funcType,genericType,handleBailout,ident,implements,implicitTypeAndValue,importPackage,incomparableCause,index,indexExpr,indexedElts,infer,inferB,initConst,initFiles,initOrder,initVar,initVars,instance,instantiateSignature,instantiatedType,interfacePtrError,interfaceType,invalidAST,invalidArg,invalidConversion,invalidOp,isBrokenAlias,isImportedConstraint,isNil,isTerminating,isTerminatingList,isTerminatingSwitch,isValidIndex,kindString,labels,langCompat,later,markImports,missingMethod,missingMethodReason,monomorph,multiExpr,multipleDefaults,needsCleanup,newAssertableTo,newError,newErrorf,newInterface,newNamed,newTypeParam,nonGeneric,objDecl,op,openScope,overflow,packageObjects,pop,processDelayed,push,qualifier,rawExpr,record,recordBuiltinType,recordCommaOkTypes,recordDef,recordImplicit,recordInstance,recordScope,recordSelection,recordTypeAndValue,recordUntyped,recordUse,rememberUntyped,reportAltDecl,reportCycle,reportInstanceLoop,representable,representation,resolveBaseTypeName,selector,shift,shortVarDecl,simpleStmt,singleIndex,singleValue,sliceExpr,softErrorf,sprintf,stmt,stmtList,structType,subst,suspendedCall,tag,trace,typ,typInternal,typeAssertion,typeDecl,typeList,typesSummary,unary,unpackRecv,unusedImports,updateExprType,updateExprType0,updateExprVal,usage,use,useLHS,validAlias,validCycle,validType,validType0,validVarType,validateTArgLen,varDecl,varType,verify,walkDecl,walkDecls"},
			"Config":        {reflect.TypeOf((*q.Config)(nil)).Elem(), "", "Check,alignof,offsetof,offsetsof,sizeof"},
			"Const":         {reflect.TypeOf((*q.Const)(nil)).Elem(), "", "String,Val,isDependency"},
			"Context":       {reflect.TypeOf((*q.Context)(nil)).Elem(), "", "getID,instanceHash,lookup,update"},
			"Error":         {reflect.TypeOf((*q.Error)(nil)).Elem(), "Error", ""},
			"Func":          {reflect.TypeOf((*q.Func)(nil)).Elem(), "", "FullName,Scope,String,hasPtrRecv,isDependency"},
			"ImportMode":    {reflect.TypeOf((*q.ImportMode)(nil)).Elem(), "", ""},
			"Info":          {reflect.TypeOf((*q.Info)(nil)).Elem(), "", "ObjectOf,TypeOf"},
			"Initializer":   {reflect.TypeOf((*q.Initializer)(nil)).Elem(), "", "String"},
			"Instance":      {reflect.TypeOf((*q.Instance)(nil)).Elem(), "", ""},
			"Interface":     {reflect.TypeOf((*q.Interface)(nil)).Elem(), "", "Complete,Embedded,EmbeddedType,Empty,ExplicitMethod,IsComparable,IsImplicit,IsMethodSet,MarkImplicit,Method,NumEmbeddeds,NumExplicitMethods,NumMethods,String,Underlying,cleanup,typeSet"},
			"Label":         {reflect.TypeOf((*q.Label)(nil)).Elem(), "", "String"},
			"Map":           {reflect.TypeOf((*q.Map)(nil)).Elem(), "", "Elem,Key,String,Underlying"},
			"MethodSet":     {reflect.TypeOf((*q.MethodSet)(nil)).Elem(), "", "At,Len,Lookup,String"},
			"Named":         {reflect.TypeOf((*q.Named)(nil)).Elem(), "", "AddMethod,Method,NumMethods,Obj,Origin,SetTypeParams,SetUnderlying,String,TypeArgs,TypeParams,Underlying,cleanup,instantiateMethod,lookupMethod,resolve,setUnderlying,under"},
			"Nil":           {reflect.TypeOf((*q.Nil)(nil)).Elem(), "", "String"},
			"Package":       {reflect.TypeOf((*q.Package)(nil)).Elem(), "", "Complete,Imports,MarkComplete,Name,Path,Scope,SetImports,SetName,String"},
			"PkgName":       {reflect.TypeOf((*q.PkgName)(nil)).Elem(), "", "Imported,String"},
			"Pointer":       {reflect.TypeOf((*q.Pointer)(nil)).Elem(), "", "Elem,String,Underlying"},
			"Qualifier":     {reflect.TypeOf((*q.Qualifier)(nil)).Elem(), "", ""},
			"Scope":         {reflect.TypeOf((*q.Scope)(nil)).Elem(), "", "Child,Contains,End,Innermost,Insert,Len,Lookup,LookupParent,Names,NumChildren,Parent,Pos,String,WriteTo,_InsertLazy,insert,squash"},
			"Selection":     {reflect.TypeOf((*q.Selection)(nil)).Elem(), "", "Index,Indirect,Kind,Obj,Recv,String,Type"},
			"SelectionKind": {reflect.TypeOf((*q.SelectionKind)(nil)).Elem(), "", ""},
			"Signature":     {reflect.TypeOf((*q.Signature)(nil)).Elem(), "", "Params,Recv,RecvTypeParams,Results,String,TypeParams,Underlying,Variadic"},
			"Slice":         {reflect.TypeOf((*q.Slice)(nil)).Elem(), "", "Elem,String,Underlying"},
			"StdSizes":      {reflect.TypeOf((*q.StdSizes)(nil)).Elem(), "", "Alignof,Offsetsof,Sizeof"},
			"Struct":        {reflect.TypeOf((*q.Struct)(nil)).Elem(), "", "Field,NumFields,String,Tag,Underlying,markComplete"},
			"Term":          {reflect.TypeOf((*q.Term)(nil)).Elem(), "", "String,Tilde,Type"},
			"Tuple":         {reflect.TypeOf((*q.Tuple)(nil)).Elem(), "", "At,Len,String,Underlying"},
			"TypeAndValue":  {reflect.TypeOf((*q.TypeAndValue)(nil)).Elem(), "Addressable,Assignable,HasOk,IsBuiltin,IsNil,IsType,IsValue,IsVoid", ""},
			"TypeList":      {reflect.TypeOf((*q.TypeList)(nil)).Elem(), "", "At,Len,list"},
			"TypeName":      {reflect.TypeOf((*q.TypeName)(nil)).Elem(), "", "IsAlias,String"},
			"TypeParam":     {reflect.TypeOf((*q.TypeParam)(nil)).Elem(), "", "Constraint,Index,Obj,SetConstraint,String,Underlying,cleanup,iface,is,underIs"},
			"TypeParamList": {reflect.TypeOf((*q.TypeParamList)(nil)).Elem(), "", "At,Len,list"},
			"Union":         {reflect.TypeOf((*q.Union)(nil)).Elem(), "", "Len,String,Term,Underlying"},
			"Var":           {reflect.TypeOf((*q.Var)(nil)).Elem(), "", "Anonymous,Embedded,IsField,String,isDependency"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"Typ":      reflect.ValueOf(&q.Typ),
			"Universe": reflect.ValueOf(&q.Universe),
			"Unsafe":   reflect.ValueOf(&q.Unsafe),
		},
		Funcs: map[string]reflect.Value{
			"AssertableTo":            reflect.ValueOf(q.AssertableTo),
			"AssignableTo":            reflect.ValueOf(q.AssignableTo),
			"CheckExpr":               reflect.ValueOf(q.CheckExpr),
			"Comparable":              reflect.ValueOf(q.Comparable),
			"ConvertibleTo":           reflect.ValueOf(q.ConvertibleTo),
			"DefPredeclaredTestFuncs": reflect.ValueOf(q.DefPredeclaredTestFuncs),
			"Default":                 reflect.ValueOf(q.Default),
			"Eval":                    reflect.ValueOf(q.Eval),
			"ExprString":              reflect.ValueOf(q.ExprString),
			"Id":                      reflect.ValueOf(q.Id),
			"Identical":               reflect.ValueOf(q.Identical),
			"IdenticalIgnoreTags":     reflect.ValueOf(q.IdenticalIgnoreTags),
			"Implements":              reflect.ValueOf(q.Implements),
			"Instantiate":             reflect.ValueOf(q.Instantiate),
			"IsInterface":             reflect.ValueOf(q.IsInterface),
			"LookupFieldOrMethod":     reflect.ValueOf(q.LookupFieldOrMethod),
			"MissingMethod":           reflect.ValueOf(q.MissingMethod),
			"NewArray":                reflect.ValueOf(q.NewArray),
			"NewChan":                 reflect.ValueOf(q.NewChan),
			"NewChecker":              reflect.ValueOf(q.NewChecker),
			"NewConst":                reflect.ValueOf(q.NewConst),
			"NewContext":              reflect.ValueOf(q.NewContext),
			"NewField":                reflect.ValueOf(q.NewField),
			"NewFunc":                 reflect.ValueOf(q.NewFunc),
			"NewInterface":            reflect.ValueOf(q.NewInterface),
			"NewInterfaceType":        reflect.ValueOf(q.NewInterfaceType),
			"NewLabel":                reflect.ValueOf(q.NewLabel),
			"NewMap":                  reflect.ValueOf(q.NewMap),
			"NewMethodSet":            reflect.ValueOf(q.NewMethodSet),
			"NewNamed":                reflect.ValueOf(q.NewNamed),
			"NewPackage":              reflect.ValueOf(q.NewPackage),
			"NewParam":                reflect.ValueOf(q.NewParam),
			"NewPkgName":              reflect.ValueOf(q.NewPkgName),
			"NewPointer":              reflect.ValueOf(q.NewPointer),
			"NewScope":                reflect.ValueOf(q.NewScope),
			"NewSignature":            reflect.ValueOf(q.NewSignature),
			"NewSignatureType":        reflect.ValueOf(q.NewSignatureType),
			"NewSlice":                reflect.ValueOf(q.NewSlice),
			"NewStruct":               reflect.ValueOf(q.NewStruct),
			"NewTerm":                 reflect.ValueOf(q.NewTerm),
			"NewTuple":                reflect.ValueOf(q.NewTuple),
			"NewTypeName":             reflect.ValueOf(q.NewTypeName),
			"NewTypeParam":            reflect.ValueOf(q.NewTypeParam),
			"NewUnion":                reflect.ValueOf(q.NewUnion),
			"NewVar":                  reflect.ValueOf(q.NewVar),
			"ObjectString":            reflect.ValueOf(q.ObjectString),
			"RelativeTo":              reflect.ValueOf(q.RelativeTo),
			"SelectionString":         reflect.ValueOf(q.SelectionString),
			"SizesFor":                reflect.ValueOf(q.SizesFor),
			"TypeString":              reflect.ValueOf(q.TypeString),
			"WriteExpr":               reflect.ValueOf(q.WriteExpr),
			"WriteSignature":          reflect.ValueOf(q.WriteSignature),
			"WriteType":               reflect.ValueOf(q.WriteType),
		},
		TypedConsts: map[string]gossa.TypedConst{
			"Bool":           {reflect.TypeOf(q.Bool), constant.MakeInt64(int64(q.Bool))},
			"Byte":           {reflect.TypeOf(q.Byte), constant.MakeInt64(int64(q.Byte))},
			"Complex128":     {reflect.TypeOf(q.Complex128), constant.MakeInt64(int64(q.Complex128))},
			"Complex64":      {reflect.TypeOf(q.Complex64), constant.MakeInt64(int64(q.Complex64))},
			"FieldVal":       {reflect.TypeOf(q.FieldVal), constant.MakeInt64(int64(q.FieldVal))},
			"Float32":        {reflect.TypeOf(q.Float32), constant.MakeInt64(int64(q.Float32))},
			"Float64":        {reflect.TypeOf(q.Float64), constant.MakeInt64(int64(q.Float64))},
			"Int":            {reflect.TypeOf(q.Int), constant.MakeInt64(int64(q.Int))},
			"Int16":          {reflect.TypeOf(q.Int16), constant.MakeInt64(int64(q.Int16))},
			"Int32":          {reflect.TypeOf(q.Int32), constant.MakeInt64(int64(q.Int32))},
			"Int64":          {reflect.TypeOf(q.Int64), constant.MakeInt64(int64(q.Int64))},
			"Int8":           {reflect.TypeOf(q.Int8), constant.MakeInt64(int64(q.Int8))},
			"Invalid":        {reflect.TypeOf(q.Invalid), constant.MakeInt64(int64(q.Invalid))},
			"IsBoolean":      {reflect.TypeOf(q.IsBoolean), constant.MakeInt64(int64(q.IsBoolean))},
			"IsComplex":      {reflect.TypeOf(q.IsComplex), constant.MakeInt64(int64(q.IsComplex))},
			"IsConstType":    {reflect.TypeOf(q.IsConstType), constant.MakeInt64(int64(q.IsConstType))},
			"IsFloat":        {reflect.TypeOf(q.IsFloat), constant.MakeInt64(int64(q.IsFloat))},
			"IsInteger":      {reflect.TypeOf(q.IsInteger), constant.MakeInt64(int64(q.IsInteger))},
			"IsNumeric":      {reflect.TypeOf(q.IsNumeric), constant.MakeInt64(int64(q.IsNumeric))},
			"IsOrdered":      {reflect.TypeOf(q.IsOrdered), constant.MakeInt64(int64(q.IsOrdered))},
			"IsString":       {reflect.TypeOf(q.IsString), constant.MakeInt64(int64(q.IsString))},
			"IsUnsigned":     {reflect.TypeOf(q.IsUnsigned), constant.MakeInt64(int64(q.IsUnsigned))},
			"IsUntyped":      {reflect.TypeOf(q.IsUntyped), constant.MakeInt64(int64(q.IsUntyped))},
			"MethodExpr":     {reflect.TypeOf(q.MethodExpr), constant.MakeInt64(int64(q.MethodExpr))},
			"MethodVal":      {reflect.TypeOf(q.MethodVal), constant.MakeInt64(int64(q.MethodVal))},
			"RecvOnly":       {reflect.TypeOf(q.RecvOnly), constant.MakeInt64(int64(q.RecvOnly))},
			"Rune":           {reflect.TypeOf(q.Rune), constant.MakeInt64(int64(q.Rune))},
			"SendOnly":       {reflect.TypeOf(q.SendOnly), constant.MakeInt64(int64(q.SendOnly))},
			"SendRecv":       {reflect.TypeOf(q.SendRecv), constant.MakeInt64(int64(q.SendRecv))},
			"String":         {reflect.TypeOf(q.String), constant.MakeInt64(int64(q.String))},
			"Uint":           {reflect.TypeOf(q.Uint), constant.MakeInt64(int64(q.Uint))},
			"Uint16":         {reflect.TypeOf(q.Uint16), constant.MakeInt64(int64(q.Uint16))},
			"Uint32":         {reflect.TypeOf(q.Uint32), constant.MakeInt64(int64(q.Uint32))},
			"Uint64":         {reflect.TypeOf(q.Uint64), constant.MakeInt64(int64(q.Uint64))},
			"Uint8":          {reflect.TypeOf(q.Uint8), constant.MakeInt64(int64(q.Uint8))},
			"Uintptr":        {reflect.TypeOf(q.Uintptr), constant.MakeInt64(int64(q.Uintptr))},
			"UnsafePointer":  {reflect.TypeOf(q.UnsafePointer), constant.MakeInt64(int64(q.UnsafePointer))},
			"UntypedBool":    {reflect.TypeOf(q.UntypedBool), constant.MakeInt64(int64(q.UntypedBool))},
			"UntypedComplex": {reflect.TypeOf(q.UntypedComplex), constant.MakeInt64(int64(q.UntypedComplex))},
			"UntypedFloat":   {reflect.TypeOf(q.UntypedFloat), constant.MakeInt64(int64(q.UntypedFloat))},
			"UntypedInt":     {reflect.TypeOf(q.UntypedInt), constant.MakeInt64(int64(q.UntypedInt))},
			"UntypedNil":     {reflect.TypeOf(q.UntypedNil), constant.MakeInt64(int64(q.UntypedNil))},
			"UntypedRune":    {reflect.TypeOf(q.UntypedRune), constant.MakeInt64(int64(q.UntypedRune))},
			"UntypedString":  {reflect.TypeOf(q.UntypedString), constant.MakeInt64(int64(q.UntypedString))},
		},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
