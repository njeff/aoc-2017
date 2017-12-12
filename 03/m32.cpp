#include <stdio.h>
#include <iostream>
#include <iomanip>
#include <fstream>
#include <vector>
#include <string>
#include <cmath>

using namespace std;

int sumadj(int x, int y, int v[][20]){
    int sum = 0;
    for(int i = x-1; i<=x+1; i++){
        for(int j = y-1; j<=y+1; j++){
            sum += v[j][i];
        }
    }
    return sum;
}

int main(int argc, char *argv[]){
    int input = 289326;
    //the number of rings out from the center
    int ring = (sqrt(input-1)+1)/2;
    //cout << ring << "\n";
    float lowerbound = (ring*2-1)*(ring*2-1);
    float upperbound = (ring*2+1)*(ring*2+1);
    
    int v[20][20] = {0};
    int center = (20)/2;
    v[center][center] = 1;
    int walls = 0;
    int x = center+1;
    int y = center+1;
    for(int i = 2; i< 150; ){
        int r2 = (sqrt(i-1)+1)/2;
        int dist = r2*2;
        for(int j = 0; j<dist; j++){
            y--;
            v[y][x] = sumadj(x,y,v);
            if(v[y][x] > input){
                cout << v[y][x];
                return 0;
            }
            i++;
        }
        for(int j = 0; j<dist; j++){
            x--;
            v[y][x] = sumadj(x,y,v);
            if(v[y][x] > input){
                cout << v[y][x];
                return 0;
            }
            i++;
        }
        for(int j = 0; j<dist; j++){
            y++;
            v[y][x] = sumadj(x,y,v);
            if(v[y][x] > input){
                cout << v[y][x];
                return 0;
            }
            i++;
        }
        for(int j = 0; j<dist; j++){
            x++;
            v[y][x] = sumadj(x,y,v);
            if(v[y][x] > input){
                cout << v[y][x];
                return 0;
            }
            i++;
        }
        x++;
        y++;
    }
    int show = 3;
    for(int i = center - show; i <= center + show; i++){
        for(int j = center - show; j <= center + show; j++){
            cout << setw(5) << v[i][j] << " ";
        }
        cout << "\n";
    }
    return 0;
}