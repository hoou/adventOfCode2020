package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := `18: 48 48
25: 48 81 | 41 7
48: "b"
4: 131 48 | 70 41
20: 61 48 | 57 41
89: 41 41 | 41 48
74: 41 107 | 48 124
98: 41 48
99: 97 48 | 92 41
91: 34 48
100: 48 41 | 67 48
6: 48 100 | 41 132
40: 81 48 | 7 41
124: 83 48 | 130 41
50: 7 41 | 7 48
68: 64 41 | 24 48
60: 30 41 | 86 48
75: 89 41 | 39 48
103: 67 67
58: 41 22 | 48 111
71: 67 34
56: 34 48 | 39 41
122: 48 120 | 41 89
12: 41 18 | 48 98
95: 34 41 | 103 48
93: 110 41 | 34 48
13: 43 41 | 69 48
44: 101 48 | 114 41
69: 106 41 | 32 48
67: 48 | 41
45: 7 48
117: 48 120 | 41 39
46: 48 29 | 41 82
121: 48 49 | 41 47
130: 103 48 | 89 41
132: 41 48 | 48 41
94: 41 4 | 48 76
14: 9 48 | 93 41
26: 41 72 | 48 81
79: 67 1
115: 67 132
15: 41 20 | 48 63
47: 120 41 | 81 48
27: 100 41 | 7 48
11: 42 31 | 42 11 31
113: 49 41 | 56 48
31: 48 133 | 41 127
131: 90 41 | 28 48
81: 48 48 | 67 41
23: 84 41 | 27 48
84: 18 48
107: 79 41 | 33 48
83: 67 89
49: 81 41 | 7 48
108: 102 41 | 60 48
37: 41 7 | 48 120
120: 48 41
32: 41 96 | 48 95
2: 48 85 | 41 128
102: 48 62 | 41 50
5: 41 110
61: 41 37 | 48 6
97: 41 120 | 48 34
85: 120 41 | 89 48
80: 120 48 | 98 41
92: 48 103 | 41 34
65: 16 48 | 58 41
112: 48 71 | 41 123
70: 87 48 | 115 41
39: 41 67 | 48 41
41: "a"
38: 2 48 | 77 41
110: 48 41 | 41 41
88: 7 48 | 89 41
52: 41 73 | 48 104
96: 100 48 | 103 41
66: 41 126 | 48 121
77: 40 48 | 47 41
3: 48 118 | 41 25
126: 75 48 | 27 41
1: 41 41 | 48 48
19: 48 72 | 41 18
42: 41 68 | 48 105
129: 48 110 | 41 120
72: 48 48 | 41 67
7: 41 41
59: 41 47 | 48 125
73: 1 48 | 103 41
114: 41 117 | 48 73
118: 89 41 | 132 48
51: 41 1 | 48 81
101: 129 48 | 93 41
133: 48 15 | 41 13
104: 48 1 | 41 81
123: 110 48 | 98 41
0: 8 11
55: 48 1 | 41 18
30: 89 41 | 18 48
76: 48 14 | 41 35
43: 52 48 | 112 41
24: 48 109 | 41 38
29: 116 48 | 3 41
106: 88 41
9: 89 41 | 100 48
125: 41 120 | 48 132
22: 103 48 | 18 41
21: 48 83 | 41 51
64: 48 44 | 41 74
111: 103 67
54: 41 36 | 48 55
119: 48 12 | 41 19
35: 123 48 | 91 41
127: 41 94 | 48 10
116: 41 93 | 48 122
8: 42 | 42 8
128: 48 72 | 41 1
28: 41 132 | 48 120
63: 41 59 | 48 17
87: 103 48 | 72 41
36: 48 81 | 41 100
17: 62 48 | 26 41
62: 100 48 | 98 41
34: 48 41 | 48 48
82: 23 48 | 21 41
78: 41 65 | 48 66
33: 120 41
109: 48 113 | 41 99
57: 41 92 | 48 80
86: 48 89 | 41 39
53: 54 48 | 119 41
10: 108 48 | 53 41
90: 81 41 | 34 48
16: 48 5 | 41 45
105: 46 41 | 78 48

bbaaabbbbaabbaababaababbbabbbaaa
bbbbbaabbabaaaaabbaaaabbabaababbbbabbbbababaabaabaaaaaabbbaaaaab
baabbaabababbabbaaaaabba
ababbbbbbbabaaaaabaabbabbbbaaaabbabbaaba
abbbbabaabbbabbbabbaaaaaaaaaabba
bbaaaaaabaabbbababbaabbabbabaaaabbbaababaaaabbbaabababaa
bbbbbaabbabbbbabaabbbbaaaaaaabba
bbbaabbaaababbbbabbbaabb
abbbbaaabbbaabababaababa
aaaababbabaaabaaabbbbbaaaabaaabaaababbbbbaabaabababbbbbbbaabaaaa
babaababaabbbbabbabbbbabbababbaa
babbaaabababbaaabbbaabab
aaaabaabbbababbbbbbbbbbb
aabbbbbabbbbabbbbbbbbababbbabababaaaababbbaaaabaaaabaaaaabbaabbaaababaaaabaaabab
aaabbaabaaaabbbababbbbbbbabaaabaababaabb
abbaaaaabbaabaabbbabbbab
baabbbaabbabbabaabbabbbb
abababbbaaaaabbbbbbbbabaabaababbabbbbbbbbbabbabbaaabaaba
abbaaaaaaabbabbaaaababab
baabababbaaaaaaabbbbabab
babbabaaaaabaabbbbbaaaab
aabaaaababbaaabbbbbbaaaa
babaaaaabaaaababbbabbabb
abbabaaaabbababbaaabbaba
aaabaaaaaabaaaababaaaabaaaaabaabaababbabaabababa
ababbabbaaabaabababaabbaaabbabababaabbaa
baaaabbbbbabaaaabbbabaabbabbbaaa
abaabaaaababababbaabaabb
aabaabbbbabaaabbaababbaabbbbbbabbbbaabababbabbaa
abaaaaaabbaaaabaaabbaaab
aaaabbbbbabaabaaababaaabbbbbaaabbbbbabbbaaaabaaabababaaa
bbabbbbaaababbbbaaababbbbbbaabbaababbabbaaaabbbaaaaabbbaaabbabbbaaababab
babaaaaaaabaaabaabbbabab
aaaabbaaabaaabaabaabbbabaaaabbbbabbaaabbbbabbbbbbbaabbabbbbbababababaabb
abbbabbabbabbbbaaabaaaaababbaabbaabaaabb
abbaaabbbbaaababbbbaaabbabbbbbbbbbbbbaababbbbbbbabbbbabbabbabbbb
babbbaababbaabaaaabbbbbb
baaaabbbbbbbaabbbabbbababaaaabba
aaabbbaababbaaaabaaaabbbbaaabaaa
abbbbaaaaababbabbbabaaabbbaaabba
babbabaabaaaabbbbbbaaaab
abbababaabbbbaababaaabba
aababbbbbababaababbaababababbbbbbaaabbbbaaabaaab
baabbaabaaabaabaababbaaaabaaaabaaaababbbbbbabbbbbbbbabab
bbababbaaaabaaaaabbabbaa
bbababbbbababbabbbaababbbabaabbaabbaaabaabababaa
baababbabababbabaababbab
abbababbbaababababaaabaaababbbabbbbaaaab
abbbbbaaaabaaaaaabababbb
abbbbaaabaaabbbaababbbaa
bbababbbbabaabaabbbbabaaaaaaaabb
aaaabbaaaabbbbaaaaababba
abaaabaaaababbabbaabaaba
babbabababbaabaaabaaabbb
babababbbaaaabbbbabaababbbbaaaaaaabbaababbaababbaaabaaababaaaaabbbabbbab
baaabbabbaababbaaabbbaaa
bbbababbbabababbaaaabaabababaabaaabbabbb
bababbabbabbbbabaababbbbabaabbabaababaabbbaaaabaaaabbaabaaabbaaa
bbabbbaabaaaabbbbbabaaabababaababaaababa
aabbbbaabaaaaaaaaabbbaabbbbabaababbaaabbbabbabbbaaabababbaabbaba
bbbbaabbbabbabaaabbbbbaaabaaabab
baabbabababaabbbabbabbbaabaabaabaabbbbaabbabbabbabbbaaaabaaaaabaabbbaaba
ababbabbbbbaaabbbbbbaaaa
aaaaaabbababbbabbaababbb
ababbabbaaabaababaababaabaababaaababbababaaaabaa
baaaababaabaaaaabaabaaba
bbabbbbaabbbaabbbabbaaabbababaababbbbbabbbabbbbb
bbbababbbaabbbaaabbbbabb
bbbababbbbbbbaababbbaabbbababbbabbabbabb
abaaaaaaaaababbbaaaabaaa
baaaababbbababbaaaabaaab
babaaabbbbbabbaaaaababbbabbbabaaabbabbaa
bbaaaabbbaaabaaaabababbbbaabaaabbbbbbbbabaaaaabbababaaaaaaaabbabbbaaabaababbbaab
aabaaaabbbabbbaaaaaaabaa
ababaaabaabbbbabbbbbaabbaabbabbababbbbbbaabbbbbb
bababaabbbbabaabbababbba
bbbbbbabababbabbbbaaaaaabaabaaaaaababaaabbbbbbaabbbabbbaaaaaabaa
abbbbaaaababaaabaaabaabaabbbbbbbbabbbaaa
abaabaabbabbaabbbabaabaabbbabbbbabbbaaabaabbabaaabbabbaabbaaabba
abbbbbaaababbaaaabbaabaabbabbaaabbababbbaaaaabaaababaabb
abbaaababbbabbbbabbbbababaabababbbbbbbabababbababbbbabbb
aababaaaaabaaaababbbaaabbbabbaabababbabababbabbb
bbbababaaabababbabaabaaabbabbbab
bbabababbbaabbababbababaabbaaaaaaaabbaba
aaabaabbbbaaabbbabababbb
abababababbbaaaaabbababbaabaaaaabbaaababaabbbababbabbbabaaaaabaa
bbbbabaaabaabaabbbaaaabaaaababbbbbaababababababa
aaaabbaaabaababbbaaabbbb
baaabbabbabbababaaaabbabaabaaaabbbbaabbbabaaaabb
aabaaaaaaaababbbabbbaabaabbbbaabbbbbbbabaababababaababaa
babaabaaaaaaabababbaaabbbbbaabaabbababab
aababbbbbbbabaaaaaabaabbbbbabbaabbabbababbbbaaba
bbbabbaababbbaabaabbaaaa
abbbaabbbbbabaaaaaababab
babbababbbbabbbbaabaaabaabbaabbaabbabaab
abbaaaabbbabbbaaabbbbbab
aabbabbbbabababbaaabbbbaaababaabbabbbbbbbbaabbbaabbbabaababbbaabbbaabbaabbbaabbb
bbabaabaabaabaaaabbababaaabaabaa
bbbbbaabaaaaaabbaabbbaababaabaaabbabaaabbbabbaaaabaaaabb
aababbababaaaaaaabbbbbbbbaabaaab
aaabbbaabbbaabbababaabbabbbbabaabbaabbaaaabababbabbbbbba
bbbbbababbabbaaabbbabaabaabbbbbb
aababaabaaaaaaabaabbaabaaaabbaaaaababbbbabbbabbbbaaabbab
baabbaaabbaaaabbaaabbbbbabaabbabbbbabbba
bababaabbabbbbbabaabbbbabbbaaaab
ababbaaabbbaabbababbbbab
babaaaababbbbbbbaaaaaababaabbbba
aaaabbabbaaabbaabaabbaba
bbbbaabbbbbbbbababababbabababbaa
baaabbaaabbaaaaababababa
aaabaaaabbbbbaababbbbababbabbbaaaabbabbaababaabbbaaabaabababaababbabaabb
babaabbababaaabbbababbbb
baabababbbaaaabbbbabbaaaabbabbaabbabbbbb
babbbbababaaaabaabbbbbaaababbaabaabbbbbabbaabaaaaabababb
babbbaaabaabaaabaabaaabb
aaabbbbbbbaaabbbbaaaaaba
baaaaaabaabbbbbbabaababa
ababaaabbbbaaabbababaaaaababbaababbbbaaaaababaaabbabbbaababababbabaabaabbabbabbbbaabaabb
abaaaaaaabbaababbbaababbbbbaabbbbbaabbaa
ababbbabbabbbaabbbbbbbbb
abbaaaabbaabbbababaaaabb
bbababbbababbbababbabbbb
baabaaaababbabaaababaaabbbbaaabbbbbbbbbaaabbaaaaaaabbbba
baaabbabbaaaaaaababaaaabbaabbbba
abaaabbbbaabbaababababbaaabbbaaabaababab
bbbaabbaaabbabaababbababaabbabbbbabbaabaaaabbaabaabaabaabaaaabbb
aabbaabababaababaaababbbbabaabbb
abbaabbabaaabbabbaabaaab
bbbbabaababbbabaabaabbbb
baabaaaaababbbabbbbaaaaabbbabababaababaa
baaaababbbaaababbbabbaabaaabbaaa
babbabaabbaabbabbbabbabaabaaabab
aababaabbabaabbabbbabbba
aaaaaaaaabbaabaaaabaabbbabbabbab
babbbaabbbabbaaabbabbbbabaaabbbb
abbaaaaaabbbabbababbbbbb
babbabbabaababbaabbbabbbaaababaa
bbabbbaaaabaabbbaababaabaababaabaaaaabba
babbabaaaaabbbbbbbbabbba
baaaaaaabbbbbababbbbbbba
bbbbabaaaaaaabbbababbbba
aababbabbabaababbaaaabbbbaaabaaa
babaabbabbaaaabbbaabaaba
abbbaaabababababbbabbbbb
bbbbaabbababbabbababababbbbaaababbbaaaba
aabaaabaabbabababbaaaabbababaabbbbaabbba
bbbaabbabababbabbbbbbbbb
abbababbbbbbbaabbaaabbababbbbaabaabbaaaaabbbbbab
ababbabbbbaaabbbababbabbaaaaaaba
bbabaabababbbbabbababaabaabaabbaaabaabab
bbaaabbbbabaabbaaabbbaba
baaaabbbbaabbaaaabaaabaabaababaabbbaabaa
aababaababbaaaaaabaaaabb
bbbbabaababaababaaaaababbbbbabaababaababbbbaaaab
abababbbbababbbbbaaaabbbaabaaaaaababbabababbbbbbbbaabbba
bbaaaaaaabbbbaaabaaaaabb
bbabaaaaabbabaaababaaaabaaabbaab
bbbbbababaabbaaaaaaaaaaaabaaaabaabaabaabbaabaaab
baabbbabaaaaaaaabbbabaaaababbabbaaababab
aabbaabaabbabaaaaabbbbaabaaabbababbabbba
aaaabaababaaabaabbaaaaab
abaabbabababbbbbabbbabbabaabbaababababaa
bbbbaabbbbbbaabbbaabbbaaaabbbaabbaabbbababbbabababbaabbb
bbaabbaababbaabaabaaaabaaaabbaaabbababaabbaabbbbbabbaaaabbbbbbababbbababaabbaaaabbaaabbb
aabbbbbaaabaaaaabbaaaaaabbabaabbabbaaaabbbbbbbbbbbababbbbbababab
bbbbaabbbbbaaabaaabbabbabbababbaababaabb
abaaaabbaaaaaaabbbbaababbabbbbaabbababaaaabbabbb
bbaabaabaabbabbababbabaaababaababaaaaaba
bababbaaabaabbabaaababaaaabbaababaabbaabaabbabbababaabbabaaabaaaaababbabaabaaabb
bbbaaabaaaababbbbbaabaabbaabaaba
abbabaaabbaaaaaabaaaabaa
bbaababbbbaaababbabaaaaaaabababbaaabbabb
abaabaaaaaabbbbbabbabaab
abbbaabbaabaabbbbaaaaabb
abbaaaabaaaaababbaaaabbbababaaabbaaabaaa
bbbaaabbbbaabaabbaaaaaab
aabbbbaababbbabaabaabbba
aababbbbbaabbbaaaaaabbabaababbba
abaaaaaaaabaaaabbabaababbbababbaaabbbaabaaaababbbbbabbba
bbaaaabaabbbbaabaaababaababbaabbbabbbaaa
aabaabbbbaaaabbbaabaabbbbbabaaabbbabbbbaaabbbababababbbbbaabbbbb
aabbaabaabbaaabaaabaaaab
bbbaabbbabaaaaabaaaababb
babaaaabbbbabbaaaaaabaabaaaababa
baaaabbbbbaababbbbaabbab
bbabbbbabbabbbbaabbbbabb
baaabbababbbaaaabaabbabbaabbabbb
abaabaaababbbababbabbaabaababaaaaaabbbbabbbaabaaaaabbaaaababbababbbbbbaa
aabbbaabbabbabaaabbaaaaaababbaaaababaabaabababbbaaabbbbaaabbabab
babbaaabbaaabbababaaaaab
abbaaabaabbaaaabaabaaabaababbaba
bbaabbbbaabbaabbbaaababa
aabaabbabababbbbabbabbabaaabaaab
ababaabbbbbbbaaaababbaba
abbbbaaaabaabbabbaabbabbaabbbabb
babbabaaaabaaaabbbbbabab
bbaaaaaaaabaaaaaababaaaa
aaabaabbaaabaabbbbaaababaaababbbaaabaaabbababaaa
bbbaaabababaaaabbabaabaabbaabaabbbbbaaaa
aaaabbbbabbaabbaaabbbbbb
abbabaaabaabbbababbbaabbbabaaaaaaaabaabbbbbbbabb
bababaabbaabbaaaaababbbbaaababababababbb
bbbbaabbabbbbaabbbaaaaaaaabbbbaababaabaaaabbaaaabbbaabbb
aaaaaaabbbbbaaabbabbaabaababbababababbbaaabbabbbbaaaabaa
babbababbabaabbabababaaa
aaaabaabbabbbbaabbbbbaabbaaaaaabbbaaaaabaabaabbbbabbbabbabbaabba
aaaabbabaabaaaaabaabaaaaaaabaababbbaabaa
aababbaaabababbaaabbaaaa
baabaaaabaaabbbaabbbaabbbbbaaaab
baabbbabaaaabaabbabbababbabaaaabaabbbabaaaabaaabbbbbbabb
ababababbbaaaaaaababbbbbbbbabbab
aabbbaabbaaabbaabaabbaaaabbabaaa
bbbaabaabaabbabbaabaaabbbbbaabaabaabaabbbbabbbaabaabbbbbbbbbbaaaabbbabab
bbaaababaaaaaaaabbbaabbb
abbbbbaababbbbabaabaaabaaabbbbaaaabbaabababbaaaaabaaaabbbbbababa
abbaaabababaababaabbbbaaaabbbbaaababaaaa
abaaaabaabbaaabababababa
abbaababbaaabbbaababababaabbaabbbbaaabaa
abbaabaabaaaabbbbbbbaabbaaababbbababaaba
abbbbbaabaabbbaaaaaaababbaaaaabb
abaabaabbbaabaabaabaabab
aabaabbbabbaaaabaabbbaba
bbbbbbababbbbbbbbbababab
babbbbabbabbbbabbbbbaabbbbbbaabbaaaaaabbbbbaabaa
bbabbabbbbbbbabaabbaaaababbabbababaabbaabbbbbbaaabbbbbbaabbabababbabaababbaaaaba
abaaaabaaaababbbaaabbbbbabbbabbaaaababaa
aababaaabaaabbababaaabba
abbaabbabbabbbaabaaaaabb
aababaabbbaaaaaabaabbbaabbaabbbbabbbbaaaaaaababb
abaabbbbaabbabaabaaabbaaababaabaabbabaab
bbbbaaabbabababbbbbabbaabbbbbabbababaaaa
bbbbaabbaaabaabbbbbbbabb
abbbaaabbaabababbabbbaaa
abbaabababbbabbbbbaaabbbaabaaaabbbbaaababbbbabba
bbbababbabbbbaabbabaabababbbaabaaabbaaab
bbabbbbbabaaabbabbbaaabbbbbaaababbbaababaabaabbaaabbaaabbabaaabbbbbbabbaaabbbaab
aabbbaababbababababbabaabaabababbbaaaaab
babbaaababababbabbbbbbba
baabbaaabaaaababbbabbaabbbabaababbbbbbaaaaaaabaa
abaabaabbbaabaabbaaabbabaabbbbaabaaabaab
babbbbababbbbbbbaaaaabaa
babaabbaabaaaaaabababbaa
aabbbbaabbabbbbaabbbbbaabbbbbababaabbbabbaabaaabbbaabaaa
abbabbaaaabaaaaaababbaaaabababbbaaaabababaaabbaa
baaaaaaabbaababbbbabaabaababbaabbbaaaabaaaababab
bababbabaabaaababbbaabab
aaababbbaaabbbaaaaabbbab
aabaaaaabbbabaababaaaaab
aabbbbababbbbaabababbbabbaabaabb
bbaababbbbababbaaaababaa
baaaaaaabbaabbbbbaabaaba
aabbabaabaabbaabbbbabbbbaaabbaababbabbbb
bbbbbaabbabaaaaaabbabbab
bbaaaabbaaaabbbbbbbbbbbbbabababa
baaabbaabbabbbbaaabaaababbbbaaabbabbabababaaabab
bbbababbabbbabbbabbaabbababaabbb
bbabaababbbbbabaabaaaabaabaababa
bbaababbbbbbbbbabbabaabbbbaababa
abbababbabaabaabbaaabbabbbbababababaabbb
aabaabaaaabaaaaaaaaabbabaaababbbaaabbbbaabaaaaab
babbbbabababbbaabbaabbbaabbaababaabbbbbaabbbbbbabbaaaaaaaaaaaabbabababbb
bbabaaabbaabbaaaaabbbaba
ababbbbbbbababbbbbaaabbbaaababba
bbbbbbabaaaaabababbbbabb
bbbababbabbbbaabaaaaaabbaaaabbbbaaaaaababbaabbab
bbabbaaaaaaabaaaaabbabbb
babbaaabbbbaaaabbaababbbabbbabaababbbbbabbbabbbaaabbbbaaaaabaaaabaaabbbb
bbabbaababbaabbabaabbaaaaabbbbbbbbbabababbabbbababababaa
abbaaababababaabaaaababa
aaabbbbbabaabbababaaaabb
ababaaababaaaaaaaabbabaabbabbbbbbabbbaaa
aabababbbbaaabbaaaababababababbb
baaabbbaabbbbaabbbabbbbababaaaaaabaabaabbbaaabbaabbabbba
aaaabbababbbaaabababbbabbbabbaba
bbaabababbbaabaaaaaaaaaababbaaabbbbababbabbaaaabababbaaababaaabbaababaababaaaabaaaaabbaa
aabbaababababbabbaabbaba
aabbaabbaabaaabababbbbabaababbabbbaababbbabbbbaa
ababababaabbbbbbbaaabaaabababababaaababbaaaaabbababaabaaaaaaaabaaababaab
baabaaaabaaaabbbabbbbbbbabbabbaabbbaabab
abababbabbaaaaaaaabbabbabaabaaba
baabbbaababaaaabbbaababa
abbbabbaababaaababbbbbab
bbaaaabaabbaaabbbbabaaabbabbbbbbbbbabbabbbaabaaa
baaabbaaabaaaaaababaabbb
baaaaaaaaaabaabbaabbbbba
bbbbbbabaaaabbbbbbbaabbabbaabbba
bbbbbababaaababbabbbbaabbabbabbbaabbaaaabaaaabaaaaaabababbbababa
aaaaaabbbabaaaabbabbaaba
aabaaaaaababbbbbaababaabbbbabaabaabbbbaa
bbaabaabababbbabbabbabbababbabababaaabab
baabbbaababaabaabaaabbaaaabaabbababaabaaabbbabaa
bbbababbbabbaaaabbabbaaaaabbbbaababababbaaaababb
ababaaabaabaabbbbaaaabba
abaaaaaabbbababbbaaaabba
aaaabbababbbaaaabbabbaabbbaaabbbabbbbababbbaaaab
babaabbaabbbbaabbbbaabbb
bbaaaabaabbbbaabbbababbbaaaabbabbbaabbab
aababaabaabbbaabaabbaaaa
aaaabbbbbabbabaabaaaaaab
aabbabbabbbbbbabaaaababa
bbabbbbaabbbabbbabaabaabababbbbbabbabbbb
bbbaabbababbabbaababbaabbbbbabaaabbaaabbbaabababaabbababbbbabbbabbaabaaabbbbaaba
bbbabbaabbaababbbabbbbabaaaabbbbababbaaababaabbb
aaaabbabbaaaababaabbbabb
abaabaaababaabbaabababbaababbaaabbbaabbaabaaaabbbaaaaabbbbbaabbbaaaabaaa
bbababbabbbbaaabbabbbbabbbabaaabaaababaa
baabbaaabbbbabaabaaaaabb
aababbaaababbbbbbaaabbbaabbbabab
aaaabaabbbabaaaabbabababbbaabbbbbabbaaaabaaaaaab
babaabbaaababbaaabaabbaa
abaaabaaabaaabaaaababaaaaaaaabbb
bbbbaabbaabbbbaabbbababa
bbaaaabbaabbaabbbbbabbba
aabaaaababbabaaaabaabbabaaaaaababbbbbbbb
aababaabaabaabaaababbaba
bbaaaabababbabaaaaaabbbb
babbbaabbbbababbbbabbaabaababaaa
baaabbbaaaabbbaabbaababa
abbbbbbbbbbabaabbabbabbb
aabaabbbaabbabbababbaaabaabbabbbabbabaab
aaabaabbaabbbbaaabaababa
abbababbaabbbbaaaaaaaaaaaaabbbbbbabaabbb
aababaaaaabbbaabaaaaabbb
aaaabbbbbbbbaabaaabababaababbbaa
aabaabaaabbaaaababbabbab
aaabbbaabbbbabaaabaabbabbabbaaaaababaaabaaabaabbbaaabaab
bbaaabaaaabbabbbaabaaaaababbababbbabbbaa
bbaabbbbaabaabbaabbbbbab
baabbbaabbaaaabaaaaababa
aaaaaaaaabaabaababbaabbb
abbbbababbaababbbaababbaabaabbabbabaabbabaaaaaab
aabbbaabbbabbbaaaababbba
ababbaaaaabaabbbabababbb
aabbbaababbbabbbaababaaaabbbbababbbaabaaabaaabbbbbaabbaa
bbbaaabbaababaabaabbabbb
baaaaaababababbaaaabbbbaaabaabaa
abbaaabababbbbabbbbaabbaabaabaabaaaaaaabbaaabbbb
baaaabaaaabaabaabaaabbaabaababbaabbbbbabaabaaababbbabaabaaaababb
ababbabbbbbabbbbaababaaa
bbbaaabbbababbabbaabaabb
babbbaabbbababbbbbbabaababbbbabbbbbbabbabbaaaabaabababbabbbbbbbbabbbaaabaaabbbab
bbbaabbababaababbaabbbba
aabaaabbbbabbbbbbaabbbaababbaaabbaaaaabaaaaaabbbbbaaaabbbaabababbabbaabaabbabaaaaabbabab
aababbbbabaababbbbaababaaabbaaab
bbaaabaaabababbaaabbababaaaabaaabaabbaaabbaababaaaaababababababbbaaaabbabaaabaaaabbaabba
abbaababbaaababbbbbaaaababbabbbbbbabbaba
aaaabbaaaababaaabaaabbbabbabaabababaabbb
baabaaaababbabaaaaabaabbbaaaaababbbbabba
babbababbbbabbbbbabaaaaabaabbbaaaaaabaabaabbbbabbbaaabaa
babbbbabbbbbbbabaabbbaaa
abbaaaababababbaababaaabbaabbbabbabbabaabbaabbabbbbbabbb
bababaababbbaaabbabbaaba
bbabbbbaababaaabbaabbbaabbbaabbbbbbaabaa
abbaaaababbaaaabbabaaabbbbabbaaabaaaabababbbbbbaaaababaaabaabbaabababaaa
babaaaabaaaabbbbbaabbbabababbaba
abbaaaaaaabaaaaaaaababaa
aabbbaabaaaaabababaaaaab
baaabbbababaaaabbbbbabab
babbbabaaababaaaaabbbabb
bbabbaabbbaaaabaabbaababbbbbabba
abaabaabbababaabbababbabbbabbaaabbaabaaa
aabbaabaaababaaabbababab
bbababbaaaabaaaaaaaabbbbbbbaaaabbbaabbab
bbaababbbabbaaaaababbbabbbaaabba
bbbaaaaabbbabaabaabaaaaaabaabbbaaabbaaab
babbbababababbabbabbbabb
bbbbbbabbbaaabbbbabaabaa
bbaabbbbbbabababbbbabbabbaabbbbaaabaabbbbbabbbbbbaabaaaabababbbaaabbbaaa
bbbbbababbbabbbbabbbbbaaabbaabaaabbbabaaabababbb
aaaabbabaabbaaabbaaabbbabbabbabaababaaabaaabbaabbabbbbbabaabbabb
babbaaabbbabbbbbbbbabbbbbbbaabbaaabaabbabbababaaaaaabbaa
abbaababababaaababaaabaaabaabaabbbababbabbbbbaaaaaaabbba
bbbbbbabbabababbabbbaabbaababababbaabbababbbbbab
baaaabababbaaabbbaabbaba
aababbbbbaabbbabbaaaaaba
abbbbaaaababbabbabbaabbaabaabbba
bbbbbaabbabaaabbbbaaabba
babaabbabaaabbabbbaaaaaaabaaaababaaabbba
bbbaabaaaabbbaaabbabaabb
aaaabbabaabaabbaabaaabba
bbbabbaaaabbaababaabaaba
bbabbbbaabbbbbbbabababbabbbababa
baabbbabbabbaaabbbbaabab
abbaaaaababababbabaaaabaaabbbabaababbbba
aaabbbaabababaaaaabaabab
baaabbaabababbabbaababbb
abbaabbaaabbbaabbaaabaaa
ababaaabbabaaaaaabbaaaaabbaabaaa
ababababbbaabaababbbbbba
bbbbaaabbbabbaaabbbaaababbabbaaabaabbabbabbbabab
ababbaaabababaabbaababbaaaabbbab
abbbabbbaabbabaaaaaabaabbbbbbabb
aababaabbbaaaabbabbaabaababbbbbbbbbbbbba
bbbaaabbbbbaaabbabbbabaa
abbabaaaababbaaabababbaa
babbbabababaaabbbbbaabab
bbbabbbbbabaabaaabbbabab
aababbabbabbaaabbbbabbaaaaabbbba
aabababbbbbbaaabababaababbbaabbabbbababbabbabaabaabbbabbabbbabbb
aababbabbbabbaaaaabbaaaa
bbbabbaabbbbbabbaababababaabaabb
baaaababaabbaabbabbaaabbbaababbabbbbabab
aababbaaabbbbababbaababa
bbabaaabababbbabbaabaaaababbaabb
abbaaaabaaaaaabbaabbabbb
babaaabbbaabbbabbbabaaaabaaaaaaaabbbabaa
aabaabbbbbaabaabbbabaaaabababbba
abbbbababaabaababbbbabbbabaaabbbaaaaabbbabababaabbbaabaa
baaabbbabbaabaabababbbba
bbbaaaaaabbbbbaabbabaaabbbbaaaab
aababaabaabbbbaaaabaaabb
aabaabbbbabaababbabbbabaabababbb
bbabaaabbaaaaaaaaaaabbbbaaaaabaa
baaabbabbbbaaababbaaabbbababbaba
abaabbabbabbabbababaababbabbbabbbbbbbbbb
bbaaaaaabbbbaabbbbbbaaaabbaabbaa
aabbabaabbababbabbbaaaab
baaabbbbaaababaabababbba
bbababbabbbabbaabbaabbbbabbaabbbbbaaabba
baababbaaabbaababaaaabbbbaabbbba
baababbababbbaababbbaabbbbabbaababaabababbbabbba
baaabbbababbaaabaabaaabb
babbabbaaabbbbbbabaaabbb
bbbaaababbbabaaaababbbbbbaababaa
aaababbbaaabbbbbbbabaaabaaabbbaabbaabbbaaaaabbbabbababaa
bbababbababababbbbaaabaa
babaabbaaaabbabbbaababbbaaababab
baaabbbaabbaaababbbbbabb
bbaaabbbbaabbbabbaababbababaaaabbaaaaababbaaabaabababbba
aabaaaababaaabaababaaaabaaaaaaba
aabbbbaaabaabaabababbabbbaabbbaabaaabaabaaaaaabaaababbba
bbabbaabbbbabaaabbbaaababbbaaaabbbababab
aababbbbababaaaababbabba
babbaaaabbbbabaabaababbbababaaabaaaabaabbababbaabbbabbaa
bbabbaaabbabbbbabaabbaaa
bbabbbaabbaaaabbaaaababb
bbabbbbaabbbbababaaaaaaaabbaabbb
abaabaaaaaaaaabbabbabbbb
aabaaabaabababababaaaababbbbbbababaabbbbbbbbbbbbbaaabbbb
babababbbbaabaaababbaaababaababa
abbbbaaaaabaabaabaababbaababbbbabbabbbab
aabaabaaabbbbaaaaaaaaaaaaababaaabbabaaabababaaaaababbaba
aaaabaabababbabbaaabaabbabaababbbaaaabbbbbbababa
bbbaabbabaaabbabbababbba
ababbbabaababbbbbabbbbba
abbaaaaabaabbaaababbbbabbbbbaaababaaaabbaabbabab
baabbaabbbabbaaaabababbb
abbbaaababbabaaababaabaabaaaabba
ababbaaaabbaaaaabbbbaabbaaabaaaabaaaaaabaaaabaaaaabbaaaababbbbaa
abaaaaaabbaaaababababbabababbabbbbaaabaabbbaabaa
aabbbbabababbbbbbaabbaba
bbbababbbaaabbaaabaabbba
abbaabaaababbaaaaaabaabbbbaaaabaabbbaabaaabaabab
bbabbbaabaaabbabaaabaabbbbabbbaabbbaababbaaaaabbaabbbbba`

	rules, words := parseInput(input)
	fmt.Println("Number of rules:", len(rules))
	fmt.Println("Number of words:", len(words))
	fmt.Println("Exists left recursion:", existsLeftRecursion(rules))

	firstSets := constructFirstSets(rules)
	table := constructTable(firstSets, rules)

	var validWords []string
	var invalidWords []string
	for _, word := range words {
		wordIndex := 0
		if recursiveDescent(word, &wordIndex, "0", table, rules) {
			validWords = append(validWords, word)
		} else {
			invalidWords = append(invalidWords, word)
		}
	}
	fmt.Println(len(validWords))
	fmt.Println(validWords)
	fmt.Println(invalidWords)
}

func recursiveDescent(word string, wordIndex *int, symbol string, table map[string]map[string][]string, rules map[string]*Rule) bool {
	result := false

	if *wordIndex >= len(word) {
		fmt.Println("eaten whole world", symbol)
		return false
	}

	fmt.Println(word[:*wordIndex+1], word[*wordIndex+1:])

	possibleRules := table[symbol]["\""+string(word[*wordIndex])+"\""]
	for _, possibleRuleRaw := range possibleRules {
		previousIndex := *wordIndex
		possibleRuleParts := strings.Split(possibleRuleRaw, "-")
		rightPartIndex, _ := strconv.Atoi(possibleRuleParts[1])
		rule := rules[possibleRuleParts[0]].right[rightPartIndex]
		if rule[0][0] == '"' {
			if rule[0] == "\""+string(word[*wordIndex])+"\"" {
				*wordIndex++
				if len(rule) > 1 {
					result = true
					for _, rulePart := range rule[1:] {
						if !recursiveDescent(word, wordIndex, rulePart, table, rules) {
							result = false
							break
						}
					}
				} else {
					fmt.Println(rules[possibleRuleParts[0]].left)
					result = true
				}
				if !result {
					*wordIndex--
				}
			}
		} else {
			result = true
			for _, rulePart := range rule {
				fmt.Println(rules[possibleRuleParts[0]].left, rule, rulePart)
				if !recursiveDescent(word, wordIndex, rulePart, table, rules) {
					result = false
					break
				}
			}
		}
		if result {
			return true
		} else {
			*wordIndex = previousIndex
			fmt.Println("scratch")
		}
	}

	return result
}

func constructTable(firstSets map[string]map[string]bool, rules map[string]*Rule) map[string]map[string][]string {
	// init table
	result := make(map[string]map[string][]string)
	for _, rule := range rules {
		result[rule.left] = make(map[string][]string)
	}

	for _, rule := range rules {
		for i, rightPart := range rule.right {
			firstSet := firstSets[rightPart[0]]
			for terminal := range firstSet {
				result[rule.left][terminal] = append(result[rule.left][terminal], rule.left+"-"+strconv.Itoa(i))
			}
		}
	}
	return result
}

func constructFirstSets(rules map[string]*Rule) map[string]map[string]bool {
	result := make(map[string]map[string]bool)

	// init for all terminals
	result["\"a\""] = make(map[string]bool)
	result["\"a\""]["\"a\""] = true
	result["\"b\""] = make(map[string]bool)
	result["\"b\""]["\"b\""] = true

	end := false
	for !end {
		end = true
		for _, rule := range rules {
			if _, ok := result[rule.left]; !ok {
				result[rule.left] = make(map[string]bool)
			}
			for _, rightPart := range rule.right {
				if firstRightPart, ok := result[rightPart[0]]; ok {
					for key := range firstRightPart {
						if _, alreadyExists := result[rule.left][key]; !alreadyExists {
							result[rule.left][key] = true
							end = false
						}
					}
				}
			}
		}
	}

	return result
}

func parseInput(input string) (map[string]*Rule, []string) {
	inputParts := strings.Split(input, "\n\n")
	wordsRaw := inputParts[1]
	rules := make(map[string]*Rule)
	for _, ruleRaw := range strings.Split(inputParts[0], "\n") {
		ruleParts := strings.Split(ruleRaw, ": ")
		var rightParts [][]string
		for _, rightPartRaw := range strings.Split(ruleParts[1], " | ") {
			rightParts = append(rightParts, strings.Split(rightPartRaw, " "))
		}
		rules[ruleParts[0]] = &Rule{ruleParts[0], rightParts}
	}
	return rules, strings.Split(wordsRaw, "\n")
}

func existsLeftRecursion(rules map[string]*Rule) bool {
	for _, rule := range rules {
		for _, rightPart := range rule.right {
			if rule.left == rightPart[0] {
				return true
			}
		}
	}
	return false
}

type Rule struct {
	left  string
	right [][]string
}
