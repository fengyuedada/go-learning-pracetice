package pipe_filter

import (
	"reflect"
	"testing"
)

func TestNewStraightPipeline(t *testing.T) {
	type args struct {
		name    string
		filters []Filter
	}
	tests := []struct {
		name string
		args args
		want *StraightPipeline
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStraightPipeline(tt.args.name, tt.args.filters...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStraightPipeline() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStraightPipeline_Process(t *testing.T) {
	type fields struct {
		Name    string
		Filters *[]Filter
	}
	type args struct {
		data Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &StraightPipeline{
				Name:    tt.fields.Name,
				Filters: tt.fields.Filters,
			}
			got, err := f.Process(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Process() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Process() got = %v, want %v", got, tt.want)
			}
		})
	}
}
