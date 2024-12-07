/*
 *   Stats:
 *   Runtime: 68ms (beats 56.25% of TS solutions)
 *   Memory: 55.22mb (beats 31.25% of TS solutions)
 */

function minimumLength(s: string): number {
    let start = 0;
    let end = s.length - 1;

    while (start < end) {
        if (s[start] !== s[end]) {
            return end - start + 1; // returns the length of the string calculated
        }

        let c = s[start];

        while (start <= end && s[start] == c) {
            start++;
        }

        while (end >= start && s[end] == c) {
            end--;
        }
    }

    return start > end ? 0 : 1;
}

// Test cases
console.log(minimumLength("ca")); // 2
console.log(minimumLength("cabaabac")); // 0
console.log(minimumLength("aabccabba")); // 3
console.log(minimumLength("bbbbbbbbbbbbbbbbbbbbbbbbbbbabbbbbbbbbbbbbbbccbcbcbccbbabbb")); // 1
