package grandpool

import (
	"glotsrans/utils/gerror"
	"sync"
)

type Pool struct {
	mu   sync.Mutex
	cap  int
	buff []byte
	head int
	avail int
}

func New(capacity int) *Pool {
	return &Pool{
		mu:    sync.Mutex{},
		cap:   capacity,
		buff:  make([]byte, capacity),
		head:  0,
		avail: 0,
	}
}

func (p *Pool) Capacity() int {
	return p.cap
}

func (p *Pool) AvailableNow() int {
	return p.avail
}

func (p *Pool) ReadTo(s []byte, nonBlocking bool) error {
	l := len(s)

	if l > p.cap {
		return gerror.RequestTooManyBytes(l, p.cap)
	}

	defer p.mu.Unlock()
	p.mu.Lock()

	if l > p.avail && nonBlocking {
		return gerror.CannotReadAsManyAsExpected(l, p.avail)
	}

 	// TODO

	return nil
}

func (p *Pool) Read(n int, nonBlocking bool) ([]byte, error) {
	temp := make([]byte, n)
	if err := p.ReadTo(temp, nonBlocking); err != nil {
		return nil, err
	}
	return temp, nil
}
