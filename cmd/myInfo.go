package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const baseURL = "https://qiita.com/api/v2/users"

type MyInfo struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	ItemsCount     int    `json:"items_count"`
	Location       string `json:"location"`
	FollowersCount int    `json:"followers_count"`
	FolloweesCount int    `json:"followees_count"`
}

//accesTokenを使用してuser_info（名前など）を取得
func FetchQiitaData(token string, userId string) (MyInfo, error) {

	req, err := http.NewRequest(http.MethodGet, baseURL+"/"+userId, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("content-type", "application/json")

	if len(token) < 0 {
		fmt.Println("--------アクセストークンがない状態でAPIリクエストします--------")
	} else {
		fmt.Println("--------アクセストークンをつけた状態でAPIリクエストします--------")
		req.Header.Set("Authorization", "Bearer "+token)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var info MyInfo
	if err := json.Unmarshal(body, &info); err != nil {
		log.Fatal(err)
	}
	return info, err
}

//コンソール出力
func OutputQiitaInfo(info MyInfo) {
	fmt.Println("==============自分のQiita情報==============")
	fmt.Printf("%-10v%v%v", "ID", ": ", info.ID)
	fmt.Printf("%-10v%v%v", "Name", ": ", info.Name)
	fmt.Printf("%-10v%v%v", "ItemsCount", ": ", info.ItemsCount)
	fmt.Printf("%-10v%v%v", "Location", ": ", info.Location)
	fmt.Printf("%-10v%v%v", "FollowersCount", ": ", info.FollowersCount)
	fmt.Printf("%-10v%v%v", "FolloweesCount", ": ", info.FolloweesCount)
}
