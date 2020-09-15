package com.abigtomato.jdk._Sync;

import java.util.Arrays;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.atomic.AtomicInteger;

public class T03_AtomicInteger {

    /*
    * CAS（compare and swap比较和交换）：
    *   1.当多线程操作共享变量时，先获取值计算结果并将此时获取的值设置为自己的期望值（记录变量的当前状态）；
    *   2.在将结果回写之前先读取变量的最新值，若和期望值不符则读取最新值并更新期望和重新计算（被其他线程修改了）；
    *   3.反复执行第2步操作直到符合期望，再将重新计算后的值回写（保证操作的是变量的最新状态）。
    * ABA问题怎么解决？
    *   1.描述：当线程将计算结果回写的时候发现最新值和期望值一样，但此时的最新值可能是其他线程修改后又再次修改回来的（A - B - A），变量可能并非该线程读取时的状态了；
    *   2.解决：加入版本号来解决，每次操作都会有递增的概念，回写之前同时比较版本号。
    * CAS修改时的原子性问题怎么解决？
    *   1.描述：多核CPU在使用CAS操作数据时会有操作的原子性问题，也就是说cmpxchg这条汇编指令会有原子性问题。
    *   2.解决：将指令修改为lock cmpxchg，相当于让CPU执行了一个锁总线的操作，本次只能有一个核心的cmpxchg指令可以被执行。
    * */
    // 轻量级锁，无锁，自旋锁
    private static final AtomicInteger m = new AtomicInteger(0);

    public static void main(String[] args) throws InterruptedException {
        Thread[] threads = new Thread[100];
        CountDownLatch latch = new CountDownLatch(threads.length);

        for (int i = 0; i < threads.length; i++) {
            threads[i] = new Thread(() -> {
                for (int j = 0; j < 1000; j++) {
                    // 原子性操作
                    m.incrementAndGet();
                }
                latch.countDown();
            });
        }

        Arrays.stream(threads).forEach(Thread::start);

        latch.await();

        System.out.println(m.get());
    }
}
