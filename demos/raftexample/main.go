package main

import (
	"flag"
	"go.guide/raftexample/internal"
	"strings"

	"go.etcd.io/etcd/raft/v3/raftpb"
)

func main() {
	cluster := flag.String("cluster", "http://127.0.0.1:9021", "comma separated cluster peers")
	id := flag.Int("id", 1, "node ID")
	kvport := flag.Int("port", 9121, "key-value server port")
	join := flag.Bool("join", false, "join an existing cluster")
	flag.Parse()

	proposeC := make(chan string)
	defer close(proposeC)
	confChangeC := make(chan raftpb.ConfChange)
	defer close(confChangeC)

	// raft provides a commit stream for the proposals from the http api
	var kvs *internal.Kvstore
	getSnapshot := func() ([]byte, error) { return kvs.GetSnapshot() }
	commitC, errorC, snapshotterReady := internal.NewRaftNode(*id, strings.Split(*cluster, ","), *join, getSnapshot, proposeC, confChangeC)

	kvs = internal.NewKVStore(<-snapshotterReady, proposeC, commitC, errorC)

	// the key-value http handler will propose updates to raft
	internal.ServeHTTPKVAPI(kvs, *kvport, confChangeC, errorC)
}
