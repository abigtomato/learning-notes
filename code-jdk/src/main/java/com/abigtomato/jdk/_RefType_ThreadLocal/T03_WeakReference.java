package com.abigtomato.jdk._RefType_ThreadLocal;

import java.lang.ref.WeakReference;

public class T03_WeakReference {

    public static void main(String[] args) {
        // 弱引用
        WeakReference<M> m = new WeakReference<>(new M());
    }
}
