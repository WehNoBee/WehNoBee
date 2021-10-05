// ЗАДАНИЕ:
// Найдите ошибки в коде! Всего ошибок : 4.
// Ответ находиться снизу.

#include "pch.h"
#include <iostream>
#include<conio.h>
using namespace std;
int main()
{
	setlocale(LC_ALL, "ru");
	int a = 0;
	int simple = 0;
	int notsimple = 0;
	while (a != 0) {
		cout << "Введите число:" << endl;
		cout << "> "
		cin >> a;
			if (a % 2 == 0 && a != 0) {
				simple += 1;
				break;
			}
			else if (a % 2 == 1 && a != 0) {
				notsimple += 2;
				continue;
			}
	}
	cout << "Простых чисел: " << simple << endl << "Не простых: " << notsimple;
	return 0;
}









// Ответы: 1. В строке 12: int = 1;
//		   2. В строке 17: cout << "> ";
//         3. В строке 24: notsimple += 1;
//		   4. В строке 21: continue;