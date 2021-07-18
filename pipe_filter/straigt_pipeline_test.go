package pipe_filter

import "testing"

func TestStraightPipeline(t *testing.T) {
	spliter := NewSplitFilter(",")
	converter := NewToIntFilter()
	sum := NewSumFilter()
	sp := NewStraightPipeline("p1", spliter, converter, sum)
	ret, err := sp.Process("3,6,9,12,0")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}
