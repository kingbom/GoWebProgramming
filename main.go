 package main
  
 import(
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
 )

 func main(){
	  router := mux.NewRouter();
	  router.HandleFunc("/", index).Methods("GET")
	  router.HandleFunc("/users/{id}", getUser).Methods("GET")
	  http.ListenAndServe(":8080", router) 
 }

 func index(w http.ResponseWriter, req *http.Request){
	 fmt.Fprintf(w, "index")
 }

 func getUser(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	fmt.Fprintf(w, "user is id %s", params["id"])
}