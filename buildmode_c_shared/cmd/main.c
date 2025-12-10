#include <stdio.h>
#include <stdlib.h>
#include "libcalculate.h"  // 包含你提供的头文件

int main() {    
    // 测试加法函数
    char method[] = "plus";
    int a = 2;
    int b = 3;
    int sum = calculate(method, a, b);
    printf("%s (%d, %d) = %d\n", method, a, b, sum);

    return 0;
}