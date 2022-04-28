package model

type Team struct {
	team_name string `json:"teamname"`
    team_location string `json:"teamlocation"`
    team_abbr string `json:"teamabbr"`
    team_id string `json:"teamid"`
    home_stadium string `json:"homestadium"`
    home_surface string `json:"homesurface"`
    home_roof string `json:"homeroof"`
    division string `json:"division"`
    conference string `json:"conference"`
    primary_color string `json:"primarycolor"`
    secondary_color string `json:"secondarycolor"`
}

type Kicker struct {
	first_name string `json:"kickerfirstname"`
	last_name string `json:"kickerlastname"`
	jersey_number string `json:"jerseynumber"`
	birthdate string `json:"birthday"`
	player_weight int `json:"playerweight"`
    player_height int `json:"playerheight"`
    birth_place string `json:"birthplace"`
    high_school string `json:"highschool"`
    college string `json:"college"`
    college_conf string `json:"collegeconf"`
    rookie_year int `json:"rookieyear"`
    player_status string `json:"playerstatus"`
    kicker_id string `json:"kickerid"`
    team_id string `json:"teamid"`
}

type KickerSeason struct {
	season_id string `json:"seasonid"`
	kicker_id string `json:"kickerid"`
	team_id string `json:"teamid"`
	season_year int `json:"seasonyear"`
	season_type string `json:"seasontype"`
	games_played int `json:"gamesplayed"`
	games_started int `json:"gamesstarted"`
    penalties int `json:"penalties"`
    penalty_yards int `json:"penaltyyards"`
    kickoffs int `json:"kickoffs"`
    kickoffs_endzone int `json:"kickoffsendzone"`
    kickoffs_inside_twenty int `json:"kickoffsinsidetwenty"`
    kickoffs_return_yards int `json:"kickoffsreturnyards"`
    kickoffs_touchbacks int `json:"kickoffstouchbacks"`
    kickoffs_yards int `json:"kickoffsyards"`
    kickoffs_out_of_bounds int `json:"kickoffsoutofbounds"`
    kickoffs_onside_attempts int `json:"kickoffsonsideattempts"`
    kickoffs_onside_success int `json:"kickoffsonsidesuccess"`
    kickoffs_squibs int `json:"kickoffssquibs"`
    fieldgoals_attempts int `json:"fieldgoalsattempts"`
    fieldgoals_made int `json:"fieldgoals_made"`
    fieldgoals_blocked int `json:" fieldgoalsblocked"`
    fieldgoal_yards int `json:"fieldgoalyards"`
    fieldgoal_longest int `json:"fieldgoallongest"`
    xp_attempts int `json:"xpattempts"`
    xp_made int `json:"xpmade"`
    xp_blocked int `json:"xpblocked"`
}

type TopKicker struct {
	first_name string `json:"kickerfirstname"`
	last_name string `json:"kickerlastname"`
	fieldgoals_attempts int `json:"fieldgoalsattempts"`
    fieldgoals_made int `json:"fieldgoals_made"`
	kickoffs_endzone int `json:"kickoffsendzone"`
    kickoffs_inside_twenty int `json:"kickoffsinsidetwenty"`
	games_played int `json:"gamesplayed"`
	games_started int `json:"gamesstarted"`
}

type PlayerProfile struct {
	first_name string `json:"kickerfirstname"`
	last_name string `json:"kickerlastname"`
	jersey_number string `json:"jerseynumber"`
	birthdate string `json:"birthday"`
	player_weight int `json:"playerweight"`
    player_height int `json:"playerheight"`
    birth_place string `json:"birthplace"`
    high_school string `json:"highschool"`
    college string `json:"college"`
	kickoffs int `json:"kickoffs"`
	kickoffs_squibs int `json:"kickoffssquibs"`
    fieldgoals_attempts int `json:"fieldgoalsattempts"`
    fieldgoals_made int `json:"fieldgoals_made"`
    fieldgoals_blocked int `json:" fieldgoalsblocked"`
    fieldgoal_yards int `json:"fieldgoalyards"`
    fieldgoal_longest int `json:"fieldgoallongest"`
    xp_attempts int `json:"xpattempts"`
    xp_made int `json:"xpmade"`
}

type SelectedPlayer struct {
	first_name string `json:"kickerfirstname"`
	last_name string `json:"kickerlastname"`
	kickoffs int `json:"kickoffs"`
	kickoffs_squibs int `json:"kickoffssquibs"`
    fieldgoals_attempts int `json:"fieldgoalsattempts"`
    fieldgoals_made int `json:"fieldgoals_made"`
    fieldgoal_yards int `json:"fieldgoalyards"`
    fieldgoal_longest int `json:"fieldgoallongest"`
    xp_attempts int `json:"xpattempts"`
    xp_made int `json:"xpmade"`
}

type KickerFieldGoal struct {
	first_name string `json:"kickerfirstname"`
	last_name string `json:"kickerlastname"`
	fieldgoals_attempts int `json:"fieldgoalsattempts"`
    fieldgoals_made int `json:"fieldgoals_made"`
    fieldgoal_yards int `json:"fieldgoalyards"`
    fieldgoal_longest int `json:"fieldgoallongest"`
}

type KickerKickoff struct {
	first_name string `json:"kickerfirstname"`
	last_name string `json:"kickerlastname"`
	kickoffs int `json:"kickoffs"`
    kickoffs_endzone int `json:"kickoffsendzone"`
    kickoffs_inside_twenty int `json:"kickoffsinsidetwenty"`
    kickoffs_return_yards int `json:"kickoffsreturnyards"`
    kickoffs_touchbacks int `json:"kickoffstouchbacks"`
    kickoffs_yards int `json:"kickoffsyards"`
    kickoffs_out_of_bounds int `json:"kickoffsoutofbounds"`
    kickoffs_onside_attempts int `json:"kickoffsonsideattempts"`
    kickoffs_onside_success int `json:"kickoffsonsidesuccess"`
    kickoffs_squibs int `json:"kickoffssquibs"`
}