package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i & 1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0 * 8))] +
		pc[byte(x>>(1 * 8))] +
		pc[byte(x>>(2 * 8))] +
		pc[byte(x>>(3 * 8))] +
		pc[byte(x>>(4 * 8))] +
		pc[byte(x>>(5 * 8))] +
		pc[byte(x>>(6 * 8))] +
		pc[byte(x>>(7 * 8))])
}

// 也可以用匿名函数初始化
var pcx [256]byte = func() (pcx [256]byt) {
	for i := range pcx {		// 这里只用了索引，也可写作for i, _ := range pcx {
		pcx[i] = pcx[i/2] + byte(i & 1)
	}
}
