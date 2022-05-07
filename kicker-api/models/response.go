package model

type JsonSearchResponse struct {
	Type string `json:"type"`
}
type JsonTeam struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Data    []Team `json:"data"`
}
type JsonTopKickerFG struct {
	Type    string        `json:"type"`
	Message string        `json:"message"`
	Data    []TopKickerFG `json:"data"`
}
type JsonTopKickerLongestFG struct {
	Type    string       `json:"type"`
	Message string       `json:"message"`
	Data    []TopKickerLongestFG `json:"data"`
}

type JsonSearchResult struct {
	Type    string           `json:"type"`
	Message string           `json:"message"`
	Data    []SelectedPlayer `json:"data"`
}

type JsonPlayerProfile struct {
	Type    string          `json:"type"`
	Message string          `json:"message"`
	Data    []PlayerProfile `json:"data"`
}

type JsonKickoffStats struct {
	Type    string          `json:"type"`
	Message string          `json:"message"`
	Data    []KickerKickoff `json:"data"`
}

type JsonFieldgoalStats struct {
	Type    string            `json:"type"`
	Message string            `json:"message"`
	Data    []KickerFieldGoal `json:"data"`
}

type JsonResponseAllKickers struct {
	Type    string       `json:"type"`
	Message string       `json:"message"`
	Data    []KickerInfo `json:"data"`
}
