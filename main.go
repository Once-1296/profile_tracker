package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	FirstName  string
	LastName   string
	Rank       string
	MaxRank    string
	Rating     int
	MaxRating  int
	Handle     string
	TitlePhoto string
}

type CFData interface {
	User
}

var BASEAPI string

func GenericQuery[T CFData](methodName string, params map[string]string, data *T) (*T, error) {
	fullAPI := BASEAPI + "/" + methodName + "?"
	for k, v := range params {
		if fullAPI[len(fullAPI)-1] == '?' {
			fullAPI += k + "=" + v
		} else {
			fullAPI += "&" + k + "=" + v
		}
	}
	var bytes []byte
	var err error
	var res *http.Response

	res, err = http.Get(fullAPI)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil, err
	}
	defer res.Body.Close()
	bytes, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil, err
	}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil, err
	}
	fmt.Printf("Fetched Successfully...\n")
	return data, err
}

func GetUserProfile(users []string) {
	user_str := ""
	for i, user := range users {
		user_str += user
		if i+1 < len(users) {
			user_str += ";"
		}
	}
	var params map[string]string
	var data *User
	params = make(map[string]string)
	params["handles"] = user_str
	data, err := GenericQuery("user.info", params, data)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Success in fetching profiles\n")
	}
}

func main() {
	BASEAPI = "https://codeforces.com/api"
	users := []string{"Awwab_coder123", "bingo", "amalverma", "hewhocodes"}
	GetUserProfile(users)
}
