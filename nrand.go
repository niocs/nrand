package nrand

import ("math/rand")

type Nrand struct {
	R         *rand.Rand
	Seed      int64
	Min       float64
	Max       float64
	mean      float64
	sd        float64
	MaxSd     float64    // ideally 3
	DecPrec   int16      // Decimal precision, default 2 assuming currencies
}

func New(seed int64) *Nrand {
	src := rand.NewSource(seed)
	return &Nrand{R:rand.New(src), Seed:seed, MaxSd:3.0, DecPrec:2}
}

func (nr *Nrand) SetRange(min, max float64) {
	nr.Min  = min
	nr.Max  = max
	nr.mean = (max - min) / 2.0 + min
	nr.sd   = (max - min) / (2.0 * nr.MaxSd)
}

func (nr *Nrand) NormFloat64() float64 {
	var rn = nr.Min - 1
	for rn <= nr.Min || rn >= nr.Max {
		rn = nr.R.NormFloat64() * nr.sd + nr.mean
	}
	return rn
}
