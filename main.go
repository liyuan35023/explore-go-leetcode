package main

import (
	"bytes"
	"fmt"
	"gopkg.in/gomail.v2"
	"math/rand"
	"strconv"
	"sync"
	"time"
)
type Person struct {
	age int
	child *Person
}

func (p *Person) modify() {
	p.age = 10
	p.child.age++
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[int(d)]
}

func TestGenBlockIdSeqId() {
	objID := uint64(177480035924445773)
	blockID, seqID := GenBlockIdSeqId(objID)
	groupID := GetGroupID(objID)
	fmt.Println(blockID, seqID, groupID)
	//assert.Equal(t, uint32(objID>>32), blockID)
	//assert.Equal(t, uint32(objID&0xffffffff), seqID)
}

func TestGenGroupId() {
	blockId := uint32(25633063)
	fmt.Println(GenGroupId(blockId))
	seqId := uint32(7355)
	objId := GenObjectId(blockId, seqId)
	fmt.Println(objId)
}

func GenBlockIdSeqId(objectId uint64) (blockId, seqId uint32) {
	blockId = uint32(objectId >> 32)
	seqId = uint32((objectId << 32) >> 32)

	return
}
func GenGroupId(storeId uint32) (groupId uint32) {
	groupId = uint32(storeId >> 10)
	return
}

func GenObjectId(blockId uint32, seqId uint32) uint64 {
	objectId := (uint64(blockId) << 32) + uint64(seqId)

	return objectId
}

func GetGroupID(objID uint64) uint32 {
	blockID, seqID := GenBlockIdSeqId(objID)
	groupID := GenGroupId(blockID)
	fmt.Println("Group ID:", groupID)
	//fmt.Println("Block ID:", blockID)
	fmt.Println("Sequence ID:", seqID)
	return groupID
}

func escapee() {
	randSize := rand.Int()
	s := make([]string, 10, 100)
	//str := "hello"
	for i:=0;i<randSize;i++ {
		s = append(s, strconv.Itoa(randSize))

	}
	_ = s
}

func reEnter(rw sync.RWMutex) {
	rw.RLock()
	defer rw.RUnlock()
	fmt.Println("keyi")
}

func sendEmail(subject string, info string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "pallas2@sohu-inc.com")
	m.SetHeader("To", "yuanli213962@sohu-inc.com")
	m.SetHeader("Subject", subject)
	timeStr := time.Now().Format("2006-1-2#15:04:05")
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("Time:%s\n", timeStr))
	buf.WriteString("Check Result:\n")
	buf.WriteString(info)
	m.SetBody("text/plain", buf.String())
	d := gomail.Dialer{Host: "transport.mail.sohu-inc.com", Port: 25}
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func main() {
	//ret := _8_four_sum.FourSum([]int{-3,-2,-1,0,0,1,2,3}, 0)
	//fmt.Println(ret)
	var mmm sync.Map
	mmm.Store("z", 26)
	mmm.Store("y", 25)
	mmm.Store("z", 24)
	go func() {
		mmm.Range(func(key, value interface{}) bool {
			k := key.(string)
			v := value.(int)
			fmt.Println(k, " ", v)
			return true
		})
	}()

	go func() {
		mmm.Store("a", 1)
		time.Sleep(2*time.Second)
		mmm.Store("b", 2)
	}()

	time.Sleep(10 * time.Second)
	//s := "9223372036854775808"
	//fmt.Println(__Atoi.MyAtoi(s))



	//s := __Longest_subString.LengthOfLongestSubstringBruteForce(" ")
	//s := __Longest_subString.LengthOfLongestSubstring("abcabcbb")
	//fmt.Println(s)
}
