module tinykap

go 1.19

replace libkapxy => ../../libkapxy

require (
	github.com/acceldata-io/libkapxy v0.0.0-20221129204120-aa3dd352cbaa
	github.com/google/gopacket v1.1.19
)

require golang.org/x/sys v0.0.0-20190412213103-97732733099d // indirect
