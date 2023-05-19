package snowflake

import (
	"go-gin/config"
	"go-gin/core/log"
	"strconv"
	"sync"
	"time"
)

type Snowflake struct {}

var (
	mutex       sync.Mutex
	// Record the timestamp of the last id
	// 记录最后一个 id 的时间戳
	timestamp   int64
	// Machine id
	// 机器 id
	machineId   int64
	// A maximum of 4096 ids can be generated in the current millisecond (starting from 0)
	// 当前毫秒内最多可以生成4096个id(从0开始)
	number      int64
	// Changing this value while the program is running may result in the same id being generated
	// 在程序运行时更改此值可能会导致生成相同的id
	startTime   int64
	typeCastErr error

	machineBit   uint8 = 10
	numberBit    uint8 = 12
	machineMax   int64 = -1 ^ (-1 << machineBit)
	numberMax    int64 = -1 ^ (-1 << numberBit)
	timeShift    uint8 = machineBit + numberBit
	machineShift uint8 = numberBit
)

// Generate an id for external invocation
// 生成 id 供外部调用
func GetId () int64 {
	
	mutex.Lock()
	defer mutex.Unlock()

	machineId, _ = strconv.ParseInt(config.App["machineId"], 10, 64)
	startTime, _ = strconv.ParseInt(config.App["startTime"], 10, 64)

	if typeCastErr != nil {
		log.Error(typeCastErr)
		return 0
	}

	return createId()
}

// Create id
// 创建 id
func createId () int64 {

	milliseconds := getMillisecond()
	if milliseconds < timestamp {
		return 0
	}

	if milliseconds == timestamp {
		number++
		if number > numberMax {
			for milliseconds <= timestamp {
				milliseconds = getMillisecond()
			}
		}
	} else {
		number = 0
	}

	timestamp = milliseconds

	return int64((milliseconds-startTime)<<timeShift | (machineId << int64(machineShift)) | (number))
}

// Get current millisecond
// 获取当前毫秒数
func getMillisecond () int64 {
	return time.Now().UnixNano() / 1e6
}
