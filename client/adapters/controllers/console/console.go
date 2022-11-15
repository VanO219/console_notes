package console

import (
	"education/console_notes/client/models/web"
)

type Viewer struct {
	usecase web.Web
}

func (cv *Viewer) View() {
	//act := ""
	//for {
	//	fmt.Println("\nВыберете действие:")
	//	fmt.Println(`n - ввести новую заметку`)
	//	fmt.Println(`e - завершить программу`)
	//	fmt.Println(`s - показать все заметки`)
	//	fmt.Scanln(&act)
	//
	//	switch act {
	//	case "n":
	//		note := dto.Note{
	//			ID:   0,
	//			User: dto.User{
	//				Name:     "sadasd",
	//				LastName: "asdasdasd",
	//			},
	//			Text: "awdasdqawd",
	//		}
	//		id, err := cv.usecase.SendNote(note)
	//		if err != nil {
	//
	//		}
	//		fmt.Println("Заметка сохранилась с идентификатором", id)
	//
	//	case "e":
	//		fmt.Println(`Всего хорошего!`)
	//		return
	//	case "s":
	//		printAll()
	//	default:
	//		continue
	//	}
	//}
}
