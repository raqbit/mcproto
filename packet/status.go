package packet

import enc "github.com/Raqbit/mcproto/encoding"

// https://wiki.vg/Server_List_Ping#Request
type Request struct{}

func (Request) ID() int {
	return 0x00
}

// https://wiki.vg/Server_List_Ping#Response
type Response struct {
	Json enc.String
}

func (Response) ID() int {
	return 0x00
}
