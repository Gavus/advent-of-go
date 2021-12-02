package types

type Ring struct {
	index uint
	size  uint
	data  []int
}

func NewRing(size uint) *Ring {
	return &Ring{0, size, make([]int, size)}
}

func MakeRing(size uint) Ring {
	return Ring{0, size, make([]int, size)}
}

func (r *Ring) Put(value int) {
	if r.size > r.index {
		r.data[r.index] = value
		r.index++
	} else {
		r.data[0] = value
		r.index = 1
	}
}

func (r Ring) Sum() int {
	result := 0
	for _, v := range r.data {
		result += v
	}
	return result
}
