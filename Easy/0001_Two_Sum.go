/*
 * Problem: Two Sum
 * Problem ID: 1
 * Difficulty: Easy
 * Language: Go
 * Runtime: 0 ms
 * Memory: 6 MB
 * Synced From: LeetCode
 * Date: 2026-07-13
 */

func twoSum(nums []int, target int) []int {
    hashmap := make(map[int]int) // number -> index

    for i, num := range nums {
        complement := target - num

        // check if complement exists
        if index, exists := hashmap[complement]; exists {
            return []int{index, i}
        }

        // store current number with its index
        hashmap[num] = i
    }

    return[]int{}
}