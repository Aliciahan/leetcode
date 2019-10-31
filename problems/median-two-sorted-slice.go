package problems

/*
The median: dividing a set into two equal length subsets, one is always greater than the other

image list A where len(A) = m, if we cut the list at position i we have :
a[0], a[1], a[2] .... a[i-1] | a[i], a[i+1], a[i+2] ... a[m-1] we have m +1 ways cutting 最左， 最右都算一种
The same way, cut B, where len(B) = n into tow parts at position j, we have n + 1 cutting.
   B[0], B[1], ..., B[j-1]  |  B[j], B[j+1], ..., B[n-1]

then we put the left parts and right parts together:
          left_part          |        right_part
    A[0], A[1], ..., A[i-1]  |  A[i], A[i+1], ..., A[m-1]
    B[0], B[1], ..., B[j-1]  |  B[j], B[j+1], ..., B[n-1]
condition meets:
len(left) = len(right) : (i-1)-0+1 + j-1-0+1 == m-1-i+1 + n-1-j+1 => i+j = m-i+n-j  OR i+j = m-i+n-j+1 得考虑总长度基数还是偶数
max(left) <= min(right) : B[j-1]<=A[i] and A[i-1]<=B[j]

ps. make sure len n>=m in this way j will not be negative
then
median = (max(left)+min(right))/2

that is to say, search i in [0,m] where B[j-1]<=A[i] and A[i-1] <= B[j] , j= (m+n+1)/2-i


Steps:
1. make sure that m>n , if not, reverse
2. set i=(iMin+iMax)/2, j = (m+n+1)/2 - i
3. if B[j-1] <= A[i] and A[i-1] <= B[j] OK
	else if B[j-1] > A[i] //A is too small
		range to [i+1, iMax], thus iMin=i+1 go to 2
	else if A[i-1]>B[j]
		range to imin, i-1 thus imax = i-1; go to 2
4. when the i is found, the median is :
max (A[i-1], B[j-1]) when m+n is odd
(max(A[i-1],B[j-1) + min(A[i],B[j]))/2

 */

func MedianTwoSortedList(nums1, nums2 []int) double {



}
