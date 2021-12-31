package snowflake

import (
	"fmt"
	"testing"
)

func TestSnowFlakeByGo(t *testing.T) {

	err := InitSnowflake("1")
	if err != nil {
		fmt.Printf("init snowflakeId has error %v", err)
	}

	worker, err := NewWorker(1)
	if err != nil {
		t.Errorf("init snowflakeId has error %v", err)
	}
	//id := worker.GetId()
	for i:=0;i<5000;i++ {
		fmt.Println("id:",i, worker.GetId())
	}

}
