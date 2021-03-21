package main

import (
	"fmt"
	"github.com/yuin/stagparser"
	"go/types"
	"strings"
)

type param struct {
	name  string
	value interface{}
}

type field struct {
	name   string
	params []param
}

type packet struct {
	name   string
	fields []field
}

func (p *packet) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("packet %s {\n", p.name))
	for _, fld := range p.fields {
		sb.WriteString(fmt.Sprintf("\t%s (", fld.name))
		for _, prm := range fld.params {
			sb.WriteString(fmt.Sprintf("%s=%s,", prm.name, prm.value))
		}
		sb.WriteString(")")
	}
	sb.WriteString("}\n")

	return sb.String()
}

func parsePacket(packetType *types.Struct, structName string) (packet, error) {
	fields := make([]field, packetType.NumFields())

	for i := 0; i < packetType.NumFields(); i++ {
		f := packetType.Field(i)
		t := packetType.Tag(i)

		// TODO: stagparser expects us to give it one "value" of a struct tag (`key:"value"`)
		// TODO: this is hard to do as the parser for it is only really present in reflect, not AST.
		// TODO: a custom tag parser will probably have to be implemented
		defs, err := stagparser.ParseTag(t, structName)

		if err != nil {
			return packet{}, fmt.Errorf("could not parse struct tag for field %s: %w", f.Name(), err)
		}

		params := make([]param, len(defs))

		for j, def := range defs {
			for name, value := range def.Attributes() {
				params[j] = param{
					name:  name,
					value: value,
				}
			}
		}

		fields[i] = field{
			name:   f.Name(),
			params: make([]param, 0),
		}

	}

	return packet{
		name:   structName,
		fields: fields,
	}, nil
}
