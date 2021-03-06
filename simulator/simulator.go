package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/prakhar0409/Distributed-Ledger/node"
	"reflect"
)



func main(){
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run simulator.go <num_nodes>")
		return
	}
	
	maxTxns := 1000
	num_nodes, err := strconv.Atoi(os.Args[1])
    if err != nil  {
        fmt.Println("Panic: Incorrect arguments")
        return
    }

    node_list := make([]node.Node, num_nodes)					//pointer to array
    fmt.Println(reflect.TypeOf(node_list))

   	quit := make(chan int)										//pointer to channel
   	quitsim := make(chan int)
	for i := 0; i < num_nodes; i++ {
		node_list[i].Initialize(i,node_list,maxTxns,quit,quitsim);		//array, channel, maps are pointers
		// fmt.Println(node_list[i].nodeid);	//cant refer to unexported field or method
	}

	for i := 0; i < num_nodes; i++ {
		go node_list[i].Run();
	}


	//channel for quitting
	num_iters := 0
	num_quits := 0
	ok := false
	for{
		select {
		case _,ok = <-quit:
			if ok {
				// fmt.Println("gjg");
				num_quits++
			}
			if num_quits >= num_nodes{
				goto end
			}
		default:
			// num_iters++				//uncomment this to allow simulator to quit after sometime
			if(num_iters > 1000000000){
				for i :=0; i<num_nodes;i++{
					node_list[i].Live = 0
				}				
			}
		}
	}
	end:
		fmt.Println("Simulator Exiting:)")
    // <- quit
}
