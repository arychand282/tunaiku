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
		Date_Data   string
		Open        int
		High        int
		Low         int
		Close       int
		Adj_Close   int
		Volume      int
		Range_Value int16
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
			data  Dataprofit
			datas []Dataprofit
		)
		rows, err := db.Query("SELECT * from data;")
		if err != nil {
			fmt.Println(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&data.Date_Data, &data.Open, &data.High, &data.Low, &data.Close, &data.Adj_Close, &data.Volume)
			datas = append(datas, data)

			if err != nil {
				fmt.Println(err.Error())
			}
		}
		defer rows.Close()

		type DataProfit struct {
			DateBuy  string
			DateSell string
			Profit   int
		}
		var dataProfit DataProfit
		var dataProfits []DataProfit

		for i := 0; i < len(datas)-1; i++ {
			dataProfit.DateSell = "minus value"
			dataProfit.Profit = 0
			for j := i + 1; j < len(datas); j++ {
				tempProfit := datas[j].Close - datas[i].Open

				if tempProfit > dataProfit.Profit {
					dataProfit.DateSell = datas[j].Date_Data
					dataProfit.Profit = tempProfit
				}
			}

			dataProfit.DateBuy = datas[i].Date_Data

			dataProfits = append(dataProfits, dataProfit)
		}
		fmt.Println(dataProfits)

		maxProfit := dataProfits[0].Profit
		maxDateBuy := dataProfits[0].DateBuy
		maxDateSell := dataProfits[0].DateSell
		for _, value := range dataProfits {
			if value.Profit > maxProfit {
				maxProfit = value.Profit
				maxDateBuy = value.DateBuy
				maxDateSell = value.DateSell
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"max_profit":    maxProfit,
			"max_date_buy":  maxDateBuy,
			"max_date_sell": maxDateSell,
			"data_profits":  dataProfits,
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
