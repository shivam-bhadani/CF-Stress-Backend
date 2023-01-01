#include <bits/stdc++.h>
using namespace std;

int main() {
  ios::sync_with_stdio(false);
  cin.tie(0);
  int tt;
  cin >> tt;
  while (tt--) {
    int n;
    cin >> n;
    vector<int> a(n);
    for (int i = 0; i < n; i++) {
      cin >> a[i];
    }
    vector<int> b(n + 1);
    for (int i = 0; i < n; i++) {
      b[i + 1] = b[i] ^ a[i];
    }
    map<int, int> last;
    vector<int> dp(n + 1);
    last[0] = 0;
    dp[0] = 0;
    for (int i = 1; i <= n; i++) {
      dp[i] = dp[i - 1];
      if (last.find(b[i]) != last.end()) {
        dp[i] = max(dp[i], dp[last[b[i]]] + 1);
      }
      last[b[i]] = i;
    }
    cout << n - dp[n] << '\n';
  }
  return 0;
}
