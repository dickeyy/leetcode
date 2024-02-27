/*
 *   Stats:
 *   Runtime: 58ms (beats 83.30% of TS solutions) :)
 *   Memory: 53.57mb (beats 6.63% of TS solutions) :(
 */

function fizzBuzz(n: number): string[] {
    let rtnStr: string[] = [];
    for (let i = 1; i <= n; i++) {
        if (i % 3 === 0 && i % 5 === 0) {
            rtnStr.push("FizzBuzz");
        } else if (i % 3 === 0) {
            rtnStr.push("Fizz");
        } else if (i % 5 === 0) {
            rtnStr.push("Buzz");
        } else {
            rtnStr.push(i.toString());
        }
    }

    return rtnStr;
}

console.log(fizzBuzz(15));
