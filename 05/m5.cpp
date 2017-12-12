#include <stdio.h>
#include <iostream>
#include <fstream>
#include <vector>
#include <algorithm>
#include <string>

using namespace std;

int main(int argc, char *argv[]){
    ifstream f;
    f.open("input.txt");

    vector<int> v;
    int temp;
    while(f >> temp){
        v.push_back(temp);
    }

    int position = 0;
    int lastP = 0;
    int steps = 0;

    cout << v.size() << endl;
    while(position >= 0 && position < v.size()){
        lastP = position;
        position += v[position];
        //v[lastP]++;
        if(v[lastP] >= 3){
            v[lastP]--;
        } else {
            v[lastP]++;
        }
        steps++;
    }
    cout << "Steps: " << steps;

    f.close();
    return 0;
}