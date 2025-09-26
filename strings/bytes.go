package strings

func ConcatByteSlices(slices [][]byte) []byte {
	// 计算总长度
	totalLen := 0
	for _, s := range slices {
		totalLen += len(s)
	}

	// 预分配内存并合并
	result := make([]byte, 0, totalLen)
	for _, s := range slices {
		result = append(result, s...)
	}
	return result
}
