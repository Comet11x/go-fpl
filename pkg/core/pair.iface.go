package core

type Pair[F any, S any] struct {
	isf bool
	f   F
	iss bool
	s   S
}

func (p *Pair[F, S]) First() Option[F] {
	if p.isf {
		return Some(p.f)
	} else {
		return None[F]()
	}
}

func (p *Pair[F, S]) Second() Option[S] {
	if p.iss {
		return Some(p.s)
	} else {
		return None[S]()
	}
}
