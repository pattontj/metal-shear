package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	//		"sync"
	"time"

	"github.com/gin-gonic/gin"
	"net/http"
	// "github.com/gin-contrib/static"

	"database/sql"
	"github.com/go-sql-driver/mysql"

	"pattontj/metal-shear/server"
)

type testClip struct {
	Link       string `json:"link"`
	TsBegin    string `json:"tsBegin"`
	TsEnd      string `json:"tsEnd"`
	StreamerID string `json:"streamerID"`
}

func getHome(c *gin.Context) {
	// c.HTML(http.StatusOK, "index.html", nil)
}

func getStreamers(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	rows, err := db.Query("SELECT * FROM streamer")
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, nil)
	}

	defer rows.Close()

	var vt = []server.Streamer{}

	for rows.Next() {
		var streamer server.Streamer
		if err := rows.Scan(&streamer.ID, &streamer.Name, &streamer.Channel, &streamer.Affiliation); err != nil {
			log.Fatal(err)
		}

		vt = append(vt, streamer)
	}

	c.IndentedJSON(http.StatusOK, vt)
}

func getStreamerPage(c *gin.Context) {
	page := c.Param("page")

	pageNum, err := strconv.Atoi(page)
	if err != nil {
		log.Fatal(err)
	}

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	rows, err := db.Query("SELECT * FROM streamer LIMIT 2 OFFSET ?", strconv.Itoa(pageNum*2))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, nil)
	}

	defer rows.Close()

	var vt = []server.Streamer{}

	for rows.Next() {
		var streamer server.Streamer
		if err := rows.Scan(&streamer.ID, &streamer.Name, &streamer.Affiliation); err != nil {
			log.Fatal(err)
		}

		vt = append(vt, streamer)
	}

	c.IndentedJSON(http.StatusOK, vt)

}

func postStreamers(c *gin.Context) {
	var newStreamer server.Streamer

	if err := c.BindJSON(&newStreamer); err != nil {
		return
	}

	query := "INSERT INTO streamer (title, affiliation) VALUES(?,?)"

	insert, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := insert.Exec(newStreamer.Name, newStreamer.Affiliation)
	if err != nil {
		log.Fatal(err)
	}

	_, rErr := resp.LastInsertId()
	if rErr != nil {
		log.Fatal(rErr)
	}

	c.IndentedJSON(http.StatusCreated, newStreamer)
}

func getStreamerByName(c *gin.Context) {
	name := c.Param("name")

	rows, err := db.Query("SELECT * FROM streamer WHERE title=?", name)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var streamer server.Streamer
	rows.Next()
	if err := rows.Scan(&streamer.ID, &streamer.Name, &streamer.Affiliation); err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, streamer)
	return

	// c.IndentedJSON(http.StatusNotFound, gin.H{"message": "vtuber not found"})
}

func getClips(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	query := "SELECT clips.id, clips.link, streamer.title, streamer.channel, streamer.affiliation " +
		"FROM clips " +
		"JOIN streamer ON streamer.id=clips.streamerID"

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var clips = []server.Clip{}

	for rows.Next() {
		var clp server.Clip
		if err := rows.Scan(&clp.ID, &clp.Link, &clp.Streamer.Name, &clp.Streamer.Channel, &clp.Streamer.Affiliation); err != nil {
			log.Fatal(err)
		}
		clips = append(clips, clp)
	}

	c.IndentedJSON(http.StatusOK, clips)
	return
}

func postClip(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	var newClip server.Clip

	if err := c.BindJSON(&newClip); err != nil {
		log.Fatal(err)
	}

	// NOTE: This query does not have to check for duplicate clips, this is handled by the database
	query := "INSERT INTO clips (link, beginTime, endTime, streamerID) VALUES(?,?,?,?)"

	insert, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := insert.Exec(&newClip.Link, &newClip.TsBegin, &newClip.TsEnd, &newClip.StreamerID)
	if err != nil {
		mes, ok := err.(*mysql.MySQLError) // grabs actual err struct
		if !ok {
			log.Fatal("something is desperately wrong: ", ok)
		}
		// if errcode is dupe key, we can just ignore it
		if mes.Number == 1062 {
			c.Status(http.StatusTeapot) // cheeky status
		} else {
			log.Fatal(err)
		}
	} else { // if there's no err
		_, rErr := resp.LastInsertId()
		if rErr != nil {
			log.Fatal(rErr)
		}

		c.IndentedJSON(http.StatusCreated, newClip)
	}
}

func example(c *gin.Context) {
	c.HTML(http.StatusOK,
		"index.html",
		nil,
	)
}

var db *sql.DB

func main() {

	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "172.22.48.1:3306",
		DBName:               "metal-shear",
		AllowNativePasswords: true,
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected to DB")

	chuubas := server.LocalGetStreamers(db)

	// spin up a light thread to run a yt scrape
	ticker := time.NewTicker(6 * time.Hour)
	go server.RunMonitorTick(ticker, chuubas)

	fmt.Println("'RunMonitorTick' running on a 6 hour interval")

	router := gin.Default()

	//	router.Use( static.Serve( "/", static.LocalFile( "./StartPage", false ) ) )
	//	router.Use( static.Serve( "/", static.LocalFile( "./Shoggoth", true ) ) )

	//	router.Static("/js", "./js")
	//	router.Static("/css", "./css")
	//	router.LoadHTMLFiles("Shoggoth/index.html")

	router.GET("/shoggoth", nil)

	api := router.Group("/api")
	{
		api.GET("/streamer", getStreamers)
		api.GET("/streamer/page", getStreamers)
		api.GET("/streamer/page/:page", getStreamerPage)
		api.GET("/streamer/:name", getStreamerByName)

		api.GET("/clips", getClips)

		api.POST("clips/post", postClip)
		api.POST("streamer", postStreamers)
	}

	router.Run("localhost:8080")
}

// TODO: JOIN tables, use vtubers for foreign key in clips
// https://dataschool.com/learn-sql/joins/
