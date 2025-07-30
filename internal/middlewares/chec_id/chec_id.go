package chec_id

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Param_chec_id struct {
	Id           string
	ModelId      any
	Name         string
	Writer       http.ResponseWriter
	ModelSecsees any
}

// This function ned for chec id from params thid id DB
func ChecId(p Param_chec_id) {

	if p.Id != fmt.Sprint(p.ModelId) {
		json.NewEncoder(p.Writer).Encode("not correct ID to " + p.Name)
	} else {
		json.NewEncoder(p.Writer).Encode(p.ModelSecsees)
	}

}
