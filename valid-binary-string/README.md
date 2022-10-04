
# Valid Binary String

- Problem from HackerRank ([link](https://www.hackerrank.com/contests/hack-the-interview-iv/challenges/good-binary-string))

You have been given a binary string containing only the characters "0" and "1". A binary string is considered valid if each of its substrings of at least a certain length contains at least one "1" character. Given the binary string, and the minimum length of substring, determine how many characters of the string need to be changed in order to make the binary string valid.

Note that string _v_ is a substring of string _w_ if it has a non-zero length and can be read starting from some position in string . For example, string "010" has six substrings: "0", "1", "0", "01", "10", "010".

For example, let's say the binary string _s_ = "0001" of length _n = 4_, and the minimum substring length is _d = 2_. This is currently not valid because one of its substrings, "000", of length 3 has no "1"s in it. By changing the second character to a "1", the string becomes "0101". In this case, all of its substrings of length 2 or more has at least one "1" character: "01", "10", and "01". Because this required 1 change, the answer is 1.

## Input Format

The first line contains the binary string, _s_.

The second line contains the integer _d_.

## Constraints

- _1 <== n <= 10<sup>6p</sup>
- _1 <= d <= n_

## Output Format

Print a single integer denoting the minimum number of changes required to make the binary string valid.

## Sample Input 0

```plain
00100
2
```

## Sample Output 0

```plain
2
```

## Explanation 0

We can make the string "10101" by changing the first and the last characters. Here, every substring with a length of at least 2 contains at least one "1". There's no way to do this in less than 2 operations.

Therefore, the answer is 2.

## Sample Input 1

```plain
101
2
```

Sample Output 1

```plain
0
```

## Explanation 1

The given string is already a valid string, so there's no need to perform any operations on the string.

Therefore, the answer is 0.
