package replication

import (
	"context"
	"fmt"
	"github.com/go-mysql-org/go-mysql/mysql"
	"testing"
)

func TestNotParseEvent(t *testing.T) {
	cfg := BinlogSyncerConfig{
		ServerID:     uint32(100),
		Flavor:       "mysql",
		Host:         "127.0.0.1",
		Port:         3306,
		User:         "root",
		Password:     "root",
		ParseEnabled: false,
	}

	syncer := NewBinlogSyncer(cfg)
	streamer, _ := syncer.StartSync(mysql.Position{Name: "binlog.000006", Pos: 4})
	byteSizes := 0
	for {
		ev, _ := streamer.GetEvent(context.Background())
		byteSizes += len(ev.RawData)
		fmt.Printf("total bytes:%v\n", byteSizes)
	}
}
