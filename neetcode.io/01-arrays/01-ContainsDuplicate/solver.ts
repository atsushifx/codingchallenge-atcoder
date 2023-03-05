
// problem solver

function containsDuplicate(nums: number[]): boolean {
	function uniq(n: number[]): number[] {
		return Array.from(new Set(n))
	}
	const nums2 = uniq(nums);

	return nums.length != nums2.length;
};

export default {
	containsDuplicate
}


