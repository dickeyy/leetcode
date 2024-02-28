/*
 *   Stats:
 *   Runtime: 63ms (beats 20.00% of TS solutions) :)
 *   Memory: 50.28mb (beats 20.00% of TS solutions) :(
 */

function poorPigs(buckets: number, minutesToDie: number, minutesToTest: number): number {
    // Find minimum x such that (T+1)^x >= N

    let T = Math.floor(minutesToTest / minutesToDie);
    let x = 0;
    while (Math.pow(T + 1, x) < buckets) {
        x++;
    }
    return x;
}

console.log(poorPigs(90, 2, 7)); // 4
