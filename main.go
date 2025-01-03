package main

import (
	"math/rand/v2"

	"github.com/andreparelho/temperature-challenge/channel"
	"github.com/andreparelho/temperature-challenge/observer"
	"github.com/andreparelho/temperature-challenge/zookeeper"
)

const MAX_ITERATOR int = 500

func main() {
	var newChannel = make(chan int)

	var channel = channel.TemperatureChannel{
		Channel: newChannel,
	}

	var observer = observer.NewObserver{
		TemperatureChannel: channel,
	}

	var zookeeperOne = zookeeper.NewZookeeperOne{
		Numbers:            []int{},
		TemperatureChannel: channel,
	}

	var zookeeperTwo = zookeeper.NewZookeeperTwo{
		Numbers:            []int{},
		TemperatureChannel: channel,
	}

	for range MAX_ITERATOR {
		var temperature = rand.IntN(100)

		go observer.Observer(temperature)
		go zookeeper.IZookeeper.Consumer(&zookeeperOne)
		go zookeeper.IZookeeper.Consumer(&zookeeperTwo)
	}
}
