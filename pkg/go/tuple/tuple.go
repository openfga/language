package tuple

import (
	"errors"
	"io"
	"strings"
	"unicode"
	"unsafe"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
)

const (
	sizeType       int = 256
	sizeRelation   int = 64
	sizeIdentifier int = 256
)

var (
	colon = []byte{':'}
	pound = []byte{'#'}

	ErrInvalidValue  = errors.New("invalid value")
	ErrValueOverflow = errors.New("value overflow")
)

func disallowed(r rune) bool {
	switch {

	case r == unicode.ReplacementChar:
		return true

	case r == 0:
		return true

	case unicode.IsControl(r):
		return true

	case unicode.IsSpace(r):
		return true

	default:
		return false
	}
}

func symbolReserved(r rune) bool {
	switch {

	case disallowed(r):
		return true

	case r == ':':
		return true

	case r == '#':
		return true

	case r == '@':
		return true

	default:
		return false
	}
}

func identifierReserved(r rune) bool {
	return disallowed(r) || r == '#'
}

func validSymbol(s string) bool {
	return !strings.ContainsFunc(s, symbolReserved)
}

func validIdentifier(s string) bool {
	return !strings.ContainsFunc(s, identifierReserved)
}

func FromKey(t *Tuple, tk *openfgav1.TupleKey) error {
	err := ParseObject(&t.Object, tk.GetObject())
	if err != nil {
		return err
	}

	err = ParseRelation(&t.Relation, tk.GetRelation())
	if err != nil {
		return err
	}

	return ParseUser(&t.User, tk.GetUser())
}

func ParseIdentifier(i *Identifier, s string) error {
	if !validIdentifier(s) {
		return ErrInvalidValue
	}

	if len(s) > sizeRelation {
		return ErrValueOverflow
	}

	clear(i.data[:])
	written := copy(i.data[:], s)
	i.size = written
	return nil
}

func ParseObject(o *Object, s string) error {
	objectType, objectIdentifier, found := strings.Cut(s, ":")
	if !found {
		return ErrInvalidValue
	}

	err := ParseType(&o.Type, objectType)
	if err != nil {
		return err
	}

	err = ParseIdentifier(&o.Identifier, objectIdentifier)
	if err != nil {
		return err
	}
	return nil
}

func ParseRelation(r *Relation, s string) error {
	if !validSymbol(s) {
		return ErrInvalidValue
	}

	if len(s) > sizeRelation {
		return ErrValueOverflow
	}
	
	clear(r.data[:])
	written := copy(r.data[:], s)
	r.size = written
	return nil
}

func ParseTuple(t *Tuple, s string) error {
	object, rest, found := strings.Cut(s, "#")
	if !found {
		return ErrInvalidValue
	}

	relation, user, found := strings.Cut(rest, "@")
	if !found {
		return ErrInvalidValue
	}

	err := ParseObject(&t.Object, object)
	if err != nil {
		return err
	}

	err = ParseRelation(&t.Relation, relation)
	if err != nil {
		return err
	}

	return ParseUser(&t.User, user)
}

func ParseType(t *Type, s string) error {
	if !validSymbol(s) {
		return ErrInvalidValue
	}

	if len(s) > sizeRelation {
		return ErrValueOverflow
	}

	clear(t.data[:])
	written := copy(t.data[:], s)
	t.size = written
	return nil
}

func ParseUser(u *User, s string) error {
	userObject, userRelation, hasRelation := strings.Cut(s, "#")
	userType, userIdentifier, found := strings.Cut(userObject, ":")
	if !found {
		return ErrInvalidValue
	}

	err := ParseType(&u.Type, userType)
	if err != nil {
		return err
	}

	err = ParseIdentifier(&u.Identifier, userIdentifier)
	if err != nil {
		return err
	}

	if !hasRelation {
		return nil
	}
	return ParseRelation(&u.Relation, userRelation)
}

type Type struct {
	data [sizeType]byte
	size int
}

func (t *Type) Equals(t2 *Type) bool {
	if t.size != t2.size {
		return false
	}
	return bytes.Equal(t.data[:], t2.data[:])
}

func (t *Type) Len() int {
	return t.size
}

func (t *Type) String() string {
	return unsafe.String(unsafe.SliceData(t.data[:t.size]), t.size)
}

func (t *Type) WriteTo(w io.Writer) (int64, error) {
	if t.size == 0 {
		return 0, nil
	}
	written, err := w.Write(t.data[:t.size])
	return int64(written), err
}

type Relation struct {
	data [sizeRelation]byte
	size int
}

func (r *Relation) Equals(r2 *Relation) bool {
	if r.size != r2.size {
		return false
	}
	return bytes.Equal(r.data[:], r2.data[:])
}

func (r *Relation) Len() int {
	return r.size
}

func (r *Relation) String() string {
	return unsafe.String(unsafe.SliceData(r.data[:r.size]), r.size)
}

func (r *Relation) WriteTo(w io.Writer) (int64, error) {
	if r.size == 0 {
		return 0, nil
	}
	written, err := w.Write(r.data[:r.size])
	return int64(written), err
}

type Identifier struct {
	data [sizeIdentifier]byte
	size int
}

func (i *Identifier) Equals(i2 *Identifier) bool {
	if i.size != i2.size {
		return false
	}
	return bytes.Equal(i.data[:], i2.data[:])
}

func (i *Identifier) Len() int {
	return i.size
}

func (i *Identifier) String() string {
	return unsafe.String(unsafe.SliceData(i.data[:i.size]), i.size)
}

func (i *Identifier) WriteTo(w io.Writer) (int64, error) {
	if i.size == 0 {
		return 0, nil
	}
	written, err := w.Write(i.data[:i.size])
	return int64(written), err
}

type Object struct {
	Type       Type
	Identifier Identifier
}

func (o *Object) Len() int {
	return o.Type.Len() + o.Identifier.Len() + 1
}

func (o *Object) String() string {
	var b strings.Builder
	_, _ = o.WriteTo(&b)
	return b.String()
}

func (o *Object) WriteTo(w io.Writer) (int64, error) {
	var totalWritten int64

	written, err := o.Type.WriteTo(w)
	totalWritten += written
	if err != nil {
		return totalWritten, err
	}

	bw, err := w.Write(colon)
	totalWritten += int64(bw)
	if err != nil {
		return totalWritten, err
	}

	written, err = o.Identifier.WriteTo(w)
	totalWritten += written
	return totalWritten, err
}

type User struct {
	Type       Type
	Identifier Identifier
	Relation   Relation
}

func (u *User) Len() int {
	total := u.Type.Len() + u.Identifier.Len() + 1
	if u.Relation.Len() == 0 {
		return total
	}
	return total + u.Relation.Len() + 1
}

func (u *User) String() string {
	var b strings.Builder
	_, _ = u.WriteTo(&b)
	return b.String()
}

func (u *User) WriteTo(w io.Writer) (int64, error) {
	var totalWritten int64

	written, err := u.Type.WriteTo(w)
	totalWritten += written
	if err != nil {
		return totalWritten, err
	}

	bw, err := w.Write(colon)
	totalWritten += int64(bw)
	if err != nil {
		return totalWritten, err
	}

	written, err = u.Identifier.WriteTo(w)
	totalWritten += written
	if err != nil {
		return totalWritten, err
	}

	if u.Relation.Len() == 0 {
		return totalWritten, err
	}

	bw, err = w.Write(pound)
	totalWritten += int64(bw)
	if err != nil {
		return totalWritten, err
	}

	written, err = u.Relation.WriteTo(w)
	totalWritten += written
	return totalWritten, err
}

type Tuple struct {
	Object   Object
	Relation Relation
	User     User
}

func (t *Tuple) Len() int {
	return t.Object.Len() + t.Relation.Len() + t.User.Len() + 2
}

func (t *Tuple) String() string {
	var b strings.Builder
	_, _ = t.WriteTo(&b)
	return b.String()
}

func (t *Tuple) WriteTo(w io.Writer) (int64, error) {
	var totalWritten int64

	written, err := t.Object.WriteTo(w)
	totalWritten += written
	if err != nil {
		return totalWritten, err
	}

	written, err = t.Relation.WriteTo(w)
	totalWritten += written
	if err != nil {
		return totalWritten, err
	}

	written, err = t.User.WriteTo(w)
	totalWritten += written
	if err != nil {
		return totalWritten, err
	}
	return totalWritten, err
}
