# First Unique Character in a String

## Problem description
Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.

### Input example
"leetcode" -> 0
"loveleetcode" -> 2
"aabb" -> -1

## Language type
- [ ] C
- [x] C++
- [ ] Java
- [ ] Python
- [ ] Go
- [ ] R

## How to solve this problem?
* Create a queue,called "order", which record the order of the characters in a word.
* Create a map, called "count", which count the visiting time of a character.
* Visit each character in the word and execute the related operations in "order" and "count".
* Sequencely Pop the elements in queue.
    * If the count of popped element is the only one, jump out the loop poping the elements to return the index.
