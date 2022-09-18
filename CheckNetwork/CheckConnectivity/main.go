package main

import (
	"fmt"
	"net/http"
)

func checkReturnCode(URL []string) {

	for _, v := range URL {

		resp, err := http.Get(v)
		if err != nil {
			fmt.Println("HTTP call failed:", err)
		}

		defer resp.Body.Close()

		// statusOK := response.StatusCode >= 200 && response.StatusCode < 300
		// if !statusOK {
		// 	fmt.Println("Non-OK HTTP status:", response.StatusCode)
		// }

		if resp.StatusCode != http.StatusOK {
			fmt.Println("Non-OK HTTP status:", resp.StatusCode, "from site", v)
		}

		fmt.Println(resp.StatusCode)
	}
}

func main() {

	checkReturnCode([]string{"https://www.metro.ca/", "https://www.rds.ca/"})

}
