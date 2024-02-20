package server

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

// Fires when interval.C channel has data
func RunMonitorTick(interval *time.Ticker, vtubers []Streamer) {

	for {
		select {
		case <-interval.C:
			fmt.Println("ticker test")

			for _, vtuber := range vtubers {
				scrape := exec.Command("py", "server/scripts/scrape_channel.py", stripChannelURL(vtuber.Channel), vtuber.ID)
				// fmt.Println( stripChannelURL(vtuber.Channel), vtuber.ID )
				_, err := scrape.Output()
				if err != nil {
					log.Fatal(err)
				}
			}

		}
	}

}

// !!WARNING!!! Hacky as fuck, "https://www.youtube.com/channel/" is exactly 32 chars long
func stripChannelURL(url string) string {
	return url[32:]
}
