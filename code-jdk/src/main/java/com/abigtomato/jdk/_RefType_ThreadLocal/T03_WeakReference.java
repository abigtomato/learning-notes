package com.abigtomato.jdk._RefType_ThreadLocal;

import java.lang.ref.WeakReference;

public class T03_WeakReference {

    public static void main(String[] args) {
        // 弱引用：可以通过引用正常访问对象，但如果一个对象只有一个弱引用，gc会直接回收
        WeakReference<M> m = new WeakReference<>(new M());

        System.out.println(m.get());
        System.gc();
        System.out.println(m.get());

        // ThreadLocalMap中的Entry就使用弱引用，其中的key就是指向ThreadLocal对象的弱引用
        /*
        * ThreadLocal为什么使用弱引用——防止内存泄漏：
        * 1.若Entry中的key使用强引用，此时外部所有的强引用断开联系，ThreadLocalMap中的key不会被gc回收，会造成内存泄漏问题；
        * 2.使用弱引用会在外部引用都断开后允许gc回收，但会造成key为null，value无人映射，也会出现内存泄漏问题；
        * 3.所以使用ThreadLocal后需要手动调用remove方法清除k-v对，防止内存泄漏。
        * */
        ThreadLocal<M> tl = new ThreadLocal<>();
        tl.set(new M());
        tl.remove();
    }
}
