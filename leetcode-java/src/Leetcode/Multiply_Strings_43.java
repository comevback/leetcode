package Leetcode;

public class Multiply_Strings_43 {
    public static void main(String[] args) {
        Multiply_Strings_43 app = new Multiply_Strings_43();
        String res = app.multiply("103", "98");
        // String res = app.addStr("1378", "9120");
        System.out.println(res);
    }

    public String multiply(String num1, String num2) {
        if (num1.equals("0") || num2.equals("0")) {
            return "0";
        }
        String res = "";
        int addNum = 0;

        for (int i = num1.length() - 1; i >= 0; i--){
            int pos = num1.length() - 1 - i;
            String tempRes = "";
            for (int x = pos; x > 0; x--){
                tempRes += "0";
            }
            int nI = num1.charAt(i) - '0';
            for (int j = num2.length() - 1; j >= 0; j--){
                int nJ = num2.charAt(j) - '0';
                int product = nJ * nI;
                int remainder = product % 10;
                int quotient = product / 10;
                int add = remainder + addNum;
                if (add > 9){
                    add = add % 10;
                    quotient += 1;
                }

                tempRes = add + tempRes;
                addNum = quotient;
            }

            if (addNum != 0){
                tempRes = Integer.toString(addNum) + tempRes;
            }

            res = addStr(res, tempRes);
            addNum = 0;
        }

        return res;
    }

    public String addStr(String str1, String str2){
        String res = "";
        int p1 = str1.length() - 1;
        int p2 = str2.length() - 1;

        int addNum = 0;

        while (p1 >= 0 && p2 >= 0){
            int n1 = str1.charAt(p1) - '0';
            int n2 = str2.charAt(p2) - '0';
            int sum = n1 + n2;
            sum += addNum;
            if (sum > 9){
                sum = sum % 10;
                addNum = 1;
            }else {
                addNum = 0;
            }

            res = Integer.toString(sum) + res;
            p1 -= 1;
            p2 -= 1;
        }

        if (p1 >= 0){
            while (addNum != 0 && p1 >= 0){
                int n1 = str1.charAt(p1) - '0';
                int sum = n1 + addNum;
                if (sum > 9){
                    sum = sum % 10;
                    addNum = 1;
                } else {
                    addNum = 0;
                }
                res = Integer.toString(sum) + res;
                p1 -=1;
            }

            if (addNum != 0){
                res = Integer.toString(addNum) + res;
                addNum = 0;
            }

            if (p1 >= 0){
                String sub = str1.substring(0, p1 + 1);
                res = sub + res;
            }
        }

        if (p2 >= 0){
            while (addNum != 0 && p2 >= 0){
                int n2 = str2.charAt(p2) - '0';
                int sum = n2 + addNum;
                if (sum > 9){
                    sum = sum % 10;
                    addNum = 1;
                } else {
                    addNum = 0;
                }
                res = Integer.toString(sum) + res;
                p2 -=1;
            }

            if (addNum != 0){
                res = Integer.toString(addNum) + res;
                addNum = 0;
            }

            if (p2 >= 0){
                String sub = str2.substring(0, p2 + 1);
                res = sub + res;
            }
        }

        if (addNum != 0){
            res = Integer.toString(addNum) + res;
        }

        return res;
    }
}
