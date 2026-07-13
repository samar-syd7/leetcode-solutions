/*
 * Problem: Longest Common Prefix
 * Problem ID: 14
 * Difficulty: Easy
 * Language: Go
 * Runtime: 0 ms
 * Memory: 4.2 MB
 * Synced From: LeetCode
 * Date: 2026-07-13
 */

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    prefix := strs[0]

    for i := 1; i < len(strs); i++ {
        for len(strs[i]) < len(prefix) || strs[i][:len(prefix)] != prefix {
            prefix = prefix[:len(prefix)-1]

            if prefix == "" {
                return ""
            }
        }
    }

    return prefix
}