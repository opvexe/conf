#include <stdlib.h>
#include <stdio.h>

int main(void)
{
	int *p = NULL;

	// 申请内存
	p = (int *)malloc(100);			// void * 可以转换为任意指针类型

	// 写数据到内存
	*p = 979;

	// 从内存读数据
	printf("*p= %d, p = %p\n", *p, p);	

	// 拓展 内存空间
	int *q = realloc(p, 200);
	*q = 748329;
	printf("*q= %d, q = %p\n", *q, q);	

	// 释放内存
	free(p);
	free(q);

	return 0;
}
