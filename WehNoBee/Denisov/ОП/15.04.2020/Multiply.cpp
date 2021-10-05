#include <iostream>
using namespace std;

int main()
{
    int a;
    int t;
    int i;
    cout << "Choose number." << endl;
    cout << "> ";
    cin >> a;
    for (a; a < 10; a++) {
        for (t = 1; t < 10; t++) {
            if (t == 9)
            cout << a * t << endl;
            else
            cout << a * t << '\t';
        }
    }
    return 0;
}
