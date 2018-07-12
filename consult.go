package main
import(
  //"encoding/json"
//"context"
    "fmt"
    //"io/ioutil"
    "net/http"
    "encoding/json"
  
//  "github.com/hashicorp/consul/api"
  // "github.com/hashicorp/consul/connect"
)
type Q struct{
  Address string `json:"address"`
}



func main() {
  // Create a Consul API client
  //client, _ := api.NewClient(api.DefaultConfig())



  req, err := http.Get("http://localhost:8500/v1/catalog/service/repairorder_local")
if err != nil {
    fmt.Println("NewRequest err: ", err)
    } else {
       var target []Q
      json.NewDecoder(req.Body).Decode(&target)
      //data,_:= ioutil.ReadAll(req.Body)
       fmt.Println("---------------------------------")
   //var q interface{}
   //    json.Unmarshal(target,&q)
        //data,_ :=json.Marshal(target)

        //n := string(data)

          fmt.Println(target[0].Address)

      
      }

     
     
    }
