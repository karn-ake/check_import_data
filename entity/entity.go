package entity

type SecurityData struct {
	StkCode  string
	StkId    string
	IsinCode string
}

type ForeignData struct {
	SecurityName string
	SecurityId   int
	DateOfRec    string
	QtyAvail     int
}

type DBConfig struct {
	Server string `json:"server"`
	Port   int    `json:"port"`
	User   string `json:"user"`
	Pass   string `json:"pass"`
	DB     string `json:"database"`
}
