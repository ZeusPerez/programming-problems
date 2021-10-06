#include <vector>
#include <iostream>
#include <string>
using namespace std;

string rotate(string str){
    string result = str;
    char last = str.back();
    result.insert(result.begin(), last);
    result.pop_back();
    return result;
}

bool check_periodic(string str, int k){
    for(int i = 0; i < str.size() - k; i=i+k){
        string prev_str = str.substr (i,k);
        string curr_str = str.substr (i+k,k);
        if (rotate(prev_str) != curr_str){
            return false;
        }
    }
    return true;
}

int main() {
    string str;


    getline(cin, str);
    int prev_index = 0;
    for (int i = 1; i <= str.size()/2; i++){
        if (str[prev_index] == str[i]){
            if (check_periodic(str, prev_index + 1)){
                cout << prev_index + 1 << endl;
                return 0;
            }
        }
        prev_index = i;
    }
    cout << str.size() << endl;
    return 0;
}
