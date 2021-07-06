# Testing with real calls
We have a func, that _pings_ if server is UP on host:port  
This func must be tested to make sure it waits for server to actually be UP in timeout AND returns an error if server is DOWN during timeout

`fake_server` can be started with a delay.

This task involves usage of paraller run of tests and pre-setup server
Criteria of acceptance:
* Given a test case for `waitfor.It` func to return nil when server is UP 
* Given a test case for `waitfor.It` func to return an error when server is DOWN
