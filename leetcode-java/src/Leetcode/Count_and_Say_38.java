package Leetcode;

public class Count_and_Say_38 {
    public static void main(String[] args) {
        Count_and_Say_38 app = new Count_and_Say_38();
        String res = app.countAndSay(4);
        System.out.println(res);
    }

    public String countAndSay(int n) {
        if (n == 1) {
            return "1";
        } else {
            return (cas(countAndSay(n - 1)));
        }
    }

    public String cas(String str) {
        int count = 1;
        char before = str.charAt(0);
        StringBuilder res = new StringBuilder();

        for (int i = 1; i< str.length(); i++){
            if (str.charAt(i) == before){
                count += 1;
            } else {
                String newStr = Integer.toString(count) + before;
                res.append(newStr);
                count = 1;
                before = str.charAt(i);
            }
        }

        String newStr = Integer.toString(count) + before;
        res.append(newStr);

        return res.toString();
    }
}
