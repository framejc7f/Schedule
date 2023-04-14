package structs

type  Week struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Lesson struct {
	Id       string `json:"id"`
	Time     string `json:"time"`
	Subgroup int `json:"subgroup"`
	Subject  string `json:"subject"`
	Сlassroom string `json:"Сlassroom"`
	Type string `json:"type"`
}

type Day struct {
	Id      string   `json:"id"`
	Data    string   `json:"data"`
	Lessons []Lesson `json:"lessons"`
}
