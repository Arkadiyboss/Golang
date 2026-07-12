package gigachat

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testServer/files"

	"github.com/google/uuid"
)

type TokenAccess struct {
	Token     string `json:"access_token"`
	TokenType string `json:"token_type"`
	ExpiresAt int64  `json:"expires_at"`
}

type GigaAcces struct {
	GigaAcces string `json:"gigaAccess"`
}

type Request struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string    `json:"role"`
	Content string `json:"content"`
}

type Respond struct {
    Choices []struct {  
        Message struct {
            Content string `json:"content"`
        } `json:"message"`
    } `json:"choices"`
}

func Access() (bool, error) {

	url := "https://ngw.devices.sberbank.ru:9443/api/v2/oauth"
	method := "POST"

	payload := strings.NewReader("scope=GIGACHAT_API_PERS")

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	req, err := http.NewRequest(method, url, payload)

	RqUID := uuid.Must(uuid.NewRandom()).String()

	if err != nil {
		fmt.Println(err)
		return false, err
	}

	authorization, err := files.Read("accessToken.json")

	if err != nil {
		fmt.Println(err)
		return false, err
	}

	var authorizationToken GigaAcces

	json.Unmarshal(authorization, &authorizationToken)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("RqUID", RqUID)
	req.Header.Add("Authorization", authorizationToken.GigaAcces)

	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return false, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		return false, err
	}

	files.Write("accessInfo.json", body)

	return true, nil

}

func Promt(city string) (string, error) {

	promt := "Расскажи интересный факт о городе (максимум до 200 символов) - " + city

	url := "https://gigachat.devices.sberbank.ru/api/v1/chat/completions"
	method := "POST"

	payloadStr := fmt.Sprintf(`{
  "model": "GigaChat-2-Max",
  "messages": [
    {
      "role": "system",
      "content": "%s"
    }
  ]
}`, promt)

	payload := strings.NewReader(payloadStr)


	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	var verifyToken TokenAccess

	token, err := files.Read("accessInfo.json")


	json.Unmarshal(token, &verifyToken)

	if err != nil {
		fmt.Println(err)
		return "", err
	}


	req, err := http.NewRequest(method, url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer " + verifyToken.Token)


	if err != nil {
		fmt.Println(err)
		return "", err
	}


	res, err := client.Do(req)


	if err != nil {
		fmt.Println(err)
		return "", err
	}


	if res.StatusCode == 401 {
		fmt.Println(err)
		Access()
		Promt(city)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var fact Respond

	json.Unmarshal(body, &fact)

	content := fact.Choices[0].Message.Content

	return content, nil
}
