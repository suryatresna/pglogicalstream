package pglogicalstream

import "github.com/suryatresna/pglogicalstream/pkg/replication"

type OnMessage = func(message replication.Wal2JsonChanges)
