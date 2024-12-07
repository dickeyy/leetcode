import assertEquals from "../../utils/assertEquals";

/*
 *   Stats:
 *   Runtime: 51ms (beats 82.69% of TS solutions)
 *   Memory: 50.14mb (beats 42.39% of TS solutions)
 */

function defangIPaddr(address: string): string {
    return address.replaceAll(".", "[.]");
}

// test cases
assertEquals("1[.]1[.]1[.]1", defangIPaddr("1.1.1.1"));
assertEquals("255[.]100[.]50[.]0", defangIPaddr("255.100.50.0"));
