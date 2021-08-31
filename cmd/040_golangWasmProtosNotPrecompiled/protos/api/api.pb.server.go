package api

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func RegisterApiHTTPMux(mux *http.ServeMux, srv ApiServer) {

	mux.HandleFunc("/api.Api/Search", func(w http.ResponseWriter, r *http.Request) {
		in := new(SearchRequest)
		inJSON, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		err = json.Unmarshal(inJSON, in)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		ret, err := srv.Search(context.Background(), in)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		retJSON, err := json.Marshal(ret)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Write(retJSON)
	})

}
