package stringx

// ParseNumber returns a Number object.
func ParseNumber(s string) Number {
	return Number{NewParser(s)}
}

// Number is a string wrapper for number.
type Number struct {
	p *Parser
}

// String returns the number string.
func (s Number) String() string {
	return s.p.s.String()
}

// Int returns the number as int.
func (s Number) Int() int {
	n, _ := s.parseInt()
	return n
}

// Int64 returns the number as int64.
func (s Number) Int64() int64 {
	n, _ := s.parseInt()
	return int64(n)
}

// Int32 returns the number as int32.
func (s Number) Int32() int32 {
	n, _ := s.parseInt()
	return int32(n)
}

// Int16 returns the number as int16.
func (s Number) Int16() int16 {
	n, _ := s.parseInt()
	return int16(n)
}

// Int8 returns the number as int8.
func (s Number) Int8() int8 {
	n, _ := s.parseInt()
	return int8(n)
}

// Float returns the number as float64.
func (s Number) Float() float64 {
	n, _ := s.parseFloat()
	return n
}

// Float32 returns the number as float32.
func (s Number) Float32() float32 {
	n, _ := s.parseFloat()
	return float32(n)
}

// Uint returns the number as uint.
func (s Number) Uint() uint {
	n, _ := s.parseInt()
	return uint(n)
}

// Uint64 returns the number as uint64.
func (s Number) Uint64() uint64 {
	n, _ := s.parseInt()
	return uint64(n)
}

// Uint32 returns the number as uint32.
func (s Number) Uint32() uint32 {
	n, _ := s.parseInt()
	return uint32(n)
}

// Uint16 returns the number as uint16.
func (s Number) Uint16() uint16 {
	n, _ := s.parseInt()
	return uint16(n)
}

// Uint8 returns the number as uint8.
func (s Number) Uint8() uint8 {
	n, _ := s.parseInt()
	return uint8(n)
}

func (s Number) parseInt() (n int, err error) {
	v, err := s.p.Int()
	return int(v), err
}

func (s Number) parseFloat() (n float64, err error) {
	return s.p.Float()
}
