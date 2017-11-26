package gigasecond

import (
	"os"
	"time"
)

const testVersion = 4
const gigasec = 1000000000

func AddGigasecond(t time.Time) time.Time {
	os.Stderr.WriteString(string(t.Add(gigasec*time.Second).String()) + "\r\n")
	return t.Add(gigasec * time.Second)

}
