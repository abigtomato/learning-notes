package com.abigtomato.jdk._Sync;

import java.util.concurrent.CountDownLatch;

public class T02_CAS_Unsafe_Sync {

    private static volatile int m = 0;

    public static void main(String[] args) {
        Thread[] threads = new Thread[100];
        CountDownLatch latch = new CountDownLatch(threads.length);

        Object o = new Object();

        for (int i = 0; i < threads.length; i++) {
            threads[i] = new Thread(() -> {
                /*
                * synchronized最早期版本的实现是重量级锁：
                *   1.锁的管理交由操作系统去完成，管理锁的获取者，锁的状态，线程间状态；
                *   2.等操作系统完成后，反馈给jvm后才能继续执行。
                * 锁升级的过程（新版的synchronized）：
                *   1.
                * */
                synchronized (o) {
                    for (int j = 0; j < 10000; j++) {
                        m++;
                    }
                    latch.countDown();
                }
            });
        }
    }
}
