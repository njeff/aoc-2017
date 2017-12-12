#include <stdio.h>
#include <iostream>
#include <fstream>
#include <vector>
#include <string>

using namespace std;

int main(int argc, char *argv[]){
    string in;
    ifstream f;
    f.open("input.txt");

    int checksum = 0;
    int v = 0;
    std::vector<int> vec;
    while(!f.eof()){
        if(f.peek()=='\n'){
            int div = 0;
            for(int i = 0; i<vec.size(); i++){
                for(int j = i+1; j<vec.size(); j++){
                    int q = vec[i];
                    int w = vec[j];
                    if(q < w){
                        int temp = w;
                        w = q;
                        q = temp;
                    }
                    if(q%w == 0){
                        div = q/w;
                        break;
                    }
                }
            }
            checksum += div;
            vec.clear();
        }
        f >> v;
        vec.push_back(v);
    }
    cout << checksum;

    f.close();
    return 0;
}