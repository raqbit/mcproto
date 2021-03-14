package channel

import (
	"github.com/Raqbit/mcproto"
	enc "github.com/Raqbit/mcproto/encoding"
)

//go:generate go run ../../tools/genpacket/genpacket.go -packet=BrandChannel -output=brand_gen.go

var BrandChannelID = mcproto.NewIdentifier("minecraft", "brand")

type BrandChannel struct {
	Brand enc.String
}
