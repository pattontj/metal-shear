package main

import(
		"os"
		"fmt"
		"log"
		"strconv"
//		"sync"

		"net/http"
		"github.com/gin-gonic/gin"
		// "github.com/gin-contrib/static"

		"database/sql"
		"github.com/go-sql-driver/mysql"
)

type vtuber struct {
	ID 			string `json:"id"` 
	Name 		string `json:"name"`
	Channel		string `json:"channel"`
	Affiliation string `json:"affiliation"`
}



type clip struct {
	ID 			string `json:"id"`
	Link 		string `json:"link"`
	TsBegin 	string `json:"tsBegin"`
	TsEnd 		string `json:"tsEnd"`
	VtuberID 	string `json:"vtuberID"`
	Vtuber 		vtuber `json:"vtuber"`
}


func getHome(c *gin.Context) {
	// c.HTML(http.StatusOK, "index.html", nil)
}

func getVtubers(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	rows, err := db.Query("SELECT * FROM vtuber")
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, nil)
	}
	
	defer rows.Close()

	var vt = []vtuber {}

	for rows.Next() {
		var chuuba vtuber
		if err := rows.Scan(&chuuba.ID, &chuuba.Name, &chuuba.Affiliation); err != nil {
			log.Fatal(err)
		}

		vt = append(vt, chuuba)
	}

	c.IndentedJSON(http.StatusOK, vt)
}


func getVtuberPage(c *gin.Context) {
	page := c.Param("page")

	pageNum, err:= strconv.Atoi(page)
	if err != nil {
		log.Fatal(err)
	}

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	rows, err := db.Query("SELECT * FROM vtuber LIMIT 2 OFFSET ?", strconv.Itoa(pageNum*2))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, nil)
	}
	
	defer rows.Close()


	var vt = []vtuber {}

	for rows.Next() {
		var chuuba vtuber
		if err := rows.Scan(&chuuba.ID, &chuuba.Name, &chuuba.Affiliation); err != nil {
			log.Fatal(err)
		}

		vt = append(vt, chuuba)
	}

	c.IndentedJSON(http.StatusOK, vt)


}

// INSERT INTO vtuber
// (title, affiliation)
// VALUES
// ('Inugami Korone', "Hololive"),


func postVtubers(c *gin.Context) {
	var newVtuber vtuber

	if err := c.BindJSON(&newVtuber); err != nil {
		return
	}

	query := "INSERT INTO vtuber (title, affiliation) VALUES(?,?)"
	
	insert, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := insert.Exec(newVtuber.Name, newVtuber.Affiliation)
	if err != nil {
		log.Fatal(err)
	}

	_, rErr := resp.LastInsertId()
	if rErr != nil {
		log.Fatal(rErr)
	}

	c.IndentedJSON(http.StatusCreated, newVtuber)
}

func getVtuberByName(c *gin.Context) {
	name := c.Param("name")

	rows, err := db.Query("SELECT * FROM vtuber WHERE title=?", name)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var chuuba vtuber
	rows.Next()
	if err := rows.Scan( &chuuba.ID, &chuuba.Name, &chuuba.Affiliation ); err != nil {
		log.Fatal(err) 
	}

	c.IndentedJSON(http.StatusOK, chuuba)
	return 
		

	// c.IndentedJSON(http.StatusNotFound, gin.H{"message": "vtuber not found"})
}


// this func is actually a fucking nightmare 
// TODO: un-thread this, original err was referring to the wrong row

/*
func getClips( c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	// wait group for threading first row query
	var wg sync.WaitGroup

	// wrap clips in mutex container
	type Container struct {
		mu sync.Mutex
		clips []clip
	}

	clipsContainer := Container{}

	wg.Add(1)

	go func(c *Container) {

		// lock mutex
		c.mu.Lock()
		defer c.mu.Unlock()

		// sends back done signal at end of func, wg.Wait picks this up
		defer wg.Done()

		rows, err := db.Query("SELECT * FROM clips")
		if err != nil {
			log.Fatal(err)
		}


		for rows.Next() {
			var clp clip
			
			if err := rows.Scan(&clp.ID, &clp.Link, &clp.TsBegin, &clp.TsEnd, &clp.VtuberID); err != nil {
				log.Fatal(err)
			}

			c.clips = append(c.clips, clp)
		}

		rows.Close()
	}(&clipsContainer)

	// wait for goroutine to finish (hopefully avoids race condition with rows)
	wg.Wait()

	actualClips := clipsContainer.clips

	var vt = []vtuber {}

	vtuberRow, err := db.Query("SELECT * FROM vtuber")
	if err != nil {
		log.Fatal(err)
	}

	for vtuberRow.Next() {
		var chuuba vtuber
		if err := vtuberRow.Scan(&chuuba.ID, &chuuba.Name, &chuuba.Channel, &chuuba.Affiliation); err != nil {
			log.Fatal(err)
		}
		vt = append(vt, chuuba)
	}

	// match vtuber info to clip and copy into struct
	for i, c := range actualClips {	
		for _, v := range vt {
			if c.VtuberID == v.ID {
			actualClips[i].Vtuber = v
			}
		} 
	}

	vtuberRow.Close()

	c.IndentedJSON(http.StatusOK, actualClips)
}
*/

func getClips( c *gin.Context ) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	query := "SELECT clips.id, clips.link, vtuber.title, vtuber.channel, vtuber.affiliation " + 
				"FROM clips " + 
				"JOIN vtuber ON vtuber.id=clips.vtuberID"

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var clips = []clip {}

	for rows.Next() {
		var clp clip
		if err := rows.Scan( &clp.ID, &clp.Link, &clp.Vtuber.Name, &clp.Vtuber.Channel, &clp.Vtuber.Affiliation ); err != nil {
			log.Fatal(err)
		}
		clips = append(clips, clp)
	}

	c.IndentedJSON(http.StatusOK, clips)
	return 

}


func example(c *gin.Context) {
	c.HTML( http.StatusOK,
		"index.html",
		nil,
	)
}


var db *sql.DB

func main() {

	test := os.Getenv("HOME")
	fmt.Println(test)

	cfg := mysql.Config{
		User: os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net: "tcp",
		Addr: "127.0.0.1:3306",
		DBName: "vtubers",
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

	router := gin.Default()


//	router.Use( static.Serve( "/", static.LocalFile( "./StartPage", false ) ) )

//	router.Use( static.Serve( "/", static.LocalFile( "./Shoggoth", true ) ) )

//	router.Static("/js", "./js")
//	router.Static("/css", "./css")
//	router.LoadHTMLFiles("Shoggoth/index.html")
	
	// router.GET("/", example)
	// router.GET("/examples", example)
	// router.GET("/about", example)
	// router.GET("/contact", example)

	router.GET("/shoggoth", nil)


	api := router.Group("/api")
	{
		api.GET("/vtubers", getVtubers)
		api.GET("/vtubers/page", getVtubers)
		api.GET("/vtubers/page/:page", getVtuberPage)
		api.GET("/vtubers/:name", getVtuberByName)


		api.GET("/clips", getClips)

		api.POST("vtubers", postVtubers)
	}

	router.Run("localhost:8080")


	fmt.Println("test")

}



// TODO: JOIN tables, use vtubers for foreign key in clips
// https://dataschool.com/learn-sql/joins/