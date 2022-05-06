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
	// "github.com/google/uuid"
	"github.com/gorilla/mux"
)

var db *sql.DB
var server = "localhost"
var port = 1433
var user = "sa"
//var password = "Interstellar"
var password = "Databases22"
var database = "KickerStats"

func main(){

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", server, user, password, port, database)
	var err error
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error connecting: ", err.Error())
	}
	// Handle api requests
	router := mux.NewRouter()
	fmt.Println("connected to router")

	//Router request paths
	router.HandleFunc("/topten", GetTopTen).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/searchplayers", GetPlayer).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/calculatefieldgoal", GetProbabilityFieldgoal).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("calculatekickoff", GetProbabilityKickoff).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/getkick", GetAllKickers).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/test", Test).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/getteams", GetTeams).Methods(http.MethodPost, http.MethodOptions)
	srv := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
	fmt.Println("Connected")
}

func Test(w http.ResponseWriter, r *http.Request){
	fmt.Println("this is a test")
}

func GetTeams(w http.ResponseWriter, r *http.Request){
	ctx := context.Background()
	err := db.PingContext(ctx)
	fmt.Println(err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tsql := fmt.Sprintf("SELECT team_name, team_location, team_abbr, team_id, home_stadium, home_surface, home_roof, division, conference, primary_color, secondary_color;")
	rows, err := db.QueryContext(ctx, tsql)
		if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
}

func GetAllKickers(w http.ResponseWriter, r *http.Request){
	ctx := context.Background()
	err := db.PingContext(ctx)
	fmt.Println(err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tsql := fmt.Sprintf("SELECT kicker_id, first_name, last_name, jersey_number, player_status FROM Kickers;")
	rows, err := db.QueryContext(ctx, tsql)
		if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var players []model.KickerInfo
	for rows.Next(){
		var p model.KickerInfo
		rows.Scan(&p.Kicker_id, &p.First_name, &p.Last_name, &p.Jersey_number, &p.Player_status)
		players = append(players, p)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	w.WriteHeader(http.StatusOK)
	resp := model.JsonResponseAllKickers{Type: "Success", Message: "", Data: players}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
func GetTopTen(w http.ResponseWriter, r *http.Request){
	//This will query the database and return a json response of the kickers that are the top 10 preformers based on a specific criteria 
	ctx := context.Background()
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//Get the top ten by taking the selected parameter and doing a join to get data from both tables using the kickerid as the primary key
	// tsql := fmt.Sprintf("SELECT TOP 10 Kickers.first_name + ' ' + last_name AS 'Name' FROM Kickers ORDER BY Kickers.kicker_id DESC")
	tsql := fmt.Sprintf("SELECT TOP 10 first_name, last_name, fieldgoals_attempts,fieldgoals_made, kickoffs_endzone, kickoffs_inside_twenty,games_played, games_started FROM Kickers INNER JOIN KickerSeason ON Kickers.kicker_id = KickerSeason.kicker_id WHERE Kickers.kicker_id = 12 ORDER BY Kickers.kicker_id DESC")
	rows, err := db.QueryContext(ctx, tsql)
		if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
		var players []model.TopKicker
	for rows.Next(){
		var p model.TopKicker
		rows.Scan(&p.First_name, &p.Last_name, &p.Fieldgoals_attempts, &p.Fieldgoals_made, &p.Kickoffs_endzone, &p.Kickoffs_inside_twenty, &p.Games_played, &p.Games_started)
		players = append(players, p)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	w.WriteHeader(http.StatusOK)
	resp := model.JsonTopKicker{Type: "Success", Message: "", Data: players}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetPlayer(w http.ResponseWriter, r *http.Request){
	//Gets a specific player and returns there stats and information as a json response 
	ctx := context.Background()
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var played model.GetPlayer
	err = json.NewDecoder(r.Body).Decode(&played)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	 fmt.Println(played)
	// var selectedId = "12"
	tsql := fmt.Sprintf("SELECT first_name,last_name,kickoffs,kickoffs_squibs,fieldgoals_attempts,fieldgoals_made,fieldgoals_blocked,fieldgoal_yards,fieldgoal_longest,xp_attempts,xp_made FROM Kickers INNER JOIN KickerSeason ON Kickers.kicker_id = KickerSeason.kicker_id WHERE Kickers.kicker_id ='%s';", played.Kicker_id)
	rows, err := db.QueryContext(ctx, tsql)
		if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var players []model.SelectedPlayer
	for rows.Next(){
		var player model.SelectedPlayer
		rows.Scan(&player.First_name, &player.Last_name, &player.Kickoffs, &player.Kickoffs_squibs, &player.Fieldgoals_attempts, &player.Fieldgoals_made, &player.Fieldgoals_made, &player.Fieldgoal_yards, &player.Fieldgoal_longest, &player.Xp_attempts, &player.Xp_made)
		players = append(players, player)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	w.WriteHeader(http.StatusOK)
	resp := model.JsonSearchResult{Type: "Success", Message: "", Data: players}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetProbabilityFieldgoal(w http.ResponseWriter, r *http.Request){
	//Will take a player and then take their stats and calculate the probability of them making a field goal based on number of success / number of attempts
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
	var players []model.KickerFieldGoal
	for rows.Next(){
		var player model.KickerFieldGoal
		rows.Scan(&player.First_name, &player.Last_name, &player.Fieldgoals_attempts, &player.Fieldgoals_made, &player.Fieldgoal_yards, &player.Fieldgoal_longest)
		players = append(players, player)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	w.WriteHeader(http.StatusOK)
	resp := model.JsonFieldgoalStats{Type: "Success", Message: "", Data: players}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetProbabilityKickoff(w http.ResponseWriter, r *http.Request){
	// takes in the information and decides if it will be a good kick or bad kick
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
	var players []model.KickerKickoff
	for rows.Next(){
		var player model.KickerKickoff
		rows.Scan(&player.First_name, &player.Last_name, &player.Kickoffs, &player.Kickoffs_endzone, &player.Kickoffs_inside_twenty, &player.Kickoffs_return_yards, &player.Kickoffs_return_yards, &player.Kickoffs_touchbacks, &player.Kickoffs_yards, &player.Kickoffs_out_of_bounds, &player.Kickoffs_onside_attempts, &player.Kickoffs_onside_success, &player.Kickoffs_squibs)
		players = append(players, player)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	w.WriteHeader(http.StatusOK)
	resp := model.JsonKickoffStats{Type: "Success", Message: "", Data: players}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}