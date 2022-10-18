#include <bits/stdc++.h>     
using namespace std;
#define ll long long
void solve(){
    ll n; cin>>n;
    vector<ll> dp(n+5,0);
    ll ans=0;
    for(ll i=1;i<=n;i++){
        ll x; cin>>x;
        dp[i]=min(dp[i-1]+1,x);
        ans+=dp[i];
    }
    cout<<ans<<"\n";
}
int main()                                                                                
{  
    ios_base::sync_with_stdio(false);                         
    cin.tie(NULL);  
    ll t; cin>>t;
    while(t--){
        solve();
    }
}  
