package upload_images

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"topsis/client/model"
)

const (
	UrlHost = "https://freeimage.host/api/1/upload"
)

func UploadImage(params *model.ParamUploadImage) (*model.ImageInfo, error) {
	logrus.Info("Get link upload image")

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile := os.Open(params.Path)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)
	part, errFile := writer.CreateFormFile("source", filepath.Base(params.Path))
	_, errFile = io.Copy(part, file)
	if errFile != nil {
		return nil, errFile
	}
	_ = writer.WriteField("key", params.ApiKey)
	err := writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", UrlHost, payload)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		if res != nil {
			bodyBytes, err := ioutil.ReadAll(res.Body)
			if err != nil {
				logrus.Warnf("response %v", string(bodyBytes))
				return nil, err
			}
			return nil, errors.New("link upload image failure")
		}
	}

	imgInfo := &model.ImageInfo{}
	err = json.Unmarshal(body, imgInfo)
	if err != nil {
		return nil, err
	}
	logrus.Info("successfully link upload image...")

	return imgInfo, nil
}
