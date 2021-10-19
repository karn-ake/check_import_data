package entity

type SecurityData struct {
	StkCode string
	StkId	string
	IsinCode string
}

type ForeignData struct {
	SecurityName string
	SecurityId int
	DateOfRec string
	QtyAvail int
}