package main

import "fmt"

/* channel基础使用 */

type Cat struct {
	Name string
	Age int
}

// channel遍历
func traversing(iChan chan int) {
	// 管道遍历之前要保证关闭，使其不能再写入
	close(iChan)
	/*
		1.可以使用range取出channel中的数据；
		2.直到取出所有数据，否则for不会退出。
	*/
	for v := range iChan {
		fmt.Printf("value: %v\n", v)
	}
}

func main() {
	// 1.创建一个 chan interface{} 空接口类型的管道（底层为队列结构）
	var iChan chan interface{}
	/*	
		1.无缓冲阻塞读写: 管道是引用类型，必须先通过make分配内存才可以使用，如果不定义缓冲区，则入队一个数据，就需要出队一个数据，否则阻塞读写；
		2.有缓冲写满缓冲区后阻塞读写: 缓冲区长度4，定义了缓冲区才可以只入队不出队，等channel缓冲的数据满了，才会出现阻塞等待。
	*/
	iChan = make(chan interface{}, 4)
	fmt.Printf("指向=%v, 类型=%T, 地址=%p\n", iChan, iChan, &iChan)

	// 2.iChan<- 向管道内部存入数据(入队操作)
	iChan<- make(map[string]int)
	iChan<- make([]float64, 10)
	iChan<- fmt.Sprintf("channel->%v", iChan)
	iChan<- &Cat{Name: "lily", Age: 3,}
	/*
		len(): 获取缓冲区未读取的数据个数
		cap(): 获取缓冲区容量
	*/
	fmt.Printf("个数: %v, 容量: %v\n", len(iChan), cap(iChan))

	// 3.使用内置close()函数关闭管道，只能读取数据无法写入数据
	close(iChan)

	// 4.<-iChan 从管道取出数据（出队操作），当channel为空时，若再次取数据则会报错
	mapper := <-iChan
	slice := <-iChan
	str := <-iChan
	cat := (<-iChan).(*Cat)	// 类型断言转换为原始类型（因为存入channel时是以空接口类型存入的）
	fmt.Printf("map=%v, slice=%v, string=%v\n", mapper, slice, str)
	fmt.Printf("cat=%v, *cat.Name=%v\n", cat, (*cat).Name)

	// 5.管道的遍历测试
	iChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		iChan2 <- i * 2
	}
	traversing(iChan2)
}