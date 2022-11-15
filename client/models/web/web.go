package web

import "education/console_notes/client/dto"

type Web interface {
	SendNote(note dto.Note) (noteID int64, err error)
	GetNote(noteID int64) (note dto.Note, err error)
	DeleteNote(noteID int64) (err error)
	GetAllNotes() (allNotes []dto.Note, err error)
	GetAllNotesByUser(user dto.User) (allNotes []dto.Note, err error)
}
