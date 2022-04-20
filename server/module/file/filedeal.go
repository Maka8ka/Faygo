package fileimport (
	"encoding/base64"
	"io"
	"os"
)func FileToHex(filelo string) string {
	f, err := os.Open(filelo)
	if err != nil {
		panic(err)
	}
	defer f.Close()	chunks := make([]byte, 0)
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
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
		panic(err)
		
	}
	var f *os.File
	var err1 error
	if CheckFileIsExist(filelo) { 
		f, err1 = os.OpenFile(filelo, os.O_APPEND, 0666) 
		if err1 != nil {
			panic(err1)
		}
		return "文件存在"
	} else {
		f, err1 = os.Create(filelo) 		if err1 != nil {
			panic(err1)
		}	}
	defer f.Close()
	f.Write(filebyte)
	
	
	f.Sync()
	return "文件写入成功"
}
