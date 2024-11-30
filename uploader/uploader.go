package uploader

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type UploadFile struct {
	Base64File string `json:"base64File"`
	BucketName string `json:"bucketName"`
	FileName   string `json:"fileName"`
	Metadata   struct {
		Description string `json:"description"`
	} `json:"metadata"`
}

type Config struct {
	BucketName  string
	BaseAddress string
}

type Uploader struct {
	opts   Config
	client *http.Client
}

func NewUploader(opts Config, client *http.Client) *Uploader {
	return &Uploader{
		opts:   opts,
		client: client,
	}
}

func (u *Uploader) UploadDocument(ctx context.Context, req UploadFile) (string, error) {
	createUrl := fmt.Sprintf("%s%s", u.opts.BaseAddress, "/file/create")

	req.BucketName = u.opts.BucketName
	bodyBytes, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	reqWithContext, err := http.NewRequestWithContext(ctx, "POST", createUrl, bytes.NewReader(bodyBytes))
	if err != nil {
		return "", err
	}

	reqWithContext.Header.Add("Content-Type", "application/json")
	resp, err := u.client.Do(reqWithContext)
	if err != nil {
		return "", err
	}
	temp := struct {
		Data struct {
			Link string `json:"link"`
		}
		Error *struct {
			UserMsg     string `json:"userMsg"`
			Descriptoin string `json:"description"`
		} `json:"error"`
	}{}

	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&temp); err != nil {
		return "", fmt.Errorf("uploader:  unmarshall to dataresult %w", err)
	}

	if temp.Error != nil {
		return "", fmt.Errorf("uploader: error description %w", errors.New(temp.Error.Descriptoin))
	}

	return temp.Data.Link, nil
}
