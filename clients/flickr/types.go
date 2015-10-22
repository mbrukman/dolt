// This file was generated by nomdl/codegen.

package main

import (
	"github.com/attic-labs/noms/clients/gen/sha1_b525f9bca5e451c21dd9af564f0960045fbaa304"

	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __mainPackageInFile_types_CachedRef = __mainPackageInFile_types_Ref()

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func __mainPackageInFile_types_Ref() ref.Ref {
	p := types.NewPackage([]types.TypeRef{
		types.MakeStructTypeRef("User",
			[]types.Field{
				types.Field{"Id", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"Name", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"OAuthToken", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"OAuthSecret", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"Albums", types.MakeCompoundTypeRef("", types.MapKind, types.MakePrimitiveTypeRef(types.StringKind), types.MakeTypeRef(ref.Ref{}, 1)), false},
			},
			types.Choices{},
		),
		types.MakeStructTypeRef("Album",
			[]types.Field{
				types.Field{"Id", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"Title", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"Photos", types.MakeCompoundTypeRef("", types.SetKind, types.MakeTypeRef(ref.Parse("sha1-b525f9bca5e451c21dd9af564f0960045fbaa304"), 0)), false},
			},
			types.Choices{},
		),
	}, []ref.Ref{
		ref.Parse("sha1-b525f9bca5e451c21dd9af564f0960045fbaa304"),
	})
	return types.RegisterPackage(&p)
}

// User

type User struct {
	m   types.Map
	ref *ref.Ref
}

func NewUser() User {
	return User{types.NewMap(
		types.NewString("$type"), types.MakeTypeRef(__mainPackageInFile_types_CachedRef, 0),
		types.NewString("Id"), types.NewString(""),
		types.NewString("Name"), types.NewString(""),
		types.NewString("OAuthToken"), types.NewString(""),
		types.NewString("OAuthSecret"), types.NewString(""),
		types.NewString("Albums"), NewMapOfStringToAlbum(),
	), &ref.Ref{}}
}

var __typeRefForUser = types.MakeTypeRef(__mainPackageInFile_types_CachedRef, 0)

func (m User) TypeRef() types.TypeRef {
	return __typeRefForUser
}

func init() {
	types.RegisterFromValFunction(__typeRefForUser, func(v types.Value) types.Value {
		return UserFromVal(v)
	})
}

func UserFromVal(val types.Value) User {
	// TODO: Do we still need FromVal?
	if val, ok := val.(User); ok {
		return val
	}
	// TODO: Validate here
	return User{val.(types.Map), &ref.Ref{}}
}

func (s User) InternalImplementation() types.Map {
	return s.m
}

func (s User) Equals(other types.Value) bool {
	if other, ok := other.(User); ok {
		return s.Ref() == other.Ref()
	}
	return false
}

func (s User) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s User) Chunks() (futures []types.Future) {
	futures = append(futures, s.TypeRef().Chunks()...)
	futures = append(futures, s.m.Chunks()...)
	return
}

func (s User) Id() string {
	return s.m.Get(types.NewString("Id")).(types.String).String()
}

func (s User) SetId(val string) User {
	return User{s.m.Set(types.NewString("Id"), types.NewString(val)), &ref.Ref{}}
}

func (s User) Name() string {
	return s.m.Get(types.NewString("Name")).(types.String).String()
}

func (s User) SetName(val string) User {
	return User{s.m.Set(types.NewString("Name"), types.NewString(val)), &ref.Ref{}}
}

func (s User) OAuthToken() string {
	return s.m.Get(types.NewString("OAuthToken")).(types.String).String()
}

func (s User) SetOAuthToken(val string) User {
	return User{s.m.Set(types.NewString("OAuthToken"), types.NewString(val)), &ref.Ref{}}
}

func (s User) OAuthSecret() string {
	return s.m.Get(types.NewString("OAuthSecret")).(types.String).String()
}

func (s User) SetOAuthSecret(val string) User {
	return User{s.m.Set(types.NewString("OAuthSecret"), types.NewString(val)), &ref.Ref{}}
}

func (s User) Albums() MapOfStringToAlbum {
	return s.m.Get(types.NewString("Albums")).(MapOfStringToAlbum)
}

func (s User) SetAlbums(val MapOfStringToAlbum) User {
	return User{s.m.Set(types.NewString("Albums"), val), &ref.Ref{}}
}

// Album

type Album struct {
	m   types.Map
	ref *ref.Ref
}

func NewAlbum() Album {
	return Album{types.NewMap(
		types.NewString("$type"), types.MakeTypeRef(__mainPackageInFile_types_CachedRef, 1),
		types.NewString("Id"), types.NewString(""),
		types.NewString("Title"), types.NewString(""),
		types.NewString("Photos"), NewSetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto(),
	), &ref.Ref{}}
}

var __typeRefForAlbum = types.MakeTypeRef(__mainPackageInFile_types_CachedRef, 1)

func (m Album) TypeRef() types.TypeRef {
	return __typeRefForAlbum
}

func init() {
	types.RegisterFromValFunction(__typeRefForAlbum, func(v types.Value) types.Value {
		return AlbumFromVal(v)
	})
}

func AlbumFromVal(val types.Value) Album {
	// TODO: Do we still need FromVal?
	if val, ok := val.(Album); ok {
		return val
	}
	// TODO: Validate here
	return Album{val.(types.Map), &ref.Ref{}}
}

func (s Album) InternalImplementation() types.Map {
	return s.m
}

func (s Album) Equals(other types.Value) bool {
	if other, ok := other.(Album); ok {
		return s.Ref() == other.Ref()
	}
	return false
}

func (s Album) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s Album) Chunks() (futures []types.Future) {
	futures = append(futures, s.TypeRef().Chunks()...)
	futures = append(futures, s.m.Chunks()...)
	return
}

func (s Album) Id() string {
	return s.m.Get(types.NewString("Id")).(types.String).String()
}

func (s Album) SetId(val string) Album {
	return Album{s.m.Set(types.NewString("Id"), types.NewString(val)), &ref.Ref{}}
}

func (s Album) Title() string {
	return s.m.Get(types.NewString("Title")).(types.String).String()
}

func (s Album) SetTitle(val string) Album {
	return Album{s.m.Set(types.NewString("Title"), types.NewString(val)), &ref.Ref{}}
}

func (s Album) Photos() SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto {
	return s.m.Get(types.NewString("Photos")).(SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto)
}

func (s Album) SetPhotos(val SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto) Album {
	return Album{s.m.Set(types.NewString("Photos"), val), &ref.Ref{}}
}

// MapOfStringToAlbum

type MapOfStringToAlbum struct {
	m   types.Map
	ref *ref.Ref
}

func NewMapOfStringToAlbum() MapOfStringToAlbum {
	return MapOfStringToAlbum{types.NewMap(), &ref.Ref{}}
}

func MapOfStringToAlbumFromVal(val types.Value) MapOfStringToAlbum {
	// TODO: Do we still need FromVal?
	if val, ok := val.(MapOfStringToAlbum); ok {
		return val
	}
	// TODO: Validate here
	return MapOfStringToAlbum{val.(types.Map), &ref.Ref{}}
}

func (m MapOfStringToAlbum) InternalImplementation() types.Map {
	return m.m
}

func (m MapOfStringToAlbum) Equals(other types.Value) bool {
	if other, ok := other.(MapOfStringToAlbum); ok {
		return m.Ref() == other.Ref()
	}
	return false
}

func (m MapOfStringToAlbum) Ref() ref.Ref {
	return types.EnsureRef(m.ref, m)
}

func (m MapOfStringToAlbum) Chunks() (futures []types.Future) {
	futures = append(futures, m.TypeRef().Chunks()...)
	futures = append(futures, m.m.Chunks()...)
	return
}

// A Noms Value that describes MapOfStringToAlbum.
var __typeRefForMapOfStringToAlbum types.TypeRef

func (m MapOfStringToAlbum) TypeRef() types.TypeRef {
	return __typeRefForMapOfStringToAlbum
}

func init() {
	__typeRefForMapOfStringToAlbum = types.MakeCompoundTypeRef("", types.MapKind, types.MakePrimitiveTypeRef(types.StringKind), types.MakeTypeRef(__mainPackageInFile_types_CachedRef, 1))
	types.RegisterFromValFunction(__typeRefForMapOfStringToAlbum, func(v types.Value) types.Value {
		return MapOfStringToAlbumFromVal(v)
	})
}

func (m MapOfStringToAlbum) Empty() bool {
	return m.m.Empty()
}

func (m MapOfStringToAlbum) Len() uint64 {
	return m.m.Len()
}

func (m MapOfStringToAlbum) Has(p string) bool {
	return m.m.Has(types.NewString(p))
}

func (m MapOfStringToAlbum) Get(p string) Album {
	return m.m.Get(types.NewString(p)).(Album)
}

func (m MapOfStringToAlbum) MaybeGet(p string) (Album, bool) {
	v, ok := m.m.MaybeGet(types.NewString(p))
	if !ok {
		return NewAlbum(), false
	}
	return v.(Album), ok
}

func (m MapOfStringToAlbum) Set(k string, v Album) MapOfStringToAlbum {
	return MapOfStringToAlbum{m.m.Set(types.NewString(k), v), &ref.Ref{}}
}

// TODO: Implement SetM?

func (m MapOfStringToAlbum) Remove(p string) MapOfStringToAlbum {
	return MapOfStringToAlbum{m.m.Remove(types.NewString(p)), &ref.Ref{}}
}

type MapOfStringToAlbumIterCallback func(k string, v Album) (stop bool)

func (m MapOfStringToAlbum) Iter(cb MapOfStringToAlbumIterCallback) {
	m.m.Iter(func(k, v types.Value) bool {
		return cb(k.(types.String).String(), v.(Album))
	})
}

type MapOfStringToAlbumIterAllCallback func(k string, v Album)

func (m MapOfStringToAlbum) IterAll(cb MapOfStringToAlbumIterAllCallback) {
	m.m.IterAll(func(k, v types.Value) {
		cb(k.(types.String).String(), v.(Album))
	})
}

type MapOfStringToAlbumFilterCallback func(k string, v Album) (keep bool)

func (m MapOfStringToAlbum) Filter(cb MapOfStringToAlbumFilterCallback) MapOfStringToAlbum {
	nm := NewMapOfStringToAlbum()
	m.IterAll(func(k string, v Album) {
		if cb(k, v) {
			nm = nm.Set(k, v)
		}
	})
	return nm
}

// SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto

type SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto struct {
	s   types.Set
	ref *ref.Ref
}

func NewSetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto() SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto {
	return SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto{types.NewSet(), &ref.Ref{}}
}

func SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhotoFromVal(val types.Value) SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto {
	// TODO: Do we still need FromVal?
	if val, ok := val.(SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto); ok {
		return val
	}
	return SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto{val.(types.Set), &ref.Ref{}}
}

func (s SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto) InternalImplementation() types.Set {
	return s.s
}

func (s SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto) Equals(other types.Value) bool {
	if other, ok := other.(SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto); ok {
		return s.Ref() == other.Ref()
	}
	return false
}

func (s SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto) Chunks() (futures []types.Future) {
	futures = append(futures, s.TypeRef().Chunks()...)
	futures = append(futures, s.s.Chunks()...)
	return
}

// A Noms Value that describes SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto.
var __typeRefForSetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto types.TypeRef

func (m SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto) TypeRef() types.TypeRef {
	return __typeRefForSetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto
}

func init() {
	__typeRefForSetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto = types.MakeCompoundTypeRef("", types.SetKind, types.MakeTypeRef(ref.Parse("sha1-b525f9bca5e451c21dd9af564f0960045fbaa304"), 0))
	types.RegisterFromValFunction(__typeRefForSetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto, func(v types.Value) types.Value {
		return SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhotoFromVal(v)
	})
}

func (s SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto) Empty() bool {
	return s.s.Empty()
}

func (s SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto) Len() uint64 {
	return s.s.Len()
}

func (s SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto) Has(p sha1_b525f9bca5e451c21dd9af564f0960045fbaa304.RemotePhoto) bool {
	return s.s.Has(p)
}

type SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhotoIterCallback func(p sha1_b525f9bca5e451c21dd9af564f0960045fbaa304.RemotePhoto) (stop bool)

func (s SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto) Iter(cb SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhotoIterCallback) {
	s.s.Iter(func(v types.Value) bool {
		return cb(v.(sha1_b525f9bca5e451c21dd9af564f0960045fbaa304.RemotePhoto))
	})
}

type SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhotoIterAllCallback func(p sha1_b525f9bca5e451c21dd9af564f0960045fbaa304.RemotePhoto)

func (s SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto) IterAll(cb SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhotoIterAllCallback) {
	s.s.IterAll(func(v types.Value) {
		cb(v.(sha1_b525f9bca5e451c21dd9af564f0960045fbaa304.RemotePhoto))
	})
}

type SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhotoFilterCallback func(p sha1_b525f9bca5e451c21dd9af564f0960045fbaa304.RemotePhoto) (keep bool)

func (s SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto) Filter(cb SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhotoFilterCallback) SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto {
	ns := NewSetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto()
	s.IterAll(func(v sha1_b525f9bca5e451c21dd9af564f0960045fbaa304.RemotePhoto) {
		if cb(v) {
			ns = ns.Insert(v)
		}
	})
	return ns
}

func (s SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto) Insert(p ...sha1_b525f9bca5e451c21dd9af564f0960045fbaa304.RemotePhoto) SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto {
	return SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto{s.s.Insert(s.fromElemSlice(p)...), &ref.Ref{}}
}

func (s SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto) Remove(p ...sha1_b525f9bca5e451c21dd9af564f0960045fbaa304.RemotePhoto) SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto {
	return SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto{s.s.Remove(s.fromElemSlice(p)...), &ref.Ref{}}
}

func (s SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto) Union(others ...SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto) SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto {
	return SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto{s.s.Union(s.fromStructSlice(others)...), &ref.Ref{}}
}

func (s SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto) Subtract(others ...SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto) SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto {
	return SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto{s.s.Subtract(s.fromStructSlice(others)...), &ref.Ref{}}
}

func (s SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto) Any() sha1_b525f9bca5e451c21dd9af564f0960045fbaa304.RemotePhoto {
	return s.s.Any().(sha1_b525f9bca5e451c21dd9af564f0960045fbaa304.RemotePhoto)
}

func (s SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto) fromStructSlice(p []SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto) []types.Set {
	r := make([]types.Set, len(p))
	for i, v := range p {
		r[i] = v.s
	}
	return r
}

func (s SetOfsha1_b525f9bca5e451c21dd9af564f0960045fbaa304_RemotePhoto) fromElemSlice(p []sha1_b525f9bca5e451c21dd9af564f0960045fbaa304.RemotePhoto) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v
	}
	return r
}
