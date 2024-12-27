package zookeeper

import (
	"fmt"

	"github.com/andreparelho/temperature-challenge/channel"
)

type NewZookeeperOne struct {
	Numbers            []int
	TemperatureChannel channel.TemperatureChannel
}

func (zookeeper *NewZookeeperOne) Consumer() {
	for numberOnChannel := range zookeeper.TemperatureChannel.Channel {
		if numberOnChannel > 50 {
			zookeeper.Numbers = append(zookeeper.Numbers, numberOnChannel)
			if len(zookeeper.Numbers) == 10 {
				zookeeper.Print()
			}
		}
	}
}

func (zookeeper *NewZookeeperOne) Print() {
	fmt.Println(">  50 = ", zookeeper.Numbers)
	zookeeper.Numbers = nil
}
