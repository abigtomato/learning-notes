# 1.类加载子系统

## 1.1.内存结构

![image-20200602194832102](assets/image-20200602194832102.png)

<img src="assets/image-20200602195847508.png" alt="image-20200602195847508"  />

## 1.2.类加载过程

1. 由类加载子系统从磁盘文件系统或网络加载.class文件（class文件在内容开头有特定的文件标识）；
2. 类加载器只负责将class文件加载进内存，决定其是否能运行的是执行引擎；
3. 类会被加载进内存中方法区，这块空间会存储类的信息和运行时的常量池信息，字符串字面量和数字常量。

![image-20200602204341276](assets/image-20200602204341276.png)

### 1.2.1.Loading

1. 通过类的全限定名获取定义此类的二进制字节流；

2. 将字节流对应的静态存储结构转化为方法区的运行时数据结构；

3. 在内存中生成一个代表该类的`java.lang.Class`对象，作为方法区中该类的数据访问入口。

### 1.2.2.Linking

<img src="assets/image-20200602204803148.png" alt="image-20200602204803148"  />

* 验证：通过工具打开.class字节码文件，16进制内容的前4位，就是文件标识

![image-20200602204954688](assets/image-20200602204954688.png)

### 1.2.3.Initialization

![image-20200602205442908](assets/image-20200602205442908.png)

## 1.3.类加载器分类

![image-20200602210753771](assets/image-20200602210753771.png)

1. JVM支持的两种类型：引导类加载器（Bootstrap ClassLoader），自定义类加载器（User-Defined ClassLoader）；

2. 所有派生于抽象类ClassLoader的类加载器都是自定义类加载器.。

   ![image-20200602211110607](assets/image-20200602211110607.png)

### 1.3.1.启动类加载器（引导类加载器，Bootstrap ClassLoader）

1. 由C/C++语言编写，嵌套在JVM内部；
2. 用来加载Java的核心库（JAVA_HOME/jre/lib/rt.jar，resources.jar，sun.boot.class.path），提供JVM自身需要的类；
3. 并不继承自java.lang.ClassLoader，没有父加载器；
4. 加载扩展类和应用程序类加载器，并指定为他们的父类加载器；
5. Bootstrap加载器只加载包名为java，javax，sun等开头的类。

### 1.3.2.扩展类加载器（Extension ClassLoader）

1. 由Java语言编写，sun.misc.Launcher$ExtClassLoader实现；
2. 派生于ClassLoader类；
3. 父类加载器是启动类加载器；
4. 从java.ext.dirs系统属性所指定的目录或JDK安装目录的jre/lib/ext子目录中加载类库（用户自定义的jar放在该目录下也会被加载）。

### 1.3.3.应用程序类加载器（系统类加载器，AppClassLoader）

1. 由Java语言编写，sun.misc.Launcher$AppClassLoader实现；
2. 派生于ClassLoader类；
3. 父类加载器是启动类加载器；
4. 负责加载环境变量classpath或系统属性java.class.path指定路径下的类库；
5. 是程序中的默认类加载器；
6. 通过ClassLoader#getSystemClassLoader()方法可以获取到该类加载器。

