/*
*   Stats:
*   Runtime: 51ms (beats 75.50% of TS solutions) :)
*   Memory: 42.56mb (beats 77.52% of TS solutions) :(
*/

// i mean come on

type JSONValue = null | boolean | number | string | JSONValue[] | { [key: string]: JSONValue };

function argumentsLength(...args: JSONValue[]): number {
	return args.length
};

console.log(argumentsLength([{},null,"3"])) // 3