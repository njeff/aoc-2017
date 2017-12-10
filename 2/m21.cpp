#include <stdio.h>
#include <iostream>
#include <fstream>
#include <string>

using namespace std;

int main(int argc, char *argv[]){
    string in;
    ifstream f;
    f.open("input.txt");

    int checksum = 0;
    int max = -1E7;
    int min = 1E7;
    int v = 0;
    while(!f.eof()){
        if(f.peek()=='\n'){
            checksum += max - min;
            max = -1E7;
            min = 1E7;
        }
        f >> v;
        if(v > max){
            max = v;
        }
        if(v < min){
            min = v;
        }
    }
    cout << checksum;

    f.close();
    return 0;
}