package message

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DipandaAser/linker/service"
	"net/http"
)

type Info struct {
	GroupName      string      `json:"group_name"`
	GroupID        string      `json:"group_id"`
	SenderNickName string      `json:"sender_nick_name"`
	MessageData    interface{} `json:"message_data"`
}

type TextMessage struct {
	Text string `json:"text"`
}

type AudioMessage struct {
	Content []byte `json:"content"`
}

type ImageMessage struct {
	Content []byte `json:"content"`
	Caption string `json:"caption"`
}

type VideoMessage struct {
	Content []byte `json:"content"`
	Caption string `json:"caption"`
}

type DocumentMessage struct {
	Content []byte `json:"content"`
}

func SendMessageToService(serviceName string, messageInfo *Info) error {
	theService, err := service.GetService(serviceName)
	if err != nil {
		return err
	}

	if theService.Status == service.StatusOffline {
		return service.ErrServiceOffline
	}

	url := fmt.Sprintf("%s/linker/%s/message", theService.Url, theService.AuthKey)
	reqData, err := json.Marshal(messageInfo)
	if err != nil {
		return err
	}

	rep, err := http.Post(url, "application/json", bytes.NewBuffer(reqData))
	if err != nil {
		return err
	}

	if rep.StatusCode != 200 {
		return errors.New(fmt.Sprintf("request fail with staus code %d", rep.StatusCode))
	}

	return nil
}
