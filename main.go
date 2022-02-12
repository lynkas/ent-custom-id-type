package main

import (
	"context"
	"ent_test/ent"
	"ent_test/ent/token"
	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	err = client.Schema.Create(ctx)
	if err != nil {
		panic(err)
	}
	a, err := client.Account.Create().SetEmail("me@me.me").Save(ctx)
	if err != nil {
		panic(err)
	}

	_, err = client.Token.Create().SetAccountID(a.ID).SetBody("token").Save(ctx)
	if err != nil {
		panic(err)
	}

	tokenWithAccount, err := client.Token.Query().Where(token.Body("token")).WithAccount().First(ctx)
	if err != nil {
		panic(err)
	}

	if tokenWithAccount.Edges.Account.ID != a.ID {
		panic("wrong account")
	}

}
