package waitfor_test

import (
	"testing_tasks/task_3/waitfor/fake_server"
)

type testServerSettings struct {
	Host string
	Port int
	Delay int
	DelayOK int
	DelayFail int
}

var settings testServerSettings

func startTestServer() {
	srv := fake_server.New(settings.Host, settings.Port)
	srv.StartWithDelay(settings.Delay)
}



func init() {
	settings = testServerSettings{
		Host:      "localhost",
		Port:      12345,
		Delay:     3,
		DelayOK:   4,
		DelayFail: 2,
	}
}