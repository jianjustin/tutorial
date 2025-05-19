package main

import (
	"bytes"
	"fmt"
	"go.etcd.io/etcd/server/v3/storage/wal"
	"go.etcd.io/etcd/server/v3/storage/wal/walpb"
	"log"
	"os"

	"go.etcd.io/raft/v3/raftpb"
	"go.uber.org/zap"
)

func main() {
	dir := "demo_etcd_wal"
	os.RemoveAll(dir)
	os.MkdirAll(dir, 0755)

	logger := zap.NewNop()

	// 创建 WAL 并写入
	w, err := wal.Create(logger, dir, nil)
	if err != nil {
		log.Fatal("创建 WAL 失败:", err)
	}
	defer w.Close()

	entry := raftpb.Entry{
		Term:  1,
		Index: 1,
		Type:  raftpb.EntryNormal,
		Data:  append([]byte("put:"), append([]byte("user:1"), []byte("Alice")...)...),
	}

	err = w.Save(raftpb.HardState{}, []raftpb.Entry{entry})
	if err != nil {
		log.Fatal("写入失败:", err)
	}

	err = w.Sync()
	if err != nil {
		log.Fatal("同步失败:", err)
	}
	fmt.Println("✅ 数据写入并同步完成")

	w.Close()

	// 再次打开读取
	wr, err := wal.Open(logger, dir, walpb.Snapshot{Index: 0})
	if err != nil {
		log.Fatal("打开 WAL 失败:", err)
	}
	defer wr.Close()

	_, _, entries, err := wr.ReadAll()
	if err != nil {
		log.Fatal("读取失败:", err)
	}

	fmt.Printf("📦 entries 数量: %d\n", len(entries))
	for _, e := range entries {
		if bytes.HasPrefix(e.Data, []byte("put:")) {
			data := e.Data[len("put:"):]
			k := data[:6]
			v := data[6:]
			fmt.Printf("🎯 Key: %s, Value: %s\n", k, v)
		}
	}
}
