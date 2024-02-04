package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Message string `json:"message"`
}
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func callApianother(c *gin.Context) {
	url := "https://ww2.mncsekuritas.id/backendweb/randomPromo"
	method := "GET"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("user_id", " 1")
	_ = writer.WriteField("level", " 1")
	_ = writer.WriteField("judulin", " ada")
	_ = writer.WriteField("judulEn", " ada")
	_ = writer.WriteField("isi", " <p>asda</p>")
	_ = writer.WriteField("isiEng", " <p>dasdad</p>")
	_ = writer.WriteField("category_donwload_id", "12")
	file, errFile8 := os.Open("/Applications/XAMPP/xamppfiles/htdocs/rapor_siswa/image.php.jpg")
	defer file.Close()
	part8,
		errFile8 := writer.CreateFormFile("gambar", filepath.Base("/Applications/XAMPP/xamppfiles/htdocs/rapor_siswa/image.php.jpg"))
	_, errFile8 = io.Copy(part8, file)
	if errFile8 != nil {
		fmt.Println(errFile8)
		return
	}
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwOlwvXC93dzIubW5jc2VrdXJpdGFzLmlkXC9hcGlcL3YxXC9sb2dpbiIsImlhdCI6MTcwNjc4MTM1MCwibmJmIjoxNzA2NzgxMzUwLCJqdGkiOiJQMzdpcEJqVWJOb2NaZnhVIiwic3ViIjoxLCJwcnYiOiIyM2JkNWM4OTQ5ZjYwMGFkYjM5ZTcwMWM0MDA4NzJkYjdhNTk3NmY3In0.cV8vjbpQkNKVnbp7Ed2dnzbbXEb3-0uUgainVuqFZm4")
	req.Header.Add("Cookie", "laravel_session=eyJpdiI6Im5ONXVvcStJUy9oTDFqbDlXUU42OUE9PSIsInZhbHVlIjoiMU4rbVh5Q0w2Mm41VW5qQ2Z4eDJSNEI3M2p4d2hpSkE1cnhoRktwSzN5VE1TN29aRk9vNlNodytydVAzZkVsREp5RlJWME1pdGtLcFhLQUh3WUNIMnFGQS9CajVIZGpyMy9tbStIUmt5REZvOXRnNHFiMW1GVkRoY3J2alZ1czQiLCJtYWMiOiI0MzExYWNkNDM0NjdmMTVlNTdkN2MyYjhkMDU2NDc4MjFmNjEyNTllZDcwNDZkY2UzNjU3NmIxOGI1NTJjNjMwIiwidGFnIjoiIn0%3D")
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var responseData ResponseData
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println("Decoded JSON:", responseData)

	c.IndentedJSON(http.StatusOK, responseData)

	// Print the response body
	fmt.Println(string(body))
	// Send JSON response using c.IndentedJSON
	// c.IndentedJSON(http.StatusOK, gin.H{

	// 	"data":    string(body),
	// 	"message": "Request processed successfully"})

	// c.IndentedJSON(http.StatusOK, gin.H{

	// 	"data":    string(body),
	// 	"message": "Request processed successfully"})

}

func listAlbum(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/v1/album", listAlbum)
	router.GET("/list/promo", callApianother)
	router.Run("localhost:8080")
}
