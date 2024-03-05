export default function assertEquals<T>(expected: T, actual: T): boolean {
    if (Array.isArray(actual) && Array.isArray(expected)) {
        if (actual.length !== expected.length || !actual.every((value, index) => value === expected[index])) {
            console.log(`Failed -> expected ${expected}, but got ${actual}`);
            return false;
        }
        console.log(`Passed`);
        return true;
    } else {
        if (actual !== expected) {
            console.log(`Failed -> expected ${expected}, but got ${actual}`);
            return false;
        }
        console.log(`Passed`);
        return true;
    }
}
