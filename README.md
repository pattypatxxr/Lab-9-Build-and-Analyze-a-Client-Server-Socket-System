How to Run
1. Run the TCP Server
    go run server.go

Expected output:

    Server running on port 8080

2. Run the Load Testing Client:
    
    go run client.go

The client will create 100 concurrent TCP connections to the server.



Expected Behavior Under Load

The server accepts multiple concurrent connections using goroutines.
Each client sends a short message and disconnects.
Under moderate load, the server responds normally.
Under high load:
Goroutine count increases rapidly
RTT increases
The server may slow down or crash if resources are exhausted
Timeouts prevent connections from hanging indefinitely.