package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/sonirico/withttp"
)

type Client struct {
	timeout time.Duration
	version string
}

func New(opts ...ClientOption) *Client {
	cli := new(Client)

	for _, opt := range opts {
		opt(cli)
	}

	return cli
}

func (c *Client) Word(ctx context.Context, word string) (WordEntry, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	res, err := GetWord(ctx, c.version, word)

	if err != nil {
		return WordEntry{}, err
	}

	if !res.Ok {
		return WordEntry{}, errors.New("word not found")

	}
	return res.Entry, nil
}

type WordEntryResponse struct {
	Ok    bool      `json:"ok"`
	Entry WordEntry `json:"data"`
	Err   string    `json:"error"`
}

var (
	raeApi = withttp.NewEndpoint("RaeAPI").
		Request(withttp.WithBaseURL("https://rae-api.com/api"))
)

func GetWord(
	ctx context.Context,
	version, word string,
) (*WordEntryResponse, error) {
	call := withttp.NewCall[*WordEntryResponse](withttp.WithFasthttp()).
		WithURI(fmt.Sprintf("/words/%s", word)).
		WithMethod(http.MethodGet).
		WithHeader("User-Agent", fmt.Sprintf("rae-api/%s See https://rae-api.com", version), false).
		WithParseJSON().
		WithExpectedStatusCodes(http.StatusOK)

	err := call.CallEndpoint(ctx, raeApi)

	return call.BodyParsed, err
}
