#include <iostream>
#include <string>
#include <vector>
#include <algorithm>

using namespace std;

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/
int main()
{
    int LX; // the X position of the light of power
    int LY; // the Y position of the light of power
    int TX; // Thor's starting X position
    int TY; // Thor's starting Y position
    cin >> LX >> LY >> TX >> TY; cin.ignore();
    
    // game loop
    while (1) {
        int E; // The level of Thor's remaining energy, representing the number of moves he can still make.
        cin >> E; cin.ignore();

        // Write an action using cout. DON'T FORGET THE "<< endl"
        // To debug: cerr << "Debug messages..." << endl;
        
        int dx = LX - TX;
        int dy = LY - TY;
    
        string dirX;
    
        if(dx > 0)
        {
            dirX = "E"; 
            TX++;
        }  
        else if(dx < 0)
        {
            dirX = "W";
            TX--;
        }
    
        string dirY;
    
        if(dy > 0)
        {
            dirY = "S"; 
            TY++;
        }
        else if(dy < 0)
        {
            dirY = "N";
            TY--;
        }
    
        string dir = dirY + dirX;

        cout << dir << endl; // A single line providing the move to be made: N NE E SE S SW W or NW
    }
}