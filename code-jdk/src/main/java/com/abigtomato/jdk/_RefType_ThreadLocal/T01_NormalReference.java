package com.abigtomato.jdk._RefType_ThreadLocal;

import java.io.IOException;

public class T01_NormalReference {

    public static void main(String[] args) throws IOException {
        // 强引用：str就是强引用一个字符串对象，若引用指向空，则之前指向的对象会被回收
        M m = new M();
        m = null;

        System.gc();    // DisableExplicitGC
        System.out.println(m);

        int read = System.in.read();    // 阻塞main线程，给gc线程执行时间
    }
}
