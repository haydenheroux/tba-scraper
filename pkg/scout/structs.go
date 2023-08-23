package scout

type Team struct {
	Number  int      `json:"number"`
	Name    string   `json:"name"`
	Region  string   `json:"region"`
	Seasons []Season `json:"seasons"`
}

type Season struct {
	Year   int     `json:"year"`
	Robots []Robot `json:"robots"`
	Events []Event `json:"events"`
}

type Robot struct {
	Name string `json:"name"`
}

type Event struct {
	Code    string  `json:"code"`
	Name    string  `json:"name"`
	Region  string  `json:"region"`
	Year    int     `json:"year"`
	Week    int     `json:"week"`
	Matches []Match `json:"matches"`
}

type Match struct {
	Set       int        `json:"set"`
	Number    int        `json:"number"`
	Type      string     `json:"type"`
	Alliances []Alliance `json:"alliances"` // TODO Must be empty
}

type Alliance struct {
	Color        string            `json:"color"`
	Metrics      map[string]string `json:"metrics"`
	Participants []Participant     `json:"participants"`
}

type Participant struct {
	TeamNumber int               `json:"teamNumber"`
	Metrics    map[string]string `json:"metrics"`
}
