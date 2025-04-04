package cassandra

import (
    "log"
    "github.com/gocql/gocql"
)

var Session *gocql.Session

func Init() {
    cluster := gocql.NewCluster("127.0.0.1")
    cluster.Keyspace = "library"
    cluster.Consistency = gocql.Quorum

    var err error
    Session, err = cluster.CreateSession()
    if err != nil {
        log.Fatalf("Failed to connect to Cassandra: %v", err)
    }
    log.Println("Connected to Cassandra")
}