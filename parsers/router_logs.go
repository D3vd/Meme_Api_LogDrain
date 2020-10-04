package parsers

import (
	"bufio"
	"io"
	"log"

	"Meme_Api_LogDrain/types"

	"github.com/bmizerany/lpx"
	"github.com/getsentry/sentry-go"
	"github.com/kr/logfmt"
)

// GetRouterLogs - Takes logs data and returns []RouterLog
func GetRouterLogs(data io.Reader) (logs []types.RouterLog) {
	lp := lpx.NewReader(bufio.NewReader(data))

	for lp.Next() {
		if string(lp.Header().Procid) == "router" {
			rl := new(types.RouterLog)
			if err := logfmt.Unmarshal(lp.Bytes(), rl); err != nil {
				sentry.CaptureException(err)
				log.Printf("Error: %v", err)
				continue
			} else {
				logs = append(logs, *rl)
			}
		}
	}

	return
}
