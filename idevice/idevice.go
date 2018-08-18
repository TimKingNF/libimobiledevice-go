package idevice

// #cgo LDFLAGS: -limobiledevice
// #include <libimobiledevice/libimobiledevice.h>
import "C"
import (
	"unsafe"

	"github.com/alyyousuf7/libimobiledevice-go/utils"
)

// IDevice is equivalent to libimobiledevice's idevice_t
type IDevice struct {
	idevice C.idevice_t
}

// DebugMode sets the debug level
func DebugMode(on bool) {
	if on {
		C.idevice_set_debug_level(1)
	} else {
		C.idevice_set_debug_level(0)
	}
}

// UDIDList returns the UDID of the connected devices
func UDIDList() ([]string, error) {
	var (
		clist **C.char
		cn    C.int
	)

	if ret := C.idevice_get_device_list(&clist, &cn); ret < 0 {
		return nil, handleError(int(ret))
	}
	defer C.idevice_device_list_free(clist)

	return utils.GoStrings(int(cn), unsafe.Pointer(clist)), nil
}

// New returns IDevice instance
func New() (*IDevice, error) {
	var idevice C.idevice_t

	if ret := C.idevice_new(&idevice, nil); ret != C.IDEVICE_E_SUCCESS {
		return nil, handleError(int(ret))
	}

	return &IDevice{
		idevice,
	}, nil
}

// NewWithUDID returns IDevice instance for a particular device
func NewWithUDID(uuid string) (*IDevice, error) {
	var idevice C.idevice_t

	if ret := C.idevice_new(&idevice, C.CString(uuid)); ret != C.IDEVICE_E_SUCCESS {
		return nil, handleError(int(ret))
	}

	return &IDevice{
		idevice,
	}, nil
}

// CPtr returns unsafe pointer for idevice_t, to be used in other packages
func (d IDevice) CPtr() unsafe.Pointer {
	return unsafe.Pointer(&d.idevice)
}

// Close frees idevice_t
func (d IDevice) Close() {
	C.idevice_free(d.idevice)
}

// UDID returns UDID of the connected device
func (d IDevice) UDID() (string, error) {
	var udid *C.char

	if ret := C.idevice_get_udid(d.idevice, &udid); ret != C.IDEVICE_E_SUCCESS {
		return "", handleError(int(ret))
	}

	return C.GoString(udid), nil
}
