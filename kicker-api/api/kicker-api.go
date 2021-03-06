package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	model "kicker-api/models"
	"log"
	"net/http"
	"strconv"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
)

var db *sql.DB
var server = "localhost"
var port = 1433
var user = "sa"

var password = "Interstellar"

// var password = "Databases22"

var database = "KickerStats"

// var database = "Kicker_Stats"

func main() {

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", server, user, password, port, database)
	var err error
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error connecting: ", err.Error())
	}
	// Handle api requests
	router := mux.NewRouter()
	fmt.Println("Welcome to the KickerAPI")

	//Router request paths
	router.HandleFunc("/toptenfg", GetTopTenFG).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/toptenlongestfg", GetTopTenLongestFG).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/toptenkintoendzone", GetTopTenEndzoneK).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/searchplayers", GetPlayer).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/probfieldgoal", GetProbabilityFieldgoal).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/probabilities", GetProbabilityKickoff).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/getkickers", GetAllKickers).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/getteams", GetAllTeams).Methods(http.MethodPost, http.MethodOptions)
	srv := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

func GetAllTeams(w http.ResponseWriter, r *http.Request) {
	//Selecting all teams
	ctx := context.Background()
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tsql := "SELECT team_name, team_location, team_abbr, team_id, home_stadium, home_surface, home_roof, division, conference, primary_color, secondary_color FROM Teams;"
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var teams []model.Team
	for rows.Next() {
		var t model.Team
		rows.Scan(&t.Team_name, &t.Team_location, &t.Team_abbr, &t.Team_id, &t.Home_stadium, &t.Home_surface, &t.Home_roof, &t.Division, &t.Conference, &t.Primary_color, &t.Secondary_color)
		teams = append(teams, t)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	w.WriteHeader(http.StatusOK)
	resp := model.JsonTeam{Type: "Success", Message: "", Data: teams}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetAllKickers(w http.ResponseWriter, r *http.Request) {
	// Returns all kickers
	ctx := context.Background()
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tsql := "SELECT kicker_id, first_name, last_name, jersey_number, player_status FROM Kickers;"
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var kickers []model.KickerInfo
	for rows.Next() {
		var k model.KickerInfo
		rows.Scan(&k.Kicker_id, &k.First_name, &k.Last_name, &k.Jersey_number, &k.Player_status)
		kickers = append(kickers, k)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	w.WriteHeader(http.StatusOK)
	resp := model.JsonResponseAllKickers{Type: "Success", Message: "", Data: kickers}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
func GetTopTenFG(w http.ResponseWriter, r *http.Request) {
	//This will query the database and return a json response of the kickers that are the top 10 performers based on a specific criteria
	ctx := context.Background()
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tsql := "SELECT TOP 10 first_name, last_name, jersey_number, k.team_id, team_name, team_location, fieldgoals_made, fieldgoal_longest, kickoffs_endzone, season_year FROM Kickers k JOIN KickerSeason ks ON ks.kicker_id = k.kicker_id JOIN Teams t ON t.team_id = k.team_id ORDER BY fieldgoals_made DESC;"
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var topkickers []model.TopKickers
	for rows.Next() {
		var tk model.TopKickers
		rows.Scan(&tk.First_name, &tk.Last_name, &tk.Jersey_number, &tk.Team_id, &tk.Team_name, &tk.Team_location, &tk.Fieldgoals_made, &tk.Fieldgoal_longest, &tk.Kickoffs_endzone, &tk.Season_year)
		topkickers = append(topkickers, tk)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	w.WriteHeader(http.StatusOK)
	resp := model.JsonTopKickers{Type: "Success", Message: "", Data: topkickers}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetTopTenLongestFG(w http.ResponseWriter, r *http.Request) {
	//This will query the database and return a json response of the kickers that are the top 10 performers based on a specific criteria
	ctx := context.Background()
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tsql := "SELECT TOP 10 first_name, last_name, jersey_number, k.team_id, team_name, team_location, fieldgoals_made, fieldgoal_longest, kickoffs_endzone, season_year FROM Kickers k JOIN KickerSeason ks ON ks.kicker_id = k.kicker_id JOIN Teams t ON t.team_id = k.team_id ORDER BY fieldgoal_longest DESC;"
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var topkickers []model.TopKickers
	for rows.Next() {
		var tk model.TopKickers
		rows.Scan(&tk.First_name, &tk.Last_name, &tk.Jersey_number, &tk.Team_id, &tk.Team_name, &tk.Team_location, &tk.Fieldgoals_made, &tk.Fieldgoal_longest, &tk.Kickoffs_endzone, &tk.Season_year)
		topkickers = append(topkickers, tk)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	w.WriteHeader(http.StatusOK)
	resp := model.JsonTopKickers{Type: "Success", Message: "", Data: topkickers}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func GetTopTenEndzoneK(w http.ResponseWriter, r *http.Request) {
	//This will query the database and return a json response of the kickers that are the top 10 performers based on a specific criteria
	ctx := context.Background()
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tsql := "SELECT TOP 10 first_name, last_name, jersey_number, k.team_id, team_name, team_location, fieldgoals_made, fieldgoal_longest, kickoffs_endzone, season_year FROM Kickers k JOIN KickerSeason ks ON ks.kicker_id = k.kicker_id JOIN Teams t ON t.team_id = k.team_id ORDER BY kickoffs_endzone DESC;"
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var topkickers []model.TopKickers
	for rows.Next() {
		var tk model.TopKickers
		rows.Scan(&tk.First_name, &tk.Last_name, &tk.Jersey_number, &tk.Team_id, &tk.Team_name, &tk.Team_location, &tk.Fieldgoals_made, &tk.Fieldgoal_longest, &tk.Kickoffs_endzone, &tk.Season_year)
		topkickers = append(topkickers, tk)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	w.WriteHeader(http.StatusOK)
	resp := model.JsonTopKickers{Type: "Success", Message: "", Data: topkickers}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetPlayer(w http.ResponseWriter, r *http.Request) {
	// Selects a player by firstname and lastname
	ctx := context.Background()
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var player model.DBSearch
	err = json.NewDecoder(r.Body).Decode(&player)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tsql := "SELECT k.kicker_id, first_name, last_name, jersey_number, player_height, player_weight, k.team_id, t.team_name, t.team_location, t.home_stadium, t.primary_color, t.secondary_color, SUM(kickoffs) kickoffs, SUM(fieldgoals_attempts) fieldgoals_attempts, SUM(fieldgoals_made) fieldgoals_made, SUM(fieldgoal_yards) fieldgoal_yards, MAX(fieldgoal_longest) fieldgoal_longest, SUM(xp_attempts) xp_attempts, SUM(xp_made) xp_made FROM Kickers k INNER JOIN KickerSeason ks ON k.kicker_id = ks.kicker_id JOIN Teams t ON t.team_id = k.team_id WHERE k.player_height BETWEEN " + strconv.Itoa(player.Height_min) + " AND " + strconv.Itoa(player.Height_max) + " AND k.player_weight BETWEEN " + strconv.Itoa(player.Weight_min) + " AND " + strconv.Itoa(player.Weight_max)
	if len(player.First_name) != 0 {
		tsql = tsql + " AND first_name = '" + player.First_name + "'"
	}
	if len(player.Last_name) != 0 {
		tsql = tsql + " AND last_name = '" + player.Last_name + "'"
	}
	if len(player.Team_id) != 0 {
		tsql = tsql + " AND k.team_id = '" + player.Team_id + "'"
	}
	tsql = tsql + "GROUP BY k.kicker_id, first_name, last_name, jersey_number, player_height, player_weight, k.team_id, t.team_name, t.team_location, t.home_stadium, t.primary_color, t.secondary_color;"
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var topkickers []model.SelectedPlayer
	for rows.Next() {
		var player model.SelectedPlayer
		rows.Scan(&player.Kicker_id, &player.First_name, &player.Last_name, &player.Jersey_number, &player.Height, &player.Weight, &player.Team_id, &player.Team_name, &player.Team_location, &player.Team_home_stadium, &player.Team_primary, &player.Team_secondary, &player.Kickoffs, &player.Fieldgoals_attempts, &player.Fieldgoals_made, &player.Fieldgoal_yards, &player.Fieldgoal_longest, &player.Xp_attempts, &player.Xp_made)
		topkickers = append(topkickers, player)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	w.WriteHeader(http.StatusOK)
	resp := model.JsonSearchResult{Type: "Success", Message: "", Data: topkickers}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetProbabilityFieldgoal(w http.ResponseWriter, r *http.Request) {
	// Gets the fieldgoal information for the selected kicker
	ctx := context.Background()
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var played model.Kicker
	err = json.NewDecoder(r.Body).Decode(&played)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tsql := fmt.Sprintf("SELECT fieldgoals_attempts, fieldgoals_made, season_year, season_type FROM KickerSeason INNER JOIN Kickers ON KickerSeason.kicker_id = Kickers.kicker_id WHERE first_name = '%s' AND last_name = '%s';", played.First_name, played.Last_name)
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var kickers []model.KickerFieldGoal
	for rows.Next() {
		var kicker model.KickerFieldGoal
		rows.Scan(&kicker.Fieldgoals_attempts, &kicker.Fieldgoals_made, &kicker.Season_year, &kicker.Season_type)
		kickers = append(kickers, kicker)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	w.WriteHeader(http.StatusOK)
	resp := model.JsonFieldgoalStats{Type: "Success", Message: "", Data: kickers}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetProbabilityKickoff(w http.ResponseWriter, r *http.Request) {
	// takes in the information and decides if it will be a good kick or bad kick
	ctx := context.Background()
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var played model.Kicker
	err = json.NewDecoder(r.Body).Decode(&played)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tsql := fmt.Sprintf("SELECT SUM(fieldgoals_attempts) AS FieldGoalAttempts, SUM(fieldgoals_made) AS FieldGoalsMade, SUM(kickoffs) AS TotalKickoffs, SUM(kickoffs_yards) AS TotalYards FROM KickerSeason INNER JOIN Kickers ON KickerSeason.kicker_id = Kickers.kicker_id WHERE Kickers.kicker_id = '%s';", played.Kicker_id)

	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var kickers []model.AllProbability
	for rows.Next() {
		var kicker model.AllProbability
		rows.Scan(&kicker.Total_fieldgoals_attempts, &kicker.Total_fieldgoals_made, &kicker.Total_kickoffs_attempts, &kicker.Total_kickoffs_yards)
		kickers = append(kickers, kicker)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	w.WriteHeader(http.StatusOK)
	resp := model.JsonProbabilityResponse{Type: "Success", Message: "", Data: kickers}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
