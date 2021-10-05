#include <stdio.h>
#include <stdlib.h>

// Допустим описана подобная структура
typedef struct Coord
{
    int x; // координата по x 
    int y; // координата по y
} Point;

// Примеры работы с указателями на атрибуты
int main(){
    // Инициализация (присвоение начальных значений) 
    // первое число - первому атрибуту (x)
    // второе число - второму атрибуту (y)

    Point coord = {7, 5}; // объект coord
    Point *directPoint;  // уаказатель на объект типа Point
    int *px, *py;       // указатели на int

    directPoint = &coord; // адрес coord присваивается directPoint

    // Присваивание адресов атрибутов coord:
    // Сначала получаем доступ к атрибуту,
    // а затем вычисляем его адрес
    px = &(coord.x);
    py = &(coord.y);

    // Отладочная наглядная печать
    printf("Object! px: %p coord.x: %d\n", px, coord.x);
    printf("Object! py: %p coord.y: %d\n", py, coord.y);

    // Присваивание адресов атрибутов через указатель directPoint:
    // Сначала получаем доступ к атрибуту,
    // а затем вычисляем его адрес
    px = &(directPoint->x);
    py = &(directPoint->y);

    // Отладочная наглядная печать
    printf("Direct! px: %p coord.x: %d\n", px, coord.x);
    printf("Direct! py: %p coord.y: %d\n", py, coord.y);

    // Получение значений атрибутов через scanf 
    scanf("%d%d", &(coord.x), &(coord.y));

    // Отладочная наглядная печать
    printf("scanf! coord.x: %d\n", coord.x);
    printf("scanf! coord.y: %d\n", coord.y);

    return 0;
}