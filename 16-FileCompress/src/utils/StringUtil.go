package utils

import "golang.org/x/text/encoding/simplifiedchinese"

// UTF8编码字符串转换成GBK编码字符串
func UTF8ToGBK(txt string) (newTxt string, err error) {
	dst := make([]byte, len(txt)*2)
	tr := simplifiedchinese.GB18030.NewDecoder()
	nDst, _, err := tr.Transform(dst, []byte(txt), true)
	if err == nil {
		newTxt = string(dst[:nDst])
	}
	return
}
