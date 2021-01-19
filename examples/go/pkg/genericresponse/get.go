package genericresponse

import (
	"context"

	"github.com/digital-dream-labs/hugh/log"

	"github.com/digital-dream-labs/vector-go-sdk/pkg/vectorpb"
)

// Get responds with text and sets off fireworks
func (w *GenericResponse) Get(ctx context.Context, arg string, m map[string]string) (string, map[string]string, error) {

	start := make(chan bool)
	stop := make(chan bool)

	go func() {
		_ = w.bot.BehaviorControl(ctx, start, stop)
	}()

	for {
		select {
		case <-start:
			_, err := w.bot.Conn.SayText(
				ctx,
				&vectorpb.SayTextRequest{
					Text:           w.response,
					UseVectorVoice: true,
					DurationScalar: 1,
				},
			)
			if err != nil {
				log.Error(err)
			}

			stop <- true
			return "", nil, nil
		}
	}

}
