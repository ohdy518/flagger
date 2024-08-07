package micro

import "encoding/base64"

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func PanicIfEmpty(str string) {
	if str == "" {
		panic("Empty string")
	}
}

func EncodeToBase64(input string) string {
	data := []byte(input)
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(dst, data)
	return string(dst)
}

func DecodeFromBase64(input string) string {
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(input)))
	n, err := base64.StdEncoding.Decode(dst, []byte(input))
	CheckError(err)
	dst = dst[:n]
	return string(dst)
}
