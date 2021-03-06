/*
 * @lc app=leetcode id=518 lang=golang
 *
 * [518] Coin Change 2
 *
 * https://leetcode.com/problems/coin-change-2/description/
 *
 * algorithms
 * Medium (40.86%)
 * Total Accepted:    37.3K
 * Total Submissions: 89.7K
 * Testcase Example:  '5\n[1,2,5]'
 *
 * You are given coins of different denominations and a total amount of money.
 * Write a function to compute the number of combinations that make up that
 * amount. You may assume that you have infinite number of each kind of
 * coin.
 *
 *
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: amount = 5, coins = [1, 2, 5]
 * Output: 4
 * Explanation: there are four ways to make up the amount:
 * 5=5
 * 5=2+2+1
 * 5=2+1+1+1
 * 5=1+1+1+1+1
 *
 *
 * Example 2:
 *
 *
 * Input: amount = 3, coins = [2]
 * Output: 0
 * Explanation: the amount of 3 cannot be made up just with coins of 2.
 *
 *
 * Example 3:
 *
 *
 * Input: amount = 10, coins = [10]
 * Output: 1
 *
 *
 *
 *
 * Note:
 *
 * You can assume that
 *
 *
 * 0 <= amount <= 5000
 * 1 <= coin <= 5000
 * the number of coins is less than 500
 * the answer is guaranteed to fit into signed 32-bit integer
 *
 *
 */
// 0-1 Knapsack problem:
// 	dp[i][j]: Number of combinations to use 'first-i' coins
// 	          to make up amount 'j'.
//
// 	dp[i][j] = Sum(dp[i-1][j - k * coins[i-1]]), where k = 0 -> (j / coins[i-1])
//	(P.S. coins[i-1] = Current coin at each iteration.)
//		ex: amount: 3, coins: [1, 2, 5]
//			=> dp[0][0] = 1				(# of coins: [] for amount: 0)
//			=> dp[1][0] = dp[0][0]		(# of coins: [] for amount: 0 + Use zero coins: 1)
//						= 1
//			=> dp[1][1] = dp[0][1] +	(# of coins: [] for amount: 1 + Use zero coins: 1)
//						  dp[0][0]		(# of coins: [] for amount: 0 + Use one coin: 1)
//						= 0 + 1
//						= 1
//			=> dp[1][2] = dp[0][2] +	(# of coins: [] for amount: 2 + Use zero coins: 1)
//						  dp[0][1] +	(# of coins: [] for amount: 1 + Use one coin: 1)
//						  dp[0][0]		(# of coins: [] for amount: 0 + Use two coins: 1)
//						= 0 + 0 + 1
//						= 1
//			=> dp[1][3] = dp[0][3] +	(# of coins: [] for amount: 3 + Use zero coins: 1)
//						  dp[0][2] +	(# of coins: [] for amount: 2 + Use one coin: 1)
//						  dp[0][1] +	(# of coins: [] for amount: 1 + Use two coins: 1)
//						  dp[0][0]		(# of coins: [] for amount: 0 + Use three coins: 1)
//						= 0 + 0 + 0 + 1
//						= 1
//			=> dp[2][0] = dp[1][0]		(# of coins: [1] for amount: 0 + Use zero coins: 2)
//						= 1
//			=> dp[2][1] = dp[1][1]		(# of coins: [1] for amount: 1 + Use zero coins: 2)
//						= 1
//			=> dp[2][2] = dp[1][2] +	(# of coins: [1] for amount: 2 + Use zero coins: 2)
//						  dp[1][0]		(# of coins: [1] for amount: 0 + Use one coins: 2)
//						= 1 + 1
//						= 2
//			=> dp[2][3] = dp[1][3] +	(# of coins: [1] for amount: 3 + Use zero coins: 2)
//						  dp[1][1] 		(# of coins: [1] for amount: 1 + Use one coin: 2)
//						= 1 + 1
//						= 2
//			=> dp[3][0] = dp[2][0]		(# of coins: [1, 2] for amount: 0 + Use zero coins: 5)
//						= 1
//			=> dp[3][1] = dp[2][1]		(# of coins: [1, 2] for amount: 1 + Use zero coins: 5)
//						= 1
//			=> dp[3][2] = dp[2][2]		(# of coins: [1, 2] for amount: 2 + Use zero coins: 5)
//						= 2
//			=> dp[3][3] = dp[2][3]		(# of coins: [1, 2] for amount: 3 + Use zero coins: 5)
//						= 2
//
// References:
//	http://bit.ly/2UDXm88
//	http://bit.ly/2UCAF4n
func change(amount int, coins []int) int {
	dp := make([][]int, len(coins)+1)
	for i := 0; i < len(coins)+1; i++ {
		dp[i] = make([]int, amount+1)
	}

	// Number of combinations to use no coins
	// to make up with amount 0: 1.
	dp[0][0] = 1

	for i := 1; i <= len(coins); i++ {
		for j := 0; j <= amount; j++ {
			for k := 0; k*coins[i-1] <= j; k++ {
				dp[i][j] += dp[i-1][j-k*coins[i-1]]
			}
		}
	}

	return dp[len(coins)][amount]
}
