package repository

type Repository struct {
	Name      string
	Challenge string
	Steps     []Step
	Version   int
}

type Step struct {
	Type    string `json:"step"`
	Log     string `json:"log"`
	Status  int    `json:"status"`
	Success bool   `json:"success"`
}
