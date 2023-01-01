#include <bits/stdc++.h>

using namespace std;

int main() {
  ios::sync_with_stdio(false); cin.tie(0);
  int t;
  cin >> t;
  while (t--) {
    int n;
    cin >> n;
    vector<int> a(n), b(n);
    for (int& x : a) cin >> x;
    for (int& x : b) cin >> x;
    
    auto check = [&](int ans) {
      map<int, int> cnt;
      for (int x : a) ++cnt[x & ans];
      for (int x : b) --cnt[~x & ans];
      bool ok = true;
      for (auto it : cnt) ok &= it.second == 0;
      return ok;
    };
    
    int ans = 0;
    for (int bit = 29; bit >= 0; --bit) 
      if (check(ans | (1 << bit)))
        ans |= 1 << bit;
    
    cout << ans << '\n';
  }
}