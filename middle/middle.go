package middle

import (
	"log"
	"net/http"
	//"time"
)

func Logger(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//request_first_called_time := time.Now()
		f(w, r)
		//time_elapsed := time.Since(request_first_called_time)
		log.Println(r.URL.Path, "   ")//, time_elapsed.Milliseconds(), "ms")
	}
}
