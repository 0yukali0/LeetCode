#include <queue>
#include <map>
class Solution {
public:
    int firstUniqChar(string s) {
        queue<char> order;
        map<char, int> count;
        for (int index = 0; index < s.length(); index++) {
            char c = s[index];
            (count.find(c) == count.end())? count[c] = 1: count[c]++;
            order.push(c);
        }
        
        int result = -1;
        int len = order.size();
        for (int index = 0; index < len; index++) {
            char c = order.front();
            if (count[c] == 1) {
                result = index;
                break;
            }
            order.pop();
        }
        
        return result;
    }
};