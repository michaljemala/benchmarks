package requestid

import (
	"fmt"
	"math/rand"
	"os"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	"github.com/oklog/ulid"
	suuid "github.com/satori/go.uuid"
)

var (
	prefix    string
	requestID uint64
)

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func init() {
	hostname, err := os.Hostname()
	if hostname == "" || err != nil {
		hostname = "localhost"
	}

	charset := "abcdefghijklmnopqrstuvwxyz0123456789"
	pid := make([]byte, 10)
	for i := range pid {
		pid[i] = charset[rnd.Intn(len(charset))]
	}

	prefix = fmt.Sprintf("%s/%s-", hostname, string(pid))
}

func ulidUUID() string {
	return ulid.MustNew(ulid.Timestamp(time.Now()), rnd).String()
}

func googleUUID() string {
	return uuid.Must(uuid.NewUUID()).String()
}

func satoriUUID() string {
	return suuid.NewV1().String()
}

func atomicUUID() string {
	dst := make([]byte, len(prefix)+20)
	copy(dst, prefix)
	buf := dst[len(prefix):]
	n := atomic.AddUint64(&requestID, 1)
	for i := 19; n > 0; i-- {
		buf[i] = byte(n%10 + '0')
		n /= 10
	}
	return string(dst)
}

func timeCounter() int64 {
	return time.Now().UnixNano()
}

var c int64

func atomicCounter() int64 {
	return atomic.AddInt64(&c, 1)
}
