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

	type Dataprofit struct {
		Date_Data string
		Open      int
		High      int
		Low       int
		Close     int
		Adj_Close int
		Volume    int
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

	// PUT - update a primenumber details
	router.PUT("/primenumber", func(c *gin.Context) {
		var buffer bytes.Buffer
		id := c.Query("id")
		numbers, err := strconv.Atoi(c.PostForm("numbers"))
		stmt, err := db.Prepare("update prime_number set numbers= ?, name_number = number_to_string(?) where id= ?;")
		if err != nil {
			fmt.Print(err.Error())
		}

		if numbers > 1000000 {
			fmt.Println("Number must less than 1000000!")
			c.JSON(http.StatusForbidden, gin.H{
				"error_message": fmt.Sprintf("Number must less than 1000000!"),
			})
			return
		}

		_, err = stmt.Exec(numbers, numbers, id)
		if err != nil {
			fmt.Print(err.Error())
		}

		// Fastest way to append strings
		buffer.WriteString(strconv.Itoa(numbers))
		buffer.WriteString(" ")
		buffer.WriteString("with its name")
		defer stmt.Close()
		messagestring := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully updated to %s", messagestring),
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

		if evennumber > 1000000 || oddnumber > 1000000 {
			fmt.Println("Number must less than 1000000!")
			c.JSON(http.StatusForbidden, gin.H{
				"error_message": fmt.Sprintf("Number must less than 1000000!"),
			})
			return
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

	// GET a evenoddnumber detail
	router.GET("/evenoddnumber/:id", func(c *gin.Context) {
		var (
			evenoddnumber Evenoddnumber
			result        gin.H
		)
		id := c.Param("id")
		row := db.QueryRow("SELECT id, number_evens, number_odds, total, name_total from even_odd_number where id = ?;", id)
		err = row.Scan(&evenoddnumber.Id, &evenoddnumber.Number_Evens, &evenoddnumber.Number_Odds, &evenoddnumber.Total, &evenoddnumber.Name_Total)
		if err != nil {
			// if no results send null
			result = gin.H{
				"result": nil,
				"count":  0,
			}
		} else {
			result = gin.H{
				"result": evenoddnumber,
				"count":  1,
			}
		}

		c.JSON(http.StatusOK, result)
	})

	// GET all evenoddnumbers
	router.GET("/evenoddnumbers", func(c *gin.Context) {
		var (
			evenoddnumber  Evenoddnumber
			evenoddnumbers []Evenoddnumber
		)
		rows, err := db.Query("SELECT * from even_odd_number;")
		if err != nil {
			fmt.Println(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&evenoddnumber.Id, &evenoddnumber.Number_Evens, &evenoddnumber.Number_Odds, &evenoddnumber.Total, &evenoddnumber.Name_Total)
			evenoddnumbers = append(evenoddnumbers, evenoddnumber)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"result": evenoddnumbers,
			"count":  len(evenoddnumbers),
		})
	})

	// DELETE evenoddnumber
	router.DELETE("/evenoddnumber", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("id"))
		fmt.Println(id)
		stmt, err := db.Prepare("DELETE from even_odd_number where id = ?;")
		if err != nil {
			fmt.Println(err.Error())
		}

		_, err = stmt.Exec(id)

		if err != nil {
			fmt.Println(err.Error())
		}

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully deleted number: %s", strconv.Itoa(id)),
		})
	})

	// ============================================================================

	// POST nominalamount details
	router.POST("/nominalamount", func(c *gin.Context) {
		nominalamount := 2000000
		total := float32(0)

		for {
			interestrate := float32(nominalamount * 3 / 100)
			total := total + float32(nominalamount) + interestrate

			stmt, err := db.Prepare("INSERT INTO nominal_amount (nominal, additional_number, total) values (?, ?, ?);")
			if err != nil {
				fmt.Println(err.Error())
			}

			_, err = stmt.Exec(nominalamount, interestrate, total)

			if err != nil {
				fmt.Print(err.Error())
				break
			}

			nominalamount = int(total)
			if total > 15000000 {
				break
			}
		}

		// Fastest way to append strings
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("interation successfully created"),
		})
	})

	// ============================================================================

	router.GET("/bestprofits", func(c *gin.Context) {
		var (
			dataprofit  Dataprofit
			dataprofits []Dataprofit
		)
		rows, err := db.Query("SELECT * from data;")
		if err != nil {
			fmt.Println(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&dataprofit.Date_Data, &dataprofit.Open, &dataprofit.High, &dataprofit.Low, &dataprofit.Close, &dataprofit.Adj_Close, &dataprofit.Volume)
			dataprofits = append(dataprofits, dataprofit)

			if err != nil {
				fmt.Println(err.Error())
			}
		}
		defer rows.Close()

		// var storingmaxprofit []int

		// for i := 0; i < len(dataprofits); i++ {
		// 	var arrayranges []int
		// 	for j := 1; j < len(dataprofits); j++ {
		// 		rangeprofit := dataprofits[j].Close - dataprofits[i].Open
		// 		arrayranges = append(arrayranges, rangeprofit)
		// 	}

		// 	_, max, _, indexmax := minmaxProfit(arrayranges)
		// 	fmt.Println("\nmaximum profit buy day ", dataprofits[i].Date_Data, ": ", max)
		// 	fmt.Println("by index Close: ", indexmax)

		// 	storingmaxprofit = append(storingmaxprofit, max)
		// }

		for i := 0; i < len(dataprofits); i++ {
			fmt.Println("index: ", i)
			for j := i + 1; j < len(dataprofits); j++ {
				fmt.Println(dataprofits[j])
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"result": dataprofits,
			"count":  len(dataprofits),
			"sell":   0,
			"buy":    0,
		})
	})

	router.Run(":3001")
}

func minmaxProfit(array []int) (int, int, int, int) {
	var max int = array[0]
	var min int = array[0]
	var indexmax int = 0
	var indexmin int = 0

	for index, value := range array {
		fmt.Print(value, " ")
		if max < value {
			max = value
			indexmax = index
		}
		if min > value {
			min = value
			indexmin = index
		}
	}
	return min, max, indexmin, indexmax
}
