package com.abigtomato.jdk._Volatile;

/**
 * 懒汉式单例
 * 虽然达到了按需初始化的目的，但却带来了线程不安全的问题
 */
public class T04_Singleton {

    /*
        对象的创建过程：
        class T {
            int m = 8;
        }
        T t = new T();
        汇编指令：
        0 new #2 <T>
        3 dup
        4 invokespecial #3 <T.<init>>
        7 astore_1
        8 return
    */
    private static volatile T04_Singleton INSTANCE;

    private T04_Singleton() {
    }

    /*
        问：使用DCL单例模式下，需不需要加volatile？
        答：需要加，因为创建对象时的汇编指令可能会发生重排序：
            0 new #2 <T> 半初始化对象，成员变量赋予初始值
            4 invokespecial #3 <T.<init>> 调用构造方法
            7 astore_1  引用和对象关联
        4和7若是发生了CPU指令重排，那会先关联引用和对象，此时INSTANCE就不为空了，此时该线程先去执行权；
        若正好进来一个新线程，外层检索 if (INSTANCE == null) 就会失效，新线程就会使用半初始化的对象，值就是默认值；
        加上了volatile会让该关键字修饰的内存空间在被指令操作时不存在乱序的情况。
    */
    /*
        问：volatile如何阻止指令的乱序执行？
        答：内存屏障
        JVM内存屏障规范：
        Hotspot虚拟机实现内存屏障：lock addl 锁总线的方式
    */
    public static T04_Singleton getInstance() throws InterruptedException {
        // DCL双重检索式（Double Check Lock）
        // 外层检索：防止大量线程直接去竞争锁带来的性能问题
        if (INSTANCE == null) {
            synchronized (T04_Singleton.class) {
                // 内层检索：防止其他通过外层检索的线程又执行一遍内部逻辑
                if (INSTANCE == null) {
                    Thread.sleep(1);
                    // 若不加锁则会出现多个线程创建多个对象的问题，单例则无从谈起
                    INSTANCE = new T04_Singleton();
                }
            }
        }
        return INSTANCE;
    }

    public void m() {
        System.out.println("m");
    }

    public static void main(String[] args) {
        for (int i = 0; i < 100; i++) {
            new Thread(() -> {
                try {
                    System.out.println(T04_Singleton.getInstance().hashCode());
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }).start();
        }
    }
}
