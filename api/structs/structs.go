package structs

type PlayerMove struct {
	Game      int    `json:"game"`
	Player    string `json:"player"`
	Move      string `json:"move"`
	MoveCount int    `json:"moveCount"`
}
