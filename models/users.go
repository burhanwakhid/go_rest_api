package models

import (
	"../db"
	"database/sql"
	"fmt"
)

type User struct {
	Id        int    `json: "id"`
	FirstName string `json: "name"`
	LastName  string `json: "akhir"`
	Email     string `json: "email"`
	Date      string `json: "date"`
}

type Users struct {
	Message string `json:"message"`
	Users   []User `json: "user"`
}

var con *sql.DB

func GetUser() Users {
	con := db.CreateCon()
	//db.CreateCon()
	sqlStatement := "SELECT id,firstname, lastname, email, reg_date FROM MyGuests order by id"

	rows, err := con.Query(sqlStatement)
	fmt.Println(rows)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		//return c.JSON(http.StatusCreated, u);
	}
	defer rows.Close()
	result := Users{}

	for rows.Next() {
		usr := User{}

		err2 := rows.Scan(&usr.Id, &usr.FirstName, &usr.LastName, &usr.Email, &usr.Date)
		// Exit if we get an error
		if err2 != nil {
			fmt.Print(err2)
		}
		result.Message = "Success"
		result.Users = append(result.Users, usr)
	}
	return result

}
