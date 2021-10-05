#include <stdio.h>
#include <stdlib.h>

typedef struct Time
{
    unsigned char hour;
    unsigned char min;
    unsigned char sec;
} Timer;

void timePlus(Timer a, struct Time b, Timer * tmp){
    int sec, min, h;
    sec = a.sec + b.sec;
    min = sec / 60 + a.min + b.min;
    h = min / 60 + a.hour + b.hour;
    tmp->sec = sec % 60;
    tmp->min = min % 60;
    tmp->hour = h % 12;

};

void printTime(Timer a)
{
    printf("%02hhu:%02hhu:%02hhu\n", a.hour, a.min, a.sec);
};

int main(){
    Timer t1, t2;
    Timer res;
    scanf("%hhu:%hhu:%hhu", &(t1.hour), &(t1.min), &(t1.sec));
    scanf("%hhu:%hhu:%hhu", &(t2.hour), &(t2.min), &(t2.sec));
    timePlus(t1, t2, &res);
    printTime(res);
    return 0;
}
