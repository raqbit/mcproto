package mcproto

import (
	"fmt"
	"strings"
)

const delimiter = ":"

type Identifier struct {
	Namespace string
	Path      string
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
