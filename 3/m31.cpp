#include <stdio.h>
#include <iostream>
#include <fstream>
#include <vector>
#include <string>
#include <cmath>

using namespace std;

int main(int argc, char *argv[]){
    int input = 289326;
    //the number of rings out from the center
    int ring = (sqrt(input-1)+1)/2;
    cout << ring << "\n";
    float lowerbound = (ring*2-1)*(ring*2-1);
    float upperbound = (ring*2+1)*(ring*2+1);
    //cout << lowerbound << " " << upperbound << "\n";
    float position = fmod((float)(input - lowerbound)/(upperbound - lowerbound),0.25)/0.25;
    //cout << position << "\n";
    float rd = ring*abs(0.5-position)/0.5;
    //cout << rd << "\n";
    cout << ring+rd;
    return 0;
}