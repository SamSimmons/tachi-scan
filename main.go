package main

import (
  "fmt"
  "os"
  "net"
  "strconv"
)

type Result struct {
	Port    int
	Success bool
	Err     error
}

func main() {
  host := ""
  minPort := 1
  maxPort := 1024

  if len(os.Args) > 1 {
    host = os.Args[1]
    results := []Result{}

    for i := minPort; i <= maxPort; i++ {
      r := ScanHost(host, i)
      fmt.Printf("%+v\n", r)
      results = append(results, r)
    }
  } else {
    fmt.Println("Please provide something to scan")
  }
}

func ScanHost(host string, port int) Result{
  fmt.Printf("Scanning: %v:%v\n", host, strconv.Itoa(port))
  conn, err := net.Dial("tcp", host + ":" + strconv.Itoa(port))

  result := Result{
    Port: port,
    Success: err == nil,
    Err: err,
  }

  if conn != nil {
    conn.Close()
  }

  return result
}
