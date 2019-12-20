package encoding

import (
	"io"
)

type Encodable interface {
	Encode(w io.Writer) error
}

type Decodable interface {
	Decode(r io.Reader) error
}
