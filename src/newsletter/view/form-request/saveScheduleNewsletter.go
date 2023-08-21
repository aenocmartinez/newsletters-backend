package formrequest

type SaveScheduleNewsletter struct {
	Id      int64  `json:"id" binding:"required"`
	Subject string `json:"subject" binding:"required"`
	Hour    string `json:"hour" binding:"required"`
	Date    string `json:"date" binding:"required"`
	From    string `json:"from"`
}

// type ScheduleNewsletter struct {
// 	Subject    string `json:"subject" binding:"required"`
// 	Date       string `json:"date" binding:"required,len=10"`
// 	Hour       string `json:"hour" binding:"required"`
// 	Message    string `json:"message" binding:"required"`
// 	From       string `json:"from" binding:"required"`
// 	Recipients string `json:"recipients" binding:"required"`
// 	Sent_by    string `json:"sent_by" binding:"required"`
// }
