package linker

import (
	"encoding/json"
	"errors"
)

var textMessageHandler func(message *TextMessage)
var audioMessageHandler func(message *AudioMessage)
var imageMessageHandler func(message *ImageMessage)
var videoMessageHandler func(message *VideoMessage)
var documentMessageHandler func(message *DocumentMessage)

var (
	ErrUnknownMessageType       = errors.New("unknown message type")
	ErrTextHandlerUndefined     = errors.New("textMessageHandler undefined")
	ErrAudioHandlerUndefined    = errors.New("audioMessageHandler undefined")
	ErrImageHandlerUndefined    = errors.New("imageMessageHandler undefined")
	ErrVideoHandlerUndefined    = errors.New("videoMessageHandler undefined")
	ErrDocumentHandlerUndefined = errors.New("documentMessageHandler undefined")
)

// ProcessMessage take a slice of byte who contain the json message and parse it into the appropriate type of message and then call their handler
func ProcessMessage(b []byte) error {

	textMsg := &TextMessage{}
	err := json.Unmarshal(b, textMsg)
	if err == nil {
		if textMessageHandler != nil {
			textMessageHandler(textMsg)
			return nil
		} else {
			return ErrTextHandlerUndefined
		}
	}

	audioMsg := &AudioMessage{}
	err = json.Unmarshal(b, audioMsg)
	if err == nil {
		if audioMessageHandler != nil {
			audioMessageHandler(audioMsg)
			return nil
		} else {
			return ErrAudioHandlerUndefined
		}
	}

	imageMsg := &ImageMessage{}
	err = json.Unmarshal(b, imageMsg)
	if err == nil {
		if imageMessageHandler != nil {
			imageMessageHandler(imageMsg)
			return nil
		} else {
			return ErrImageHandlerUndefined
		}
	}

	videoMsg := &VideoMessage{}
	err = json.Unmarshal(b, videoMsg)
	if err == nil {
		if videoMessageHandler != nil {
			videoMessageHandler(videoMsg)
			return nil
		} else {
			return ErrVideoHandlerUndefined
		}
	}

	docMsg := &DocumentMessage{}
	err = json.Unmarshal(b, docMsg)
	if err == nil {
		if documentMessageHandler != nil {
			documentMessageHandler(docMsg)
			return nil
		} else {
			return ErrDocumentHandlerUndefined
		}
	}

	return ErrUnknownMessageType
}

// SetTextMessageHandler set the handler for TextMessage
func SetTextMessageHandler(handler func(message *TextMessage)) {
	textMessageHandler = handler
}

// SetAudioMessageHandler set the handler for AudioMessage
func SetAudioMessageHandler(handler func(message *AudioMessage)) {
	audioMessageHandler = handler
}

// SetImageMessageHandler set the handler for ImageMessage
func SetImageMessageHandler(handler func(message *ImageMessage)) {
	imageMessageHandler = handler
}

// SetVideoMessageHandler set the handler for VideoMessage
func SetVideoMessageHandler(handler func(message *VideoMessage)) {
	videoMessageHandler = handler
}

//SetDocumentMessageHandler set the handler for DocumentMessage
func SetDocumentMessageHandler(handler func(message *DocumentMessage)) {
	documentMessageHandler = handler
}
