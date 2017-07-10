package nhl

import (
	"log"
	"testing"

	"io/ioutil"

	"strings"

	mysportsfeeds "github.com/delaneyj/mysportsfeeds-go"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNHL(t *testing.T) {
	Convey("When using web client", t, func() {
		bytes, err := ioutil.ReadFile("../auth.txt")
		So(err, ShouldBeNil)

		rows := strings.Split(string(bytes), "\r\n")
		So(len(rows), ShouldEqual, 2)
		username := rows[0]
		password := rows[1]
		wc := mysportsfeeds.NewWebClient(username, password)
		nhl := NewNHL(wc, 2017, true)

		Convey("Can get Cumulative Player Stats", func() {
			cps, err := nhl.CumulativePlayerStats()
			So(err, ShouldBeNil)
			log.Printf("%+v", cps)
		})

	})

}
