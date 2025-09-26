package quoptedprintable

import "encoding/hex"

// isHexChar 判断字符是否为十六进制字符
func isHexChar(c byte) bool {
	return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')
}

// decodeHeaderQP 解码邮件头中的Quoted-Printable编码（RFC 2047规范）
// 特性：处理=XX格式编码、下划线转空格、忽略非法编码保留原始字符
func DecodeMimeHeader(data []byte) []byte {
	result := make([]byte, 0, len(data)) // 预分配内存提升性能
	i := 0
	for i < len(data) {
		// 处理=XX编码（两位十六进制）
		if data[i] == '=' && i+2 < len(data) {
			// 提取十六进制字符（允许大小写）
			hexStr := string(data[i+1 : i+3])
			if len(hexStr) == 2 && isHexChar(hexStr[0]) && isHexChar(hexStr[1]) {
				// 解码十六进制为字节
				val, err := hex.DecodeString(hexStr)
				if err == nil && len(val) == 1 {
					result = append(result, val[0])
					i += 3 // 跳过整个=XX序列
					continue
				}
			}
		}

		// 处理下划线（邮件头QP中_代表空格）
		if data[i] == '_' {
			result = append(result, ' ')
			i++
			continue
		}

		// 普通字符直接保留（包括合法的ASCII可见字符和未识别的编码）
		result = append(result, data[i])
		i++
	}
	return result
}
func DecodeMimeBody(data []byte) []byte {
	result := make([]byte, 0, len(data)) // 预分配内存提升性能
	i := 0
	for i < len(data) {
		// 处理=XX编码（两位十六进制）
		if data[i] == '=' && i+2 < len(data) {
			// \r\n
			if data[i+1] == '\r' && data[i+2] == '\n' {
				// 软换行，跳过
				i += 3
				continue
			}
			if data[i+1] == '\n' {
				// 软换行，跳过
				i += 2
				continue
			}
			// 提取十六进制字符（允许大小写）
			hexStr := string(data[i+1 : i+3])
			if len(hexStr) == 2 && isHexChar(hexStr[0]) && isHexChar(hexStr[1]) {
				// 解码十六进制为字节
				val, err := hex.DecodeString(hexStr)
				if err == nil && len(val) == 1 {
					result = append(result, val[0])
					i += 3 // 跳过整个=XX序列
					continue
				}
			}
		}

		// 普通字符直接保留（包括合法的ASCII可见字符和未识别的编码）
		result = append(result, data[i])
		i++
	}
	return result
}
