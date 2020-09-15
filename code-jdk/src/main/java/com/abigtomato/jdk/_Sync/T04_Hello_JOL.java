package com.abigtomato.jdk._Sync;

import org.openjdk.jol.info.ClassLayout;

public class T04_Hello_JOL {

    // JOL：JAVA对象内存布局
    public static void main(String[] args) {
        /*
            java.lang.Object object internals:
             OFFSET  SIZE   TYPE DESCRIPTION                               VALUE
                  0     4        (object header)                           01 00 00 00 (00000001 00000000 00000000 00000000) (1)
                  4     4        (object header)                           00 00 00 00 (00000000 00000000 00000000 00000000) (0)
                  8     4        (object header)                           e5 01 00 f8 (11100101 00000001 00000000 11111000) (-134217243)
                 12     4        (loss due to the next object alignment)
            Instance size: 16 bytes 空对象16个字节
            Space losses: 0 bytes internal + 4 bytes external = 4 bytes total

            1.偏移量0开始的和偏移量4开始的共8个字节代表对象的MarkWord（）；
            2.偏移量8开始的4个字节代表对象的ClassPoint（对象所属的类）；
            3.因为没有成员变量，所有偏移量12开始的4个字节是填充字节（字节对齐），为了让整个对象的字节大小符合被8整除。
        */
        Object o = new Object();
        System.out.println(ClassLayout.parseInstance(o).toPrintable());

        /*
            java.lang.Object object internals:
             OFFSET  SIZE   TYPE DESCRIPTION                               VALUE
                  0     4        (object header)                           28 f7 a8 02 (00101000 11110111 10101000 00000010) (44627752)
                  4     4        (object header)                           00 00 00 00 (00000000 00000000 00000000 00000000) (0)
                  8     4        (object header)                           e5 01 00 f8 (11100101 00000001 00000000 11111000) (-134217243)
                 12     4        (loss due to the next object alignment)
            Instance size: 16 bytes
            Space losses: 0 bytes internal + 4 bytes external = 4 bytes total

            使用synchronized后，此时从偏移量4开始的4个字节发生了变化，是因为偏向锁的信息被添加到了对象的MarkWord上（说白了偏向锁就是将对象的id信息直接贴到锁上，无需争抢竞争）
        */
        synchronized (o) {
            System.out.println(ClassLayout.parseInstance(o).toPrintable());
        }
    }
}
