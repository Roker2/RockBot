package ext

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

const ApiUrl = "https://api.telegram.org/bot"

var DefaultTgBotGetter = TgBotGetter{
	Client: &http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       time.Millisecond * 1500,
	},
	ApiUrl: ApiUrl,
}

type Response struct {
	Ok          bool
	Result      json.RawMessage
	ErrorCode   int `json:"error_code"`
	Description string
	Parameters  json.RawMessage
}

type TgBotGetter struct {
	Client *http.Client
	ApiUrl string
}

type TgBotGetterInterface interface {
	Get(bot Bot, method string, params url.Values) (*Response, error)
	Post(bot Bot, fileType string, method string, params url.Values, file io.Reader, filename string) (*Response, error)
}

func Get(bot Bot, method string, params url.Values) (*Response, error) {
	return DefaultTgBotGetter.Get(bot, method, params)
}

func Post(bot Bot, fileType string, method string, params url.Values, file io.Reader, filename string) (*Response, error) {
	return DefaultTgBotGetter.Post(bot, fileType, method, params, file, filename)
}

func (tbg *TgBotGetter) Get(bot Bot, method string, params url.Values) (*Response, error) {
	req, err := http.NewRequest("GET", tbg.ApiUrl+bot.Token+"/"+method, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to build GET request to %v", method)
	}
	req.URL.RawQuery = params.Encode()

	bot.Logger.Debug("executing GET: %+v", req)
	resp, err := tbg.Client.Do(req)
	if err != nil {
		bot.Logger.WithError(err).Debugf("failed to execute GET request to %v", method)
		return nil, errors.Wrapf(err, "unable to execute GET request to %v", method)
	}
	defer resp.Body.Close()
	bot.Logger.Debugf("successful GET request: %+v", resp)

	var r Response
	if err = json.NewDecoder(resp.Body).Decode(&r); err != nil {
		bot.Logger.WithError(err).Debugf("failed to deserialize GET response body for %s", method)
		return nil, errors.Wrapf(err, "could not decode in GET %v call", method)
	}
	bot.Logger.Debugf("received result: %+v", r)
	bot.Logger.Debugf("result response: %v", string(r.Result))
	return &r, nil
}

func (tbg *TgBotGetter) Post(bot Bot, fileType string, method string, params url.Values, file io.Reader, filename string) (*Response, error) {
	if filename == "" {
		filename = "unnamed_file"
	}
	b := bytes.Buffer{}
	w := multipart.NewWriter(&b)
	part, err := w.CreateFormFile(fileType, filename)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	err = w.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", tbg.ApiUrl+bot.Token+"/"+method, &b)
	if err != nil {
		bot.Logger.WithError(err).Debugf("failed to execute POST request to %v", method)
		return nil, errors.Wrapf(err, "unable to execute POST request to %v", method)
	}
	req.URL.RawQuery = params.Encode()
	req.Header.Set("Content-Type", w.FormDataContentType())

	bot.Logger.Debug("POST request with body: %+v", b)
	bot.Logger.Debug("executing POST: %+v", req)
	resp, err := tbg.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bot.Logger.Debugf("successful POST request: %+v", resp)

	var r Response
	if err = json.NewDecoder(resp.Body).Decode(&r); err != nil {
		bot.Logger.WithError(err).Debug("failed to deserialize POST response body for %s", method)
		return nil, errors.Wrapf(err, "could not decode in POST %v call", method)
	}
	bot.Logger.Debugf("received result: %+v", r)
	bot.Logger.Debugf("result response: %v", string(r.Result))
	return &r, nil
}
