package banggood

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/cubixle/go-banggood/endpoints"
)

var baseURL = "https://api.banggood.com"

func NewClient(APIToken string) *Client {
	return &Client{
		APIToken: APIToken,
		BaseURL:  baseURL,
	}
}

type Client struct {
	APIToken string
	BaseURL  string
}

func (c *Client) Execute(ep endpoints.Request) endpoints.Response {
	return c.execute(ep)
}

func (c *Client) execute(ep endpoints.Request) endpoints.Response {
	var EPResponse endpoints.ResponseHandler
	ep.SetAccessToken(c.APIToken)

	req, _ := http.NewRequest(ep.GetType(), c.buildFullURL(ep.GetURL()), ep.GetBody())
	req.Header.Add("User-Agent", "cubixle-go-banggood")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)

	if urlErr, ok := err.(*url.Error); ok {
		log.Printf("Error while trying to send req: %v", urlErr)
		EPResponse.SetError(urlErr)
	} else if resp.StatusCode != 200 {
		rspBody, _ := ioutil.ReadAll(resp.Body)
		err = fmt.Errorf("<%d> - %s", resp.StatusCode, rspBody)
		log.Println(err)
		EPResponse.SetError(err)
	}

	rspBody, _ := ioutil.ReadAll(resp.Body)
	EPResponse.SetBody(rspBody)

	return EPResponse
}

func (c *Client) buildFullURL(ep string) string {
	return fmt.Sprintf(
		"%s/%s",
		strings.TrimRight(c.BaseURL, "/"),
		strings.TrimLeft(ep, "/"),
	)
}
