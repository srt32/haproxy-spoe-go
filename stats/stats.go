package stats

// Statter is an type for emitting metrics
type Statter interface {
	Count(name string, amount int, tags Tags)
}

// Tags is a map of string tags to include with metrics
type Tags map[string]string

// NullStatter is a no-op Statter
type NullStatter struct{}

// Count is a no-op
func (n *NullStatter) Count(name string, amount int, tags Tags) {}

// NewNullStatter returns a new NullStatter
func NewNullStatter() *NullStatter {
	return &NullStatter{}
}
