package charset

import (
	"github.com/saintfish/chardet"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/htmlindex"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

// DetectCharset 检测字节流的字符集
func DetectCharset(data []byte) (string, int) {
	detector := chardet.NewTextDetector()
	result, err := detector.DetectBest(data)
	if err != nil {
		return "", 0
	}
	return result.Charset, result.Confidence
}

// ConvertToUTF8 将指定字符集的字节流转换为UTF-8编码
func ConvertToUTF8(data []byte, fromCharset string, defaultFromCharset string) string {
	var encodingObj encoding.Encoding

	fromCharset = NormalizeCharset(fromCharset)
	if fromCharset == "" {
		// 未指定字符集，尝试检测
		detectedCharset, _ := DetectCharset(data)
		fromCharset = detectedCharset
	}
	if fromCharset == "" {
		fromCharset = defaultFromCharset
	}
	if fromCharset == "" {
		fromCharset = "UTF-8" // 默认使用UTF-8
	}
	fromCharset = NormalizeCharset(fromCharset)

	// 使用htmlindex获取编码器，支持更多标准字符集名称
	encodingObj, err := htmlindex.Get(fromCharset)
	if err != nil {
		// 尝试处理一些特殊情况
		switch fromCharset {
		case "UTF-16BE":
			encodingObj = unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)
		case "UTF-16LE":
			encodingObj = unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
		default:
			return string(data) // 无法识别的字符集，返回原始数据
		}
	}

	// 转换为UTF-8
	result, _, err := transform.Bytes(encodingObj.NewDecoder(), data)
	if err != nil {
		return string(data) // 转换失败，返回原始数据
	}
	return string(result)
}
