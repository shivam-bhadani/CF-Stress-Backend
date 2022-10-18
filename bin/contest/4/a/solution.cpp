#include<bits/stdc++.h>
using namespace std;

int main() {
    int waterMelon;
    cin >> waterMelon;
    for (int pete = 1 ; pete < waterMelon ; pete++) {
        int billy = waterMelon - pete;
        if (pete % 2 == 0 && billy % 2 == 0) {
            cout << "YES" << endl;
            return 0;
        }
    }
    cout << "NO" << endl;
}