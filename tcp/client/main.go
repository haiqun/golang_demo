package main

import (
	"context"
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
