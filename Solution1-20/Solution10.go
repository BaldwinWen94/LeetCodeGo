package main

import "regexp"

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	var dp [][]bool

	var matches = func(i int, j int) bool {
		if i == 0 {
			return false
		}

		if p[j-1] == '.' {
			return true
		} else {
			return p[j-1] == s[i-1]
		}
	}

	dp[0][0] = true

	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j] || dp[i][j-2]
				if matches(i, j-1) {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else if matches(i, j) {
				dp[i][j] = dp[i][j] || dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}

func isMatch2(s string, p string) bool {
	return regexp.MustCompile("^" + p + "$").MatchString(s)
}
