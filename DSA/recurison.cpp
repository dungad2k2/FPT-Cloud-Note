//Tail Recursion 
#include<bits/stdc++.h>
using namespace std;
// int tailRecur(int n, int res){
//     if (n == 0)
//         return res;
//     return tailRecur(n - 1, res + n);
// }
//Recursion Tree
int fib(int n){
    if (n == 1 || n == 2){
        return n - 1;
    }
    int res = fib(n - 1) + fib(n - 2);
    return res;
}
int main(){
    cout << fib(10);
}
//Recursion Tree
//Recursion terms like 'call stack' and 'stack frame space''
//1. Calling: when a function is called, the system allocates a new stack frame on the 'call stack'
//for that function, storing local variables, parameters, return addresses and other data
//2. Returning: when a function completes execution and returns, the corresponding stack frame removed
//from the 'call stack', restoring the execution environment of the previous function.


