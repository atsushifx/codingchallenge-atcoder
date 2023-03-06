
// problem solver
function isAnagram(s: string, t: string): boolean {
	const sHash = s.split("").sort().toString()
	const tHash = t.split("").sort().toString()
	return sHash == tHash;
};

export default {
	isAnagram
}


