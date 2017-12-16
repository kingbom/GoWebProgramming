 package main
  
 import(
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
 )

 func main(){
	  router := mux.NewRouter();
	  router.HandleFunc("/", index)
	  router.HandleFunc("/login", login)
	  router.HandleFunc("/signup", signup)
	  router.HandleFunc("/file", file)
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

func file(w http.ResponseWriter, req *http.Request){
	http.ServeFile(w, req, "file.txt")
}

 func getUser(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	fmt.Fprintf(w, "user is id %s", params["id"])
}