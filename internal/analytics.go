package internal

import (
	"log"
	"os"

	"github.com/getsentry/sentry-go"
	"github.com/rs/xid"
	segment "github.com/segmentio/analytics-go"
)

type Analytics struct {
	ID            string
	SegmentClient segment.Client
	SentryClient  sentry.Client
}

func (a *Analytics) Init() {
	a.ID = xid.New().String()

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              "https://adb527b644304373a8b045474a9f19dc@o1267000.ingest.sentry.io/4504684284805120",
		TracesSampleRate: 1.0,
		Debug:            false,
		Release:          "komiser@" + Version,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	if os.Getenv("SEGMENT_WRITE_KEY") != "" {
		a.SegmentClient = segment.New(os.Getenv("SEGMENT_WRITE_KEY"))

		a.TrackEvent("engine_launched", map[string]interface{}{
			"version": Version,
		})
	}
}

func (a *Analytics) TrackEvent(event string, properties map[string]interface{}) {
	if a.SegmentClient != nil {
		eventProperties := segment.NewProperties()

		for key, values := range properties {
			eventProperties.Set(key, values)
		}

		a.SegmentClient.Enqueue(segment.Track{
			UserId:     a.ID,
			Event:      event,
			Properties: eventProperties,
		})
	}
}
