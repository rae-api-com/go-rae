package rae

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
	cli := &Client{
		timeout: 5 * time.Second,
		version: "dev",
	}

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
	return res.Data, nil
}

func (c *Client) Random(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	res, err := GetRandom(ctx, c.version)

	if err != nil {
		return "", err
	}

	if !res.Ok {
		return "", errors.New("word not found")

	}
	return res.Data.Word, nil
}

func (c *Client) Daily(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	res, err := GetDaily(ctx, c.version)

	if err != nil {
		return "", err
	}

	if !res.Ok {
		return "", errors.New("word not found")

	}
	return res.Data.Word, nil
}

type ApiResponse[T any] struct {
	Ok   bool   `json:"ok"`
	Data T      `json:"data"`
	Err  string `json:"error"`
}

type WordEntryResponse = ApiResponse[WordEntry]

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

type WordSingle struct {
	Word string `json:"word"`
}

type WordResponse = ApiResponse[WordSingle]

func GetDaily(
	ctx context.Context,
	version string,
) (*WordResponse, error) {
	call := withttp.NewCall[*WordResponse](withttp.WithFasthttp()).
		WithURI("daily").
		WithMethod(http.MethodGet).
		WithHeader("User-Agent", fmt.Sprintf("rae-api/%s See https://rae-api.com", version), false).
		WithParseJSON().
		WithExpectedStatusCodes(http.StatusOK)

	err := call.CallEndpoint(ctx, raeApi)

	return call.BodyParsed, err
}

func GetRandom(
	ctx context.Context,
	version string,
) (*WordResponse, error) {
	call := withttp.NewCall[*WordResponse](withttp.WithFasthttp()).
		WithURI("/random").
		WithMethod(http.MethodGet).
		WithHeader("User-Agent", fmt.Sprintf("rae-api/%s See https://rae-api.com", version), false).
		WithParseJSON().
		WithExpectedStatusCodes(http.StatusOK)

	err := call.CallEndpoint(ctx, raeApi)

	return call.BodyParsed, err
}
