package com.abigtomato.jdk._Volatile;

/**
 * 饿汉式单例
 * 类加载到内存后就实例化一个单例，JVM保证线程安全
 */
public class T03_Singleton {

    private static final T03_Singleton INSTANCE = new T03_Singleton();

    private T03_Singleton() {}

    // 类加载的时候直接初始化，永远只会存在一个对象
    public static T03_Singleton getInstance() {
        return INSTANCE;
    }

    public void m() {
        System.out.println("m");
    }

    public static void main(String[] args) {
        T03_Singleton m1 = T03_Singleton.getInstance();
        T03_Singleton m2 = T03_Singleton.getInstance();
        System.out.println(m1 == m2);
    }
}
