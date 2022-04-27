package main
import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	model "kicker-api/models"
	"log"
	"net/http"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var db *sql.DB
var server = "localhost"
var port = 1433
var user = "sa"
var password = "Interstellar"
// var database = ""

func main(){
	fmt.Printf("Hello World")
	connString := fmt.Sprintf("server='%s';user id='%s';password='%s';port=%d;database='%s';", server, user, password, port, database)
	var err error
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error connecting: ", err.Error())
	}
	// Handle api requests
	router := mux.NewRouter()

	//Router request paths

	srv := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
	
}