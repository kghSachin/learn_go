package main 
import "net/http"


//function signature to define http in go 
func handlerReadiness(w http.ResponseWriter, r *http.Request){
	respondWithJson(w,200, struct{}{})
}