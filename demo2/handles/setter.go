package handles

import (
	"github.com/waltcow/golangProject/demo2/storage"
	"net/http"
	"io/ioutil"
)

func SetKey(db storage.Database) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "missing key name in path", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		val, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "error reading PUT body", http.StatusBadRequest)
			return
		}
		if err := db.Set(key, val); err != nil {
			http.Error(w, "error setting value in DB", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}