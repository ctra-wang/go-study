package byte_rune

// GetStringSize 如何优化下面这个方法
func GetStringSize(s string) int {
	size := 0
	for i := 0; i < len(s); i++ {
		size++
	}
	return size
}
