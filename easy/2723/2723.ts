/*
*   Stats:
*   Runtime: 57ms (beats 84.29% of TS solutions) :)
*   Memory: 43.50mb (beats 9.87% of TS solutions) :(
*/

type P = Promise<number>

async function addTwoPromises(promise1: P, promise2: P): P {
    let arr = await Promise.all([promise1, promise2])
	return new Promise((resolve, reject) => {
        resolve(arr[0] + arr[1])
    })
};

addTwoPromises(Promise.resolve(2), Promise.resolve(2)).then(console.log); // 4