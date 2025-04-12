#include<bits/stdc++.h>
using namespace std;

void insert(int *nums, int n, int index ,int target){
    for(int i = n; i > index; i--){
        nums[i] = nums[i-1];
    }
    nums[index] = target;
}
void deleteElement(int *nums, int n, int index){
    for (int i = index; i < n; i++){
        nums[i] = nums[i+1];
    }
}

int main(){
    // stored on stack
    int arr[5];
    int nums[4] = {1, 2, 5, 4};
    insert(nums, 4, 1, 3);
    deleteElement(nums, 4, 1);
    for (int i = 0; i <4; i++ ){
        cout << nums[i] << endl;
    }
    // stored on heap
    // int* arr1 = new int[5];
    // int* num1 = new int[5] {1, 2, 4, 5, 6};  
}