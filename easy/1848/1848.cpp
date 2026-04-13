#include <iostream>
using namespace std;

/*
 *   Stats:
 *   Runtime: 0ms (beats 100.00% of C++ solutions)
 *   Memory: 13.37mb (beats 16.70% of C++ solutions)
*/


class Solution {
    public:
        int getMinDistance(vector<int>& nums, int target, int start) {
            int idx = nums.size();
            for (int i=0; i < nums.size(); i++) {
                if (nums[i] == target) {
                    idx = std::min(idx, std::abs(i-start));
                }
            }
            return idx;
        }
};