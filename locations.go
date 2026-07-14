type Location struct {
	Id     string     `json:"id"`
	Name   string     `json:"name"`
	Region struct {
		Name     string     `json:"name"`
		URL      string     `json:"url"`
	} `json:"region"`
	Names []string {

	}
}