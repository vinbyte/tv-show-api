package forms

// SearchForm is
type SearchForm struct {
	Score float32    `json:"score"`
	Show  searchShow `json:"show"`
}

type searchShow struct {
	ID           int            `json:"id"`
	URL          string         `json:"url"`
	Name         string         `json:"name"`
	Type         string         `json:"type"`
	Language     string         `json:"language"`
	Genres       []string       `json:"genres"`
	Status       string         `json:"status"`
	Runtime      int            `json:"runtime"`
	Premiered    string         `json:"premiered"`
	OfficialSite string         `json:"officialSite"`
	Schedule     searchSchedule `json:"schedule"`
}

type searchSchedule struct {
	Time string   `json:"time"`
	Days []string `json:"days"`
}
