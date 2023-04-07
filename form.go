package main

import (
	"html/template"
	"net/http"
	_ "strconv"
	"fmt"
	"github.com/openrdap/rdap"
	"encoding/json"
)

type Lookup struct {
	Address string
	Network *rdap.IPNetwork

}

func (l Lookup) RawString() (string, error){
	if b, err := json.Marshal(l); err == nil{
		return string(b), nil
	} else {
		return "",err
	}
}

func main() {
    tmpl := template.Must(template.ParseFiles("forms.html"))

 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            tmpl.Execute(w, nil)
            return
        }
		
        details := Lookup{
            Address:   r.FormValue("address"),
        }

		if details.Address != ""{
			client := &rdap.Client{}
			if network, err := client.QueryIP(details.Address); err == nil{
				fmt.Println("found")
				details.Network = network

			} else {
				fmt.Println(err)
			}
			

		}

        // do something with details
        fmt.Printf("%#v\n", details)

        tmpl.Execute(w, details)
    })

    http.ListenAndServe(":8080", nil)
}