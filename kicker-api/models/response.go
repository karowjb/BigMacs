package model

type JsonSearchResponse struct {
	Type String `json:"type"`
}

type JsonTopKicker struct {
	Type string `json:"type"`
	Message string `json:"message"`
	Data []TopKicker `json:"data"`
}

type JsonSearchResult struct {
	Type string `json:"type"`
	Message string `json:"message"`
	Data []SelectedPlayer `json:"data"`
}

type JsonPlayerProfile struct {
	Type string `json:"type"`
	Message string `json:"message"`
	Data []PlayerProfile `json:"data"`
}

type JsonKickoffStats struct {
	Type string `json:"type"`
	Message string `json:"message"`
	Data []KickerKickoff `json:"data"`
}

type JsonFieldgoalStats struct {
	Type string `json:"type"`
	Message string `json:"message"`
	Data []KickerFieldGoal `json:"data"`
}