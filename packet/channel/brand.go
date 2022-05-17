package channel

import (
	enc "github.com/Raqbit/mcproto/encoding"
	"github.com/Raqbit/mcproto/game"
)

//go:generate go run ../../tools/genpacket -packet=BrandChannel -output=brand_gen.go

var BrandChannelID = game.NewIdentifier("minecraft", "brand")

// BrandChannel informs the client of the server brand
type BrandChannel struct {
	Brand enc.String
}
