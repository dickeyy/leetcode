/*
*   Stats:
*   Runtime: 63ms (beats 36.93% of TS solutions)
*   Memory: 42.98mb (beats 98.10% of TS solutions)
*/

function longestCommonPrefix(strs: string[]): string {
    
    let prefix = "";

    for (let i=0; i<strs.length; i++) {
        if (strs[i].length === 0) return prefix="";
        if (i === 0) {
            prefix = strs[i];
        } else {
            let j = 0;
            while (j < prefix.length && j < strs[i].length) {
                if (prefix[j] !== strs[i][j]) {
                    prefix = prefix.slice(0, j);
                    break;
                }
                j++;
            }
            if (j < prefix.length) {
                prefix = prefix.slice(0, j);
            }
        }
    }

    return prefix

};

console.log(longestCommonPrefix(["flower","flow","flight"]))