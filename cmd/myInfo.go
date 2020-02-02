package cmd

const baseURL = "https://qiita.com/api/v2/users"

type MyInfo struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	ItemsCount     int    `json:"items_count"`
	Location       string `json:"location"`
	FollowersCount int    `json:"followers_count"`
	FolloweesCount int    `json:"followees_count"`
}

func FetchQiitaData(token string, user string) {

}
