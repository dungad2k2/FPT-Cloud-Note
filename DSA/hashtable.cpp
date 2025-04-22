#include<bits/stdc++.h>
using namespace std;
struct Pair{
    public:
        int key;
        string val;
        Pair(int key, string val){
            this->key = key;
            this->val = val;
        }
};

class ArrayHashMap{
    private:
        vector<Pair *> buckets;
    public:
        ArrayHashMap(){
            buckets = vector<Pair *>(100);
        }
        ~ArrayHashMap(){
            for (const auto &bucket : buckets){
                delete bucket;
            }
            buckets.clear();
        }
        int hashFunc(int key){
            int index = key % 100;
            return index;
        }
        string get(int key){
            int index = hashFunc(key);
            Pair *pair = buckets[index];
            if (pair == NULL){
                return "";
            }
            return pair->val;
        }
        void put(int key, string val){
            Pair *pair = new Pair(key, val);
            int index = hashFunc(key);
            buckets[index] = pair;
        }
        void remove(int key){
            int index = hashFunc(key);
            delete buckets[index];
            buckets[index] = NULL;
        }
        vector<Pair *> pairSet(){
            vector<Pair *> pairSet;
            for (Pair *pair : buckets){
                if (pair != NULL){
                    pairSet.push_back(pair);
                }
            }
            return pairSet;
        }
        vector<int> keySet(){
            vector<int> keySet;
            for (Pair *pair : buckets){
                if (pair != NULL){
                    keySet.push_back(pair->key);
                }
            }
            return keySet;
        }
        vector<string> valueSet(){
            vector<string> valueSet;
            for(Pair *pair : buckets){
                if (pair != NULL){
                    valueSet.push_back(pair->val);
                }
            }
            return valueSet;
        }
        void print(){
            for (Pair *kv : pairSet()){
                cout << kv->key << "->" << kv->val << endl;
            }
        }
};
class HashMapChaining{
    private:
        int size;
        int capacity; //hash table capacity
        double loadThres; // load factor threshold for triggering expansion
        int extendRatio; // Expansion multiplier
        vector<vector<Pair *>> buckets;
    public:
    HashMapChaining() : size(0), capacity(4), loadThres(2.0 / 3.0), extendRatio(2){
        buckets.resize(capacity);
    }
    ~HashMapChaining(){
        for(auto &bucket : buckets){
            for (Pair *pair : bucket){
                delete pair;
            }
        }
    }
    int hashFunction(int key){
        return key % capacity;
    }
    double loadFactor(){
        return (double)size / (double)capacity;
    }
    string get(int key){
        int index = hashFunction(key);
        for (Pair *pair: buckets[index]){
            if (pair->key == key){
                return pair->val;
            }
        }
        return "";
    }
    void put(int key, string val){
        if (loadFactor() > loadThres){
            extend();
        }
        int index = hashFunction(key);
        for (Pair *pair : buckets[index]){
            if (pair->key == key){
                pair->val = val;
                return;
            }
        }
        buckets[index].push_back(new Pair(key, val));
        size++;
    }

    void remove(int key){
        int index = hashFunction(key);
        auto &bucket = buckets[index];
        for (int i = 0; i < bucket.size(); i++){
            if (bucket[i]->key == key){
                Pair *tmp =  bucket[i];
                bucket.erase(bucket.begin() + i);
                delete tmp;
                size --;
                return;
            }
        }
    }
    void extend(){
        vector<vector<Pair *>> bucketsTmp = buckets;
        capacity *= extendRatio;
        buckets.clear();
        buckets.resize(capacity);
        size = 0;
        for (auto &bucket : bucketsTmp){
            for (Pair *pair : bucket){
                put(pair->key, pair->val);
                delete pair; 
            }
        }
    }
    void print(){
        for (auto &bucket : buckets){
            cout << "[";
            for (Pair *pair : bucket){
                cout << pair->key << " -> " << pair->val << ", ";
            }
            cout << "]\n";
        }
    }
};
class HashMapOpenAddressing {
    private:
        int size;
        int capacity = 4;
        const double loadThres = 2.0 / 3.0;
        const int extendRatio = 2;
        vector<Pair *> buckets;
        Pair *TOMBSTONE = new Pair(-1, "-1");
    public:
        HashMapOpenAddressing() : size(0), buckets(capacity, NULL){

        }
        ~HashMapOpenAddressing(){
            for (Pair *pair : buckets){
                if (pair != NULL && pair != TOMBSTONE){
                    delete pair;
                }
            }
            delete TOMBSTONE;
        }
        int hashFunc(int key){
            return key % capacity;
        }
        double loadFactor(){
            return (double)size / capacity;
        }
        int findBucket(int key){
            int index = hashFunc(key);
            int firstTombstone = -1; 
            while (buckets[index] != NULL){
                if (buckets[index]->key == key){
                    if (firstTombstone != -1){
                        buckets[firstTombstone] = buckets[index];
                        buckets[index] = TOMBSTONE;
                        return firstTombstone;
                    }
                    return index;
                }
                if (firstTombstone == -1 && buckets[index] == TOMBSTONE) {
                    firstTombstone = index;
                }
                index = (index + 1) % capacity;
            }
            return firstTombstone == -1 ? index : firstTombstone;
        }
        string get(int key){
            int index = findBucket(key);
            if (buckets[index] != NULL && buckets[index] != TOMBSTONE){
                return buckets[index]->val;
            }
            return "";
        }
        void put(int key, string val){
            if (loadFactor() > loadThres){
                extend();
            }
            int index = findBucket(key);
            if (buckets[index] != NULL && buckets[index] != TOMBSTONE){
                buckets[index]->val = val;
                return;
            }
            buckets[index] = new Pair(key, val);
            size++;
        }
        void remove(int key){
            int index = findBucket(key);
            if (buckets[index] != NULL && buckets[index] != TOMBSTONE){
                delete buckets[index];
                buckets[index] = TOMBSTONE;
                size--;
            }
        }
        void extend(){
            vector<Pair *> bucketsTmp = buckets;
            capacity *= extendRatio;
            buckets = vector<Pair *>(capacity, NULL);
            size=0;
            for (Pair *pair : bucketsTmp){
                if (pair != NULL && pair != TOMBSTONE){
                    put(pair->key, pair->val);
                    delete pair;
                }
            }
        }
        void print(){
            for(Pair *pair : buckets){
                if (pair == NULL){
                    cout << "NULLPTR" << endl;
                }else if (pair == TOMBSTONE){
                    cout << "TOMBSTONE" << endl;
                }else{
                    cout << pair->key << " -> " << pair->val << endl;
                }
            }
        }
};
int main(){
    ArrayHashMap test ;
    test.put(1000, "hi");
    test.put(1002, "aaa");
    cout << test.get(100000) << endl;
    test.print();
}