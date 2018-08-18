package lockdownd

// #cgo LDFLAGS: -limobiledevice -lplist
// #include <libimobiledevice/libimobiledevice.h>
// #include <libimobiledevice/lockdown.h>
// #include <plist/plist.h>
import "C"
import (
	"unsafe"

	"github.com/alyyousuf7/libimobiledevice-go/idevice"
)

// Client is equivalent to libimobiledevice's lockdownd_client_t
type Client struct {
	client C.lockdownd_client_t
}

// NewClient returns Lockdownd client for a particular device
func NewClient(device *idevice.IDevice, label string) (*Client, error) {
	var client C.lockdownd_client_t

	if ret := C.lockdownd_client_new(
		*(*C.idevice_t)(device.CPtr()),
		&client,
		C.CString(label),
	); ret != C.LOCKDOWN_E_SUCCESS {
		return nil, handleError(int(ret))
	}

	return &Client{
		client,
	}, nil
}

// NewClientWithHandshake returns Lockdownd client with handshake for a
// particular device
func NewClientWithHandshake(device *idevice.IDevice, label string) (*Client, error) {
	var client C.lockdownd_client_t

	if ret := C.lockdownd_client_new_with_handshake(
		*(*C.idevice_t)(device.CPtr()),
		&client,
		C.CString(label),
	); ret != C.LOCKDOWN_E_SUCCESS {
		return nil, handleError(int(ret))
	}

	return &Client{
		client,
	}, nil
}

// Close frees lockdownd_client_t
func (c Client) Close() {
	C.lockdownd_client_free(c.client)
}

// DeviceName returns the device's name
func (c Client) DeviceName() (string, error) {
	var deviceName *C.char

	if ret := C.lockdownd_get_device_name(c.client, &deviceName); ret != C.LOCKDOWN_E_SUCCESS {
		return "", handleError(int(ret))
	}

	return C.GoString(deviceName), nil
}

// Query makes an empty query, used to prevent connection from closing
func (c Client) Query() error {
	if ret := C.lockdownd_query_type(c.client, nil); ret != C.LOCKDOWN_E_SUCCESS {
		return handleError(int(ret))
	}

	return nil
}

// QueryType returns lockdownd query type
func (c Client) QueryType() (string, error) {
	var queryType *C.char

	if ret := C.lockdownd_query_type(c.client, &queryType); ret != C.LOCKDOWN_E_SUCCESS {
		return "", handleError(int(ret))
	}

	return C.GoString(queryType), nil
}

// Get returns data for a domain and a key in plist binary form
// Empty string sends NULL to lockdownd_get_value
func (c Client) Get(domain, key string) ([]byte, error) {
	var (
		cplist    C.plist_t
		domainPtr *C.char
		keyPtr    *C.char
	)

	if domain != "" {
		domainPtr = C.CString(domain)
	}

	if key != "" {
		keyPtr = C.CString(key)
	}

	if ret := C.lockdownd_get_value(c.client, domainPtr, keyPtr, &cplist); ret != C.LOCKDOWN_E_SUCCESS {
		return nil, handleError(int(ret))
	}

	var (
		bin    *C.char
		length C.uint
	)
	C.plist_to_bin(cplist, &bin, &length)
	defer C.plist_free(cplist)

	return C.GoBytes(unsafe.Pointer(bin), C.int(length)), nil
}
