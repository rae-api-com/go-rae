package rae

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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
		Request(withttp.BaseURL("https://rae-api.com/api"))
)

func GetWord(
	ctx context.Context,
	version, word string,
) (*WordEntryResponse, error) {
	call := withttp.NewCall[*WordEntryResponse](withttp.Fasthttp()).
		URI(fmt.Sprintf("/words/%s", word)).
		Method(http.MethodGet).
		Header("User-Agent", fmt.Sprintf("rae-api/%s See https://rae-api.com", version), false).
		ParseJSON().
		ExpectedStatusCodes(http.StatusOK)

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
	call := withttp.NewCall[*WordResponse](withttp.Fasthttp()).
		URI("/daily").
		Method(http.MethodGet).
		Header("User-Agent", fmt.Sprintf("rae-api/%s See https://rae-api.com", version), false).
		ParseJSON().
		ExpectedStatusCodes(http.StatusOK)

	err := call.CallEndpoint(ctx, raeApi)

	return call.BodyParsed, err
}

func GetRandom(
	ctx context.Context,
	version string,
) (*WordResponse, error) {
	call := withttp.NewCall[*WordResponse](withttp.Fasthttp()).
		URI("/random").
		Method(http.MethodGet).
		Header("User-Agent", fmt.Sprintf("rae-api/%s See https://rae-api.com", version), false).
		ParseJSON().
		ExpectedStatusCodes(http.StatusOK)

	err := call.CallEndpoint(ctx, raeApi)

	return call.BodyParsed, err
}

type searchResponseDoc struct {
	Word string `json:"id"`
	Raw  string `json:"raw"`
}

type searchResponse struct {
	Doc  searchResponseDoc `json:"doc"`
	Hits int               `json:"hits"`
}

func GetSearch(
	ctx context.Context,
	version string,
	engine string,
	terms string,
) ([]WordEntry, error) {
	terms = url.QueryEscape(terms)
	uri := fmt.Sprintf("/search?q=%s", terms)
	if engine != "" {
		uri += fmt.Sprintf("&eng=%s", engine)
	}

	call := withttp.NewCall[[]searchResponse](withttp.Fasthttp()).
		URI("/search").
		Query("q", terms)

	if engine != "" {
		call = call.Query("eng", engine)
	}

	call.Method(http.MethodGet).
		Header("User-Agent", fmt.Sprintf("rae-api/%s See https://rae-api.com", version), false).
		ParseJSON().
		ExpectedStatusCodes(http.StatusOK)

	err := call.CallEndpoint(ctx, raeApi)

	if err != nil {
		return nil, errors.Wrapf(err, "failed to search for terms %s", terms)
	}

	res := make([]WordEntry, 0, len(call.BodyParsed))

	for _, hit := range call.BodyParsed {
		var doc WordEntry
		if err := json.Unmarshal([]byte(hit.Doc.Raw), &doc); err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal doc %s", hit.Doc.Word)
		}
		doc.Word = hit.Doc.Word
		res = append(res, doc)
	}

	return res, nil
}
