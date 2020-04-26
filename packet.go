package mcproto

import (
	"bytes"
	"errors"
	"fmt"
	enc "github.com/Raqbit/mcproto/encoding"
	"io"
	"math"
)

const MaxPacketLength = 1048576

var (
	ErrInvalidPacketLength = errors.New("packet has a malformed length")
)

type (
	// A packet type
	Packet interface {
		fmt.Stringer
		Info() PacketInfo
		Marshal(buf PacketWriter) error
		Unmarshal(buf PacketReader) error
	}

	// A reader with functions for reading Minecraft protocol types
	PacketReader interface {
		io.Reader
		ReadBytes(maxLength int64) (PacketReader, error)
		ReadResourceLocation() (*ResourceLocation, error)
		ReadMaxString() (string, error)
		ReadString(maxLength int32) (string, error)
		ReadFloat() (float32, error)
		ReadDouble() (float64, error)
		ReadLong() (int64, error)
		ReadVarInt() (int32, error)
		ReadInt() (int32, error)
		ReadUnsignedShort() (uint16, error)
		ReadByte() (int8, error)
		ReadUnsignedByte() (uint8, error)
		ReadBool() (bool, error)
	}

	// A writer with functions for writing Minecraft protocol types
	PacketWriter interface {
		io.Writer
		WriteBytes(data PacketReader) error
		WriteResourceLocation(value *ResourceLocation) error
		WriteString(value string) error
		WriteFloat(value float32) error
		WriteDouble(value float64) error
		WriteLong(value int64) error
		WriteVarInt(value int32) error
		WriteInt(value int32) error
		WriteUnsignedShort(value uint16) error
		WriteByte(value int8) error
		WriteUnsignedByte(value uint8) error
		WriteBool(value bool) error
	}

	// The info which identifies a packet
	PacketInfo struct {
		ID              int32
		Direction       PacketDirection
		ConnectionState ConnectionState
	}
)

// A buffer containing packet data
type PacketBuffer struct {
	r io.Reader
	w io.Writer
}

func (b *PacketBuffer) Write(p []byte) (n int, err error) {
	return b.w.Write(p)
}

func (b *PacketBuffer) Read(p []byte) (n int, err error) {
	return b.r.Read(p)
}

func NewPacketBuffer(rw io.ReadWriter) *PacketBuffer {
	return &PacketBuffer{r: rw, w: rw}
}

func NewPacketReader(r io.Reader) *PacketBuffer {
	return &PacketBuffer{r: r}
}

func (b *PacketBuffer) WriteVarInt(value int32) error {
	_, err := b.w.Write(enc.WriteVarInt(value))
	return err
}

func (b *PacketBuffer) WriteInt(value int32) error {
	_, err := b.w.Write(enc.WriteInt(value))
	return err
}

func (b *PacketBuffer) WriteString(value string) error {
	_, err := b.w.Write(enc.WriteString(value))
	return err
}

func (b *PacketBuffer) WriteUnsignedShort(value uint16) error {
	_, err := b.w.Write(enc.WriteUnsignedShort(value))
	return err
}

func (b *PacketBuffer) WriteLong(value int64) error {
	_, err := b.w.Write(enc.WriteLong(value))
	return err
}

func (b *PacketBuffer) WriteFloat(value float32) error {
	_, err := b.w.Write(enc.WriteFloat(value))
	return err
}

func (b *PacketBuffer) WriteDouble(value float64) error {
	_, err := b.w.Write(enc.WriteDouble(value))
	return err
}

func (b *PacketBuffer) WriteByte(value int8) error {
	_, err := b.w.Write(enc.WriteByte(value))
	return err
}

func (b *PacketBuffer) WriteBool(value bool) error {
	_, err := b.w.Write(enc.WriteBool(value))
	return err
}

func (b *PacketBuffer) WriteUnsignedByte(value uint8) error {
	_, err := b.w.Write(enc.WriteUnsignedByte(value))
	return err
}

func (b *PacketBuffer) WriteResourceLocation(value *ResourceLocation) error {
	_, err := b.w.Write(enc.WriteString(value.String()))
	return err
}

func (b *PacketBuffer) WriteBytes(data PacketReader) error {
	// Copy from the passed buffer into the packet buffer to send
	_, err := io.Copy(b.w, data)
	return err
}

func (b *PacketBuffer) ReadVarInt() (int32, error) {
	return enc.ReadVarInt(b.r)
}

func (b *PacketBuffer) ReadUnsignedShort() (uint16, error) {
	return enc.ReadUnsignedShort(b.r)
}

func (b *PacketBuffer) ReadMaxString() (string, error) {
	return enc.ReadString(b.r, math.MaxInt16)
}

func (b *PacketBuffer) ReadString(maxLength int32) (string, error) {
	return enc.ReadString(b.r, maxLength)
}

func (b *PacketBuffer) ReadLong() (int64, error) {
	return enc.ReadLong(b.r)
}

func (b *PacketBuffer) ReadFloat() (float32, error) {
	return enc.ReadFloat(b.r)
}

func (b *PacketBuffer) ReadDouble() (float64, error) {
	return enc.ReadDouble(b.r)
}

func (b *PacketBuffer) ReadInt() (int32, error) {
	return enc.ReadInt(b.r)
}

func (b *PacketBuffer) ReadUnsignedByte() (uint8, error) {
	return enc.ReadUnsignedByte(b.r)
}

func (b *PacketBuffer) ReadBool() (bool, error) {
	return enc.ReadBool(b.r)
}

func (b *PacketBuffer) ReadResourceLocation() (*ResourceLocation, error) {
	str, err := b.ReadMaxString()

	if err != nil {
		return nil, err
	}

	return NewResourceLocationFromString(str), nil
}

func (b *PacketBuffer) ReadBytes(maxLength int64) (PacketReader, error) {
	var buff bytes.Buffer
	_, _ = io.CopyN(&buff, b.r, maxLength)

	nextByte := make([]byte, 1)
	nRead, _ := buff.Read(nextByte)
	if nRead > 0 {
		// There's too much data
		return nil, fmt.Errorf("payload may not be larger than %d bytes", maxLength)
	}

	return NewPacketBuffer(&buff), nil
}

func (b *PacketBuffer) ReadByte() (int8, error) {
	return enc.ReadByte(b.r)
}

// The direction of a packet
type PacketDirection byte

func (p PacketDirection) String() string {
	names := []string{
		"client-bound",
		"server-bound",
	}

	return names[p]
}

const (
	ClientBound PacketDirection = iota
	ServerBound
)
