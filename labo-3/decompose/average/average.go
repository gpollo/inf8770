package average

type Signal interface {
	Value() float64
}

type averageValue struct {
	signal Signal
	weight float64
}

type Moving struct {
	size   int
	next   int
	ready  bool
	value  float64
	values []averageValue
}

func NewMoving(size uint) *Moving {
	return &Moving{
		size:   int(size),
		next:   0,
		ready:  false,
		values: make([]averageValue, int(size)),
	}
}

func (m *Moving) Add(s Signal) (Signal, float64, bool) {
	last := m.values[m.next].signal

	if m.ready {
		m.value -= m.values[m.next].weight
	}

	m.values[m.next].signal = s
	m.values[m.next].weight = s.Value() / float64(m.size)

	m.value += m.values[m.next].weight
	m.next = (m.next + 1) % m.size

	if !m.ready {
		if m.next == 0 {
			m.ready = true
		}

		return nil, 0, false
	} else {
		return last, m.value, true
	}
}
