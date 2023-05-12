package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type song struct{
	ID     string `json:"id"`
	Title  string `json :"title"`
	Singer string `json :"singer"`
}


var songs = []song{
	{ID: "1", Title: "Have a Nice Day", Singer: "Bon Jovi"},
	{ID: "2", Title: "The Nights", Singer: "Avicii"},
	{ID: "3", Title: "One Way Ticket", Singer: "ONE OK ROCK"},	
}

func getSongs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, songs)
}


func createBook(c *gin.Context){
	var newSong song

	if err := c.BindJSON(&newSong); err != nil{
		return
	}

	songs = append(songs, newSong)
	c.IndentedJSON(http.StatusCreated, newSong)

}

func songById(c *gin.Context){
	id := c.Param("id")
	song, err := getSongById(id)
	if  err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Song not found."})
		return
	}
	c.IndentedJSON(http.StatusOK, song)

}


func getSongById(id string) (*song, error) {
	for i,s := range songs {
		if s.ID == id {
			return &songs[i], nil
		}
	}
//登録されていない場合、*songはnilを返し、第二引数としてエラーを返す
	return nil, errors.New("songs not found")
	
}



func main(){
	router := gin.Default()

	router.GET("/songs", getSongs)
	router.GET("/songs/:id",songById )
	router.POST("/songs", createBook)

	router.Run()
}