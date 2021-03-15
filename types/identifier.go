package types

import (
	"fmt"
	enc "github.com/Raqbit/mcproto/encoding"
	"io"
	"strings"
)

const delimiter = ":"

type Identifier struct {
	Namespace string
	Path      string
}

func (i *Identifier) Write(w io.Writer) error {
	str := enc.String(i.String())
	return str.Write(w)
}

func (i *Identifier) Read(r io.Reader) error {
	var str enc.String

	if err := str.Read(r); err != nil {
		return err
	}

	*i = NewIdentifierFromString(string(str))
	return nil
}

func NewIdentifierFromString(resourceName string) Identifier {
	parts := strings.Split(resourceName, delimiter)
	return NewIdentifier(parts[0], parts[1])
}

func NewIdentifier(namespace string, path string) Identifier {
	if namespace == "" {
		namespace = "minecraft"
	}
	return Identifier{Namespace: namespace, Path: path}
}

func (r Identifier) String() string {
	return fmt.Sprintf("%s%s%s", r.Namespace, delimiter, r.Path)
}
