 package main
  
 import(
	"fmt"
	"net/http"
	"database/sql"
	"github.com/gorilla/mux"
  _ "github.com/go-sql-driver/mysql"
 )

 /*type Product struct {
	 Name string
	 Price int
 }*/

 /*type Cookie struct {
	 Name string
	 Value string
	 Expires time.Time
 }*/

 func main(){
	  router := mux.NewRouter()
	  /*router.HandleFunc("/", index)
	  router.HandleFunc("/login", login)
	  router.HandleFunc("/signup", signup)
	  router.HandleFunc("/file", file)
	  router.HandleFunc("/product", getProduct)
	  router.HandleFunc("/upload", upload).Methods("GET")
	  router.HandleFunc("/upload", uploadHandle).Methods("POST")
	  router.HandleFunc("/cookie", cookie)*/
	  //router.HandleFunc("/db", connectDb)

	  connectDb()
	  http.ListenAndServe(":8080", router) 
 }

 func connectDb(){
	 //connect to db
	 db, err := sql.Open("msql","root:password@dbname")
	 if err != nil {
		fmt.Println(err) 
	 }
	 defer db.Close()

	 //query
	 rows, err := db.Query("select id,name from product")
	 if err != nil {
		fmt.Println(err) 
	 }

	 //loop data
	 for rows.Next(){
		var id int
		var name string
		err=rows.Scan(&id,&name)
		fmt.Printf("id : %d name : %s \n", id, name)
	 }
}

 /*func cookie(w http.ResponseWriter, req *http.Request){
	expiration := time.Now().Add(time.Hour * 24 * 365)
	cookie := http.Cookie{Name:"user",Value:"jaruwit", Expires: expiration}
	http.SetCookie(w, &cookie)
	fmt.Fprintf(w, "Create cookie")
}*/

 /*func index(w http.ResponseWriter, req *http.Request){
	 http.ServeFile(w, req, "index.html")
 }

 func login(w http.ResponseWriter, req *http.Request){
	fmt.Println("method : ", req.Method)
	req.ParseForm()
	fmt.Println("username : ", req.Form["username"])
	fmt.Println("password : ", req.Form["password"]) 
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

func upload(w http.ResponseWriter, req *http.Request){
	http.ServeFile(w, req, "upload.html")
}

func uploadHandle(w http.ResponseWriter, req *http.Request){
	file,handle,err := req.FormFile("file")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

    fmt.Fprintf(w, "%v", handle.Header)
	f, err := os.OpenFile("./upload/" + handle.Filename, os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	fmt.Fprintf(w,"Upload complete")
	http.ServeFile(w, req, "upload.html")	
}*/

