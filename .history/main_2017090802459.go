package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/tunaiku")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}

	type Primenumber struct {
		Id          int
		Numbers     int
		Name_Number string
	}

	type Evenoddnumber struct {
		Id           int
		Number_Evens int
		Number_Odds  int
		Total        int
		Name_Total   string
	}

	router := gin.Default()

	// POST new primenumber details
	router.POST("/primenumber", func(c *gin.Context) {
		var buffer bytes.Buffer
		numbers, err := strconv.Atoi(c.PostForm("numbers"))
		stmt, err := db.Prepare("INSERT INTO prime_number (numbers, name_number) values (?, number_to_string(?));")
		if err != nil {
			fmt.Println(err.Error())
		}

		if numbers > 1000000 {
			fmt.Println("Number must less than 1000000!")
			c.JSON(http.StatusForbidden, gin.H{
				"error_message": fmt.Sprintf("Number must less than 1000000!"),
			})
			return
		}
		_, err = stmt.Exec(numbers, numbers)

		if err != nil {
			fmt.Print(err.Error())
		}

		// Fastest way to append strings
		buffer.WriteString(strconv.Itoa(numbers))
		defer stmt.Close()
		successmess := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf(" %s successfully created", successmess),
		})
	})

	// GET a primenumber detail
	router.GET("/primenumber/:id", func(c *gin.Context) {
		var (
			primenumber Primenumber
			result      gin.H
		)
		id := c.Param("id")
		row := db.QueryRow("SELECT id, numbers, name_number from prime_number where id = ?;", id)
		err = row.Scan(&primenumber.Id, &primenumber.Numbers, &primenumber.Name_Number)
		if err != nil {
			// if no results send null
			result = gin.H{
				"result": nil,
				"count":  0,
			}
		} else {
			result = gin.H{
				"result": primenumber,
				"count":  1,
			}
		}
		c.JSON(http.StatusOK, result)
	})

	// GET all primenumbers
	router.GET("/primenumbers", func(c *gin.Context) {
		var (
			primenumber  Primenumber
			primenumbers []Primenumber
		)
		rows, err := db.Query("SELECT id, numbers, name_number from prime_number;")
		if err != nil {
			fmt.Println(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&primenumber.Id, &primenumber.Numbers, &primenumber.Name_Number)
			primenumbers = append(primenumbers, primenumber)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"result": primenumbers,
			"count":  len(primenumbers),
		})
	})

	// ============================================================================

	// POST new evenoddnumber details
	router.POST("/evenoddnumber", func(c *gin.Context) {
		var buffer bytes.Buffer
		evennumber, err := strconv.Atoi(c.PostForm("evennumber"))
		oddnumber, err := strconv.Atoi(c.PostForm("oddnumber"))
		totalnumber := evennumber + oddnumber
		stmt, err := db.Prepare("INSERT INTO even_odd_number (number_evens, number_odds, total, name_total) values (?, ?, ?, number_to_string(?));")
		if err != nil {
			fmt.Println(err.Error())
		}
		_, err = stmt.Exec(evennumber, oddnumber, totalnumber, totalnumber)

		if err != nil {
			fmt.Print(err.Error())
		}

		// Fastest way to append strings
		buffer.WriteString("even number: " + strconv.Itoa(evennumber) + ", ")
		buffer.WriteString("odd number: " + strconv.Itoa(oddnumber) + ", ")
		buffer.WriteString("total: " + strconv.Itoa(totalnumber))
		defer stmt.Close()
		successmess := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf(" %s successfully created", successmess),
		})
	})

	router.Run(":3001")
}