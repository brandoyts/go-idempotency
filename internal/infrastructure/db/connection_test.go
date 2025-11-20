package db

import (
	"fmt"
	"testing"
)

func TestConnection(t *testing.T) {
	credentials := Credentials{
		Host:         "localhost",
		User:         "secretuser",
		Password:     "secretpassword",
		DatabaseName: "go-idempotency",
		Port:         "5432",
	}
	conn, err := New(&credentials)

	if err != nil {
		t.Error(err)
	}

	fmt.Println(conn)
}
