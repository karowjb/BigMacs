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
var database = "KickerStats"

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
	router.HandleFunc("/topten", GetTopTen).Methods(http.MethodPost)
	router.HandleFunc("/searchplayers", GetPlayer).Methods(http.MethodPost)
	router.HandleFunc("/calculatefieldgoal", GetProbabilityFieldgoal).Methods(http.MethodPost)
	router.HandleFunc("calculatekickoff", GetProbabilityKickoff).Methods(http.MethodPost)

	srv := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
	
}

func GetTopTen(w http.ResponseWriter, r *http.Request){
	return 
}

func GetPlayer(w http.ResponseWriter, r *http.Request){
	ctx := context.Background()
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// tsql := fmt.Sprintf("SELECT first_name,last_name,kickoffs,kickoffs_squibs,fieldgoals_attempts,fieldgoals_made,fieldgoals_blocked,fieldgoal_yards,fieldgoal_longest,xp_attempts,xp_made FROM ")
	tsql := fmt.Sprintf("SELECT * FROM Teams")
	rows, err := db.QueryContext(ctx, tsql)
		if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var player []model.SelectedPlayer
	for rows.Next(){
		var player model.SelectedPlayer
		rows.Scan(&player.first_name, &player.last_name, &player.kickoffs, &player.kickoffs_squibs, &player.fieldgoals_attempts, &player.fieldgoals_made, &player.fieldgoals_made, &player.fieldgoal_yards, &player.fieldgoal_longest, &player.xp_attempts, &player.xp_made)
		players = append(players, player)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := model.JsonSearchResponse{Type: "Success", Message: "", Data: players}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetProbabilityFieldgoal(w http.ResponseWriter, r *http.Request){
	ctx := context.Background()
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tsql := fmt.Sprintf("SELECT * FROM Teams")
	rows, err := db.QueryContext(ctx, tsql)
		if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var player []model.KickerFieldGoal
	for rows.Next(){
		var player model.KickerFieldGoal
		rows.Scan(&player.first_name, &player.last_name, &player.fieldgoals_attempts, &player.fieldgoals_made, &player.fieldgoal_yards, &player.fieldgoal_longest)
		players = append(players, player)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := model.JsonFieldgoalStats{Type: "Success", Message: "", Data: players}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetProbabilityKickoff(w http.ResponseWriter, r *http.Request){
	ctx := context.Background()
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// tsql := fmt.Sprintf("SELECT first_name,last_name,kickoffs,kickoffs_squibs,fieldgoals_attempts,fieldgoals_made,fieldgoals_blocked,fieldgoal_yards,fieldgoal_longest,xp_attempts,xp_made FROM ")
	tsql := fmt.Sprintf("SELECT * FROM Teams")
	rows, err := db.QueryContext(ctx, tsql)
		if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var player []model.KickerKickoff
	for rows.Next(){
		var player model.KickerKickoff
		rows.Scan(&player.first_name, &player.last_name, &player.kickoffs, &player.kickoffs_endzone,, &player.kickoffs_inside_twenty, &player.kickoffs_return_yards, &player.kickoffs_return_yards, &player.kickoffs_touchbacks, &player.kickoffs_yards, &player.kickoffs_out_of_bounds, &player.kickoffs_onside_attempts, &player.kickoffs_onside_success, &player.kickoffs_squibs)
		players = append(players, player)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := model.JsonKickoffStats{Type: "Success", Message: "", Data: players}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}