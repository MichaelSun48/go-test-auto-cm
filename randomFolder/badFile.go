package main
import (
	"log"
	"time"
	"os"
 	"github.com/getsentry/sentry-go"
)

func main() { 
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://efd5cea9ec7c1f7b0e1bdfddaf88da92@sun-dev.ngrok.dev/12",
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
  	}
	// Flush buffered events before the program terminates.
  	defer sentry.Flush(2 * time.Second)

	_, e := os.Open("filename.ext")
	if e != nil {
		sentry.CaptureException(e)
	}


}
