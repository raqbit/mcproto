package nbt

//go:generate stringer -type=TagType -output tags_string.go -linecomment

// TagType is the type of an NBT tag
type TagType byte

const (
	TagTypeEnd       TagType = 0  // End
	TagTypeByte      TagType = 1  // Byte
	TagTypeShort     TagType = 2  // Short
	TagTypeInt       TagType = 3  // Int
	TagTypeLong      TagType = 4  // Long
	TagTypeFloat     TagType = 5  // Float
	TagTypeDouble    TagType = 6  // Double
	TagTypeByteArray TagType = 7  // Byte Array
	TagTypeString    TagType = 8  // String
	TagTypeList      TagType = 9  // List
	TagTypeCompound  TagType = 10 // Compound
	TagTypeIntArray  TagType = 11 // Int Array
	TagTypeLongArray TagType = 12 // Long Array
)
