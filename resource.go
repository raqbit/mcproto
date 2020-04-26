package mcproto

import (
	"fmt"
	"strings"
)

const delimiter = ":"

type ResourceLocation struct {
	Namespace string
	Path      string
}

func NewResourceLocationFromString(resourceName string) *ResourceLocation {
	parts := decompose(resourceName)
	return NewResourceLocation(parts[0], parts[1])
}

func NewResourceLocation(namespace string, path string) *ResourceLocation {
	if namespace == "" {
		namespace = "minecraft"
	}
	return &ResourceLocation{Namespace: namespace, Path: path}
}

func (r ResourceLocation) String() string {
	return fmt.Sprintf("%s%s%s", r.Namespace, delimiter, r.Path)
}

func decompose(resourceName string) []string {
	return strings.Split(resourceName, delimiter)
}
