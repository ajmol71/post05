package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/mactsouk/post05"
)

var MIN = 0
var MAX = 26

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func getString(length int64) string {
	startChar := "4"
	temp := ""
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == length {
			break
		}
		i++
	}
	return temp
}

func main() {
	post05.Hostname = "localhost"
	post05.Port = 5432
	post05.Username = "postgres"
	post05.Password = "pogsdata"
	post05.Database = "MSDS"

	data, err := post05.ListCourses()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range data {
		fmt.Println(v)
	}

	SEED := time.Now().Unix()
	rand.Seed(SEED)
	random_id := getString(5)

	t := post05.MSDSCourse{
		CID:    	random_id,
		CNAME:      "Database Systems",
		CPREREQ:   "400"}

	id := post05.AddCourse(t)
	if id == -1 {
		fmt.Println("There was an error adding course", t.CID)
	}

	err = post05.DeleteCourse(id)
	if err != nil {
		fmt.Println(err)
	}

	// Trying to delete it again!
	err = post05.DeleteCourse(id)
	if err != nil {
		fmt.Println(err)
	}

	id = post05.AddCourse(t)
	if id == -1 {
		fmt.Println("There was an error adding course", t.CID)
	}

	t = post05.MSDSCourse{
		CID:    random_id,
		CNAME:     "Statistics with R",
		CPREREQ:    "None"}

	err = post05.UpdateCourse(t)
	if err != nil {
		fmt.Println(err)
	}
}
