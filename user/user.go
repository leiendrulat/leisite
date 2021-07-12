package user

type User struct {
	ID       string `param:"id" query:"id" form:"id" json:"id" xml:"id"`
	Fname    string `param:"fname" query:"fname" form:"fname" json:"fname" xml:"fname"`
	Lname    string `param:"lname" query:"lname" form:"lname" json:"lname" xml:"lname"`
	Email    string `param:"email" query:"email" form:"email" json:"email" xml:"email"`
	Password string `param:"password" query:"password" form:"password" json:"password" xml:"password"`
}
