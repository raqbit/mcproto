package mcproto

//go:generate stringer -type=Dimension -output dimension_string.go -linecomment
type Dimension int32

const (
	DimensionNether    Dimension = -1 // Nether
	DimensionOverworld Dimension = 0  // Overworld
	DimensionEnd       Dimension = 1  // End
)
