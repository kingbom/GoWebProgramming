 package main
  
 import(
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
 )

 func main(){
	  router := mux.NewRouter();
	  router.HandleFunc("/", index).Methods("GET")
	  router.HandleFunc("/login", login).Methods("GET")
	  router.HandleFunc("/signup", signup).Methods("GET")
	  http.ListenAndServe(":8080", router) 
 }

 func index(w http.ResponseWriter, req *http.Request){
	 http.ServeFile(w, req, "index.html")
 }

 func login(w http.ResponseWriter, req *http.Request){
	http.ServeFile(w, req, "login.html")
}

func signup(w http.ResponseWriter, req *http.Request){
	http.ServeFile(w, req, "signup.html")
}

 func getUser(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	fmt.Fprintf(w, "user is id %s", params["id"])
}