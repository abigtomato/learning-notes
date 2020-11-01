package com.abigtomato.jdk;

import java.util.Random;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

public class test {

    private static final int threadNum = 10;

    public static void main(String[] args) throws InterruptedException {
        CountDownLatch countDownLatch = new CountDownLatch(10);
        ExecutorService threadPool = Executors.newFixedThreadPool(100);

        for (int i = 0; i < threadNum; i++) {
            threadPool.execute(() -> {
                try {
                    Thread.sleep(5000 + new Random().nextInt(10000));
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
                countDownLatch.countDown();
                System.out.println("countDownLatch.getCount() = " + countDownLatch.getCount());
            });
        }
        countDownLatch.await();
        System.out.println("end!");

        threadPool.shutdown();
        for (;;) {
        }
    }
}
