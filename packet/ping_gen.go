// Code generated by "genpacket -packet=PingPacket -output=ping_gen.go"; DO NOT EDIT.

package packet

import "io"

func (p *PingPacket) Write(w io.Writer) error {
	var err error
	if err = p.Payload.Write(w); err != nil {
		return err
	}
	return nil
}
func (p *PingPacket) Read(r io.Reader) error {
	var err error
	if err = p.Payload.Read(r); err != nil {
		return err
	}
	return nil
}