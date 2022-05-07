package model

type Team struct {
	Team_name       string `json:"teamname"`
	Team_location   string `json:"teamlocation"`
	Team_abbr       string `json:"teamabbr"`
	Team_id         string `json:"teamid"`
	Home_stadium    string `json:"homestadium"`
	Home_surface    string `json:"homesurface"`
	Home_roof       string `json:"homeroof"`
	Division        string `json:"division"`
	Conference      string `json:"conference"`
	Primary_color   string `json:"primarycolor"`
	Secondary_color string `json:"secondarycolor"`
}

type Kicker struct {
	First_name    string `json:"kickerfirstname"`
	Last_name     string `json:"kickerlastname"`
	Jersey_number string `json:"jerseynumber"`
	Birthdate     string `json:"birthday"`
	Player_weight int    `json:"playerweight"`
	Player_height int    `json:"playerheight"`
	Birth_place   string `json:"birthplace"`
	High_school   string `json:"highschool"`
	College       string `json:"college"`
	College_conf  string `json:"collegeconf"`
	Rookie_year   int    `json:"rookieyear"`
	Player_status string `json:"playerstatus"`
	Kicker_id     string `json:"kickerid"`
	Team_id       string `json:"teamid"`
}

type KickerSeason struct {
	Season_id                string `json:"seasonid"`
	Kicker_id                string `json:"kickerid"`
	Team_id                  string `json:"teamid"`
	Season_year              int    `json:"seasonyear"`
	Season_type              string `json:"seasontype"`
	Games_played             int    `json:"gamesplayed"`
	Games_started            int    `json:"gamesstarted"`
	Penalties                int    `json:"penalties"`
	Penalty_yards            int    `json:"penaltyyards"`
	Kickoffs                 int    `json:"kickoffs"`
	Kickoffs_endzone         int    `json:"kickoffsendzone"`
	Kickoffs_inside_twenty   int    `json:"kickoffsinsidetwenty"`
	Kickoffs_return_yards    int    `json:"kickoffsreturnyards"`
	Kickoffs_touchbacks      int    `json:"kickoffstouchbacks"`
	Kickoffs_yards           int    `json:"kickoffsyards"`
	Kickoffs_out_of_bounds   int    `json:"kickoffsoutofbounds"`
	Kickoffs_onside_attempts int    `json:"kickoffsonsideattempts"`
	Kickoffs_onside_success  int    `json:"kickoffsonsidesuccess"`
	Kickoffs_squibs          int    `json:"kickoffssquibs"`
	Fieldgoals_attempts      int    `json:"fieldgoalsattempts"`
	Fieldgoals_made          int    `json:"fieldgoals_made"`
	Fieldgoals_blocked       int    `json:" fieldgoalsblocked"`
	Fieldgoal_yards          int    `json:"fieldgoalyards"`
	Fieldgoal_longest        int    `json:"fieldgoallongest"`
	Xp_attempts              int    `json:"xpattempts"`
	Xp_made                  int    `json:"xpmade"`
	Xp_blocked               int    `json:"xpblocked"`
}

type TopKickerFG struct {
	First_name             string `json:"kickerfirstname"`
	Last_name              string `json:"kickerlastname"`
	Fieldgoals_attempts    int    `json:"fieldgoalsattempts"`
	Fieldgoals_made        int    `json:"fieldgoals_made"`
	Season_year            int    `json:"seasonyear"`
	Kickoffs_endzone       int    `json:"kickoffsendzone"`
	Kickoffs_inside_twenty int    `json:"kickoffsinsidetwenty"`
	Games_played           int    `json:"gamesplayed"`
	Games_started          int    `json:"gamesstarted"`
}
type TopKickerK struct {
	First_name             string `json:"kickerfirstname"`
	Last_name              string `json:"kickerlastname"`
	Kickoffs               int    `json:"kickoffs"`
	Kickoffs_yards         int    `json:"kickoffsyards"`
	Season_year            int    `json:"seasonyear"`
	Kickoffs_endzone       int    `json:"kickoffsendzone"`
	Kickoffs_inside_twenty int    `json:"kickoffsinsidetwenty"`
	Games_played           int    `json:"gamesplayed"`
	Games_started          int    `json:"gamesstarted"`
}

type PlayerProfile struct {
	First_name          string `json:"kickerfirstname"`
	Last_name           string `json:"kickerlastname"`
	Jersey_number       string `json:"jerseynumber"`
	Birthdate           string `json:"birthday"`
	Player_weight       int    `json:"playerweight"`
	Player_height       int    `json:"playerheight"`
	Birth_place         string `json:"birthplace"`
	High_school         string `json:"highschool"`
	College             string `json:"college"`
	Kickoffs            int    `json:"kickoffs"`
	Kickoffs_squibs     int    `json:"kickoffssquibs"`
	Fieldgoals_attempts int    `json:"fieldgoalsattempts"`
	Fieldgoals_made     int    `json:"fieldgoals_made"`
	Fieldgoals_blocked  int    `json:" fieldgoalsblocked"`
	Fieldgoal_yards     int    `json:"fieldgoalyards"`
	Fieldgoal_longest   int    `json:"fieldgoallongest"`
	Xp_attempts         int    `json:"xpattempts"`
	Xp_made             int    `json:"xpmade"`
}

type SelectedPlayer struct {
	First_name          string `json:"kickerfirstname"`
	Last_name           string `json:"kickerlastname"`
	Kickoffs            int    `json:"kickoffs"`
	Kickoffs_squibs     int    `json:"kickoffssquibs"`
	Fieldgoals_attempts int    `json:"fieldgoalsattempts"`
	Fieldgoals_made     int    `json:"fieldgoals_made"`
	Fieldgoal_yards     int    `json:"fieldgoalyards"`
	Fieldgoal_longest   int    `json:"fieldgoallongest"`
	Xp_attempts         int    `json:"xpattempts"`
	Xp_made             int    `json:"xpmade"`
}

type KickerFieldGoal struct {
	Fieldgoals_attempts int    `json:"fieldgoalsattempts"`
	Fieldgoals_made     int    `json:"fieldgoals_made"`
	Season_year         int    `json:"seasonyear"`
	Season_type         string `json:"seasontype"`
}

type KickerKickoff struct {
	Kickoffs                 int `json:"kickoffs"`
	Kickoffs_endzone         int `json:"kickoffsendzone"`
	Kickoffs_inside_twenty   int `json:"kickoffsinsidetwenty"`
	Kickoffs_return_yards    int `json:"kickoffsreturnyards"`
	Kickoffs_touchbacks      int `json:"kickoffstouchbacks"`
	Kickoffs_yards           int `json:"kickoffsyards"`
	Kickoffs_out_of_bounds   int `json:"kickoffsoutofbounds"`
	Kickoffs_onside_attempts int `json:"kickoffsonsideattempts"`
	Kickoffs_onside_success  int `json:"kickoffsonsidesuccess"`
	Kickoffs_squibs          int `json:"kickoffssquibs"`
}

type KickerInfo struct {
	Kicker_id     string `json:"kickerid"`
	First_name    string `json:"kickerfirstname"`
	Last_name     string `json:"kickerlastname"`
	Jersey_number string `json:"jerseynumber"`
	Player_status string `json:"playerstatus"`
}

type GetPlayer struct {
	First_name string `json:"kickerfirstname"`
	Last_name  string `json:"kickerlastname"`
	Team_id string `json:"kickerteamid"`
	HeightMin int `json:"kickerheightmin"`
	HeightMax int `json:"kickerheightmax"`
	WeightMin int `json:"kickerweightmin"`
	WeightMax int `json:"kickerweightmax"`
}
type GetParameters struct {
	Fieldgoals_made int `json:"fieldgoals_made"`
}
