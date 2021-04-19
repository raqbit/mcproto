package channel

import (
	enc "github.com/Raqbit/mcproto/encoding"
	"github.com/Raqbit/mcproto/types"
)

//go:generate go run ../../tools/genpacket/genpacket.go -packet=BrandChannel -output=brand_gen.go

var BrandChannelID = types.NewIdentifier("minecraft", "brand")

// BrandChannel informs the client of the server brand
type BrandChannel struct {
	Brand enc.String
}
