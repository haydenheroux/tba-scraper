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
	Name    string  `json:"name"`
	Region  string  `json:"region"`
	Year    int     `json:"year"`
	Week    int     `json:"week"`
	Matches []Match `json:"matches"`
}

type Match struct {
	Set          int           `json:"set"`
	Number       int           `json:"number"`
	Type         string        `json:"type"`
	Participants []Participant `json:"participants"` // TODO Must be empty
}

type Participant struct {
	Alliance   string   `json:"alliance"`
	TeamNumber int      `json:"teamNumber"`
	Metrics    []Metric `json:"metrics"`
}

type Metric struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
