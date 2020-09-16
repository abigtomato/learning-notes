package com.abigtomato.jdk._Volatile;

import java.util.concurrent.TimeUnit;

public class T01_Hello_Volatile {

    // 线程可见性
    volatile boolean running = true;

    void m() {
        System.out.println("m start");
        while (running) {
        }
        System.out.println("m end!");
    }

    public static void main(String[] args) throws InterruptedException {
        T01_Hello_Volatile t = new T01_Hello_Volatile();

        new Thread(t::m, "t1").start();

        TimeUnit.SECONDS.sleep(1);

        t.running = false;
    }
}
