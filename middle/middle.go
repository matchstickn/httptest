package middle

import (
	"log"
	"net/http"
	//"time"
)

func Logger(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//request_first_called_time := time.Now()
		f.ServeHTTP(w, r)
		//time_elapsed := time.Since(request_first_called_time)
		log.Println(r.URL.Path, "   ") //, time_elapsed.Milliseconds(), "ms")
	}
}

func CORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
