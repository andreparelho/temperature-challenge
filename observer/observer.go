package observer

import (
	"github.com/andreparelho/temperature-challenge/channel"
)

type NewObserver struct {
	TemperatureChannel channel.TemperatureChannel
}

func (observer *NewObserver) Observer(temperature int) {
	observer.TemperatureChannel.Channel <- temperature
}
