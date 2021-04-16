package snowflake

import (
	"fmt"
	"testing"
)

func TestSnowFlakeByGo(t *testing.T) {

	err := InitSnowflake("host1")
	if err != nil {
		fmt.Print("init snowflakeId has error %v", err)
	}

	worker, err := NewWorker(1)
	if err != nil {
		t.Errorf("init snowflakeId has error %v", err)
	}
	//id := worker.GetId()
	fmt.Println("id1:", worker.GetId())
	fmt.Println("id2:", worker.GetId())
	fmt.Println("id3:", worker.GetId())

}
