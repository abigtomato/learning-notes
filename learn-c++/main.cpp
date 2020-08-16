#include <iostream>

// 指定命名空间std，包含c++标准库的变量和函数
using namespace std;

// namespace定义，类似java包
namespace test01 {
    int a = 10;
    int fun(int a) {
        cout << "Hello test01!" << endl;
        return a + 10;
    }
}

namespace test02 {
    int a = 20;
    int fun(int b) {
        cout << "Hello test02!" << endl;
        return b + 20;
    }
}

// 当一个形参有默认值，其后的所有形参都必须有默认值
void fun001(int a, int b = 2, int c = 3) {
    cout << a << endl;
    cout << b << endl;
    cout << c << endl;
}

// 形参拷贝
void swap001(int a, int b) {
    int temp = a;
    a = b;
    b = temp;
    cout << a << "\t" << b << endl;
}

// 指针传递
void swap002(int *a, int *b) {
    int *temp = a;
    a = b;
    b = temp;
    cout << *a << "\t" << *b << endl;
}

int & fun002(int &r) {
    return r;
}

// 引用传递
void swap003(int &a, int &b) {
    int temp = a;
    a = b;
    b = temp;
    cout << a << "\t" << b << endl;
}

// 内联函数，编译时会将其内容转入调用方
inline int square(int x) {
    int j;
    j = x * x;
    return j;
}

int main() {
    cout << "Hello World!" << endl;

    // 变量类型，内存byte大小
    int i = 10;
    int j(20);
    cout << "int: " << sizeof(i) << endl;
    float f = 3.14f;
    cout << "float: " << sizeof(f) << endl;
    double d = 2.345;
    cout << "double: " << sizeof(d) << endl;
    char c = 'c';
    cout << "char: " << sizeof(c) << endl;
    char arr[] = "abcd";
    cout << "arr: " << sizeof(arr) << endl;
    int *p = &i;
    cout << "p: " << sizeof(p) << endl;
    bool b = true;
    cout << "b: " << sizeof(b) << endl;

    // cin >> i;
    // cout << "i: " << i << endl;

    // namespace调用，类似java包
    cout << test01::fun(10) << endl;
    cout << test02::fun(10) << endl;

    // 默认参数测试
    fun001(1);

    // 引用变量
    int &ref = i;
    cout << ref << endl;   
    cout << &ref << endl;   // 引用的地址就是其引用变量的地址（引用变量不会申请内存空间）
    cout << &i << endl;
    cout << &p << endl;     // 指针的地址（指针变量会在内存中申请4byte的空间）

    // const修饰符
    ref = 20;
    cout << i << endl;      // 通过引用修改变量的值
    const int &cref = i;
    cout << cref << endl;   // 通过常引用不能修改变量的值

    const double PI = 3.1415926;    // 常量定义
    cout << PI << endl;

    int a = 10;
    const int *ap = &a;             // 指针指向常量，不能修改指针指向的值
    int *const cap = &a;            // 常量指针，指针保存的地址不能被修改
    const int *const ccap = &a;     // 指针指向的值不能改变，指针的值也不能改变

    // 形参拷贝，指针传递，引用传递
    int x = 10;
    int y = 20;
    swap001(x, y);
    swap002(&x, &y);
    swap003(x, y);

    // 函数的返回值是引用类型，可以放在赋值号左边
    int aa = 1000;
    fun002(aa) = 30;
    cout << aa << endl;

    // 内联函数，调用时无需开辟新的栈空间，适用于代码少的函数
    int ii = 10, temp;
    temp = square(ii);
    cout << temp << endl;

    return 0;
}