#include<bits/stdc++.h>

using namespace std;

#define fore(i, l, r) for(int i = int(l); i < int(r); i++)
#define sz(a) int((a).size())

#define x first
#define y second

typedef long long li;
typedef long double ld;
typedef pair<int, int> pt;

const int INF = int(1e9);
const li INF64 = li(1e18);
const ld EPS = 1e-9;

int n;

inline bool read() {
	if(!(cin >> n))
		return false;
	return true;
}

void calc(int n, vector<pt> &ans) {
	if (n == 2)
		return;
	
	int y = max(1, (int)sqrt(n) - 1);
	while (y < (n + y - 1) / y)
		y++;
	
	fore (pos, y + 1, n)
		ans.emplace_back(pos, n);
	ans.emplace_back(n, y);
	ans.emplace_back(n, y);
	
	calc(y, ans);
}

inline void solve() {
	vector<pt> ans;
	calc(n, ans);
	
	cout << sz(ans) << endl;
	for(auto p : ans)
		cout << p.first << " " << p.second << '\n';
}

int main() {
	ios_base::sync_with_stdio(false);
	cin.tie(0), cout.tie(0);
	cout << fixed << setprecision(15);
	
	int tc; cin >> tc;
	while(tc--) {
		read();
		solve();
	}
	return 0;
}