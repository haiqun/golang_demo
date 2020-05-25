package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
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
				str := Krand(10,1000)
				_, err := io.WriteString(conn, "hello"+" "+string(str))
				if err != nil {
					log.Fatalln(err.Error())
				}
				len, err := conn.Read(buf)
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("Receive From Server:", string(buf[:len]))
				}
				// 执行完,换链接
				conn = m.getConn()
			}
		}
	}
}

type tlsPoolConnInfo struct {
	maxConn int
	conn map[int]*tls.Conn// 连接池设置
}

var m tlsPoolConnInfo

func init()  {
	m = tlsPoolConnInfo{
		maxConn:20,
		conn:make(map[int]* tls.Conn,20),
	}
	m.createConn()
}

func (m * tlsPoolConnInfo) createConn()  {
	fileCrt := "./../server/ca.crt";
	rootPEM := Read3(fileCrt)
	if len(rootPEM) == 0 {
		log.Fatalf("证书读取有误 %s",rootPEM)
	}
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(rootPEM))
	if !ok {
		panic("failed to parse root certificate")
	}
	for j:=0;j<=m.maxConn;j++ {
		conn, err := tls.Dial("tcp", "fhq.com:8999", &tls.Config{
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
	n := RandInt(1,m.maxConn)
	log.Println(n)
	if len(m.conn) == 0 {
		return nil
	}
	conn, ok := m.conn[n];
	if ok {
		return conn
	}else{
		return  nil
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

func Read3(fileName string)  (string){
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("read file fail", err)
		return ""
	}
	defer f.Close()
	fd, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("read to fd fail", err)
		return ""
	}
	return string(fd)
}
