package callback

import (
	"fmt"
	"testing"
)

type Actor struct {
}

func (a *Actor) OnEvent(param interface{}) {
	fmt.Println("actor event:", param)
}

func GlobalEvent(param interface{}) {
	fmt.Println("global event:", param)
}

func TestCallEvent(t *testing.T) {
	type args struct {
		name  string
		param interface{}
	}
	a := new(Actor)
	var tests = []struct {
		name string
		args args
	}{}
	RegisterEvent("OnSkill", a.OnEvent)
	RegisterEvent("OnSkill", GlobalEvent)

	CallEvent("OnSkill", 99)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}

func TestRegisterEvent(t *testing.T) {
	type args struct {
		name     string
		callback func(interface{})
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
