package ticTacToe

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"net/http"

	"github.com/api/structs"

	_ "github.com/go-sql-driver/mysql"
)

// HomeLink establishes a quick endpoint testing function
func HomeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

// PlayerMove receives the player and their move
func PlayerMove(w http.ResponseWriter, r *http.Request) {

	var p structs.PlayerMove
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error: "+err.Error())
	}

	json.Unmarshal(reqBody, &p)

	fmt.Println("GAME: ", p.Game)
	fmt.Println("PLAYER: ", p.Player)
	fmt.Println("MOVE: ", p.Move)
	fmt.Println("MOVECOUNT: ", p.MoveCount)

	db, err := sql.Open("mysql", "gameWriter:user_password@tcp(my_db:3306)/tic_tac_toe")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Database connection successful")
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO PlayerMoves(Game, Player, Move, MoveCount) VALUES (?, ?, ?, ?)")
	if err != nil {
		fmt.Println("DB ERROR: ", err.Error())
	}

	_, err = stmt.Exec(p.Game, p.Player, p.Move, p.MoveCount)
	if err != nil {
		fmt.Println("DB ERROR: ", err.Error())
	}
}
