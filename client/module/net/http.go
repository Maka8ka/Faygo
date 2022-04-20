package netimport (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)func HttpGet(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		
		s := strings.Split(fmt.Sprintln(err), " ")
		
		return s[len(s)-1] 
	} else {
		defer resp.Body.Close()
		
		
		
		
		return string(resp.Status)
	}}func HttpPost(url string, body string) string {
	resp, err := http.Post(url, "text/html", strings.NewReader(body))
	if err != nil {
		fmt.Println(err)
		time.Sleep(time.Duration(30) * time.Second) 
		return "error2"
	} else {
		defer resp.Body.Close()
		response, _ := ioutil.ReadAll(resp.Body)
		return string(response)
	}
}
