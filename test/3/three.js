// 修bug

function solution(A) {
    var n = A.length;
    var result = 0;
    var i;
    for (i = 0; i < n - 1; i++) {
        if (A[i] == A[i + 1])
            result = result + 1;
    }
    // 这里的r不应该为0，因为如果[0,0,0]这种情况，本身有两个相邻数，但是任意翻转一个数，都会减少一个相邻数，所以r应该为一个极小的数。例如-Infinity
    var r = 0;
    for (i = 0; i < n; i++) {
        var count = 0;
        if (i > 0) {
            if (A[i - 1] != A[i])
                count = count + 1;
            else
                count = count - 1;
        }
        if (i < n - 1) {
            if (A[i + 1] != A[i])
                count = count + 1;
            else
                count = count - 1;
        }
        r = Math.max(r, count);
    }
    return result + r;
}
