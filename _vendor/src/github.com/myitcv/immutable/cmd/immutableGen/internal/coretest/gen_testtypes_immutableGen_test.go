// Code generated by immutableGen. DO NOT EDIT.

// My favourite license

package coretest

//go:generate echo "hello world"
//immutableVet:skipFile

import (
	"github.com/myitcv/immutable"
)

// a comment about MyMap
//
// MyTestMap is an immutable type and has the following template:
//
// 	map[string]int
//
type MyTestMap struct {
	theMap  map[string]int
	mutable bool
	__tmpl  _Imm_MyTestMap
}

var _ immutable.Immutable = new(MyTestMap)
var _ = new(MyTestMap).__tmpl

func NewMyTestMap(inits ...func(m *MyTestMap)) *MyTestMap {
	res := NewMyTestMapCap(0)
	if len(inits) == 0 {
		return res
	}

	return res.WithMutable(func(m *MyTestMap) {
		for _, i := range inits {
			i(m)
		}
	})
}

func NewMyTestMapCap(l int) *MyTestMap {
	return &MyTestMap{
		theMap: make(map[string]int, l),
	}
}

func (m *MyTestMap) Mutable() bool {
	return m.mutable
}

func (m *MyTestMap) Len() int {
	if m == nil {
		return 0
	}

	return len(m.theMap)
}

func (m *MyTestMap) Get(k string) (int, bool) {
	v, ok := m.theMap[k]
	return v, ok
}

func (m *MyTestMap) AsMutable() *MyTestMap {
	if m == nil {
		return nil
	}

	if m.Mutable() {
		return m
	}

	res := m.dup()
	res.mutable = true

	return res
}

func (m *MyTestMap) dup() *MyTestMap {
	resMap := make(map[string]int, len(m.theMap))

	for k := range m.theMap {
		resMap[k] = m.theMap[k]
	}

	res := &MyTestMap{
		theMap: resMap,
	}

	return res
}

func (m *MyTestMap) AsImmutable(v *MyTestMap) *MyTestMap {
	if m == nil {
		return nil
	}

	if v == m {
		return m
	}

	m.mutable = false
	return m
}

func (m *MyTestMap) Range() map[string]int {
	if m == nil {
		return nil
	}

	return m.theMap
}

func (m *MyTestMap) WithMutable(f func(mi *MyTestMap)) *MyTestMap {
	res := m.AsMutable()
	f(res)
	res = res.AsImmutable(m)

	return res
}

func (m *MyTestMap) WithImmutable(f func(mi *MyTestMap)) *MyTestMap {
	prev := m.mutable
	m.mutable = false
	f(m)
	m.mutable = prev

	return m
}

func (m *MyTestMap) Set(k string, v int) *MyTestMap {
	if m.mutable {
		m.theMap[k] = v
		return m
	}

	res := m.dup()
	res.theMap[k] = v

	return res
}

func (m *MyTestMap) Del(k string) *MyTestMap {
	if _, ok := m.theMap[k]; !ok {
		return m
	}

	if m.mutable {
		delete(m.theMap, k)
		return m
	}

	res := m.dup()
	delete(res.theMap, k)

	return res
}
func (s *MyTestMap) IsDeeplyNonMutable(seen map[interface{}]bool) bool {
	if s == nil {
		return true
	}

	if s.Mutable() {
		return false
	}
	return true
}

// a comment about Slice
//
// MyTestSlice is an immutable type and has the following template:
//
// 	[]*string
//
type MyTestSlice struct {
	theSlice []*string
	mutable  bool
	__tmpl   _Imm_MyTestSlice
}

var _ immutable.Immutable = new(MyTestSlice)
var _ = new(MyTestSlice).__tmpl

func NewMyTestSlice(s ...*string) *MyTestSlice {
	c := make([]*string, len(s))
	copy(c, s)

	return &MyTestSlice{
		theSlice: c,
	}
}

func NewMyTestSliceLen(l int) *MyTestSlice {
	c := make([]*string, l)

	return &MyTestSlice{
		theSlice: c,
	}
}

func (m *MyTestSlice) Mutable() bool {
	return m.mutable
}

func (m *MyTestSlice) Len() int {
	if m == nil {
		return 0
	}

	return len(m.theSlice)
}

func (m *MyTestSlice) Get(i int) *string {
	return m.theSlice[i]
}

func (m *MyTestSlice) AsMutable() *MyTestSlice {
	if m == nil {
		return nil
	}

	if m.Mutable() {
		return m
	}

	res := m.dup()
	res.mutable = true

	return res
}

func (m *MyTestSlice) dup() *MyTestSlice {
	resSlice := make([]*string, len(m.theSlice))

	for i := range m.theSlice {
		resSlice[i] = m.theSlice[i]
	}

	res := &MyTestSlice{
		theSlice: resSlice,
	}

	return res
}

func (m *MyTestSlice) AsImmutable(v *MyTestSlice) *MyTestSlice {
	if m == nil {
		return nil
	}

	if v == m {
		return m
	}

	m.mutable = false
	return m
}

func (m *MyTestSlice) Range() []*string {
	if m == nil {
		return nil
	}

	return m.theSlice
}

func (m *MyTestSlice) WithMutable(f func(mi *MyTestSlice)) *MyTestSlice {
	res := m.AsMutable()
	f(res)
	res = res.AsImmutable(m)

	return res
}

func (m *MyTestSlice) WithImmutable(f func(mi *MyTestSlice)) *MyTestSlice {
	prev := m.mutable
	m.mutable = false
	f(m)
	m.mutable = prev

	return m
}

func (m *MyTestSlice) Set(i int, v *string) *MyTestSlice {
	if m.mutable {
		m.theSlice[i] = v
		return m
	}

	res := m.dup()
	res.theSlice[i] = v

	return res
}

func (m *MyTestSlice) Append(v ...*string) *MyTestSlice {
	if m.mutable {
		m.theSlice = append(m.theSlice, v...)
		return m
	}

	res := m.dup()
	res.theSlice = append(res.theSlice, v...)

	return res
}
func (s *MyTestSlice) IsDeeplyNonMutable(seen map[interface{}]bool) bool {
	if s == nil {
		return true
	}

	if s.Mutable() {
		return false
	}
	return true
}

// a comment about myStruct
//
// MyTestStruct is an immutable type and has the following template:
//
// 	struct {
// 		Name, surname	string
// 		age		int
//
// 		fieldWithoutTag	bool
// 	}
//
type MyTestStruct struct {
	//somethingspecial

	_Name, _surname  string `tag:"value"`
	_age             int    `tag:"age"`
	_fieldWithoutTag bool

	mutable bool
	__tmpl  _Imm_MyTestStruct
}

var _ immutable.Immutable = new(MyTestStruct)
var _ = new(MyTestStruct).__tmpl

func (s *MyTestStruct) AsMutable() *MyTestStruct {
	if s.Mutable() {
		return s
	}

	res := *s
	res.mutable = true
	return &res
}

func (s *MyTestStruct) AsImmutable(v *MyTestStruct) *MyTestStruct {
	if s == nil {
		return nil
	}

	if s == v {
		return s
	}

	s.mutable = false
	return s
}

func (s *MyTestStruct) Mutable() bool {
	return s.mutable
}

func (s *MyTestStruct) WithMutable(f func(si *MyTestStruct)) *MyTestStruct {
	res := s.AsMutable()
	f(res)
	res = res.AsImmutable(s)

	return res
}

func (s *MyTestStruct) WithImmutable(f func(si *MyTestStruct)) *MyTestStruct {
	prev := s.mutable
	s.mutable = false
	f(s)
	s.mutable = prev

	return s
}
func (s *MyTestStruct) IsDeeplyNonMutable(seen map[interface{}]bool) bool {
	if s == nil {
		return true
	}

	if s.Mutable() {
		return false
	}

	if seen == nil {
		return s.IsDeeplyNonMutable(make(map[interface{}]bool))
	}

	if seen[s] {
		return true
	}

	seen[s] = true
	return true
}

// my field comment
//somethingspecial
/*

	Heelo

*/
func (s *MyTestStruct) Name() string {
	return s._Name
}

// SetName is the setter for Name()
func (s *MyTestStruct) SetName(n string) *MyTestStruct {
	if s.mutable {
		s._Name = n
		return s
	}

	res := *s
	res._Name = n
	return &res
}

// my field comment
//somethingspecial
/*

	Heelo

*/
func (s *MyTestStruct) surname() string {
	return s._surname
}

// setSurname is the setter for Surname()
func (s *MyTestStruct) setSurname(n string) *MyTestStruct {
	if s.mutable {
		s._surname = n
		return s
	}

	res := *s
	res._surname = n
	return &res
}
func (s *MyTestStruct) age() int {
	return s._age
}

// setAge is the setter for Age()
func (s *MyTestStruct) setAge(n int) *MyTestStruct {
	if s.mutable {
		s._age = n
		return s
	}

	res := *s
	res._age = n
	return &res
}
func (s *MyTestStruct) fieldWithoutTag() bool {
	return s._fieldWithoutTag
}

// setFieldWithoutTag is the setter for FieldWithoutTag()
func (s *MyTestStruct) setFieldWithoutTag(n bool) *MyTestStruct {
	if s.mutable {
		s._fieldWithoutTag = n
		return s
	}

	res := *s
	res._fieldWithoutTag = n
	return &res
}
