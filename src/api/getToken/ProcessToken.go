package processToken

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func ProcessToken (token string) string {
	//accessToken := model.Token{}

	apiUrl := "https://11.11.5.146:9443/oauth2/userinfo?schema=openid"

	bearer := "Bearer" + token

	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		fmt.Println("Error : ", err.Error())
		os.Exit(1)
	}

	req.Header.Add("Authorization", bearer)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport:tr}
	resp, err := client.Do(req)
	if err != nil{
		fmt.Println("Error : ", err.Error())
	}
	body , err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println("Error :", err.Error())
	}
	fmt. Println(string([]byte(body)))

	return token
}