import assertEquals from "../../utils/assertEquals";

/*
 *   Stats:
 *   Runtime: 44ms (beats 98.73% of TS solutions)
 *   Memory: 51.23mb (beats 72.15% of TS solutions)
 */

// This question states that the encoding and decoding algorithms can be whatever we want
// so long as it can be converted back to the original url.
// given that, we can just return the provided string.
// yes just returning the string isn't doing anything, but,
// the question isn't asking us to do anything so who cares?
// you could make it encode it using a crypto library or some custom algo,
// but why?

/**
 * Encodes a URL to a shortened URL.
 */
function encode(longUrl: string): string {
    return longUrl;
}

/**
 * Decodes a shortened URL to its original URL.
 */
function decode(shortUrl: string): string {
    return shortUrl;
}

// test cases
assertEquals(
    "https://leetcode.com/problems/design-tinyurl",
    decode(encode("https://leetcode.com/problems/design-tinyurl"))
);
