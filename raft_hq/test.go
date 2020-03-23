package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*

	0 n个队列，随机时间，发起选举
	1 开始选举，谁快谁当选
	2 选举完之后 ，投票通知 【判断是不是比自己发起的早】
	3 票数超过一半，成为领导
	4 领导发起心跳监控
	5

*/
const raftCount1  = 10

// 选举候选人接锅
var flag2 = false

// 关闭通道
var loadOnce sync.Once
var wg sync.WaitGroup
var look sync.RWMutex


var (
	// 参选渠道
	ch1 = make(chan int)
	// 候选人渠道
	ch2 = make(chan int,1)
	// 投票通道
	ch3 = make(chan int,1)

	)

// 存储节点信息
var nodeTable = make(map[int]*Raft1,raftCount1)

// 声明leader 对象
type Leader1 struct {
	// 任期
	Term int
	// LeaderId编号
	LeaderId int
}
// 0 还没上任  -1 没有编号
var leader1 = Leader1{0, -1}

// 声明raft
type Raft1 struct {
	// 锁
	mu sync.Mutex
	// 节点编号
	me int
	// 当前任期
	currentTerm int
	// 为哪个节点投票
	votedFor int
	// 3个状态
	// 0 follower   1 candidate  2 leader
	state int
	// 发送最后一条数据的时间
	lastMessageTime int64
	// 设置当前节点的领导
	currentLeader int
	// 节点间发信息的通道
	message chan bool
	// 选举通道
	electCh chan bool
	// 心跳信号的通道
	heartBeat chan bool
	// 返回心跳信号的通道
	heartbeatRe chan int
	// 超时时间
	timeout int
}

func main() {
	for i:=1;i<=raftCount1;i++ {
		go raft(i)
	}
	for  {
		if leader1.LeaderId > 0 {
			fmt.Printf("大佬是:%d,term:%d \n",leader1.LeaderId,leader1.Term)
			time.Sleep(time.Second * 3)
			for k,v :=range nodeTable{
				fmt.Printf("key:%d value:%v\n",k,v.votedFor)
			}
			// 3秒钟后置空 大佬id 重新选择
			//leader1.LeaderId = 0
			//flag2 = false
		}
	}

}

// 发起选举
func raft(i int)  {
	// 实例化选举人
	rf := Raft1{
		mu:              sync.Mutex{},
		me:              i,
		currentTerm:     0,
		votedFor:        0,
		state:           0,
		lastMessageTime: 0,
		currentLeader:   0,
		message:         make(chan bool),
		electCh:         make(chan bool,raftCount1),
		heartBeat:       make(chan bool,1000000),
		heartbeatRe:     make(chan int,10000),
		timeout:         0,
	}
	look.Lock()
	nodeTable[i] = &rf
	look.Unlock()
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())
	// 随机时间
	n := randRange1(150,300)
	fmt.Printf("i=%d 的等待时间是%d\n",i,n)
	// 延期n，是否有人发起选票，如果有就不发起了
	for  {
		if leader1.LeaderId >0  {
			continue
		}
		select {
		// 等待
		case <-time.After(time.Duration(n) * time.Millisecond):
		}
		var  result = false
		for !result {
			// 选主逻辑
			result = rf.election_one_round1(&leader1)
		}
	}


	fmt.Println("选举大佬：",leader1.LeaderId)

	// 开启监控心跳渠道
	//rf.r1()
	////已经有领导了 , 开始保持心跳
	//for  {
	//	if leader1.LeaderId > 0 {
	//		if leader1.LeaderId == rf.me{
	//			fmt.Printf("主节点%d发起心跳\n",rf.me)
	//			// 循环发心跳
	//			for i:=1;i<=raftCount1;i++ {
	//				if i == rf.me {
	//					continue
	//				}
	//				rf.heartBeat <- true
	//			}
	//			// 设置回复心跳的时间 1s 秒钟的心跳回复时间
	//			rf.lastMessageTime = time.Now().Unix() + 1
	//			time.Sleep(time.Second)
	//		}
	//		//else{
	//		//	fmt.Printf("子节点 %d，心跳回复\n",rf.me)
	//		//	// 回复心跳
	//		//	ret := <-nodeTable[leader1.LeaderId].heartBeat
	//		//	if ret {
	//		//		nodeTable[leader1.LeaderId].heartbeatRe <- rf.me
	//		//	}
	//		//}
	//	}else{
	//		fmt.Println("没有大佬")
	//		time.Sleep(time.Millisecond * 50)
	//	}
	//}
}


// 监控回复心跳
func (rf *Raft1)r1()  {
	for {
		// 备点
		if leader1.LeaderId != rf.me {
			select {
			case nodeTable[leader1.LeaderId].heartbeatRe <- rf.me:
				fmt.Printf("%d 已经回复了心跳\n",rf.me)
			}
		}else{
		// leader

		}
	}
}

// 选举逻辑
func (rf *Raft1)election_one_round1(leader *Leader1) bool  {
	// 发起选票的选举人
	go func() {
		for  {
			// 判断是否已经有了领导
			if leader1.LeaderId > 0 {
				time.Sleep(time.Second * 3)
				continue
			}
			// 开始监控候选人渠道
			select {
			case q:=<-ch1:
				// 获取候选人
				fmt.Println("成为备选人：",q)
				// 关闭候选人渠道
				//close(ch1)
				if flag2 || leader1.LeaderId > 0 {
					nodeTable[q].votedFor = leader1.LeaderId
					continue
				}
				flag2 = true
				//ch2 <- q
				// 投票给自己
				nodeTable[q].votedFor = q
				// 遍历所有节点拉选票
				for i := 0; i <= raftCount1; i++ {
					// 拉选票
					fmt.Println("开始拉选票：",i);
					go func(i int) {
						if i != nodeTable[q].me {
							// 如果没有leader
							if _,ok := nodeTable[i] ; ok {
								nodeTable[i].votedFor = q
							}
							if leader.LeaderId < 0 {
								// 设置投票
								nodeTable[q].electCh <- true
							}
						}
					}(i)
				}
			}
		}
	}()
	// 统计选票
	go rf.t1()
	select {
	case <-ch1:
	case ch1 <- rf.me:
	}
	fmt.Println("选举逻辑结束")
	return  true
}
// 统计票数
func (rf *Raft1)t1()  {
	n := 1
	for {
		if leader1.LeaderId > 0 {
			time.Sleep(time.Millisecond * 50)
			continue
		}
		ok := <- rf.electCh
		if ok {
			n++
			if n >= raftCount1 / 2 {
				fmt.Printf("票数为%d,大于等于%d的半数 \n",n,raftCount1)
				rf.setL(rf.me)
			}
		}
	}
}

// 成为leader
func (rf *Raft1)setL(i int)  {
	leader1.LeaderId = i
	leader1.Term ++
}

// 随机值
func randRange1(min, max int64) int64 {
	return rand.Int63n(max-min) + min
}
