// http://www.golangbootcamp.com/book/types
package ch7

import "fmt"

type User struct {
	Id             int
	Name, Location string
}

func (u *User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s", u.Name, u.Location)
}

type Player struct {
	*User
	GameId int
}

func (p *Player) String() string {
	return fmt.Sprintf("Player(Id=%d Name=%s Location=%s GameId=%d)", p.Id, p.Name, p.Location, p.GameId)
}

func NewPlayer(id int, name string, location string, gameId int) *Player {
	return &Player{User: &User{Id: id, Name: name, Location: location}, GameId: gameId}
}

func main() {
	player1 := Player{
		User:   &User{Id: 23, Name: "Tim", Location: "Cologne"},
		GameId: 4,
	}
	println(player1.Greetings())

	// Soluton
	player2 := NewPlayer(42, "Matt", "LA", 90404)
	println(player2.Greetings())
}
