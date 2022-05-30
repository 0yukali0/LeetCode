# Missing Number

## Problem description
Given an array ”nums” containing n distinct numbers in the range ”[0, n]”, return the only number in the range that is missing from the array.
* An input in this problem is a array, which is “nums”.
* Elements in this array isn’t sorted.

### Input example
Input: [0,1,3,2,5]
Output: 4

## Language type
- [ ] C
- [x] C++
- [ ] Java
- [ ] Python
- [ ] Go
- [ ] R

## How to solve this problem?
The order of the elements in the input is not important.
1. Create a array, which is bool type and has n + 1 elements, called "visited".
2. Initialize "visited" array with "false" value.
3. We mark a number that we can find in the inputs.
4. Check which element in "visited" is "false" and return this number.
