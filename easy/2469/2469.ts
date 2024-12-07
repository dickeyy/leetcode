/*
*   Stats:
*   Runtime: 52ms (beats 71.62% of TS solutions) :)
*   Memory: 43.46mb (beats 10.13% of TS solutions) :(
*/

// this may be the easiest one ever

function convertTemperature(celsius: number): number[] {
    return [ (celsius + 273.15), ((celsius * 1.80) + 32.00) ]
};

console.log(convertTemperature(36.50))
console.log(convertTemperature(122.11))