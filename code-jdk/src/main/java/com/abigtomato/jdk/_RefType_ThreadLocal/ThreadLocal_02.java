package com.abigtomato.jdk._RefType_ThreadLocal;

import java.util.concurrent.TimeUnit;

public class ThreadLocal_02 {

    // ThreadLocal.ThreadLocalMap threadLocals = null;
    // ThreadLocal类中定义了ThreadLocalMap这个类型
    // Thread类中维护一个ThreadLocalMap threadLocals对象，以ThreadLocal的引用为key
    static ThreadLocal<Person> tl = new ThreadLocal<>();

    public static void main(String[] args) {
        new Thread(() -> {
            try {
                TimeUnit.SECONDS.sleep(2);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
            /*
                public T get() {
                    Thread t = Thread.currentThread();
                    ThreadLocalMap map = getMap(t);
                    if (map != null) {
                        ThreadLocalMap.Entry e = map.getEntry(this);
                        if (e != null) {
                            @SuppressWarnings("unchecked")
                            T result = (T)e.value;
                            return result;
                        }
                    }
                    return setInitialValue();
                }
            */
            System.out.println(tl.get());
        }).start();

        new Thread(() -> {
            try {
                TimeUnit.SECONDS.sleep(2);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
            /*
                public void set(T value) {
                    Thread t = Thread.currentThread();
                    ThreadLocalMap map = getMap(t);
                    if (map != null)
                        map.set(this, value);
                    else
                        createMap(t, value);
                }

                ThreadLocalMap getMap(Thread t) {
                    return t.threadLocals;
                }
             */
            // set()其实是在当前线程的map集合中存储tl（key）和person（value）
            Person person = new Person();
            tl.set(person);
        }).start();

        /*
            void createMap(Thread t, T firstValue) {
                t.threadLocals = new ThreadLocalMap(this, firstValue);
            }

            ThreadLocalMap(ThreadLocal<?> firstKey, Object firstValue) {
                table = new Entry[INITIAL_CAPACITY];
                int i = firstKey.threadLocalHashCode & (INITIAL_CAPACITY - 1);
                table[i] = new Entry(firstKey, firstValue);
                size = 1;
                setThreshold(INITIAL_CAPACITY);
            }

            private static final int INITIAL_CAPACITY = 16;

            static class Entry extends WeakReference<ThreadLocal<?>> {
                Object value;

                Entry(ThreadLocal<?> k, Object v) {
                    super(k);
                    value = v;
                }
            }

            private void setThreshold(int len) {
                threshold = len * 2 / 3;
            }
        */
    }

    static class Person {
    }
}
