// This file was generated by nomdl/codegen.

package test

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __testPackageInFile_struct_optional_CachedRef = __testPackageInFile_struct_optional_Ref()

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func __testPackageInFile_struct_optional_Ref() ref.Ref {
	p := types.NewPackage([]types.TypeRef{
		types.MakeStructTypeRef("OptionalStruct",
			[]types.Field{
				types.Field{"s", types.MakePrimitiveTypeRef(types.StringKind), true},
				types.Field{"b", types.MakePrimitiveTypeRef(types.BoolKind), true},
			},
			types.Choices{},
		),
	}, []ref.Ref{})
	return types.RegisterPackage(&p)
}

// OptionalStruct

type OptionalStruct struct {
	m   types.Map
	ref *ref.Ref
}

func NewOptionalStruct() OptionalStruct {
	return OptionalStruct{types.NewMap(
		types.NewString("$type"), types.MakeTypeRef(__testPackageInFile_struct_optional_CachedRef, 0),
	), &ref.Ref{}}
}

type OptionalStructDef struct {
	S string
	B bool
}

func (def OptionalStructDef) New() OptionalStruct {
	return OptionalStruct{
		types.NewMap(
			types.NewString("$type"), types.MakeTypeRef(__testPackageInFile_struct_optional_CachedRef, 0),
			types.NewString("s"), types.NewString(def.S),
			types.NewString("b"), types.Bool(def.B),
		), &ref.Ref{}}
}

func (s OptionalStruct) Def() (d OptionalStructDef) {
	if v, ok := s.m.MaybeGet(types.NewString("s")); ok {
		d.S = v.(types.String).String()
	}
	if v, ok := s.m.MaybeGet(types.NewString("b")); ok {
		d.B = bool(v.(types.Bool))
	}
	return
}

var __typeRefForOptionalStruct = types.MakeTypeRef(__testPackageInFile_struct_optional_CachedRef, 0)

func (m OptionalStruct) TypeRef() types.TypeRef {
	return __typeRefForOptionalStruct
}

func init() {
	types.RegisterFromValFunction(__typeRefForOptionalStruct, func(v types.Value) types.Value {
		return OptionalStructFromVal(v)
	})
}

func OptionalStructFromVal(val types.Value) OptionalStruct {
	// TODO: Do we still need FromVal?
	if val, ok := val.(OptionalStruct); ok {
		return val
	}
	// TODO: Validate here
	return OptionalStruct{val.(types.Map), &ref.Ref{}}
}

func (s OptionalStruct) InternalImplementation() types.Map {
	return s.m
}

func (s OptionalStruct) Equals(other types.Value) bool {
	if other, ok := other.(OptionalStruct); ok {
		return s.Ref() == other.Ref()
	}
	return false
}

func (s OptionalStruct) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s OptionalStruct) Chunks() (futures []types.Future) {
	futures = append(futures, s.TypeRef().Chunks()...)
	futures = append(futures, s.m.Chunks()...)
	return
}

func (s OptionalStruct) S() (v string, ok bool) {
	var vv types.Value
	if vv, ok = s.m.MaybeGet(types.NewString("s")); ok {
		v = vv.(types.String).String()
	}
	return
}

func (s OptionalStruct) SetS(val string) OptionalStruct {
	return OptionalStruct{s.m.Set(types.NewString("s"), types.NewString(val)), &ref.Ref{}}
}

func (s OptionalStruct) B() (v bool, ok bool) {
	var vv types.Value
	if vv, ok = s.m.MaybeGet(types.NewString("b")); ok {
		v = bool(vv.(types.Bool))
	}
	return
}

func (s OptionalStruct) SetB(val bool) OptionalStruct {
	return OptionalStruct{s.m.Set(types.NewString("b"), types.Bool(val)), &ref.Ref{}}
}
