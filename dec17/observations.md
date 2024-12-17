Register A: 22571680
Register B: 0
Register C: 0

Program: 2,4,1,3,7,5,0,3,4,3,1,5,5,5,3,0

#

0 2 bst 4 -> B = A & 0x7
1 1 bxl 3 -> B = B ^ 0x3
2 7 cdv 5 -> C = A >> B
3 0 adv 3 -> A = A >> 3
4 4 bxc 3 -> B = B ^ C
5 1 bxl 5 -> B = B ^ 5
6 5 out 5 -> out += B & 0x7
7 3 jnz 0 -> jump 0

First digit:
| # | inst | inst | op | A | B | C |
| - | ---- | ---- | -- | - | - | - |
|0|2|bst|4| x | x&7 | 0 |
|1|1|bxl|3| x | (x&7)^3 | 0 |
|2|7|cdv|5| x |  (x&7)^3 | x >> ((x&7)^3) |
|3|0|adv|3| x>>3 |  (x&7)^3 | x >> ((x&7)^3) |
|4|4|bxc|3| x>>3 |  ((x&7)^3) ^ (x >> ((x&7)^3)) | x >> ((x&7)^3) |
|5|1|bxl|5| x>>3 |  (((x&7)^3) ^ (x >> ((x&7)^3))) ^5 | x >> ((x&7)^3) |
|6|5|out|5| -> B & 0x7 |
|7|3|jnz|0| jump to zero if A is not zero | ||

So we consume 3 bits of A for each digit.
Last time we will break the loop if A == 0.