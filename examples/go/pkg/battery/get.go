package battery

import (
	"context"
	"fmt"

	"github.com/digital-dream-labs/hugh/log"

	"github.com/digital-dream-labs/vector-go-sdk/pkg/vectorpb"
)

// Get responds with text and sets off fireworks
func (w *Battery) Get(ctx context.Context, arg string, m map[string]string) (string, map[string]string, error) {

	// You'll have to overwrite the incoming context, or it'll time out before this completes.
	ctx = context.Background()

	start := make(chan bool)
	stop := make(chan bool)

	go func() {
		_ = w.bot.BehaviorControl(ctx, start, stop)
	}()

	for {
		select {
		case <-start:
			w.speak(
				ctx,
				w.getBatteryLevel(ctx),
				1,
			)
			stop <- true

			return "", nil, nil
		}
	}

}

func (w *Battery) getBatteryLevel(ctx context.Context) string {
	bs, err := w.bot.Conn.BatteryState(
		context.Background(),
		&vectorpb.BatteryStateRequest{},
	)
	if err != nil {
		return "something went wrong"
	}

	var response string

	if bs.IsOnChargerPlatform {
		response += "I'm on the charger."
	}

	switch bs.BatteryLevel {
	case vectorpb.BatteryLevel_BATTERY_LEVEL_FULL:
		response += fmt.Sprintf("My battery is full, with an output of around %.2f volts", bs.BatteryVolts)
	case vectorpb.BatteryLevel_BATTERY_LEVEL_NOMINAL:
		response += fmt.Sprintf("My battery is pretty okay, with an output of around %.2f volts", bs.BatteryVolts)
	case vectorpb.BatteryLevel_BATTERY_LEVEL_LOW:
		response += fmt.Sprintf("My battery is getting low, with an output of around %.2f volts", bs.BatteryVolts)
	case vectorpb.BatteryLevel_BATTERY_LEVEL_UNKNOWN:
		response += "I can't see my battery.  I'm scared.  Hold me."
	}

	return response

}

func (w *Battery) speak(ctx context.Context, arg string, speed float32) {
	_, err := w.bot.Conn.SayText(
		ctx,
		&vectorpb.SayTextRequest{
			Text:           arg,
			UseVectorVoice: true,
			DurationScalar: speed,
		},
	)
	if err != nil {
		log.Error(err)
	}
}
