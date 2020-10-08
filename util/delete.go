package util

import (
	"fmt"
	"github.com/Srijan-Sengupta/Purulia-Pollution-Check-Admin/ui"
	"io/ioutil"
	"net/http"
)

func delete(fWindow *ui.FirstWindow){

	
	url := host()+"/rest/admin/haljha1245cxjcbucx/user="+fWindow.Username.Text+"&&password="+fWindow.Password.Text+"/delete="+fWindow.Date.Text+""
  method := "DELETE"

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, nil)

  print(method,url)

  if err != nil {
    fmt.Println(err)
  }
  res, err := client.Do(req)
  if err!=nil {
	  fmt.Printf(err.Error())
  }
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)

  fmt.Println(string(body))
}
