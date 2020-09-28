package es

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"hidevops.io/hiboot/pkg/at"
	"hidevops.io/hiboot/pkg/log"
	"net/http"
	"time"
)

type Properties struct {
	at.ConfigurationProperties `value:"es"`
	Url                        string `json:"url"`
	Port                       int    `json:"port" default:"5672"`
	Host                       string `json:"host" default:"127.0.0.1"`
	Username                   string `json:"username" default:""`
	Password                   string `json:"password" default:""`
	Token                      string `json:"token" default:""`
}

type Client struct {
	*elastic.Client
}

func newClient() (client *Client) {
	return new(Client)
}

func (c *Client) Connect(p *Properties) (err error) {
	if p.Username != "" && p.Password != "" {
		return c.ConnectToBasicAuth(p)
	} else if p.Token != "" {
		return c.ConnectToToken(p)
	}
	return c.ConnectNoAuth(p)
}

func (c *Client) ConnectNoAuth(p *Properties) (err error) {
	esUrl := fmt.Sprintf("http://%s:%d", p.Host, p.Port)
	if p.Url != "" {
		esUrl = p.Url
	}
	client, err := elastic.NewSimpleClient(
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetURL(esUrl),
	)
	if err != nil {
		log.Errorf("elastic connection errors:%v", esUrl)
		return
	}
	c.Client = client
	log.Debugf("elastic connection success")
	return
}

func (c *Client) ConnectToBasicAuth(p *Properties) (err error) {
	esUrl := fmt.Sprintf("http://%s:%d", p.Host, p.Port)
	if p.Url != "" {
		esUrl = p.Url
	}
	client, err := elastic.NewSimpleClient(
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetURL(esUrl),
		elastic.SetBasicAuth(p.Username, p.Password),
	)
	if err != nil {
		log.Errorf("elastic connection errors:%v", esUrl)
		return
	}
	c.Client = client
	log.Debugf("elastic connection success")
	return
}

func (c *Client) ConnectToToken(p *Properties) (err error) {
	esUrl := fmt.Sprintf("http://%s:%d", p.Host, p.Port)
	if p.Url != "" {
		esUrl = p.Url
	}
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+p.Token)
	client, err := elastic.NewSimpleClient(
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetURL(esUrl),
		elastic.SetHeaders(headers),
	)
	if err != nil {
		log.Errorf("elastic connection errors:%v", esUrl)
		return
	}
	c.Client = client
	log.Debugf("elastic connection success")
	return
}
