package core

func init() {
	// TODO: seed with temp sensor
}

var r = struct {
	w, x, y, z uint32
}{12345678, 4185243, 776511, 45411}

// Seed seeds the random number generator with s.
func Seed(s uint32) {
	r.w, r.x, r.y, r.z = s, 4185243, 776511, 45411
}

// Random returns a pseudorandom number from
// 0 to 4294967294.
func Random() uint32 {
	t := r.x ^ (r.x << 11)
	r.x, r.y, r.z = r.y, r.z, r.w
	r.w = (r.w ^ (r.w >> 19)) ^ (t ^ (t >> 8))
	return r.w
}
