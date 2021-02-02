package core


type Filter  struct {
    Status	bool	`json:"status"`
    IP		string	`json:"ip"`
    PORT	string	`json:"port"`
    USERNAME	string	`json:"username"`
    PASSWORD	string	`json:"password"`
    Desp	string	`json:"desp"`
    Message	string	`json:"message"`
}
