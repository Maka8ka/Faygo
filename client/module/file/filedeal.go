package fileimport (
	"encoding/base64"
	"fmt"
	"io"
	"os"
)func FileToHex(filelo string) string {
	f, err := os.Open(filelo)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()	chunks := make([]byte, 0)
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}
	
	return base64.StdEncoding.EncodeToString(chunks)
}func CheckFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}func HexToFile(fileencode string, filelo string) string {
	filebyte, err := base64.StdEncoding.DecodeString(fileencode)
	if err != nil {
		fmt.Println(err)
	}
	f, err := os.OpenFile(filelo, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	f.Write(filebyte)
	return "文件写入成功"}
