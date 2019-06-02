// Problem link:
// https://www.hackerearth.com/practice/basic-programming/bit-manipulation/basics-of-bit-manipulation/practice-problems/algorithm/mystery-30/

#include<bits/stdc++.h>
using namespace std;

long long countBits(long long n){
    int answer=0;
    
    while(n){
        n=n&(n-1);
        answer++;
    }
    
    return answer;
}

int main(){
    ios_base::sync_with_stdio(false);
    cin.tie(NULL);

    long long number;
    
    while(cin>>number){
        cout<<countBits(number)<<"\n";
    }
}
