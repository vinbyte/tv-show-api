package forms

// ScheduleForm is
type ScheduleForm struct {
	ID      int           `json:"id"`
	Name    string        `json:"name"`
	Season  int           `json:"season"`
	Number  int           `json:"number"`
	Airdate string        `json:"airdate"`
	Airtime string        `json:"airtime"`
	Runtime int           `json:"runtime"`
	Image   scheduleImage `json:"image"`
	Summary string        `json:"summary"`
	Show    scheduleShow  `json:"show"`
}

type scheduleImage struct {
	Medium   string `json:"medium"`
	Original string `json:"original"`
}

type scheduleShow struct {
	ID           int              `json:"id"`
	URL          string           `json:"url"`
	Name         string           `json:"name"`
	Type         string           `json:"type"`
	Language     string           `json:"language"`
	Genres       []string         `json:"genres"`
	Status       string           `json:"status"`
	Runtime      int              `json:"runtime"`
	Premiered    string           `json:"premiered"`
	OfficialSite string           `json:"officialSite"`
	Schedule     scheduleSchedule `json:"schedule"`
	Rating       scheduleRating   `json:"rating"`
	Weight       int              `json:"weight"`
	Network      scheduleNetwork  `json:"network"`
	Image        scheduleImage    `json:"image"`
	Summary      string           `json:"summary"`
}

type scheduleNetwork struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type scheduleRating struct {
	Average float32 `json:"average"`
}

type scheduleSchedule struct {
	Time string   `json:"time"`
	Days []string `json:"days"`
}
