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
	"os/exec"
	"strings"
)

var version string = "0.1.0"

func main() {
	isWrongCommand := true
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
		if len(arguments) == 2 {
			if arguments[1] == "configure" {
				configure()
				return
			}
			if os.Args[1] == "version" {
				fmt.Println(version)
				return
			}
		}
	} else if len(arguments) == 4 {
		if arguments[1] == "start" {
			if arguments[2] == "ticket" {
				ticket := os.Args[3]
				if isValidTicket(ticket) {
					fmt.Println(ticket)

					// get ticket data
					_, ticketSummary, ticketDescription := getTicket(URL, TOKEN, ticket)
					fmt.Println("#--------------------#")
					fmt.Println("ID:", ticket)
					fmt.Println("Title:", ticketSummary)
					fmt.Println("Description:", ticketDescription)

					// check if current directory is a git repo
					if !isGitRepo() {
						fmt.Println("This command must be run inside a Git repository.")
					}

					_, err := getCurrentBranch()
					if err != nil {
						fmt.Println("Could not get current Git branch:", err)
					}

					branchName := createBranchName(ticket, ticketSummary)
					// fmt.Printf("Branch: %s for ticket %s", branchName, ticket)
					err = createBranch(branchName)
					if err != nil {
						fmt.Println("Could not create branch:", err)
					}
					fmt.Println("Created Branch:", branchName, "for ticket:", ticket)
					return
				}
			}
		}
	}
	// fmt.Println(isWrongCommand)
	if isWrongCommand {
		docString()
		return
	}
}

func docString() {
	fmt.Println("Hello, Platform user!")
	fmt.Println("If you face any difficulty feel free to reach out to the platform team")
	fmt.Println("Usage: platform start ticket EPD-123")
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

func getTicket(baseURL string, token string, ticketID string) (err error, ticketSummary string, ticketDescription string) {
	fmt.Println("URL:", baseURL)
	fmt.Println("Ticket:", ticketID)

	url := baseURL + "/api/issues/" + ticketID + "?fields=idReadable,summary,description"

	body, status, err := get(url, token)
	if err != nil {
		fmt.Println("Error:", err)
		return err, "", ""
	}

	fmt.Println("Status:", status)
	// fmt.Println(reflect.TypeOf(body))
	// fmt.Println(string(body))

	// unpack json
	var data map[string]interface{}

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Could not read ticket:", err)
		return err, "", ""
	}

	ticketSummary = data["summary"].(string)
	ticketDescription = data["description"].(string)

	// fmt.Println("ID:", ticketID)
	// fmt.Println("Title:", ticketSummary)
	// fmt.Println("Description:", ticketDescription)

	// fmt.Println(reflect.TypeOf(ticketSummary))
	updateState(ticketID, ticketSummary, ticketDescription)

	return nil, ticketSummary, ticketDescription
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

func isGitRepo() bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")

	err := cmd.Run()
	if err != nil {
		return false
	}

	return true
}

func getCurrentBranch() (string, error) {
	cmd := exec.Command("git", "branch", "--show-current")

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}

func createBranchName(ticketID string, ticketSummary string) string {
	// creates branch name with 4 words after ticket id
	words := strings.Fields(ticketSummary)

	if len(words) > 4 {
		words = words[:4]
	}

	summary := strings.ToLower(strings.Join(words, "-"))

	return "feature/" + ticketID + "-" + summary
}

func createBranch(branchName string) error {
	cmd := exec.Command("git", "checkout", "-b", branchName)

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
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
