package com.abigtomato.jdk._Sync;

import java.util.Arrays;
import java.util.concurrent.CountDownLatch;

public class T00_CAS_Unsafe {

    private static int m = 0;

    public static void main(String[] args) throws InterruptedException {
        Thread[] threads = new Thread[100];
        CountDownLatch latch = new CountDownLatch(threads.length);

        for (int i = 0; i < threads.length; i++) {
            threads[i] = new Thread(() -> {
                // 发生并发问题
                for (int j = 0; j < 10000; j++) {
                    m++;
                }
                latch.countDown();
            });
        }

        Arrays.stream(threads).forEach(Thread::start);

        latch.await();

        System.out.println(m);
    }
}
