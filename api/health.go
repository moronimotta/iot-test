package healthpackage healthpackage api



import (

	"encoding/json"

	"net/http"import (import (

)

	"encoding/json"	"encoding/json"

func Handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")	"net/http"	"net/http"

	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")))

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")



	if r.Method == "OPTIONS" {

		w.WriteHeader(http.StatusOK)func Handler(w http.ResponseWriter, r *http.Request) {func Handler(w http.ResponseWriter, r *http.Request) {

		return

	}	w.Header().Set("Content-Type", "application/json")	w.Header().Set("Content-Type", "application/json")



	response := map[string]interface{}{	w.Header().Set("Access-Control-Allow-Origin", "*")	w.Header().Set("Access-Control-Allow-Origin", "*")

		"status":  "ok",

		"message": "Server is running on Vercel",	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

	}

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)

}
	if r.Method == "OPTIONS" {	if r.Method == "OPTIONS" {

		w.WriteHeader(http.StatusOK)		w.WriteHeader(http.StatusOK)

		return		return

	}	}



	response := map[string]interface{}{	response := map[string]interface{}{

		"status":  "ok",		"status":  "ok",

		"message": "Server is running on Vercel",		"message": "Server is running on Vercel",

	}	}



	w.WriteHeader(http.StatusOK)	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)	json.NewEncoder(w).Encode(response)

}}