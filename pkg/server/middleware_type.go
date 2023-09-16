package server

type EtcdConf struct {
	Prefix   string `json:"prefix"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type EtcdMsg struct {
	Url           string
	Value         string
	Etcdbasicauth string
	TaskID        string
	Uuid          string
	Ip            string
	Cmd           string
}
