package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	firestore "cloud.google.com/go/firestore"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mssola/user_agent"
)

var db *sql.DB
var fs *firestore.Client
var ctx context.Context

type Visit struct {
	When           string
	Browser        string
	BrowserVersion string
	Platform       string
	OS             string
	Mobile         bool
}

func main() {

	var err error

	// db, err = sql.Open("mysql", "root@tcp(localhost:3306)/masons")
	// bug(err)
	// bug(db.Ping())
	// log.Println("mysql conn √")

	ctx = context.Background()
	fs, err = firestore.NewClient(ctx, "cwybra123")
	defer fs.Close()
	bug(err)
	log.Println("firestore conn √")

	http.HandleFunc("/", home)
	log.Println("http serving on 100")
	log.Fatal(http.ListenAndServe(":100", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	a := user_agent.New(r.UserAgent())
	doc := make(map[string]interface{})
	doc["browser"], _ = a.Browser()
	_, err := fs.Collection("visits").Doc("one").Set(ctx, doc)
	bug(err)

	// t := template.New("index.gohtml")
	// b, c := t.ParseFiles("index.gohtml")
	// bug(c)
	// b.Execute(w, doc["browser"])
}

func bug(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
