// @Title uuid.go
// @Description
// @Author Hunter 2024/9/3 18:10

package uuid

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"syscall"
	"time"
)

const (
	// Bits is the number of bits in a UUID
	Bits = 128

	// Size is the number of bytes in a UUID
	Size = Bits / 8

	format = "%08x%04x%04x%04x%012x"
)

var (
	// Loggerf can be used to override the default logging target.
	// Log messages in this library should be logged at warning level or higher.
	Loggerf = func(format string, args ...interface{}) {}
)

// UUID represents a UUID value. UUIDs can be compared and set to other values and accessed by byte.
type UUID [Size]byte

// GenerateUUID creates a new UUID
// @return u UUID
func GenerateUUID() (u UUID) {
	const (
		maxretries = 9
		backoff    = time.Millisecond * 10
	)

	var (
		totalBackoff time.Duration
		count        int
		retries      int
	)

	for {
		b := time.Duration(retries) * backoff
		time.Sleep(b)
		totalBackoff += b

		n, err := io.ReadFull(rand.Reader, u[count:])
		if err != nil {
			if retryOnError(err) && retries < maxretries {
				count += n
				retries++
				Loggerf("error generating version 4 uuid, retrying: %v", err)
				continue
			}

			panic(fmt.Errorf("error reading random number generator, retried for %v: %v", totalBackoff.String(), err))
		}

		break
	}

	// Set the version (4) and variant fields
	u[6] = (u[6] & 0x0f) | 0x40
	u[8] = (u[8] & 0x3f) | 0x80

	return u
}

// String formats the UUID as a string
// @receiver u UUID
// @return string formatted UUID string
func (u UUID) String() string {
	return fmt.Sprintf(format, u[:4], u[4:6], u[6:8], u[8:10], u[10:])
}

// retryOnError attempts to detect if a retry would be effective
// @param err error
// @return bool whether retry might be effective
func retryOnError(err error) bool {
	switch err := err.(type) {
	case *os.PathError:
		return retryOnError(err.Err) // unpack the target error
	case syscall.Errno:
		if err == syscall.EPERM {
			return true
		}
	}

	return false
}
