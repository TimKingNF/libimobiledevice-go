package idevice

// #cgo LDFLAGS: -limobiledevice
// #include <libimobiledevice/libimobiledevice.h>
//
// void idevice_event_subscribe_proxy(idevice_event_t *event, void *user_data);
//
// static void idevice_event_subscribe_proxy_with_const(const idevice_event_t *event_const, void *user_data) {
// 	idevice_event_t event = *event_const;
// 	idevice_event_subscribe_proxy(&event, user_data);
// }
//
// static idevice_error_t _idevice_event_subscribe(void* ch) {
// 	return idevice_event_subscribe(idevice_event_subscribe_proxy_with_const, ch);
// }
import "C"
import (
	"unsafe"
)

var eventChs []chan Event

// EventType is an integer
type EventType int

// Event types
const (
	EventAdd EventType = iota + 1
	EventRemove
	EventPair
)

// Event describes device activity
type Event struct {
	Type     EventType
	UDID     string
	ConnType int
}

// EventSubscribe returns a channel for events
func EventSubscribe() (chan Event, error) {
	var ret C.idevice_error_t

	ret = C._idevice_event_subscribe(nil)
	if ret != C.IDEVICE_E_SUCCESS {
		return nil, handleError(int(ret))
	}

	// create a new channel to return
	ch := make(chan Event)
	eventChs = append(eventChs, ch)

	return ch, nil
}

// EventUnsubscribe closes all channels
func EventUnsubscribe() error {
	var ret C.idevice_error_t

	ret = C.idevice_event_unsubscribe()
	if ret != C.IDEVICE_E_SUCCESS {
		return handleError(int(ret))
	}

	// close all channels
	for _, ch := range eventChs {
		close(ch)
	}

	return nil
}

//export idevice_event_subscribe_proxy
func idevice_event_subscribe_proxy(e *C.idevice_event_t, u unsafe.Pointer) {
	event := Event{
		Type:     EventType(e.event),
		UDID:     C.GoString(e.udid),
		ConnType: int(e.conn_type),
	}

	// publish to all channels
	for _, ch := range eventChs {
		ch <- event
	}
}
