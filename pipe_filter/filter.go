package pipe_filter

//spliFilter——tointFilter——sumFilter

type Request interface {
}

type Response interface {

}

type Filter interface {
	Process(data Request) (Response, error)
}
