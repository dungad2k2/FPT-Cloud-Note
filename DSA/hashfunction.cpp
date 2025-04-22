#include<bits/stdc++.h>
using namespace std;
const int MODULUS = 1000000007;
int addHash(string key){
    long long hash = 0;
    for (unsigned char c : key){
        hash = (hash + (int)c)  % MODULUS;
    }
    return (int)hash;
}
int mulHash(string key){
    long long hash = 0;
    for (unsigned char c : key){
        hash = (31 * hash + (int)c) % MODULUS;
    }
    return (int)hash;
}
int xorHash(string key){
    int hash = 0;
    for (unsigned char c : key){
        hash ^= (int)c;
    }
    return hash & MODULUS;
}
int rotHash(string key){
    long long hash = 0;
    for (unsigned char c : key){
        hash = ((hash << 4) ^ (hash >> 20) ^(int(c))) % MODULUS;
    }
    return (int)hash;
}