package keeper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	Login string `json:"login"`
}

func GetFollowers() []User {
	followers := getLogin("https://api.github.com/users/sapvs/followers?per_page=100")
	for _, follower := range followers {
		fmt.Println(follower.Login)

	}

	return followers

}

func GetFollowing() []User {
	following := getLogin("https://api.github.com/users/sapvs/following?per_page=100")
	for _, follower := range following {
		fmt.Println(follower)

	}

	return following
}

func getLogin(url string) []User {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	logins := []User{}
	err = json.Unmarshal(body, &logins)
	if err != nil {
		panic(err)
	}

	return logins
}
