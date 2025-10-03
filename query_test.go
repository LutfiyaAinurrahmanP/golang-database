package golangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
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
	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer;"
	// QueryContext berlaku untuk query yang membutuhkan pengambilan data
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success select data to database")

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var createdAt time.Time
		var birthDate sql.NullTime
		var married bool
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id : ", id)
		fmt.Println("Name : ", name)
		if email.Valid {
			fmt.Println("Email : ", email.String)
		}
		fmt.Println("Balance : ", balance)
		fmt.Println("Rating : ", rating)
		if birthDate.Valid {
			fmt.Println("Birth_date : ", birthDate.Time.String())
		}
		fmt.Println("Married : ", married)
		fmt.Println("Created_at : ", createdAt)
	}
	defer rows.Close()
}

func TestQueryParams(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	// sql injection jika concat atau menyambungkan variabel dalam string query
	// username := "admin'; #"

	username := "admin"
	password := "admin"

	ctx := context.Background()
	sqlQuery := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, sqlQuery, username, password)
	if err != nil {
		panic(err)
	}
	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Success login", username)
	} else {
		fmt.Println("Gagal login")
	}
}

func TestInsertWithoutSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	username := "staff"
	password := "staff"

	ctx := context.Background()

	// sql injection jika menggunakan/menggabungkan value dalam string
	// script := "INSERT INTO user(username, password) VALUES ('upi', 'Upi');"

	// gunakan tanda "?" untuk menghindari sql injection
	script := "INSERT INTO user(username, password) VALUES (?, ?);"

	// ExexContext berlaku untuk query yang tidak mengembalikan data ( insert, update, delete)
	// _, err := db.ExecContext(ctx, script)

	// Mencegah sql injection dengan memasukkan pada context
	_, err := db.ExecContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert data to database")
}
