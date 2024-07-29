package events

import "fmt"

func CreateEventListener(callback func(event Event), release ...func()) EventListener {

	self := &_EventListener{
		id:       fmt.Sprintf("%p", &callback),
		callback: callback,
	}

	var rfn func()
	if len(release) == 0 {
		rfn = self.__release
	} else {
		rfn = release[0]
	}

	self.release = rfn

	return self
}
