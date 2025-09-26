package strings

func TrimRightBytes(s []byte, cutset []byte) []byte {
	// 快速处理边界情况
	if len(s) == 0 || len(cutset) == 0 {
		return s
	}

	// 使用 map 存储 cutset 中的字节，实现 O(1) 查找
	cutsetMap := make(map[byte]bool, len(cutset))
	for _, c := range cutset {
		cutsetMap[c] = true
	}

	// 从右向左查找第一个不在 cutset 中的字符位置
	end := len(s)
	for end > 0 && cutsetMap[s[end-1]] {
		end--
	}

	// 如果所有字符都在 cutset 中，返回空切片
	// 否则返回修剪后的切片
	return s[:end]
}

func TrimLeftBytes(s []byte, cutset []byte) []byte {
	// 快速处理边界情况
	if len(s) == 0 || len(cutset) == 0 {
		return s
	}

	// 使用 map 存储 cutset 中的字节，实现 O(1) 查找
	cutsetMap := make(map[byte]bool, len(cutset))
	for _, c := range cutset {
		cutsetMap[c] = true
	}

	// 从左向右查找第一个不在 cutset 中的字符位置
	start := 0
	for start < len(s) && cutsetMap[s[start]] {
		start++
	}

	// 如果所有字符都在 cutset 中，返回空切片
	// 否则返回修剪后的切片
	return s[start:]
}
func TrimBytes(s []byte, cutset []byte) []byte {
	// 快速处理边界情况
	if len(s) == 0 || len(cutset) == 0 {
		return s
	}

	// 使用 map 存储 cutset 中的字节，实现 O(1) 查找
	cutsetMap := make(map[byte]bool, len(cutset))
	for _, c := range cutset {
		cutsetMap[c] = true
	}

	// 从左向右查找第一个不在 cutset 中的字符位置
	start := 0
	for start < len(s) && cutsetMap[s[start]] {
		start++
	}

	// 从右向左查找第一个不在 cutset 中的字符位置
	end := len(s)
	for end > 0 && cutsetMap[s[end-1]] {
		end--
	}

	// 如果所有字符都在 cutset 中，返回空切片
	// 否则返回修剪后的切片
	if start >= end {
		return []byte{}
	}
	return s[start:end]
}
