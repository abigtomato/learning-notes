package com.abigtomato.jdk._RefType_ThreadLocal;

public class M {

    @Override
    protected void finalize() throws Throwable {
        System.out.println("finalize");
    }
}
