package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := `(3 * (4 * 8) * 5 * 7 * 3) + 8
4 * (6 * 4 * 6) + 8 + 4
(5 + 7) * 2 * 4 + 6 + 7 + (2 + 2 + 6 + (9 + 3 + 7 * 3))
(4 * 8 + 2 * (7 + 6 + 7) * 2) + ((4 * 2 + 4) + (6 + 7 * 2 * 3) + 5 * 2) * (4 + (4 + 5 + 4) * 7 * 8 * 4 + (7 * 7 + 5 * 3 * 6)) + 8
(3 + 3 + (6 * 4 * 9 * 3 * 6) + (5 * 8 + 7 * 9) * 8 * 7) + 9 + 3 * 8 * 4 + (7 + 4 * 7)
7 + (6 * (6 * 8 * 2) * (8 + 4 + 7 + 5 + 4 * 6) * 5) + 8 * (2 * (9 + 3 * 2) + 3 * 9 * 9 + 7) * 8 + 3
3 + (2 + 8 * 2 + 3 + 6) * 8 * (3 + 7 + 8 + (2 + 9 * 4 * 6) * 3) * 5 + 7
6 * 2 * 7 * (8 + (9 * 4 * 7 + 7 * 4) * 9) + 9
3 + 4 + (8 * 4) + 3 + ((7 * 9 + 8 + 5 * 5) + 8) + (2 + 7)
9 * 3 + ((8 + 2) * 5 + 2) * 6
((5 + 9 * 4 + 9 * 4 + 3) + 4) + ((3 * 7 * 5 + 7) + (8 + 6) + (6 + 9 * 4 + 7)) + 9 + 5 * 3 * 5
(6 * 7 + 5) + 9 * 9 + 4 * 7 * (4 + 9 + 9 * 2 * (4 * 8 * 5 * 7))
((5 + 7 + 7 * 8) + 3 + 4 * (2 + 2 + 9 + 3 * 7 * 9)) + 2 * 9 + 2 * 4
9 * 2 + (3 * 3 + 5 + 7) * (2 * 3 * 9 * 8) * 2 + 3
5 + 5 + (7 + 9 * (9 + 9 * 2) + 4)
(8 + (4 * 7 + 2 * 8) * 3 * 8 * 3) * (7 + 4 + 4 + 8 * 5 * 9) * 2 + 4
3 * 3 * ((2 * 3 * 8 + 4 + 3 * 3) + 4 * 7 + 3 * 5) * 9
3 * (8 * 5 + 9)
9 + 8 + (8 + 5 * 2 * (6 + 8) * 4)
6 + 8 + 3 * 5 * 2
7 * 9 + ((8 + 8 * 5 * 9) * 3 * 3 + 7)
5 + 9 + (3 + 6 * 4)
(6 + 2 + 7 + 6 * 8) + (3 * 6 + 7 + 2 * 4 * 6)
(8 * (4 * 2 * 2 * 5 + 6) * 8 * 8) + (9 * 4 + 9 * 2 + 9 * 5) + 7 * 2 * ((6 + 6 + 7 + 6 * 6) + 7)
(6 * 3 * 4 * (5 * 5 * 9 + 5 + 9)) + 2 + (5 + (8 + 9 * 5 * 3) * 2 * 3 * 8) + 7
8 * (6 + 6)
7 + 9 + (8 * (7 + 3 * 8 + 3 * 8 * 5) * 8 * 7)
(8 + 6 + 6 + 9 + 3) * ((2 * 6 + 8 + 7 * 2) + 9 * 7 * 7) + 5
((5 + 8) + (4 + 5 * 3 + 2) + (4 * 4 * 7 + 8 + 6) * 2 * 9) + 7 * 3 * 7 + 9 + 5
4 * 2 * (3 * 7) * 5 + 8
8 + (4 * 7 * 6 + 8 * (2 + 4 + 5 + 2)) + 5 * 7 * 9 * 3
(5 + 6 * 4) + 4 * 9
(3 * 2 * 2) * (5 * (4 + 7 + 7)) * 3
8 + (5 * 4 * 3 + (4 * 5 + 6))
3 * 5 + 3 + 3 + 6 * 2
(6 + 9 * 4 + (5 * 9 * 6) * 7) + 5
7 * 6 + 7 * (2 + 7)
4 * (2 + 6 * 9 + 7 + 7) + 6 + 8 * (8 * 7 + 7)
5 * 9 + (6 * 4 * 5) + 4 + (4 + 6 + 4 * 4 * 9) + 9
(5 + (4 * 8 + 6) * 7) + 6 + (3 * 6) * (3 * 7 * 3 * 7 * 3 * 9) * (4 + 8 * 5 + 7 * (8 + 2 * 9 + 4 + 8)) + 4
6 + 6 * (3 + (4 * 8 * 2 * 7) * 2 + 9)
2 + 2 + (9 + 8 * 7 + 5) + 7
2 * 8 + ((5 + 7 + 5) * 3 + 9) + 7
3 * (7 + 4 * 7 * 3 * 3 * 8) + 9 + 2 + 4
8 + 4 + 5 * 2 + 3
((6 * 2 + 2) * 9 + 7) + 9 * 2 * (7 * 6 + 4 * 3 * 3) + (8 + 8 + (6 + 5 * 8 * 6) + 6 + (9 * 9 + 4 + 2 + 6) + 2)
9 + 9 + 6 + 9 * (8 + 4 * 7 + 2 * 9 + 6)
(4 + 8) * 6 + 5
(5 * 3 * 2) + 8
7 + 7 + (6 * 9) + (3 + 4 + (9 + 3 * 6) + 5 * (4 * 6 * 6) + 9) + (6 + 8 + 7) * (3 + 3)
(2 * 8) + 3 + (4 * (7 * 9 * 9 + 6) * (7 + 2 * 5 * 6) * 3 * 7 + 2) * 7
3 * 4 * 6 * 2 + 9 + (8 + (4 * 9 + 8 * 9) * 6)
3 + 7 + (3 * 6 + 2 * 6 + 6 * 4) * (3 * (7 + 4) * (4 * 9 + 7 * 9) + 5 * 8 + 9)
4 + (7 * 8 * 8 * 7 * 5) + 7 * 6 * 7
(5 + 3 * 6 * 8 * 9 + 4) * 9 + (2 + 4 + 4 * 6 + 4 * 7) + (7 + (9 + 8 + 3)) * 6 + 4
(7 * 8 + 8 * 7 + 5 + 4) * 2 + (6 + 4 + 7) + (8 * 6)
9 * 8 + 8 * 3
9 + (7 + (8 * 5 * 5) * 7 + 6 + 2) * 7 + 7 + (3 + 2 + 5 + (3 + 4) + 9) + 2
9 * 7 + 3 * 9 * 6 * (3 + (2 * 6 + 5 * 3 + 7) * 4 * (8 * 8 * 2))
((4 * 6 + 9) * 6 + (4 + 5 + 8) + 8) + 9 * 6 * 8 + 3
6 + 4 * (2 * 3 * (7 + 8 * 6) + 8 + 4)
(8 * 8 + 4 * 8 + 7) + 2 * 4 * 3 * 7
(3 * (7 * 4 * 6 * 5 + 7 * 7)) * 7 + 6
7 * 3 * 3 * 2 + (2 + (3 + 7 * 8 + 4 * 4) + (4 + 3 + 5) * (5 + 2 * 3 * 3 + 9) + 6 + (4 * 3 + 4 + 6))
5 * (4 * 7 + 6) * 8
9 + 7 + 7 + 2 * 2
8 * (9 * 2 * 2 * 3) + (9 + (9 + 3 + 6 + 2 * 7 + 8)) + 6 + 4 + 7
7 * 2 + ((6 * 7 * 8) * 6 * 7) + 6 + (9 * 8 + 7) + (3 * 3 * (9 + 8) + 8 * 6 * 2)
(9 * (2 + 5 * 2) + 3 * 7) * (8 * 3 + 2) + 3 + (6 * 9 + 9 * 8 + 8 * 2) + ((2 + 5 + 3 + 8 + 6) + 4 * 6)
(8 * (9 + 4) * 4 + 8 * 7 + 3) + 2 * (3 + 2 * 2) * (6 + 8) + (5 * 9 * 7 + 7 * 2) + 4
6 * 8 * 7 + ((6 + 6 + 8 * 4 + 8) * 3) + 4
2 * 8 * (7 + 7)
(9 + 8) + 3 + (5 * 7 + 9 + 6 + 6)
2 * 6 + 9 * 6 * 3
3 + (5 + 2 + 7 * 3 * 9) * 8 * 6 * 4 + 3
(3 * 9 + 2 * 7) + 6 + (9 + (9 + 9 + 4 * 4)) + (7 * 7 + 8 + 4 * 4)
7 * 2 + 7 * ((8 + 9 + 7 + 4) * 7 + (7 + 4 + 7 * 3 * 4 + 2) + 8)
(9 + 3 * 9) + (4 * 9 + 7 + 2) + 7 + (7 * (3 + 4 + 8 * 9) + 8) + 9
8 + (4 + 3 + 4 + 9 + 7) * 7 + 6
8 * (9 + (5 + 8)) + 8
4 * (7 + 8 * (8 * 5 + 9 + 3)) * 4
2 + 5 + 7 * (7 * 6 + 6 * 5) * (8 * (7 + 3 + 6 + 4 * 8) + (8 + 6) * (4 * 7 + 2)) * 9
4 * 2 + 7 + (2 * 2 + 8 + 2 * (7 + 8 + 9)) * 5
(6 + (9 + 9 * 2 + 4 + 2)) * 6
(2 + 4 * 8 + 3) * 9 + 7 * 6 + 3
2 * 6 * ((8 * 3) * 8) * 4 * 5
4 + 5 * (5 + (8 + 8 * 2) * 8 * 5) + 3
9 * 7 * 7 + 9 * (7 * 8 + (7 * 4 * 4 * 2 + 5 * 3) + 5 + 5 + (3 + 2 * 4 + 8 * 9)) + 9
2 * ((5 * 8 + 8 + 5 + 6 * 3) * 3 * (3 * 6) * 2) + 8 * 9
9 + (6 + (4 + 3) + 5 + 2 + (6 + 4 * 8 + 7 * 2) * (8 * 6 * 8 + 6 * 4 * 5)) + 6
6 + (9 * 9 + 7 + 3) * 7 * (3 * 5 * 5 * 8) * 8
6 * 7 * 4 * 9 * 6 + (7 + 3 + 9)
4 * 3 + 7
8 * 4 + (6 + 8) + 8
9 * 7 + ((7 * 6 * 7 * 6 * 3 * 5) * 7 * 6 + 3 * 2) + ((4 * 5 * 6) * 5 * 7 + (2 * 6 * 2)) * 8 + 3
7 * 6 + 4 * (4 * 6 + 6 * 4 * 8)
7 + 8 + 3 + 3
6 + 8 + (8 * (3 * 2 * 3 + 5 * 8 + 3) + 7) * ((4 * 4 * 7 * 7 * 9 + 5) * 6 + 7) * 2 * (2 + 3)
7 + 7 * 2 * (4 * 3 + (4 * 2)) + 6 + 4
(5 * 6 + 3) * 8 + (2 + 8)
(8 + 8 * 3 + 9 + 5 * 9) + 4 * 9
9 + 4 * 8 + 5 + 6 + 4
4 * 4 + 5 + 3 + ((4 + 3) * 7 * (3 + 9))
(6 * 9 + 3 + 4 * 8 * 2) * (6 * 5) * 5 * 9 * 5 + 3
9 * 2 + 3 + ((3 * 5 * 5 + 8) + 7 * 8 * (2 + 3 + 9) * (6 + 2 + 8)) + 3 + 8
(3 + 9 * 6 + 4) + (7 * (9 * 6 * 2 * 4 + 8) * 4 * 4 + 9 * 3) + 7 * 6 * (7 * 6 + 2 * 2 + (5 + 8 * 4 + 3 * 7) + 3)
6 * (6 * 5 * 7 + (3 + 5)) + 9 + (3 * 2 + 6) + 4
8 * 9 + 8 + (5 + 8) * 8 + (2 + 7 + 7 * 8 * 3 * 5)
9 * 8 * 4 * 3
6 * (3 + 5 * 9 + (3 * 5 * 4 + 6 + 6 + 3) * 8)
(7 + 8 + 7 + 4 * 5) * 8 * (7 * 8 + 9 * 2)
9 * 4 * (2 + 2 + (8 * 4 + 8 + 9 * 4 + 4) * 6 + 4 + 5) * (3 * 4 * 5) * 6
5 * 6
5 * 3 * 4 + (9 * (2 + 9 * 8 + 8 + 3)) + 4 + 7
8 + 9 + (7 * (6 * 9 * 2 + 8) * 7 * (2 * 3) * 6) + 3 + (8 + 4 + (7 + 2 * 2 + 2 + 3) + 3 + (3 * 3 * 9 + 7 * 5 * 3) * 4) + 2
8 * 8 * (8 + (3 + 7 + 9 + 2) + 4)
(2 + 7 + 8 * 8 * 7 * 6) + 2 + 7 * 7
5 + (4 * 9) * 7
7 + 5 + 4
(3 * 3 * (3 * 6 + 9)) * (6 + 6 * 4 + 7 + 6)
(6 + 7 + (3 + 8 * 9 * 7) * 8) + 3 * 6 * 9 * 3 + (7 + 3 * (9 * 3 * 5 + 5))
3 * (7 * 2 + 8 + (2 * 4 + 4 * 3)) * 7 * (4 * 8)
(3 * 9 + 6 * 9 * 7) + 7 * 4 * 4 + ((5 * 7 + 8 + 4 * 8 * 8) + 6 * 3 * 8 + 5 + 4)
4 + (2 + 7 * 4 * (5 + 3 * 4 * 3 + 3 + 6)) * 6
3 + (7 * 4 + (5 + 7 + 2 + 8) * (4 * 5) * 6) + ((4 + 5) + (7 * 4) + (4 + 3) * (5 * 6)) + 5
(7 * 9 + (7 * 3) + 5 + (7 + 3 + 3 * 9 + 2 * 6)) + 2 * 6 * 9 + 3 + 4
((8 + 5 + 3 * 5 * 5) * 4 + (6 * 5 + 8 + 6 + 4 + 5) + 2 * 6 * 6) * (2 + 2 + 3) * 7 + 8 * (7 + 8 + 7 * 6 + 7 * (2 + 5 * 2 + 3)) * 4
3 + 2 + 7 + (9 * 7 * 3 * (9 * 2)) * 2
5 + 6 + 7 * 4 * ((6 + 8 * 2 + 6 + 3 + 3) + (3 + 3 + 8) * 5 * 5 * 9 + 5)
3 + 5 * 2 + (3 + (6 + 5 * 4 + 2) * 2 + 4 + 4 + 2) * 2
3 + 5
4 + (9 * (9 + 7 + 4 * 3 * 3) * 8 + 6 + (7 + 8)) + 3 * 6 * 5 * 6
(2 * 4 * 6 + 6) + (2 + 6 * 3 + 4 * 7 * 6) * 8 * 6
5 + (3 * 9 * 9 + 4 + (4 + 6) + (5 + 9 + 4)) + 9 * 3 * 6
((6 + 3) * 8 + 7 * 4) + 3
5 + (5 * (5 + 8 + 2 + 5 + 7 * 8) + (5 * 4 * 5 + 9) + (6 * 3) + (7 * 7)) * 5
((4 + 4 * 8 + 6 * 8) * (9 + 6 * 9) * (6 * 9)) * 9 + 9 * 4
5 * 6
8 + (5 * 5 + 4 * 5 * 9 * 2)
8 + 5 + (5 + 6 + 4 * 6)
3 * 5 * 6 + 9 + (4 + 2 * 2 * 2 + 4)
8 * 6 + 4 + (5 * (4 * 9) + 2 * (8 + 3 * 2) + 2) * 7
7 * (8 * 6) + 7 * (2 * (4 + 5 * 5 + 6) + (9 + 4 * 4 + 6 * 7) * 3 * 9 * 4) * (3 * (3 * 8 + 8 + 7) + (9 * 2 + 8 + 8 * 2)) * 7
9 + 3 * 9 * (6 + 5 * 6 * 9 + 8) + (9 * 4 * 9 * 5 + 9) + 2
4 + 5 + 7
5 * 5 + ((2 + 5 + 4 + 3 + 4) * 8 + 4 * 3) * 5 * 9 * (8 * 2)
(6 * (2 * 2 + 7 + 7)) * 2 + (9 + 2 * 2 + 2 + (7 * 3 * 5 * 9 + 2 + 7)) + 9 + 7 * 2
(9 + 2 + 6) + (6 * 6 * (5 * 6 + 9 * 6 + 8) + 8 + 7 + 2) * 4 * ((2 + 2 * 5 + 7) * 9)
6 + 9 + (8 * 3) + (7 + 6 + 6 * 3 * 2 + 9) * 7
6 + (3 * (4 * 6 + 3 * 4 + 7 + 8))
3 + (3 + 2) + 9 * (9 + (2 * 3 * 2 * 6 + 9) + 8) + (5 * 7) + (7 + 4 * 9 * 5 + 4)
2 * 5 * (2 + 8 * (4 * 9)) * 8 + (5 + 7 * (5 + 8 * 9 * 9 * 8) * 4 * 8 * 5)
2 * 8 + (2 + (5 + 3) + (2 * 3 + 3 + 6 * 4))
6 + 3 + 9 + 5 + 5 + ((3 + 3 + 4 * 8) * 4)
5 * ((3 + 6 + 7 * 6 * 8) * 3 + 8 + 7) + 2 + (2 * 6 + (5 + 5) * 9) + 3
9 * ((5 + 6 * 4 + 2) * 5 * 7 + (6 * 9 + 2 + 7)) + 4 + 4 * 2
6 * ((3 * 5 + 6 * 5 + 8 * 3) * 4 * (6 + 7 + 8 + 5 + 3) + 8 + 4 * (6 * 5 + 7 * 9 * 5)) + 5 * 5 * 7 + 4
(7 + 2) + 9 * 8 + 3 * 8 * 7
2 * 5 + 3 + (2 + 9) * 3
(5 * 6 + 9 + 5 + 9 + 3) * (3 + 5 + 6 * 2 + 2) * 5 * 6 * 2
(6 + 8 * 3 * 9 + 2 + 2) + 2 + 2 * 6 * 3
8 + (2 * 4 * (7 * 9) + 4) * 4 * 8 * 4 * 5
3 * (9 + (6 * 4 * 9 * 3 * 5) + 9 + (7 * 3 + 4) * 2 * 8) * 5 + 3 + 3 * (4 * 5)
3 * (2 + 2 + 2 + (9 + 7 + 6 * 3)) + 7
8 * ((3 * 5 * 8 + 9) * (4 * 6 * 6 + 6 + 7) + 6 * 7 + 8 + (3 + 7 * 8 + 6 + 4)) + 4 * 2 * (5 * 2)
3 * 2 * 4 * 7 * 9 * ((2 * 7 + 3 * 4 * 3 * 4) * 9 * 3 * 5)
(9 * 5) * ((7 + 4 + 4 + 5 * 4) * 5 + 8 + 8 * 2 * 2) * 2
3 + (4 + 7 * 2 * (9 + 9 * 9 * 5 + 7) * 5) * 5
(3 + (7 * 5 + 6 * 4) * 4 * 3 * 3 + 7) * 9 * 7
(2 + (3 * 3 * 6 + 4) * 6 * 8) + 2 + 9
9 * 4 + 4 + 7 * (6 + 6 * (9 * 2 + 6 + 5) + (3 + 5 * 9)) * 6
((9 + 6 + 7) * 2 + 2 * (6 + 9) * 8) * 7 * 6 + 6
3 * 5 + (5 * 2 + 8 + 5) + (8 + 9 * 9 + 5)
6 * 2 + (6 * 2 * 6 + (8 * 8 + 7 + 7))
3 * 6 + 3 * (4 + 9 * 7 + (8 + 6 * 3 + 4) * 7 * 9)
8 + (4 * 9 + 7 + (9 + 9 * 4 + 9))
8 * ((7 * 6 * 4 + 4) + 3)
6 + 7 * 6 + (6 + 6 + 4 * (2 * 2) + (2 + 7 + 2)) * 3
6 + 8
(6 + 5 + 6 * (3 * 2) * (5 * 7 * 3) * 5) + (6 * 3 + 8) * 6 * 2 + 6 + 6
3 * (7 + (6 + 8 + 9 * 4 + 5 + 4) * 4) * ((9 * 9 * 8 * 5 * 7 + 5) + 7 * 5) * 8 + 8 * 6
(8 * (3 * 5 * 9 + 3 + 4 * 2) * 5 + 6 * 7) * 5 + (6 + (5 + 3 * 3 + 3 * 2) * 4 * 4)
5 * (6 * (7 + 9)) + (9 + 9 + 2) + (2 * 2 + (2 * 9 * 6 + 8 * 2) + 4) * 6 + 2
((4 + 8) * 4 + 2 + 6 * 7 + (7 * 4 * 8 * 6 + 3)) * 8 * 2 + 5 + 9
5 + (5 * 9 * 4 * 3) * (7 * 3 * 7) * 6
(7 * 5 + 2 + 7) * (2 + 3 + 5 + 6 * (4 + 4 + 9)) + 9 + 2 * 5
((5 * 9 * 6 + 5 * 9) + 2 + 7 + 9 * (4 + 6 + 8)) + ((8 * 9) + 4 * 3 + 7 * 7 * 2)
9 + 9 + ((7 + 8 + 6 * 3) * 7)
((6 * 2 + 3 + 8 * 5) * 4 * (5 + 7 * 5)) + 4 + 5 + 6 * 2 * (6 + 2)
8 * ((4 * 8 * 3 * 4 + 5) + 3 + 6 + 9 * 9 + 3) + (7 * (8 * 2) * 9) * (8 * 9 + 2 + 7) + (4 * (2 * 8) + 3 + 2 + (7 * 2 * 6))
2 + 8 * ((2 * 2 + 3 * 6 + 5) + 8) + 6
(2 * (5 + 7 * 3 * 9 * 6) * (4 + 9) * (8 + 2 * 9)) + 9 * 2
2 * 4 * 4 * 9 + 7
4 + 7 + 6 + (6 + (8 + 8 + 7 + 5 * 2) * 7 * 2) + 8
3 + 4 + (4 * (9 + 4 + 5)) + 7 * 8 * 7
9 + 8 * 7 * 4 * ((8 + 2 * 4 + 6 * 9 * 2) + 3 * 3 + (2 * 5 * 5 + 2 + 9 * 9)) * 7
(8 * (5 * 7 + 7 + 2 + 7) + 3 * 9) + 9 * ((6 * 2 + 8 * 8 + 9 + 7) * 2 + (3 * 3) + (7 + 4) + 9 * 5) + (4 * 8 * 4) + 2 + 9
(3 * 5 + (7 + 4 * 8 + 7)) + ((8 + 5 * 6 * 5 + 3) * 2 * (3 * 9 + 4 * 2 + 3))
5 * 6 * 3 + (8 + 6 * 6 + 6)
6 * (5 * 7) + 2 * (3 + 7 * 5)
(9 + 6 + 2 * 7) * 2 * 7 + 6
5 + 3 * ((8 * 5 + 8 + 4) * 7 * 8 * 7 * (2 * 6 * 4 + 5 * 9)) + 3 + 4
3 + 9 + 6 + (4 + 3) * 2 * 7
((6 + 7 + 8 * 3) * 6 + 9) + 4 * 9 * 8
9 * 2 + (6 + 9 * 5) * (8 + 5 * 9 + 9 * 7 * 9) * 6
9 + 8 * 5 * (4 + (9 * 4 + 5 * 8 + 5) * 7 * (7 * 3 * 4 * 7) + 2) + 5
2 + 6 * (6 + 5 * 7 * 2 * 2) + 5
3 + (5 + 5) + 8 * (8 * 4)
(3 * 2) + 5 + 4 * 6 + 3 * (3 + 8)
8 + (2 * (5 * 7 + 4 + 7) * 4 + 5) + (8 * 9 * (7 + 2) + 5) + (6 + 7 * (6 * 8 * 8 + 9) + 6 * 8)
(2 + (8 * 3 + 3 + 5 * 8 * 7) + (7 * 6 + 4) + 8) + (2 + 7 * 9) + 9 * 8 * (4 + 5)
8 + (6 * 7 * 8 + (5 + 3 * 9 * 7) * 2) + (2 + 3 + 5 * 9) * 8 * 6 * 5
8 * 3 * 8 + (5 * (6 * 6) * 3 + 6) + ((8 * 5 + 8 * 4 * 7 + 5) * 6 * 7 + 9 * 6 + 8)
3 + 3 * 4 + (4 * 3 * 2) + 2 + 5
6 + 4 + 3 + (2 + 3 * (2 + 6 + 9 * 8 + 4) * 5 + (2 * 5 + 5 * 6 + 6) + 5)
3 + 2 * 4 + 6 + 4 * (6 + 2 + 9 * 2 + 8 + 4)
(6 * 9 + 2 * 7) + 7
(8 + 5 * 2) + 7 * (5 * (8 + 5 * 8 + 6 * 7) + (3 * 9 * 9 + 9 * 6) * 5 * 7 * 9) + 6 * 2 * (9 * 9 + (5 + 7 * 4 * 4 * 9 + 6) + (6 * 3 + 7 + 6 + 4 + 8) + 8 * 5)
6 * 7 * (4 * 9 * (2 * 5 * 3 * 3 * 8 * 9) * 7) + 9 * (2 + 5) + 4
8 * 8 + 7 + (7 + 9 * 7)
((3 * 5 * 3) * (6 * 8 * 5 + 8) * 7 + 6) + 7 * 9 + (5 * 3 + 2 * 5 * 8 + 6) + (5 * 7)
7 + 7 * (2 * 9 + 6) * 4 * 4 + 9
(5 * 8 + (8 * 9 + 3 * 2 * 7 * 9)) * 8 * 8 + 3 + 7
6 + 5 + 8 * 2 * (7 * 6 * (5 * 8 * 4)) + 6
(9 + 4 * 2 * (6 + 2 * 2 * 4 + 7 * 9)) + 8
(5 * 3 * 2) * 2 + 2 * 6 * 8 * 8
4 * (2 * (5 + 6 + 6) + 6 + (7 * 2 * 9 * 9) + 4 + (4 * 9 * 7))
8 + (8 * 6 + 2) * 3 * 6 * 3 + 7
3 + 8 * ((6 + 6) * 9 * (9 * 4) + 4 * 4 + 5) + (3 * (7 + 3 * 4 * 6 * 4) * 8 + 7 * (6 + 7 + 7 * 5 + 9)) + (6 + 6 * 9 + (6 + 6 * 2 * 4 * 5))
8 + 4 * 7 * ((9 + 5 + 9 + 7) + 5) + 7
8 * 7 + (8 * 9 + 3) + 7
9 * (9 * 2 * (3 + 3 + 7))
((9 + 9 * 8) * (4 + 8 + 7 * 6 * 5 + 3) + 5 + 3 * 4 + (7 + 7 * 3 * 6 * 3)) * 6 + 5 * 9
2 + 7 * ((3 * 6) * (6 * 5) + 4 + 9 + (4 + 2 + 3 * 3 * 5))
(9 * (7 + 7) + 5 * 2 + 5) + 6 * 3 * 3 * 8
(7 + (8 * 6 + 7) * 7 + (2 * 7 + 3) * 7 * 9) * 7 + 5 * ((7 * 5 + 2 + 5 * 2) + 6 + (7 + 9 * 3 + 8 * 6))
(9 * 2 + 5 * 3 * 8 + 4) * ((4 + 9 + 8) + 6 * 7 * 3 * 8 + 3) * 6 + 5 + 4 * ((9 + 4 + 6 * 4) + (2 + 7) * 5 * 5 * 3 + 5)
((9 * 6 * 3 * 6) * 3 * 9 + (6 + 7 + 5) * 6) * (7 * (4 * 9 + 7 + 8 + 6)) * 7 * 6 + (2 * 4 + 9 * 5 * 2)
7 * (8 + 5 + 4 + 7)
((5 + 8 * 2 * 8 + 6) * 8) + 8 + 2 * 2
(2 * 5 * 5 + (5 * 5 + 5 * 9 + 2)) * 9 * 2 + (6 * 7 + 5)
(9 * 5 * 9) + 4 * 8 + 2 * (5 * 7 * 4 + (8 * 4 + 3 + 8) + 9)
7 + 3 * (5 + (5 * 4 * 4 + 9 + 9) * (6 * 8 + 3 * 8) * 5)
7 * 6 + 7 + 3
(5 + 6) + 5 * (6 + 7)
9 * 2 * 9 + (9 + 6 * 4 + 4 * (5 + 4 + 8 * 2 * 9 * 3))
7 + (8 + 2 * 6 + 9 * (2 * 6 * 3 + 4 + 2 + 6) + 9) + 3 + (9 * 8 + (9 + 3 * 6 * 5) + 6 * 8 * 2) * (2 * 4 + 5) + ((2 + 7 + 2 * 6) * 3 + 7 * 6)
7 + 9 * ((3 * 2) * (7 + 5 + 8)) + 8 * 5
2 + 3 * 6 + 8 + 2 * (9 * 9 + 8 + (6 + 3 + 4 * 4 + 4 * 6))
4 * 2 + 6 * ((6 * 9 * 5 * 7 + 8 + 7) * 5 + (4 + 2 + 6) * 9 + 3 * 3)
8 * (2 + 3 * 5) + 7 + 8 * (9 + 2 * 8 + 4 * 9) + 8
(9 * 2 + 6) + 3 * ((5 + 6 * 3) + 4) + 6
6 * 3 * (3 + (4 * 4 + 5 * 7) * 4 + (4 + 5 + 2 + 8 + 4 + 3)) * 9 + 7
8 * 3 * 6 + 6 + 7 + 2
(6 + 4 * (6 * 2 * 9 * 4) * 2) + 9 + 4 * 7 + 8 * 4
5 + (7 * (4 + 5 + 9 * 5) * 2 * 5)
(4 * (6 * 3 * 6) + 7 + (4 + 6 + 8 + 3 + 3 + 9)) + 6 + (3 * 6)
3 * 8 * (7 * 8 + 2 + 4 + 5 + 4) * 8 + 9 + 3
3 + 4 + 7 + 6
(6 * (9 * 5 + 2 + 8 + 7 * 6) + 4) * 3 * 2 + 3
8 * 6 + (6 + 2 + 7 + (7 * 5)) * 6 * (8 * 4 * (4 * 5 + 4) + 2)
7 * 5 + 2
(6 + 3) + 2 * (3 * 6 + 5 + 6) * 9
6 * (2 * (8 + 8 * 6 * 5 + 8 + 3) + 2 * 9 + 4 + 4) + 5 + 6
5 + 2 + (4 + 2) * (4 * (8 + 3) * 9) * 2
9 + 2 * 4 * 5 + 6 + (9 * 8 + 2 * (8 * 4) * 7 + 8)
2 * 2 * 2 * 6 + 6 * 8
6 + (4 + 7 * 4 * 5 + 2 + 4) * 2 * 8
(5 * 6 + (9 * 7 + 2 * 2 * 7) * 8 * 3 + 5) + 5
(9 + (6 * 8 + 9)) + 7 * 7 * ((8 + 7 + 6 * 9 + 4) * 3 + 6 * (4 + 5) * 2) + (6 * 5 + 9 * 2 + 5 + (7 + 5 * 5)) + 9
((5 + 8 + 7 * 8 * 6) + 9 + 2) + 9 + 6 + (6 + 2 + 4)
(8 + 5 + 2 * (8 * 6 * 2 * 3 + 7 * 5) + 8 + 8) * 7 * 7 + 2
6 * ((7 * 8 + 6 * 8) * 2 + 5 + (3 + 8)) * 2 + 9
4 + 3 + 8 * (6 * 7 * 8 * 6 * 3 + 2) * 9 + 6
3 * ((9 + 7 + 4 * 7) + 3 * 3) + 8
2 + ((5 * 8 * 6 + 4 * 5 + 6) + 2 + 3 * 9 * 7) + 6 * 7
3 * 3 * (5 + 8 * 4 + (3 * 3 + 6) + 2) * 6 + ((8 + 3 + 7 + 5 + 7 * 6) * (8 * 9 + 8 + 2) * 8 + 2) * 3
(6 + 3 * 9 * 9 * 4) + 7 * 6 + 3 + 6 + 8
6 * 2 + (7 + 9 + 6 + 4 + 9)
(6 + 8 + 8 + (6 + 4 * 2 * 2 * 4) + 7) * 8 * 7
3 * 7 + 9
(6 + 7 * (2 * 3 * 6 * 8 + 6) + 8 + 3) * 4 + (4 * 9 + (5 + 8 * 4 + 6 + 6) * 6 * 5)
9 * ((7 * 2) + 9 * 2 * 5) * (6 * 4 * 3 + (8 + 3 + 7 * 9 + 8) + 4) + (2 + 4 + 4 + 9)
3 + (9 * 6) * 4 * 8
6 * (3 + (4 * 7 * 7 * 4 + 7) * 5 + (9 + 5) * (9 + 9 + 5 + 7 + 2 + 2))
7 + ((6 + 2 + 3 * 9) + 4 + 5 + 2 + 8) * (8 + 9 + 9 + 9 * 6) + 9
7 * (6 + (7 + 5 * 3 + 4 * 9 + 3) + 2)
6 * (7 * 7 + (4 * 3 + 2 + 7 + 5 + 6) + 4 + 7) + 4 * 6
(2 * 2 * 9 + (9 * 7 + 8)) + 6 + (9 + 7) + 8 + 4 + (7 * 5 + 8 + 8 * (2 + 5 + 7 + 5 * 9))
2 * (6 * (8 * 8 * 2 * 3 + 8 * 9) + 3) + 3 * 5 + 7 + (7 + 2 * 2 * 4)
6 * (8 * (4 * 3 + 3 * 4) * (7 + 4 + 5 * 5 + 2) * (6 * 4 * 6 + 9 * 9) + 2 * 5)
8 + (4 + 5 + 6 + (4 + 7) + 3 * 8) + 3 + 4
((9 * 7 + 2 * 6 + 7) * 2 * 8) * 7 + 2 + (6 + (2 * 6) + 7 * (5 * 4 * 6 * 7 + 9) + 8 + 2) * (8 + 2 + 5 * (6 * 8 + 5 * 4) * 7 + 4)
(5 + 2 + (8 * 2 + 2 * 9) + (4 * 8) * 8 + 9) * 9 + 9 * 2 * (2 + (3 * 3) * (5 * 2 * 4 * 2) * 3 + 7) + 8
4 * (4 + (8 + 6) + 9 * 8) * 4 * 6 + (5 * 3 * 3 * 5 * (8 * 2) + 8)
8 + 5 + 2 + (9 * 7 * (4 * 5 + 6 * 5 * 6 + 8)) + 8 + 7
(8 + 9 * 4) + 8 * ((5 * 6 + 9 * 2 * 4 + 8) + 6 * 7 * 4 + 4 * 2) + 8 * (5 + (2 * 7 + 6 * 4 * 8) + 2) + 2
(7 + 3 * 8 * 2) * (6 * 5 * (2 * 3 * 5 * 7 + 8 * 8) + 7 + 7 * (2 * 3)) + 9
7 + 4 + (3 + 2 + 4) * 4 * 4
((3 * 2) + 3) * (2 * 6) + 2 + 6 * 6 + (3 + (2 * 5 * 3 * 2 * 9 + 2) * 3 * 8 + 3)
9 * ((6 + 2) + 3 + (9 + 7 + 9)) * 8 + 8 + 5 * 3
(8 + 4) * 6 + 2 * (6 * 6 + 4 + 9 + 9 + 6) * (2 * 7) * 9
(2 + 9 + 9 * 6 * 2 * 5) * 2 + 2 + 2 * (7 * 3 + 6 * 7 + 5 * 7)
4 + 8 * 7 * (6 * 7 + (9 * 2 * 8 + 8) + 4 + 5 + 2) * 5 * 7
5 + (7 * 6 + (3 * 2 * 7 * 5) + 4 * 6 + 2) * 2
7 * (4 * 6 * 2 * 3 + (6 * 8 * 4) * 6) + 3 + 6
4 + 3 * ((3 + 3 * 5 + 4 * 3 * 2) * (3 * 6 + 7 * 7 + 9) * 7) + (8 + 2 + 9 + 5 * 8 * 8) + (7 + 8 + 9 + 4 + 2 * 4) + 7
6 * (5 + 2 + 6 + (7 + 9 + 4 * 3) + 8) + 6 * 5 * 8 * 7
7 + 9 * (2 * 9) * 8 + (2 * 2 * 8 * 4 * (7 * 3 * 4 * 2 * 9 * 3) * 5) * (9 + 2 + 3 + 3 + 4)
6 * 3 * (9 * 6 + 7 + 9 * 2 + 2)
5 * ((3 + 5 + 4 + 7) + (5 + 7) * (2 * 5 + 5 + 9 + 7)) + (9 + 7) * 8
(7 + 7 * 5 + 2 + 7 + 3) * 9
9 + (3 * 7 * 4) + 8 * 2 + 2
(3 + 7 + 3 * 6 * (2 + 2) * 8) + 3
3 * (9 + (8 + 4) + (3 * 7 * 6 + 9 * 7 + 7) * (8 * 7 * 9 + 6) * 9 * 7) + 4
5 * 2 * 4 + (4 * 5 + 6 + (6 * 8 + 4 + 2) * 5) + (7 + 2) + ((6 + 7) + 5 * 5 * 5 + 5 * 6)
7 * (7 * 3 * (8 * 2 * 8 * 2 + 4) * 8) * (6 + (4 * 7 * 6 * 8 + 2) + 2 + (2 + 6) * 4) + 7 * 6 + 4
8 * 3 + ((5 + 4 + 4 + 2) * (6 * 8 * 6) * (5 * 7 * 2) + 4) + 8 + 5 * 2
5 * 5 * (6 * 2 * 5 + (8 + 9 * 5 + 5) + 5)
(4 * 5 * 2) + 5 * 9 + 3 * 9
7 * 8 + ((3 + 9 * 5 + 4) + 9 + (3 + 7 * 8 + 7) + (2 + 9 * 3) * 5) + 6 * 6
(4 * (8 + 4) + (8 + 2 + 9 + 3 + 2 * 9) + 4 * 7 * 6) + 9
7 * (5 + 5 * 2 * 3) + 4
(8 + (6 * 9 * 7 + 4 * 3) * 6 * 9) * 3 * ((2 * 4 * 5 + 2 * 7) + 4 + (8 * 9 * 9 + 9 + 6 + 7) * 3 + 9 + 3)
(3 + (7 + 5 + 4 + 9)) * 4 * 5 * 8 + 6 + 6
(6 * 9 * 8) * (2 + 9 + (6 + 2 * 9 * 4) * (3 * 6 + 5 * 6 * 9) * 9 + 8) + 3 * 3
((4 * 4 + 4) * 9 * 9 * (3 * 5 + 5 * 4 * 5) + 8) + 8 * 7 * 8
4 + 4 * (2 + 5 * 7 + 2 * 7 + 3) + 5 * 2
6 * (3 * (6 + 8 * 7 + 5 * 3 * 4) * 8 * (4 * 3)) * 7 + 6 + 3
2 + 4 + 5 + 6 + 9 + ((3 * 7 + 4 + 8 + 6) + 5 * 2 + 4 * 6)
(8 + 5 * 8 * 7 * 9) + 9 * 5 + 3
5 + 3 * 4 + (6 + 2 * (6 * 2 + 6 + 3) * 2 * 6 + (9 * 9 * 3)) + 2 + 6
3 * 5 * (5 + (2 + 5) * (6 + 8) + 2 + 4 * 5) * 7
4 + 7 + 4 + 9 + 2 * (3 + (7 + 2 + 4 * 9 * 3) * 6 * 7)
((9 * 9) * 5) + ((6 * 3 + 6 * 6 + 2 + 3) + (9 + 7 + 8) + 5 * (3 + 4)) * (8 + 4 + 9)
3 + ((5 * 6 + 2 * 3 + 9) * (6 * 8)) * 7 + 6 * 8 + 8
8 * 5 * 5 * 9 * ((5 * 6) * 6 * (2 * 8 * 9 * 3) + 7 * (9 * 3) * 3) * 9
2 * (5 + 2 + (4 + 6 + 7 + 9 + 7 + 4)) * 5 + 9 * 9
(5 + 6) + 7 + 8 * (4 + 3 * 3 + 5 * (5 * 3 + 8))
4 * ((6 * 7 + 8 * 3) + 8 * (8 * 2 * 7 * 8 * 2 + 7) * 4 + 9) + 7 + 3 + 9 * 5
7 + (5 * 7 * 4 + (9 + 7 + 2 + 5 * 4 * 2)) + 5 + 2
(9 + 7 * 8) + 7 * 3 * 3
3 * 6 + 9 * 6 + 8 + 5
2 * 3 * 8 + ((6 * 6 + 2) + 2)
5 + 3 + 7 + (2 + (4 + 6 * 7 + 7 + 5 * 7) + 4)
4 * 6
(7 * 6) * 6 * 4 + 2
6 + ((7 + 9 * 3 + 9) * 2 * (6 * 4 * 5) + 3) + 5 + 9
(5 + 3 + (4 + 8 * 6 + 6 + 6) + 4 * 6) * (3 + 6 * (9 + 3 + 6) + 7 + 4 * 3) * 3 * (2 + 5 * 2 + (6 + 5 + 8 * 9 * 8 * 3) + 2 * 6)
(6 * 5 * 5 * 6) + 3 * (5 + 8) * (4 * 7 + 8 * 6 + (4 + 9 + 6) + 4) + 4
3 + (2 * 6 * 5 + 9 * 3) + 6
9 * (3 * 4 * 8 * 4 + 4)
(9 * (2 + 9 * 6 + 5 * 6 + 2) * (8 * 4 * 6 * 8) * 5) * 7 + (7 + 5 + 3 + 5 + 7 * 9) + (6 * 7) + 8
(4 + 3 * 7) + (4 + 5 * 9 + (8 * 6 * 3) + (4 + 2 * 2) + 3) + 4 + 7 * (9 * 9 + 7 + (8 * 4) * 4 * 6)
((3 * 2 + 2) * 4 + 3 * 2 + 9) * 4 + 7
(7 + 2 + 6 + 7) + 8 + 4 + (3 + 3 + 9) + 8
(7 * 4 * (9 * 2 + 4 + 6)) + 6 * 5 + 9 * 7
7 + 3 + 3 + 8 + (4 * 5 * 9 * 9 * (8 * 4 + 9 + 2 + 4 * 8))
4 * 3 + 7 * 7 + (4 * 8 * 9) + 2
(4 + 2 + 8 * 9 + 5) * 5 + 8 * (7 * 5 * 5) * 9 * 9
9 * (2 + 5 + (2 * 5 * 5) * 6 * 8 * (8 * 9 + 6 * 2 * 4)) + 5 + 3
4 + 2 + 2 + 3 + (7 * 5 + 3 + 7 + (6 + 6 * 7))
(9 * 3 + 6 + 5) + (3 + 3 + 5 + 9)
3 * ((8 * 6) * 7 * 4) + 9 + 6 + 9
2 + (6 * 7 + 7 + 6 * (7 * 6 * 2 + 4 + 6 + 7))
(4 * 6 * 7) * 6 * (3 * 7) + 2
(3 * 7) * ((9 * 3 * 5) + 9 + 8 + 4 + 4 * 7) + 3 + 8 + 6 + 2
(5 * 5) + (3 * 4 + 8 * (6 + 8 * 5 + 7 * 3)) + 8 * (9 * 2)
(3 + (2 * 7 + 9) * 2) * 3 * 7 + 2 + 4 * 9
9 * 9 + 4 * 3 * (8 * 4 * (2 * 9 + 6 + 6 + 2))
7 * ((5 * 7 + 5 * 4 + 9 + 9) * 4 + 4 + 6 + 2) + 6 * (7 * 5 + 5 + (2 + 9 * 4 + 7) + (8 * 7 + 4 + 5 + 5) + 2) * 7
(8 + (4 * 3 + 9 * 2 + 2 * 7) + 7 * 2 * 5) * 3
5 * 2 * 9 * (5 + 5) + (6 + 3) * (2 + 3 * (5 + 5 + 6) + (9 + 2 + 7) + 5)
(9 * 8 + 6 + 9) * 8 * 8
(4 * 4 + 7 + 2) + 7 * 2 + (4 + 3)
2 * 8 + 5 + (2 + (3 * 9 * 8 * 8 + 2 * 5) * 8 * 2 * 2 + 4) + ((3 * 7 + 7 * 5 * 5 + 5) * 8 * 5 * 8 + 6) + 5
(2 * (9 + 3 + 8 * 8) + 6) * 3 * 9 * (7 + 7) * (2 * 2) * 5
((6 * 3 + 3) * 6 + 5) + 3 + 4`

	expressions := strings.Split(input, "\n")
	sum := 0
	for _, expression := range expressions {
		sum += Calculate(expression)
	}
	fmt.Println(sum)
}

func Calculate(expression string) int {
	expression = removeWhitespaces(expression)
	stack := Stack{}
	for {
		token := stack.top()
		if token == nil || (token.name == "operand" && stack.tokensInBracket() < 3) || token.name == "operator" || (token.name == "bracket" && token.value == "(") {
			// shift
			token, expression = getFirstToken(expression)
			if token == nil {
				break
			}
			stack.push(token)
		} else if (token.name == "operand" && stack.tokensInBracket() >= 3) || (token.name == "bracket" && token.value == ")") {
			// reduce
			token = stack.pop()
			if token.name == "bracket" && token.value == ")" {
				tokenInsideBracket := stack.pop()
				stack.pop() // left bracket
				stack.push(tokenInsideBracket)
			} else if token.name == "operand" {
				operator := stack.pop()
				operand1 := stack.top()
				operand1.value = doOperation(operator.value.(string), operand1.value.(int), token.value.(int))
			}
		}
	}
	return stack.top().value.(int)
}

func doOperation(operator string, a, b int) int {
	switch operator {
	case "+":
		return a + b
	case "*":
		return a * b
	}
	panic(fmt.Sprintf("Unknown operator %s", operator))
}

func removeWhitespaces(expression string) string {
	return strings.ReplaceAll(expression, " ", "")
}

func getFirstToken(expression string) (*Token, string) {
	if len(expression) <= 0 {
		return nil, ""
	}
	token := new(Token)
	if expression[0] == '+' || expression[0] == '*' {
		token.name = "operator"
		token.value = string(expression[0])
		expression = expression[1:]
	} else if expression[0] == '(' || expression[0] == ')' {
		token.name = "bracket"
		token.value = string(expression[0])
		expression = expression[1:]
	} else {
		regex := regexp.MustCompile(`\d+`)
		operandRaw := regex.FindString(expression)
		operand, _ := strconv.Atoi(operandRaw)
		token.name = "operand"
		token.value = operand
		expression = expression[len(operandRaw):]
	}
	return token, expression
}

type Token struct {
	name  string
	value interface{}
}
