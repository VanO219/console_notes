package client_http

import (
	"education/console_notes/client/dto"
	"net/http"
)

type WebClient struct {
	client *http.Client
}

func NewWebClient() *WebClient {
	return &WebClient{}
}

func (w WebClient) SendNote(note dto.Note) (noteID int64, err error) {
	//TODO implement me
	panic("implement me")
}

func (w WebClient) GetNote(noteID int64) (note dto.Note, err error) {
	//TODO implement me
	panic("implement me")
}

func (w WebClient) DeleteNote(noteID int64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (w WebClient) GetAllNotes() (allNotes []dto.Note, err error) {
	//TODO implement me
	panic("implement me")
}

func (w WebClient) GetAllNotesByUser(user dto.User) (allNotes []dto.Note, err error) {
	//TODO implement me
	panic("implement me")
}
