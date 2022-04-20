package mainimport (
	cipher "Faygo/client/module/cipher"
	command "Faygo/client/module/command"
	getmac "Faygo/client/module/getmac"	
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
)func (c *Client) Runcommand() {
	_, out, err := command.NewCommand().Exec(c.Task)
	if err != nil {
		log.Panic(err)
	}
	c.Result = out
	
	
}func AesEnCode(str string) string {
	bytestr := []byte(str)
	cryptstr, err := cipher.AesCbcEncrypt(bytestr, []byte(AesKey))
	if err != nil {
		fmt.Println(err)
		return "err"
	}
	return base64.StdEncoding.EncodeToString(cryptstr)
}func AesDeCode(str string) string {
	bytestr, _ := base64.StdEncoding.DecodeString(str)
	
	decryptstr, err := cipher.AesCbcDecrypt(bytestr, []byte(AesKey))
	if err != nil {
		fmt.Println(err)
		return "err"
	}
	return string(decryptstr)
}func Struct2Json(clientstruct interface{}) string {
	json, _ := json.Marshal(clientstruct) 
	return string(json)
}
func (client *Client) Json2Struct(clientjson string) {
	jsontmp := string(clientjson)
	json.Unmarshal([]byte(jsontmp), client)
}func GetIP() string {
	return getmac.GetIPs()[0]
}func GetMac() string {
	return getmac.GetMacAddrs()[0]
}
