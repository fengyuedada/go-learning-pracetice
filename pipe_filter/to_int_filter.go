package pipe_filter

import (
	"errors"
	"strconv"
)

var ToIntFilterWrongFormatError = errors.New("input data should be []string")


type ToIntFilter struct {
}

func NewToIntFilter() *ToIntFilter {
	return &ToIntFilter{}
}

func (sf *ToIntFilter) Process(data Request)(Response,error)  {
	parts,ok := data.([]string) //检查数据格式类型,或者强转
	if !ok {
		return nil,ToIntFilterWrongFormatError
	}
	ret := []int{}
	for _, part := range parts {
		s, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		ret = append(ret,s)
	}
	return ret,nil
}


