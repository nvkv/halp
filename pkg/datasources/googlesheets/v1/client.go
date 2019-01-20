package googlesheets

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"

	"github.com/nvkv/halp/pkg/types/data/v1"
)

func createClient(tokenfile string, config *oauth2.Config) (*http.Client, error) {
	token, err := tokenFromFile(tokenfile)
	if err != nil {
		token, err = getTokenFromWeb(tokenfile, config)
		if err != nil {
			return nil, err
		}
		saveToken(tokenfile, token)
	}
	return config.Client(context.Background(), token), nil
}

func getTokenFromWeb(tokenfile string, config *oauth2.Config) (*oauth2.Token, error) {
	authURL := config.AuthCodeURL(
		"state-token",
		oauth2.AccessTypeOffline,
	)

	fmt.Printf("Get token by opening this URL and paste it here: \n%v\n", authURL)
	var authCode string

	if _, err := fmt.Scan(&authCode); err != nil {
		return nil, fmt.Errorf("Unable to read authorization code: %v", err)
	}

	token, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve token from web: %v", err)
	}
	return token, nil
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	token := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(token)
	return token, err
}

func saveToken(path string, token *oauth2.Token) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(token)
}

func parseRow(row []interface{}) (data.Meal, error) {
	if len(row) < 4 {
		return data.Meal{}, fmt.Errorf("Cant parse malformed row: %v", row)
	}

	name := row[0].(string)
	if len(name) == 0 {
		return data.Meal{}, fmt.Errorf("Nameless meal: %v", row)
	}

	typeStr := row[1].(string)
	lentenStr := row[2].(string)
	lavishStr := row[3].(string)

	meal := data.Meal{}
	meal.Name = name

	switch typeStr {
	case "Breakfast":
		meal.Type = data.Breakfast
	case "Lunch":
		meal.Type = data.Lunch
	case "Dinner":
		meal.Type = data.Dinner
	case "Snack":
		meal.Type = data.Snack
	default:
		return data.Meal{}, fmt.Errorf("Unknown meal type: %v", typeStr)
	}

	var err error

	meal.IsLenten, err = strconv.ParseBool(lentenStr)
	if err != nil {
		return data.Meal{}, fmt.Errorf("Can't parse lent status: %v", row)
	}

	meal.IsLavish, err = strconv.ParseBool(lavishStr)
	if err != nil {
		return data.Meal{}, fmt.Errorf("Can't parse lavish status: %v", row)
	}

	return meal, nil
}

func fetchAll(credentials, tokenfile, sheetId string) ([]data.Meal, error) {
	credBuff, err := ioutil.ReadFile(credentials)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(
		credBuff,
		"https://www.googleapis.com/auth/spreadsheets.readonly",
	)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse client secret file to config: %v", err)
	}
	client, err := createClient(tokenfile, config)
	if err != nil {
		return nil, err
	}

	srv, err := sheets.New(client)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve Sheets client: %v", err)
	}

	readRange := "Halp!A2:D"
	resp, err := srv.Spreadsheets.Values.Get(sheetId, readRange).Do()
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve data from sheet: %v", err)
	}

	meals := []data.Meal{}

	for _, row := range resp.Values {
		if len(row) == 0 || len(row[0].(string)) == 0 {
			continue
		}
		meal, err := parseRow(row)
		if err != nil {
			return nil, err
		}
		meals = append(meals, meal)
	}

	return meals, nil
}

func Test(credentials, tokenfile, sheetId string) {
	meals, err := fetchAll(credentials, tokenfile, sheetId)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%#v\n", meals)
}
