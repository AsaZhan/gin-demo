package login

type Response struct {
	Message string
	Status  string
	Secret  interface{}
}

type ProtoBufExample struct {
	Label string
	Resp  []int64
}
