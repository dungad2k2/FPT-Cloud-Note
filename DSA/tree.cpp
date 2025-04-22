#include<bits/stdc++.h>
using namespace std;
struct TreeNode{
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL){}
};
vector<int> levelOrder(TreeNode *root){
    queue<TreeNode *> queue;
    queue.push(root);
    vector<int> vec;
    while(!queue.empty()){
        TreeNode *node = queue.front();
        queue.pop();
        vec.push_back(node->val);
        if(node->left != NULL){
            queue.push(node->left);
        }
        if(node->right != NULL){
            queue.push(node->right);
        }
    }
    return vec;
}
void preorder(TreeNode *root){
    if(root == NULL){
        return;
    }
    cout << root->val << " ";
    preorder(root->left);
    preorder(root->right);
}
void inorder(TreeNode *root){
    if(root == NULL){
        return;
    }
    inorder(root->left);
    cout << root->val << " ";
    inorder(root->right);
}
void postorder(TreeNode *root){
    if(root == NULL){
        return;
    }
    postorder(root->left);
    postorder(root->right);
    cout<< root->val << " ";
}

int main(){
    TreeNode  *n1 = new TreeNode(1);
    TreeNode *n2 = new TreeNode(2);
    TreeNode* n3 = new TreeNode(3);
    TreeNode* n4 = new TreeNode(4);
    TreeNode* n5 = new TreeNode(5);
    n1->left = n2;
    n1->right = n3;
    n2->left = n4;
    n2->right = n5;
    vector<int> result = levelOrder(n1);
    for (auto x : result){
        cout << x << endl;
    }
    preorder(n1);
    inorder(n1);
}