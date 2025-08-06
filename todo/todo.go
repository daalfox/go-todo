package todo

type Todo struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func (t *Todo) Validate() map[string]string {
	problems := make(map[string]string)
	if t.Title == "" {
		problems["title"] = "cannot be empty"
	}

	return problems
}
