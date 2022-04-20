package commandimport "golang.org/x/text/encoding/simplifiedchinese"type Charset stringconst (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)
type Commander interface {
	
	
	
	Exec(args ...string) (int, string, error)	
	
	
	
	
	ExecAsync(stdout chan string, args ...string) int	
	
	
	ExecIgnoreResult(args ...string) error
}func ConvertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}
