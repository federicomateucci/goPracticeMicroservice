package Dtos

type CurrentFile struct {
	CurrentFile string `json:"current_time"`
}

func (cr *CurrentFile) SetCurrentFileString(time string) {
	cr.CurrentFile = time
}
