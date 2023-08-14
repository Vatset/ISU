package entities

type GroupInfo struct {
	GroupNumber string `json:"number"`  // P321211
	Faculty     string `json:"faculty"` // Piikt
	Course      int    `json:"course"`  // 3
}

type Student struct {
	Name        string `json:"name"`
	Lastname    string `json:"lastname"`
	ISU         int    `json:"ISU"`
	GroupNumber string `json:"number"`
}

type Group struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	ISU      int    `json:"ISU"`
}
