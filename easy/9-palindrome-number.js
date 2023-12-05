// Given an integer x, return true if x is a palindrome, and false otherwise
// an integer is a palindrome if it is the same forwards as it is backwards

/**
 * @param {number} x
 * @return {boolean}
 */
var isPalindrome = function(x) {
    x = String(x)
    let split = x.split("");
    let reverse = split.reverse()
    let revJoin = reverse.join("")

    if (x == revJoin) {
        return true
    } else {
        return false
    }
};