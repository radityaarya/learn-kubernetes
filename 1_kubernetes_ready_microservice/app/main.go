package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber"
)

type Quote struct {
	ID    int    `json:"id"`
	Quote string `json:"quote"`
}

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(data-mysql:3306)/database")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	app := fiber.New()

	app.Get("/get_quote", getRandomQuote)

	app.Listen(3000)
}

func getRandomQuote(c *fiber.Ctx) {
	db, err := connectDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM quotes ORDER BY RAND() LIMIT 1")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []Quote

	for rows.Next() {
		var each = Quote{}
		var err = rows.Scan(&each.Quote, &each.ID)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	c.JSON(result[0])
}
