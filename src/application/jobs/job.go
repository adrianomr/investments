package jobs

import "time"

type Job interface {
	Execute()
	ExecuteAfter() time.Duration
}
