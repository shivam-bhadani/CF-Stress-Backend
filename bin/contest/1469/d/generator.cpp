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
	int n = randomNumberBetween(1, 200000);
	cout << n << endl;
}