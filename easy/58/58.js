/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLastWord = function(s) {
    
    // s = "Kello World"
    // out = 5

    let split = s.split(" ")

    split = split.filter((str) => str != "");

    return split[split.length - 1].length
};

console.log(lengthOfLastWord('Poop in my ass '))