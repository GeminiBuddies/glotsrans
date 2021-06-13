package grand

const multiplier uint64 = 6364136223846793005

type PCG struct {
	inc   uint64
	state uint64
}

func NewPCG(seq uint64) *PCG {
	return &PCG{
		inc:   (seq << 1) | 1,
		state: 0,
	}
}

func NewPCGWithSeed(seq, seed uint64) *PCG {
	p := NewPCG(seq)
	p.SeedU(seed)
	return p
}

func (p *PCG) UInt32() uint32 {
	s := p.state

	p.state = multiplier*p.state + p.inc
	x := uint32(((s >> 18) ^ s) >> 27)
	r := uint32(s >> 59)
	return (x >> r) | (x << (-r & 31))
}

func (p *PCG) SeedU(seed uint64) {
	p.state = 0
	p.UInt32()
	p.state += seed
	p.UInt32()
}

func (p *PCG) Int63() int64 {
	return int64(p.Uint64() & 0x7fffffffffffffff)
}

func (p *PCG) Seed(seed int64) {
	p.SeedU(uint64(seed))
}

func (p *PCG) Uint64() uint64 {
	return (uint64(p.UInt32()) << 32) | (uint64(p.UInt32()))
}
