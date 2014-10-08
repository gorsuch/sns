package sns

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/bmizerany/aws4"
)

type ErrorResponse struct {
	Error struct {
		Type    string
		Code    string
		Message string
	}
	RequestId string
}

type PublishResponse struct {
	PublishResult struct {
		MessageId string
	}
	ResponseMetadata struct {
		RequestId string
	}
}

type Client struct {
	endpoint string
}

func New(endpoint string) *Client {
	return &Client{
		endpoint: endpoint,
	}
}

func (c *Client) Publish(topic, subject, message string) error {
	v := url.Values{}
	v.Set("Action", "Publish")
	v.Set("TopicArn", topic)
	v.Set("Subject", subject)
	v.Set("Message", message)

	res, err := aws4.PostForm(c.endpoint, v)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		var errorResponse ErrorResponse
		err := xml.Unmarshal(body, &errorResponse)
		if err != nil {
			return err
		}
		return fmt.Errorf(errorResponse.Error.Message)
	}

	var publishResponse PublishResponse
	err = xml.Unmarshal(body, &publishResponse)
	if err != nil {
		return err
	}
	return nil
}
