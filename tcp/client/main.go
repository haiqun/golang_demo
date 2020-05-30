package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io"
	"math/rand"
	"sync"
	"time"
	"fmt"
	"log"
)

var ws sync.WaitGroup

func main() {
	log.Println(m.conn)
	for j:=0;j<2;j++ {
		ws.Add(1)
		go test()
	}
	ws.Wait()
}

func test()  {
	defer ws.Done()
	conn := m.getConn()
	log.Println("Client Connect To ", conn.RemoteAddr())
	status := conn.ConnectionState()
	log.Printf("%#v\n", status)
	buf := make([]byte, 1024)
	ticker := time.NewTicker(1 * time.Millisecond * 5000)
	for {
		select {
		case <-ticker.C:
			{
				m.send(conn)
				conn = m.getConn()
			}
		case <-m.receiveChan:

		}
	}
}


type tlsPoolConnInfo struct {
	maxConn int
	conn map[int]*tls.Conn// 连接池设置
	receiveChan chan string // 接收
	sendChan chan string // 发送
}

var m tlsPoolConnInfo

func init()  {
	m = tlsPoolConnInfo{
		maxConn:20,
		conn:make(map[int]* tls.Conn,20),
		receiveChan:make(chan string,100),
		sendChan:make(chan string,100),
	}
	m.createConn()
}

func (m * tlsPoolConnInfo) createConn()  {
	// Connecting with a custom root-certificate set.
	const rootPEM = `
-----BEGIN CERTIFICATE-----
MIIDdDCCAlwCCQCh7qp2iYvLlDANBgkqhkiG9w0BAQsFADB7MRAwDgYDVQQIDAdm
aHEuY29tMRAwDgYDVQQHDAdmaHEuY29tMRAwDgYDVQQKDAdmaHEuY29tMRAwDgYD
VQQLDAdmaHEuY29tMRAwDgYDVQQDDAdmaHEuY29tMR8wHQYJKoZIhvcNAQkBFhAy
ODExMjczNjFAcXEuY29tMCAXDTIwMDUyMzA1MDcwOVoYDzIxMjAwNDI5MDUwNzA5
WjB7MRAwDgYDVQQIDAdmaHEuY29tMRAwDgYDVQQHDAdmaHEuY29tMRAwDgYDVQQK
DAdmaHEuY29tMRAwDgYDVQQLDAdmaHEuY29tMRAwDgYDVQQDDAdmaHEuY29tMR8w
HQYJKoZIhvcNAQkBFhAyODExMjczNjFAcXEuY29tMIIBIjANBgkqhkiG9w0BAQEF
AAOCAQ8AMIIBCgKCAQEArTxy0whc51X6BW5BmadOaHaeWPFuG49aCs+IjZl/uUrz
MHXHoA5NN78lgqyo6xlr94jUFdnRTUPzPFA+S17ZJTQ/Pry7+YVgyhDb7oPA3cm3
y51Tb6mSNw+sju4NNBOiDFbF/NR9KJafeQ8zbzo4MJlYeDMnMyHmX1UgXgp7ddZR
I1McBHDis9/a2+GJ2Jr8GLEUAmwxygdAjARbkIrAocbshDpjfWfNydcmtntLjSou
UMADj3EEjrCPhWncmG6+Su/+ctCqI5oGjZhtGaf5O/0TuGb92jffTQrEdYNWGF9g
EPfUWdhmU78AW0TSakvvtgUQ+fUt3U8+bSy+En0IqQIDAQABMA0GCSqGSIb3DQEB
CwUAA4IBAQCX8X1rNblRcKV7fH1N243SIXcfz502589e5AarydSggGHCMU7OLwP3
Kc6Vl7PpM9wk9E3oUQlXSKCwKyEKy7u2yntZb1mAc4yACGuJYdlItvdN6aRopsTv
LHwV0Xv8ZKWvWEKf0nrJEZqiprKDY4ihH9rOJS7PcZq3XT/imvQKv0S3s5RCkHDP
euNEay/jBmbaRQ//uOOx+Lq8TSFSNxy1peE2A6GLhEedIYfURL8AdAGCZNt1VQQ/
Nj8+tl5tKnXIYbz2osoDvBLVkbsrdMClACWLV0sRczyx3zmTqsAOpx3pdK+RXRWw
F5Q9LDgiZTBgAHXQvuOg0Clt/jFFZY0e
-----END CERTIFICATE-----`
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(rootPEM))
	if !ok {
		panic("failed to parse root certificate")
	}
	for j:=0;j<=m.maxConn;j++ {
		conn, err := tls.Dial("tcp", "fhq.com:8888", &tls.Config{
			RootCAs: roots,
		})
		if err != nil {
			panic("failed to connect: " + err.Error())
		}
		m.conn[j] = conn
	}
	//defer conn.Close()
}

func (m * tlsPoolConnInfo) getConn() *tls.Conn  {
	ctx, cancel := context.WithCancel(context.Background())
	n := RandInt(1,m.maxConn)
	log.Println(n)
	if len(m.conn) == 0 {
		return nil
	}
	conn, ok := m.conn[n];
	if ok {
		return conn
	}else{
		// 开启接收通道
		go m.receive(conn,ctx)
		return  nil
	}
}

func (m * tlsPoolConnInfo) closeConn () {

}

func (t *tlsPoolConnInfo)send(c *tls.Conn )  {
	select {
	case mgs := <-t.sendChan:
		_, err := io.WriteString(c, "hello"+" "+mgs)
		if err != nil {
			log.Fatalln(err.Error())
		}
	}
}

func (t *tlsPoolConnInfo)receive(c *tls.Conn,ctx context.Context)  {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			buf := make([]byte, 1024)
			len, err := c.Read(buf)
			msg := ""
			if err != nil {
				panic(err.Error())
			} else {
				msg = string(buf[:len])
			}
			t.receiveChan<- msg
		}
	}
}




// 随机字符串
func Krand(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i :=0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base+rand.Intn(scope))
	}
	return result
}

// 随机
func RandInt(min, max int) int {
	if min >= max {
		return max
	}
	return rand.Intn(max-min) + min
}