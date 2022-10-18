#include<bits/stdc++.h>
using namespace std;

int randomNumberBetween(int a, int b) {
	if (a > b)
		swap(a, b);
	return a + rand() % (b - a + 1);
}
int main(int argc, char* argv[]) {
	srand(atoi(argv[1]));
    cout << "1" << endl;
	int n = randomNumberBetween(1, 20);
	cout << n << endl;
    for(int i=1; i<=n; i++) {
        int x = randomNumberBetween(1, n);
        cout << x;
        if(i!=n) {
            cout << " ";
        }
    }
    cout << endl;
}