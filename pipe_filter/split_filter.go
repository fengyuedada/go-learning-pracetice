package pipe_filter

import (
	"errors"
	"strings"
)

var SplitFilterWrongFormatError=errors.New("input data should be string")

type SplitFilter struct {
	delimiter string
}

func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{delimiter}
}

func (sf *SplitFilter) Process(data Request)(Response,error)  {
	str,ok := data.(string) //检查数据格式类型,或者强转
	if !ok {
		return nil,SplitFilterWrongFormatError
	}
	parts := strings.Split(str, sf.delimiter)
	return parts,nil
}