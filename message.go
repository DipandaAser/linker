package linker

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type MessageInfo struct {
	GroupName      string `json:"group_name"`
	GroupID        string `json:"group_id"`
	SenderNickName string `json:"sender_nick_name"`
}

type TextMessage struct {
	Text string      `json:"text"`
	Info MessageInfo `json:"info"`
}

type AudioMessage struct {
	Content []byte      `json:"content"`
	Info    MessageInfo `json:"info"`
}

type ImageMessage struct {
	Content []byte      `json:"content"`
	Caption string      `json:"caption"`
	Info    MessageInfo `json:"info"`
}

type VideoMessage struct {
	Content []byte      `json:"content"`
	Type    string      `json:"type"`
	Caption string      `json:"caption"`
	Info    MessageInfo `json:"info"`
}

type DocumentMessage struct {
	Content []byte      `json:"content"`
	Type    string      `json:"type"`
	Info    MessageInfo `json:"info"`
}

func SendMessageToService(serviceName string, message interface{}) error {
	theService, err := GetService(serviceName)
	if err != nil {
		return err
	}

	if theService.Status == StatusOffline {
		return ErrServiceOffline
	}

	reqData, err := json.Marshal(message)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPost, theService.Url, bytes.NewBuffer(reqData))
	if err != nil {
		return err
	}
	request.Header.Set("contentType", "application/json")
	request.Header.Set(HeaderAuthKey, theService.AuthKey)
	rep, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	if rep.StatusCode != 200 {
		return errors.New(fmt.Sprintf("request fail with status code %d", rep.StatusCode))
	}

	return nil
}
