package models

type PussMsgRequest struct {
	Dest     string `json:"dest"`
	DestType string `json:"dest_type"`
	Msg      string `json:"msg"`
}
