package loader

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
)

func GetInput(fileName, year, day, sessionID string) string {
	url := fmt.Sprintf("https://adventofcode.com/%s/day/%s/input", year, day)

	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("%v", err.Error())
	}
	info, err := f.Stat()
	if err != nil {
		panic(err)
	}

	if info.Size() == 0 {
		f.Seek(0, 0)
		r, err := http.NewRequest("GET", url, nil)
		if err != nil {
			panic(err)
		}

		c := &http.Cookie{
			Name:  "session",
			Value: sessionID,
		}
		r.AddCookie(c)

		res, err := http.DefaultClient.Do(r)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		scanner := bufio.NewScanner(res.Body)

		for scanner.Scan() {
			input := scanner.Text()
			f.WriteString(input + "\n")
		}
	}

	f.Seek(0, 0)
	scanner := bufio.NewScanner(f)
	buffer := bytes.Buffer{}
	for scanner.Scan() {
		buffer.Write(scanner.Bytes())
		buffer.Write([]byte("\n"))
	}

	return buffer.String()
}
