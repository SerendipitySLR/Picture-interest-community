package command

type CreatPost struct {
	PictureUrl string `json:"pictureurl"`
	Contents   string `json:"contents"`
	Date       string `json:"date"`
	UID        int    `json:"uid"`
	Location   string `json:"location"`
}

type CreatForward struct {
	PostID   int    `json:"postid"`
	OtherID  int    `json:"otherid"`
	Contents string `json:"contents"`
	Date     string `json:"date"`
	Location string `json:"location"`
	UID      int    `json:"uid"`
}
