// 🧠 Problem
// Given an m x n grid, a robot starts at the top-left corner and can only move:
// right ➡️
// down ⬇️
// Find how many unique paths exist to reach the bottom-right corner.

package main

import "fmt"

func uniquePaths(m int, n int) int {
	// Create DP table
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// Initialize first row and first column
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	// Fill DP table
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

func UniquePath() {
	fmt.Println(uniquePaths(3, 7)) // Output: 28
}
