import assertEquals from "../utils/assertEquals";

/*
 *   Stats:
 *   Runtime: 77ms (beats 89.64% of TS solutions)
 *   Memory: 53.33mb (beats 99.36% of TS solutions)
 */

function lengthOfLongestSubstring(s: string): number {
    // make a left and right pointer
    // start at s[0]
    // make a map and a length int.
    // go through the s and see if range [l, r] has any repeating chars
    // if it doesnt, check if [l,r].length > length
    // if it is, update length

    let left = 0,
        right = 0,
        len = 0;
    const n = s.length;

    let map = new Map<string, number>();

    while (right < n) {
        if (map.has(s.charAt(right))) {
            left = Math.max((map.get(s.charAt(right)) || 0) + 1, left);
        }

        map.set(s.charAt(right), right);
        len = Math.max(len, right - left + 1);

        right++;
    }

    return len;
}

// test cases
assertEquals(3, lengthOfLongestSubstring("abcabcbb"));
assertEquals(1, lengthOfLongestSubstring("bbbbb"));
assertEquals(3, lengthOfLongestSubstring("pwwkew"));
