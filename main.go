package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Animal interface {
	modify()
	defercall()
}

type Person struct {
	age   int
	child *Person
	m map[uint32]int
}

func (p *Person) modify() {
	if p == nil {
		return
	}
	p.age = 10
	//p.child.age++
}

type Person2 struct {
	age   int
	child *Person
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

func TestGenBlockIdSeqId() *Person {
	objID := uint64(341499936400999360)
	blockID, seqID := GenBlockIdSeqId(objID)
	groupID := GetGroupID(objID)
	fmt.Println(blockID, seqID, groupID)
	//assert.Equal(t, uint32(objID>>32), blockID)
	//assert.Equal(t, uint32(objID&0xffffffff), seqID)
	p1 := new(Person)
	p2 := new(Person)
	//p1 := &Person{}
	//p2 := &Person{}
	println(p1)
	println(p2)
	//fmt.Printf(" p2:%p", p2)
	return p2
}

func TestGenGroupId() {
	blockId := uint32(2836495)
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
	for i := 0; i < randSize; i++ {
		s = append(s, strconv.Itoa(randSize))

	}
	_ = s
}

func reEnter(rw sync.RWMutex) {
	rw.RLock()
	defer rw.RUnlock()
	fmt.Println("keyi")
}

func finbina(n int) int {
	dp := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		if i < 2 {
			dp[i] = i
			continue
		}
		dp[i] = (dp[i-1]%1000000007 + dp[i-2]%1000000007) % 1000000007
	}
	return dp[n]
}

func dataRace() (ret int) {
	ret = 5
	go func() {
		i := ret
		fmt.Println(i)
	}()
	return 6
}

func goMaxProcs() {
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	wg.Add(19)
	for i := -1; i < 10; i++ {
		go func() {
			defer wg.Done()
			fmt.Printf("A: %d\n", i)
		}()
	}

	for i := -1; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Printf("B: %d\n", i)
		}(i)
	}
	wg.Wait()

}

func structCompare() {
	type com struct {
		a int
		_ string
		s [3]int
		sli []int
	}
	c1 := com{sli: []int{3,4}}
	c2 := com{sli: []int{1,2,3}}
	fmt.Println(reflect.DeepEqual(c1, c2))

	v1 := struct {
		A int
		B string
	} {
		A: 10,
		B: "bb",
	}
	v2 := struct {
		A int
		B string
	} {
		A: 10,
		B: "bbc",
	}
	fmt.Println(v1 == v2)
}

func gmp() {
	runtime.GOMAXPROCS(1)
	go func() {
		for {
		}
	}()

	time.Sleep(time.Millisecond)
	println("OK")
}

func regex() {
	bucketUpper := regexp.MustCompile("^http://(.*?)\\.bjcnc.*?/(.*)")
	//bucketLower := regexp.MustCompile("http://bjcnc.*?/(.*?)/(.*)")
	res := bucketUpper.FindStringSubmatch("http://app-content.bjcnc.scs.sohucs.com/article/6784756376276041728")
	match := bucketUpper.Match([]byte("http://app-content.bjcnc.scs.sohucs.com/article/6784756376276041728"))
	fmt.Println(match)
	fmt.Println(len(res))
	fmt.Println(res)
	//
	//res2 := bucketLower.FindAllStringSubmatch("http://bjcnc.scs-internal.sohucs.com/app-content/article/6783718609412685839", -1)
	//
	//fmt.Println(res2)

	//numberRegex := regexp.MustCompile("abc([0-9]*)(f)")
	//res3 := numberRegex.FindAllStringSubmatch("abc44fabc6f", 1)
	//fmt.Println(res3)


}

func tt(stack []int) []int {
	helperStack := make([]int, 0)
	for len(stack) != 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for len(helperStack) != 0 && top < helperStack[len(helperStack)-1] {
			stack = append(stack, helperStack[len(helperStack)-1])
			helperStack = helperStack[:len(helperStack)-1]
		}
		helperStack = append(helperStack, top)
	}

	return helperStack
}


func testLock() {
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()
}

func walkPath() {
	filepath.Walk("/data2/p2online/datanode/data/trash", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		cells := strings.Split(info.Name(), ".")
		if len(cells) != 5 {
			_ = os.Remove(path)
			return nil
		}
		generateTimeStamp := cells[4]
		i, _ := strconv.ParseInt(generateTimeStamp, 10, 64)
		t := time.Unix(i, 0)
		if time.Now().Sub(t) > 24 * time.Hour {
			_ = os.Remove(path)
		}
		return nil
	})
}

type ctxStr string

func ctxKV() {
	ctx := context.WithValue(context.Background(), "liyuan", "123")
	ctx = context.WithValue(ctx, "liyuan", "456")
	v := ctx.Value("liyuan")
	fmt.Printf("%v\n", v)
}


func TestMapAssign() {
	// Map 整体替换会不会race
	mut := sync.Mutex{}
	p := &Person{
		age:   0,
		child: nil,
		m:     nil,
	}
	p.m = make(map[uint32]int)
	p.m[3] = 3

	go func() {
		mut.Lock()
		age := p.age
		m := p.m
		fmt.Println(m)
		fmt.Println(age)
		mut.Unlock()
	}()
	tmp := make(map[uint32]int)
	mut.Lock()
	p.m = tmp
	p.age = 4
	mut.Unlock()
}

func (p *Person) defercall() {
	if p == nil {
		return
	}
	p.age--
}

func TestRangeDefer(s []int) []Animal {
	ans := make([]Animal, 0)
	f := func() {
		for _, v := range s {
			var a Animal
			a = &Person{age: v}
			ans = append(ans, a)
			defer a.defercall()
		}
	}
	f()
	fmt.Println(ans)
	return ans
}

func defer_call()  {
	defer func(){
		fmt.Println("11111")
	}()
	defer func(){
		fmt.Println("22222")
	}()


	defer func(){
		fmt.Println("33333")
	}()

	fmt.Println("111 Helloworld")

	panic("Panic 1!")

	panic("Panic 2!")

	fmt.Println("222 Helloworld")
}
func callDefer() {
	defer func() {
		if r := recover(); r!= nil {
			fmt.Println("Recover from r : ",r)
		}
	}()
	defer_call()

	fmt.Println("call defer world")
}
func simplifyPath(path string) string {
	stack := make([]string, 0)
	var str string
	for i := 0; i < len(path); i++ {
		if path[i] == '/' {
			if len(str) != 0 {
				stack = append(stack, str)
				str = ""
			}
			continue
		}
		fmt.Println(string(path[i]))
		if path[i] == '.' && (i == len(path)-1 || path[i+1] == '/') {
			continue
		}
		if path[i] == '.' && path[i+1] == '.' && (i == len(path)-2 || path[i+2] == '/') {
			if len(stack) > 0 {

				stack = stack[:len(stack)-1]
			}
			continue
		}
		str += string(path[i])
	}
	if len(str) != 0 {
		stack = append(stack, str)
	}
	ans := "/"
	for k, s := range stack {
		ans += s
		if k < len(stack) - 1 {
			ans += "/"
		}
	}
	return ans
}

func main() {
	callDefer()
	fmt.Println("333 Hello world")

	fmt.Println(simplifyPath("/home/"))


	//TestMapAssign()
	//time.Sleep(5 * time.Second)
	//sort_alogrithm.QuickSort([]int{3, 7, 3, 4, 3, 2, 1, 8, 6}, 0, 8)
	//ctxKV()
	//walkPath()
	//testLock()
	//regex()
	//gmp()
	//structCompare()
	//dataRace()
	//goMaxProcs()

	//ans := finbina(45)
	//fmt.Println(ans)

	//TestGenGroupId()
	//p := TestGenBlockIdSeqId()
	//fmt.Println(p)
	//s := "9223372036854775808"
	//fmt.Println(__Atoi.MyAtoi(s))

	//var inter *Person
	//fmt.Println(reflect.TypeOf(inter))
	//var yy *Person2 = nil
	//fmt.Println(inter == nil)

	//s := __Longest_subString.LengthOfLongestSubstringBruteForce(" ")
	//s := __Longest_subString.LengthOfLongestSubstring("abcabcbb")
	//fmt.Println(s)
}
