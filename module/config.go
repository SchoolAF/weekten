package module

import (
	"github.com/aiteung/atdb"
	"os"
)

var MongoString string = os.Getenv("MONGOSTRCONNECT")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "WSDB",
}

var MongoConn = atdb.MongoConnect(MongoInfo)
