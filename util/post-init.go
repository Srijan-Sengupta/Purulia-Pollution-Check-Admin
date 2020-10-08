package util

import (
	"fmt"
	"github.com/Srijan-Sengupta/Purulia-Pollution-Check-Admin/ui"
	"io/ioutil"
	"net/http"
	"strings"
)

func postInit(fWindow *ui.FirstWindow){
	url := host()+"/rest/admin/haljha1245cxjcbucx/user/init"
  method := "POST"

  payload := strings.NewReader("{\n    \"userName\":\""+fWindow.Username.Text+"\",\n    \"password\":\""+fWindow.Password.Text+"\"\n}")

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
  }
  req.Header.Add("Content-Type", "application/json")

  res, err := client.Do(req)
  if err != nil {
    fmt.Printf(err.Error())
  }
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)

  fmt.Println(string(body))

  print(method,url,payload)
}
