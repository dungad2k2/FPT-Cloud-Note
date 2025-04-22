#include<bits/stdc++.h>
using namespace std;
struct TreeNode{
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL){}
};
TreeNode *root;
TreeNode *search(int num){
    TreeNode *cur = root;
    while(cur != NULL){
        if (cur->val < num){
            cur = cur->right;
        }
        else if (cur->val > num){
            cur = cur->left;
        }
        else{
            break;
        }
    }
    return cur;
}
void insert(int num){
    if (root == NULL){
        root = new TreeNode(num);
        return;
    }
    TreeNode *cur = root, *pre = NULL;
    while(cur != NULL){
        if (cur->val == num){
            return;
        }
        pre = cur;
        if (cur->val < num){
            cur = cur->right;
        }
        else{
            cur = cur->left;
        }
    }
    TreeNode *node = new TreeNode(num);
    if (pre->val < num){
        pre->right = node;
    }
    else {
        pre->left = node;
    }
}

void remove(int num){
    if (root == NULL){
        return;
    }
    TreeNode *cur = root, *pre = NULL;
    while(cur != NULL){
        if (cur->val == num){
            break;
        }
        pre = cur;
        if (cur->val < num){
            cur = cur->right;
        }
        else{
            cur = cur->left;
        }
    }
    if (cur == NULL){
        return;
    }
    if (cur->left == NULL || cur->right == NULL){
        TreeNode *child = cur->left != NULL ? cur->left : cur->right;
        if (cur != root){
            if (pre->left == cur){
                pre->left = child;
            }
            else{
                pre->right = child;
            } 
        } else{
            root = child;
        }
        delete cur;
    }else{
        TreeNode *tmp = cur->right;
        while(tmp->left != NULL){
            tmp = tmp->left;
        }
        int tmpVal = tmp->val;
        remove(tmp->val);
        cur->val = tmpVal;
    }
}
int main(){
     
}