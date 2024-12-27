package zookeeper

import (
	"fmt"

	"github.com/andreparelho/temperature-challenge/channel"
)

type NewZookeeperTwo struct {
	Numbers            []int
	TemperatureChannel channel.TemperatureChannel
}

func (zookeeper *NewZookeeperTwo) Consumer() {
	for numberOnChannel := range zookeeper.TemperatureChannel.Channel {
		if numberOnChannel <= 50 {
			zookeeper.Numbers = append(zookeeper.Numbers, numberOnChannel)
			if len(zookeeper.Numbers) == 10 {
				zookeeper.Print()
			}
		}
	}
}

func (zookeeper *NewZookeeperTwo) Print() {
	fmt.Println("<= 50 = ", zookeeper.Numbers)
	zookeeper.Numbers = nil
}
