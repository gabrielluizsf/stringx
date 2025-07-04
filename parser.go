package stringx

import (
	"strconv"
)

// Parser is a string parser.
type Parser struct {
	s       String
	bitSize int
}

// NewParser creates a [github.com/i9si-sistemas/stringx.Parser].
func NewParser(s string) *Parser {
	return &Parser{String(s), 64}
}

// Bool parses a bool-comparison or a plain bool literal.
func (p *Parser) Bool() (bool, error) {
	s := p.s
	type cmpOp struct {
		token   string
		compare func(a, b string) bool
	}
	compare := func(
		a, b string,
		compareN func(a, b float64) bool,
		compareS func(a, b string) bool,
	) bool {
		aFloat, aErr := NewParser(a).Float()
		bFloat, bErr := NewParser(b).Float()
		if aErr == nil && bErr == nil {
			return compareN(aFloat, bFloat)
		}
		return compareS(a, b)
	}

	ops := []cmpOp{
		{"!=", func(a, b string) bool { return a != b }},
		{"==", func(a, b string) bool { return a == b }},
		{">", func(a, b string) bool {
			return compare(
				a, b,
				func(a, b float64) bool {
					return a > b
				},
				func(a, b string) bool {
					return a > b
				},
			)
		}},
		{"<", func(a, b string) bool {
			return compare(
				a, b,
				func(a, b float64) bool {
					return a < b
				},
				func(a, b string) bool {
					return a < b
				},
			)
		}},
	}

	for _, op := range ops {
		if s.Includes(op.token) {
			parts := s.SplitN(op.token, 2)
			if len(parts) != 2 {
				return false, ErrInvalidExpression
			}
			clean := func(x String) string {
				return x.Trim(Space.String()).Trim(`"`).Trim("'").String()
			}
			left, right := clean(String(parts[0])), clean(String(parts[1]))
			return op.compare(left, right), nil
		}
	}
	return strconv.ParseBool(s.String())
}

// Int parses an integer.
func (p *Parser) Int() (n int64, err error) {
	s := p.s
	parseInt := func(s string) (int64, error) {
		return strconv.ParseInt(s, 10, p.bitSize)
	}
	ops := []mathOp[int64]{
		{"+", func(a, b string) (int64, error) {
			return execute(a, b, parseInt, func(a, b int64) (int64, error) {
				return a + b, nil
			})
		}},
		{"-", func(a, b string) (int64, error) {
			return execute(a, b, parseInt, func(a, b int64) (int64, error) {
				return a - b, nil
			})
		}},
		{"*", func(a, b string) (int64, error) {
			return execute(a, b, parseInt, func(a, b int64) (int64, error) {
				return a * b, nil
			})
		}},
		{"/", func(a, b string) (int64, error) {
			return execute(a, b, parseInt, func(a, b int64) (int64, error) {
				return a / b, nil
			})
		}},
		{"%", func(a, b string) (int64, error) {
			return execute(a, b, parseInt, func(a, b int64) (int64, error) {
				return a % b, nil
			})
		}},
	}
	for _, op := range ops {
		if v, err := op.run(s); err == nil {
			return v, nil
		}
	}
	return parseInt(p.s.String())
}

// Float parses a float.
func (p *Parser) Float() (n float64, err error) {
	s := p.s
	parseFloat := func(s string) (float64, error) {
		return strconv.ParseFloat(s, p.bitSize)
	}
	ops := []mathOp[float64]{
		{"+", func(a, b string) (float64, error) {
			return execute(a, b, parseFloat, func(a, b float64) (float64, error) {
				return a + b, nil
			})
		}},
		{"-", func(a, b string) (float64, error) {
			return execute(a, b, parseFloat, func(a, b float64) (float64, error) {
				return a - b, nil
			})
		}},
		{"*", func(a, b string) (float64, error) {
			return execute(a, b, parseFloat, func(a, b float64) (float64, error) {
				return a * b, nil
			})
		}},
		{"/", func(a, b string) (float64, error) {
			return execute(a, b, parseFloat, func(a, b float64) (float64, error) {
				return a / b, nil
			})
		}},
		{"%", func(a, b string) (float64, error) {
			return execute(a, b, parseFloat, func(a, b float64) (float64, error) {
				return float64(int64(a) % int64(b)), nil
			})
		}},
	}
	for _, op := range ops {
		if v, err := op.run(s); err == nil {
			return v, nil
		}
	}
	for _, op := range ops {
		if v, err := op.run(s); err == nil {
			return v, nil
		}
	}
	return parseFloat(p.s.String())
}

// MathError is a math error.
type MathError string

func (e MathError) Error() string {
	return String("Math Error:").Concat(NewLine, String(e)).String()
}

// ErrInvalidExpression is returned when the expression is invalid.
var ErrInvalidExpression = MathError("Invalid expression")

type mathOp[T int64 | float64] struct {
	token   string
	execute func(a, b string) (T, error)
}

func (op mathOp[T]) run(s String) (T, error) {
	if s.Includes(op.token) {
		v := s.Replace(Space.String(), Empty.String()).Split(op.token)
		if len(v) != 2 {
			return 0, ErrInvalidExpression
		}
		a, b := v[0], v[1]
		n, err := op.execute(a, b)
		if err != nil {
			return 0, err
		}
		return n, nil
	}
	return 0, ErrInvalidExpression
}

func execute[T int64 | float64](
	a, b string,
	parse func(v string) (T, error),
	fn func(a, b T) (T, error),
) (T, error) {
	aInt, aErr := parse(a)
	if aErr != nil {
		return 0, ErrInvalidExpression
	}
	bInt, bErr := parse(b)
	if bErr != nil {
		return 0, ErrInvalidExpression
	}
	return fn(aInt, bInt)

}
