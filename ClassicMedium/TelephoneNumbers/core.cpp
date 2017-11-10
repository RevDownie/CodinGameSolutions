#include <iostream>
#include <string>
#include <vector>
#include <algorithm>
#include <memory>

using namespace std;

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/
 
struct Node
{
    Node(int in_num) : m_num(in_num){}
    int m_num;
    std::vector<std::unique_ptr<Node>> m_children;
};

Node* GetNodeForNumber(const std::vector<std::unique_ptr<Node>>& in_nodes, int in_num)
{
    for(const auto& node : in_nodes)
    {
        if(node->m_num == in_num)
        {
            return node.get();
        }
    }
    
    return nullptr;
}

void AddNumToTree(int in_level, const string& in_num, std::vector<std::unique_ptr<Node>>& in_nodes)
{
    if(in_level == in_num.length())
        return;
        
    auto num = in_num[in_level] - '0';
    auto node = GetNodeForNumber(in_nodes, num);
    if(node == nullptr)
    {
        in_nodes.push_back(std::unique_ptr<Node>(new Node(num)));
        node = in_nodes.back().get();
    }
    
    AddNumToTree(in_level + 1, in_num, node->m_children);
}

int GetNumNodes(const std::vector<std::unique_ptr<Node>>& in_nodes)
{
    int counter = 0;
    
    for(const auto& node : in_nodes)
    {
        counter++;
        counter += GetNumNodes(node->m_children);
    }
    
    return counter;
}

int main()
{
    std::vector<std::unique_ptr<Node>> trees;
    
    int N;
    cin >> N; cin.ignore();
    for (int i = 0; i < N; i++) 
    {
        string telephone;
        cin >> telephone; cin.ignore();
        cerr << telephone << endl;
        
        AddNumToTree(0, telephone, trees);
    }

    cout << GetNumNodes(trees) << endl; // The number of elements (referencing a number) stored in the structure.
}