package com.abigtomato.jdk._RefType_ThreadLocal;

import java.util.concurrent.TimeUnit;

public class ThreadLocal_01 {

    volatile static Person p = new Person();

    public static void main(String[] args) {
        new Thread(() -> {
            try {
                TimeUnit.SECONDS.sleep(2);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
            System.out.println(p.name);
        }).start();

        new Thread(() -> {
            try {
                TimeUnit.SECONDS.sleep(2);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
            p.name = "lisi";
        }).start();
    }

    static class Person {
        String name = "zhangsan";
    }
}
