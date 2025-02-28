package control


import "C"
import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"time"
	"unsafe"
	"strings"

	"gerrit.o-ran-sc.org/r/ric-plt/xapp-frame/pkg/clientmodel"
	"gerrit.o-ran-sc.org/r/ric-plt/xapp-frame/pkg/xapp"

	//	"bytes"
	//	"encoding/binary"

	//"encoding/base64"
	//"strings"

	influxdb2 "github.com/influxdata/influxdb-client-go"
)



func NewControl() Control {
	xapp.Logger.Info("In new control")
	create_db()
	xapp.Logger.Info("returning control")
	return Control{
		make(chan *xapp.RMRParams),
		influxdb2.NewClient("http://ricplt-influxdb.ricplt:8086", "client"),
	}
}
func create_db() {
	//Create a database named kpimon in influxDB
	xapp.Logger.Info("In create_db")
	_, err := http.Post("http://ricplt-influxdb.ricplt:8086/query?q=create%20database%20kpimon", "", nil)
	if err != nil {
		xapp.Logger.Error("Create database failed!")
	}
	xapp.Logger.Info("exiting create_db")
}