package zookeeper_test

import (
	"testing"
	"time"

	"github.com/andreparelho/temperature-challenge/channel"
	"github.com/andreparelho/temperature-challenge/observer"
	"github.com/andreparelho/temperature-challenge/zookeeper"
	"github.com/stretchr/testify/assert"
)

func TestZookeeperOne(test *testing.T) {
	test.Run("Deve validar corretamente os numeros acima de 50 na struct do Zookeeper quando os numeros estiverem disponiveis no canal", func(test *testing.T) {
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

		var numbers = []int{1, 32, 76, 34, 90, 91, 54, 50, 2, 45, 98, 76, 82, 12, 55, 25, 77, 100}

		for i := 0; i < len(numbers); i++ {
			time.Sleep(time.Second)
			go observer.Observer(numbers[i])
			go zookeeper.IZookeeper.Consumer(&zookeeperOne)
		}

		for _, numberZookeeper := range zookeeperOne.Numbers {
			assert.NotEmpty(test, zookeeperOne.Numbers)
			assert.True(test, numberZookeeper > 50)
		}
	})
}
