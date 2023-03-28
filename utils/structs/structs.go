package structs

type Users struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Category struct {
	Id   int `json:"id"`
	Name int `json:"name"`
}

type Game struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category_id int    `json:"category_id"`
}

type Comment struct {
	Id       int    `json:"id"`
	Text     string `json:"text"`
	Users_id int    `json:"users_id"`
	Game_id  int    `json:"game_id"`
}

type Rating struct {
	Id       int `json:"id"`
	Rate     int `json:"rate"`
	Users_id int `json:"users_id"`
	Game_id  int `json:"game_id"`
}
