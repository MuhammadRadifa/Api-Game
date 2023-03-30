package structs

type Users struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Roles    string `json:"role"`
}

type Category struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Game struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category_id int    `json:"category_id"`
	Category    []Category
}

type Comment struct {
	Id       string `json:"id"`
	Text     string `json:"text"`
	Users_id string `json:"users_id"`
	Game_id  int    `json:"game_id"`
}

type Rating struct {
	Id       string `json:"id"`
	Rate     int    `json:"rate"`
	Users_id string `json:"users_id"`
	Game_id  int    `json:"game_id"`
}
