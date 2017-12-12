#include <stdio.h>
#include <iostream>
#include <fstream>
#include <vector>
#include <set>
#include <algorithm>
#include <string>

using namespace std;

int main(int argc, char *argv[]){
    string in;
    ifstream f;
    f.open("input.txt");

    vector<string> v;
    set<string> s;
    int valid = 0;
    while(!f.eof()){
        if(f.peek()=='\n'){
            if(s.size() == v.size()){
                valid++;
            }
            s.clear();
            v.clear();
        }
        string temp;
        f >> temp;
        sort(temp.begin(), temp.end()); //comment out for pt1
        v.push_back(temp);
        s.insert(temp);
    }
    cout << valid;
    f.close();
    return 0;
}