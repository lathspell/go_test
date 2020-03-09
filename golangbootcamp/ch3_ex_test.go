package ch7

import (
	"testing"
)

func Test1(t *testing.T) {
	description := "NewPlayer creation"

	var expectedUser = User{
		Id:       42,
		Name:     "Matt",
		Location: "LA",
	}
	var expectedPlayer = Player{
		User:   &expectedUser,
		GameId: 90404,
	}

	if actualPlayer := NewPlayer(42, "Matt", "LA", 90404); actualPlayer.String() != expectedPlayer.String() {
		t.Fatalf("FAIL: %s\nExpected: %s\nActual:   %s", description, expectedPlayer.String(), actualPlayer.String())
	}
	t.Logf("PASS: %s", description)
}
