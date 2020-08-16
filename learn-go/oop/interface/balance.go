package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Balancer interface {
	DoBalance([]*Instance) (*Instance, error)
}

type Instance struct {
	host string
	port int
}

func (i *Instance) String() string {
	return i.host + ":" + port
}

type RandomBalance struct {
}

func (r *RandomBalance) DoBalance(insts []*Instance) (inst *Instance, err error) {
	lens := len(insts)
	
	if lens == 0 {
		err = errors.New("No instance")
		return
	}

	index := rand.Intn(lens)
	inst = insts[index]

	return
}

type RoundRobinBalance struct {
	curIndex int
}

func (r *RoundRobinBalance) DoBalance(insts []*Instance) (inst *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("No instance")
		return
	}

	inst = insts[r.curIndex]
	r.curIndex = (r.curIndex + 1) % len(insts)

	return
}

type BalanceMgr struct {
	allBalance map[string]Balancer
}

func (b *BalanceMgr) RegisterBalancer(name string, balancer *Balancer) {
	b.allBalance[name] = balancer
}

func (b *BalanceMgr) DoBalance(name string, insts []*Instance) (inst *Instance, err error) {
	if balancer, ok := b.allBalance[name]; !ok {
		err = fmt.Errorf("Not found %s balancer", name)
		return
	}

	inst, err = balancer.DoBalance(insts)
	return 
}

var mgr = &BalanceMgr{
	allBalance: make(map[string]*Balancer),
}

func init() {
	mgr = RegisterBalancer("random", &RandomBalance{})
	mgr = RegisterBalancer("roundrobin", &RoundRobinBalance{})
}

func main() {
	var insts []*Instance
	for i := 0; i < 16; i++ {
		insts = append(insts, &Instance{
			host: fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255)),
			port: 8080,
		})
	}

	var balancer Balancer
	var balanceName = "random"
	
	if len(os.Args) > 1 {
		balanceName = os.Args[1]
	}

	fmt.Printf("use %s balancer.\n", balanceName)

	for {
		inst, err := mgr.DoBalance(conf, insts)
		if err != nil {
			fmt.Println("do balance err: ", err)
			continue
		}

		fmt.Println(inst)
		time.Sleep(time.Second)
	}
}