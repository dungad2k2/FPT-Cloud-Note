#include<bits/stdc++.h>
using namespace std; 
struct linkedList{
    int val;
    linkedList *next;
    linkedList(int x) : val(x), next(nullptr) {}
};
void insert(linkedList *n0, linkedList *P){
    linkedList *n1 = n0->next;
    P->next=n1;
    n0->next=P;
}

void remove(linkedList *n0, linkedList *P){
    if (n0->next == nullptr){
        return;
    }
    // n0->next = P 
    linkedList *n1 = P->next;
    n0->next=n1;
    delete P;
}

linkedList *access(linkedList *head, int index){
    for (int i = 0; i < index; i++){
        if (head == nullptr){
            return NULL;
        }
        head = head->next;
    }
    return head;
}

int find(linkedList *head, int target){
    int index;
    while(head != NULL){
        if (head->val == target){
            return index;
        }
        head = head->next;
        index++;
    }
    return -1;
}

int main(){
    linkedList *head = new linkedList(1);
    linkedList *temp = new linkedList(2);
    insert(head, temp);
    remove(head,head->next);
    while(head != NULL){
        cout << head->val << " ";
        head = head->next;
    }
}