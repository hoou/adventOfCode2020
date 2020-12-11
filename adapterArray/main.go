package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//		input := `16
	//10
	//15
	//5
	//1
	//11
	//7
	//19
	//6
	//12
	//4`

	//	input := `28
	//33
	//18
	//42
	//31
	//14
	//46
	//20
	//48
	//47
	//24
	//23
	//49
	//45
	//19
	//38
	//39
	//11
	//1
	//32
	//25
	//35
	//8
	//17
	//7
	//9
	//4
	//2
	//34
	//10
	//3`

	input := `99
128
154
160
61
107
75
38
15
11
129
94
157
84
121
14
119
48
30
10
55
108
74
104
91
45
134
109
164
66
146
44
116
89
79
32
149
1
136
58
96
7
60
23
31
3
65
110
90
37
43
115
122
52
113
123
161
50
95
150
120
101
126
151
114
127
73
82
162
140
51
144
36
4
163
85
42
59
67
64
86
49
2
145
135
22
24
33
137
16
27
70
133
130
20
21
83
143
100
41
76
17`

	adaptersRaw := strings.Split(input, "\n")
	var adapters []int
	for _, adapterRaw := range adaptersRaw {
		adapter, _ := strconv.Atoi(adapterRaw)
		adapters = append(adapters, adapter)
	}
	sort.Ints(adapters)
	if adapters != nil {
		adapters = append(adapters, adapters[len(adapters)-1]+3)
	}
	lastAdapter := 0
	oneJoltDifferencesSum := 0
	threeJoltsDifferencesSum := 0
	distinctWays := 1
	freshOnes := 0
	firstInGroup := 0
	for _, adapter := range adapters {
		switch adapter - lastAdapter {
		case 1:
			oneJoltDifferencesSum++
			freshOnes++

			if adapter-firstInGroup == 4 {
				if freshOnes == 2 {
					distinctWays *= 2
				} else if freshOnes == 3 {
					distinctWays *= 4
				} else if freshOnes == 4 {
					distinctWays *= 7
				}
				freshOnes = 0
				firstInGroup = adapter
			}

		case 3:
			freshOnes--
			threeJoltsDifferencesSum++
			firstInGroup = adapter
			if freshOnes == 1 {
				distinctWays *= 2
			} else if freshOnes == 2 {
				distinctWays *= 4
			} else if freshOnes == 3 {
				distinctWays *= 7
			}
			freshOnes = 0
		}
		lastAdapter = adapter
	}

	fmt.Println(oneJoltDifferencesSum, threeJoltsDifferencesSum, oneJoltDifferencesSum*threeJoltsDifferencesSum, distinctWays)
}
