package callback

import "sync"

//map存事件对应方法
var eventByName map[string][]func(interface{})

var eventMap sync.Mutex

func init() {
	eventByName = make(map[string][]func(interface{}))
}

func RegisterEvent(name string, callback func(interface{})) {
	defer eventMap.Unlock()
	eventMap.Lock()
	list := eventByName[name]

	list = append(list, callback)

	eventByName[name] = list

}

func CallEvent(name string, param interface{}) {
	list := eventByName[name]

	for _, callback := range list {
		callback(param)
	}
}
