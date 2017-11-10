#include <iostream>
#include <string>
#include <vector>
#include <algorithm>

using namespace std;

/**
 * The code below will read all the game information for you.
 * On each game turn, information will be available on the standard input, you will be sent:
 * -> the total number of visible enemies
 * -> for each enemy, its name and distance from you
 * The system will wait for you to write an enemy name on the standard output.
 * Once you have designated a target:
 * -> the cannon will shoot
 * -> the enemies will move
 * -> new info will be available for you to read on the standard input.
 **/
int main()
{
    struct Enemy
    {
        std::string m_name;
        int m_distance;
    };
    std::vector<Enemy> enemies;
    
    // game loop
    while (1) {
        int count; // The number of current enemy ships within range
        cin >> count; cin.ignore();
        
        enemies.clear();
        enemies.reserve(count);
        
        for (int i = 0; i < count; i++) {
            string name; // The name of this enemy
            int dist; // The distance to your cannon of this enemy
            cin >> name >> dist; cin.ignore();
            
            Enemy enemy;
            enemy.m_name = name;
            enemy.m_distance = dist;
            enemies.push_back(enemy);
        }

        // Write an action using cout. DON'T FORGET THE "<< endl"
        // To debug: cerr << "Debug messages..." << endl;
        string closestName;
        int closestDistance = 999999;
        
        for(int i=0; i<enemies.size(); ++i)
        {
            if(enemies[i].m_distance < closestDistance)
            {
                closestDistance = enemies[i].m_distance;
                closestName = enemies[i].m_name;
            }
        }

        cout << closestName << endl; // The name of the most threatening enemy (HotDroid is just one example)
    }
}