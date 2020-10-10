package util

import (
	"fmt"
	"github.com/Srijan-Sengupta/Purulia-Pollution-Check-Admin/ui"
	"io/ioutil"
	"net/http"
	"strings"
)

func post(fWindow *ui.FirstWindow){
	url := host()+"/rest/admin/haljha1245cxjcbucx/user="+fWindow.Username.Text+"&&password="+fWindow.Password.Text
  method := "POST"

  payload := strings.NewReader("{\n    \"time\":\""+fWindow.Date.Text+"\",\n    \"temperature\":"+(fWindow.Temperature.Text)+",\n    \"pm2_5\":"+fWindow.Pm2_5Level.Text+",\n    \"humidity\":"+fWindow.Humidity.Text+"\n}")

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
  fWindow.Log.SetText(string(body))
  print(method,url,payload)
}
