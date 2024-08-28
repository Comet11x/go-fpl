package core

type pair[F any, S any] struct {
	f F
	s S
}

func (p *pair[F, S]) First() F {
	return p.f
}

func (p *pair[F, S]) Second() S {
	return p.s
}

func (p *pair[F, S]) ToTuple() (F, S) {
	return p.f, p.s
}

func (p *pair[F, S]) SwapFirst(v F) F {
	f := p.f
	p.f = v
	return f
}

func (p *pair[F, S]) SwapFirstFrom(fn func(v F) F) F {
	f := p.f
	p.f = fn(f)
	return f
}

func (p *pair[F, S]) SwapSecond(v S) S {
	s := p.s
	p.s = v
	return s
}

func (p *pair[F, S]) SwapSecondFrom(fn func(s S) S) S {
	s := p.s
	p.s = fn(s)
	return s
}
