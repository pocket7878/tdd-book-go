package money

type Pair struct {
	from string
	to   string
}

func NewPair(from, to string) *Pair {
	return &Pair{from, to}
}

func (p *Pair) From() string {
	return p.from
}

func (p *Pair) To() string {
	return p.to
}
