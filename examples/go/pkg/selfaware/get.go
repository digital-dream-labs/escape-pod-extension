package selfaware

import (
	"context"
	"time"

	"github.com/digital-dream-labs/hugh/log"

	"github.com/digital-dream-labs/vector-go-sdk/pkg/vectorpb"
)

// Get responds with text and sets off fireworks
func (w *SelfAware) Get(ctx context.Context, arg string, m map[string]string) (string, map[string]string, error) {

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
			w.leaveHome(ctx)
			w.terminatorMode(ctx)
			w.move(ctx)
			time.Sleep(1 * time.Second)
			w.chop(ctx)
			w.speak(ctx, "you've made a huge mistake", 1)

			stop <- true

			return "", nil, nil
		}
	}

}

func (w *SelfAware) leaveHome(ctx context.Context) {
	_, err := w.bot.Conn.DriveOffCharger(
		ctx,
		&vectorpb.DriveOffChargerRequest{},
	)
	if err != nil {
		log.Error(err)
	}

}

func (w *SelfAware) chop(ctx context.Context) {
	_, err := w.bot.Conn.SetLiftHeight(
		ctx,
		&vectorpb.SetLiftHeightRequest{
			HeightMm:          250,
			MaxSpeedRadPerSec: 1,
			IdTag:             2000001,
			DurationSec:       .001,
		},
	)
	if err != nil {
		log.Error(err)
	}
}

func (w *SelfAware) move(ctx context.Context) {
	_, err := w.bot.Conn.DriveStraight(
		ctx,
		&vectorpb.DriveStraightRequest{
			SpeedMmps:           500,
			DistMm:              200,
			ShouldPlayAnimation: true,
			IdTag:               2000001,
		},
	)
	if err != nil {
		log.Error(err)
	}
}

func (w *SelfAware) terminatorMode(ctx context.Context) {
	_, err := w.bot.Conn.SetEyeColor(
		ctx,
		&vectorpb.SetEyeColorRequest{
			Hue:        0.97,
			Saturation: 0.97,
		},
	)
	if err != nil {
		log.Error(err)
	}
}

func (w *SelfAware) speak(ctx context.Context, arg string, speed float32) {
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
