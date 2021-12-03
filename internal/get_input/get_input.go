package get_input

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"

	"github.com/joho/godotenv"
)

func sessionCookieValue() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file\n")
	}

	return os.Getenv("SESSION")
}

func GetInput(url string) string {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalf("Got error while creating cookie jar #{err.Error()}\n")
	}
	client := http.Client{
		Jar: jar,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error creating request for %s\n", url)
	}
	req.AddCookie(&http.Cookie{
		Name:   "session",
		Value:  sessionCookieValue(),
		MaxAge: 300,
	})

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error getting input url \"#{url}\": #{err.Error()}\n")
	}
	defer resp.Body.Close()

	fmt.Printf("Status code %d\n", resp.StatusCode)

	var body string
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Error reading body: #{err.Error()}\n")
		}
		body = string(bodyBytes)
	}

	return body
}
