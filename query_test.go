package golangdatabase

import (
	"context"
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO customer(id, name) VALUES ('upi', 'Upi');"
	// ExexContext berlaku untuk query yang tidak mengembalikan data ( insert, update, delete)
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert data to database")
}

func TestSelect(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT id, name FROM customer;"
	// QueryContext berlaku untuk query yang membutuhkan pengambilan data
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	
	fmt.Println("Success select data to database")

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id : ", id)
		fmt.Println("Name : ", name)
	}
	defer rows.Close()
}