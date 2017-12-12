#include <stdio.h>
#include <iostream>
#include <fstream>
#include <string>

using namespace std;

int main(int argc, char *argv[]){
    string in;
    ifstream f;
    f.open("input.txt");

    getline(f,in);
    int total = 0;
    for(int i = 0; i<in.length(); i++){
        if(in[i] == in[(i+in.length()/2)%in.length()]){
            total += in[i] - '0';
        }
    }
    cout << total;

    f.close();
    return 0;
}