/*
Building a developer platform

author: ashraf-minhaj
mail: ashraf_minhaj@yahoo.com
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func docString() {
	fmt.Println("Hello, Platform user!")
	fmt.Println("If you face any difficulty feel free to reach out to the platform team")
	fmt.Println("Usage: platform start ticket EPD-123")
}

func main() {
	isWrongCommand := false
	URL, TOKEN := loadConfig()
	// godotenv.Load()

	// // load env
	// URL := os.Getenv("YOUTRACK_URL")
	// TOKEN := os.Getenv("YOUTRACK_TOKEN")

	// fmt.Println(URL)
	// fmt.Println(TOKEN)

	arguments := os.Args
	// fmt.Println(arguments)

	// check number of arguments
	if len(arguments) < 4 {
		if len(arguments) == 2 && arguments[1] == "configure" {
			configure()
		}
		isWrongCommand = true
	} else if arguments[2] == "ticket" {
		ticket := os.Args[3]
		if isValidTicket(ticket) {
			fmt.Println(ticket)

			// get ticket data
			getTicket(URL, TOKEN, ticket)
		} else {
			isWrongCommand = true
		}
	} else {
		isWrongCommand = true
	}

	// fmt.Println(isWrongCommand)
	if isWrongCommand {
		docString()
		return
	}
}

func configure() {
	homeDir, err := os.UserHomeDir()

	configDir := homeDir + "/.platform"

	err = os.MkdirAll(homeDir+"/.platform", 0755)
	if err != nil {
		fmt.Println("Could not create .platform directory:", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("YouTrack URL: ")
	URL, _ := reader.ReadString('\n')

	fmt.Print("YouTrack Token: ")
	TOKEN, _ := reader.ReadString('\n')

	URL = strings.TrimSpace(URL)
	TOKEN = strings.TrimSpace(TOKEN)

	// fmt.Print(URL, TOKEN)

	config := map[string]string{
		"YOUTRACK_URL":   URL,
		"YOUTRACK_TOKEN": TOKEN,
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Println("Could not create config:", err)
		return
	}

	configFile := configDir + "/config.json"

	err = os.WriteFile(configFile, data, 0600)
	if err != nil {
		fmt.Println("Could not save config:", err)
		return
	}

	fmt.Println("Configuration saved.")
}

func loadConfig() (string, string) {
	homeDir, _ := os.UserHomeDir()

	configFile := homeDir + "/.platform/config.json"

	data, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Println("Could not read config:", err)
		return "", ""
	}

	var config map[string]string

	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Could not read config:", err)
		return "", ""
	}

	return config["YOUTRACK_URL"], config["YOUTRACK_TOKEN"]
}

func isValidTicket(ticket string) bool {
	// ticketStartsWith := []string{"EP", "EPD"}

	parts := strings.Split(ticket, "-")
	if len(parts) == 2 {
		if parts[0] == "EP" || parts[0] == "EPD" {
			// if strconv.Atoi(parts[1])
			return true
		}
	}

	fmt.Println(parts)

	return false
}

func getTicket(baseURL string, token string, ticketID string) {
	fmt.Println("URL:", baseURL)
	fmt.Println("Ticket:", ticketID)

	url := baseURL + "/api/issues/" + ticketID + "?fields=idReadable,summary,description"

	body, status, err := get(url, token)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Status:", status)
	// fmt.Println(reflect.TypeOf(body))
	// fmt.Println(string(body))

	// unpack json
	var data map[string]interface{}

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Could not read ticket:", err)
		return
	}

	ticketSummary := data["summary"].(string)
	ticketDescription := data["description"].(string)

	fmt.Println("ID:", ticketID)
	fmt.Println("Title:", ticketSummary)
	fmt.Println("Description:", ticketDescription)

	// fmt.Println(reflect.TypeOf(ticketSummary))
	updateState(ticketID, ticketSummary, ticketDescription)
}

func updateState(ticketID string, ticketSummary string, ticketDescription string) {
	state := map[string]string{
		"ticketID":          ticketID,
		"ticketSummary":     ticketSummary,
		"ticketDescription": ticketDescription,
	}

	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		fmt.Println("Could not create JSON:", err)
		return
	}

	// dir, _ := os.Getwd()
	// fmt.Println(dir)
	// err = os.WriteFile(dir+"/.platformState.json", data, 0644)
	err = os.WriteFile(".platformState.json", data, 0644)
	if err != nil {
		fmt.Println("Could not save state:", err)
		return
	}

	fmt.Println("State updated.")
}

func get(url string, token string) ([]byte, int, error) {
	// AI generated f if I know
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, 0, err
	}

	request.Header.Set("Authorization", "Bearer "+token)

	client := http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, 0, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, response.StatusCode, err
	}

	return body, response.StatusCode, nil
}
