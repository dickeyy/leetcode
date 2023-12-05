/**
 * @param {string} s
 * @return {number}
 */
var romanToInt = function(s) {
    let map = {
        I: 1,
        V: 5,
        X: 10,
        L: 50,
        C: 100,
        D: 500,
        M: 1000
    }

    let split = s.split("")
    let start = 0

    for(let i = 0; i < split.length; i++) {
        if(map[split[i]] < map[split[i+1]]) {
            start += map[split[i+1]] - map[split[i]]
            i++
        } else {
            start += map[split[i]]
        }
    }

    return start
};

console.log(romanToInt('MCMXCIV'))