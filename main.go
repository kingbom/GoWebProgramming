 package main
  
 import(
	"fmt"
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
 )

 type Product struct {
	 Name string
	 Price int
 }

 func main(){
	  router := mux.NewRouter()
	  router.HandleFunc("/", index)
	  router.HandleFunc("/login", login)
	  router.HandleFunc("/signup", signup)
	  router.HandleFunc("/file", file)
	  router.HandleFunc("/product", getProduct)
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

func getProduct(w http.ResponseWriter, req *http.Request){
	var templates = template.Must(template.ParseFiles("product.html"))
	product := Product{"Iphone", 22000}
	templates.ExecuteTemplate(w, "product.html", product)
}

 func getUser(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	fmt.Fprintf(w, "user is id %s", params["id"])
}