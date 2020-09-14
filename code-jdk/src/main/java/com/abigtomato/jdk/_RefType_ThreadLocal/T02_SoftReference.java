package com.abigtomato.jdk._RefType_ThreadLocal;

import java.lang.ref.SoftReference;
import java.util.Arrays;

public class T02_SoftReference {

    public static void main(String[] args) throws InterruptedException {
        // 实验的前置条件：-Xmx: 20 将堆内存设置为20mb
        // 软引用：m指向SoftReference是强引用，SoftReference对象内的成员变量指向的10mb的字节数组是软引用
        // 适用场景：缓存，比如大文件，写入内存中后使用软引用指向，需要使用时直接从内存获取，不需要时软引用自动断开，释放内存
        SoftReference<byte[]> m = new SoftReference<>(new byte[1024 * 1024 * 10]);
        System.out.println(Arrays.toString(m.get()));

        // 一次gc后软引用指向的对象未被回收，因为此时内存足够
        System.gc();
        Thread.sleep(500);
        System.out.println(Arrays.toString(m.get()));

        // 再次分配15mb的字节数组，超出堆内存，这时软引用指向的对象会被释放
        byte[] b = new byte[1024 * 1024 * 15];
        System.out.println(Arrays.toString(m.get()));
    }
}
