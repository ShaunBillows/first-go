package models

type Question struct {
    ID     int    `json:"id"`
    Topic  string `json:"topic"`
    Ques   string `json:"ques"`
    Ans    string `json:"ans"`
}
