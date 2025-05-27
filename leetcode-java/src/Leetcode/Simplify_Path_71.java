package Leetcode;

import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

/**
 * Leetcode 71. Simplify Path
 * https://leetcode.com/problems/simplify-path/
 *
 * Given an absolute path for a file (Unix-style), simplify it.
 *
 * Example:
 * Input: "/a/./b/../../c/"
 * Output: "/c"
 */
public class Simplify_Path_71 {
    public static void main(String[] args){
        String input = "/../";
        String input1 = "/a/./b/../../c/";
        Simplify_Path_71 app = new Simplify_Path_71();
        String res = app.simplifyPath(input);
        System.out.println(res);
    }


    public String simplifyPath(String path) {
        String res = "";
        String[] strArr = path.split("/");
        List<String> stack = new ArrayList<>();
        for (String str : strArr){
            if (!Objects.equals(str, "")){
                stack.add(str);
            }
        }

        int pop = 0;
        while (stack.size() > 0){
            if (Objects.equals(stack.getLast(), ".")){
                stack.removeLast();
            } else if (Objects.equals(stack.getLast(), "..")){
                pop += 1;
                stack.removeLast();
            } else {
                if (pop>0){
                    pop -= 1;
                    stack.removeLast();
                }else{
                    res = stack.getLast() + res;
                    res = "/" + res;
                    stack.removeLast();
                }
            }
        }

        if (Objects.equals(res, "")){
            res = "/";
        }

        return res;
    }

//    public String simplifyPath(String path) {
//        if (Objects.equals(path,"/")){
//            return path;
//        }
//        String res = "";
//        List<String> strArr = new ArrayList<>();
//        strArr.add("/");
//        int head = 1;
//        int tail = 1;
//
//        while (tail < path.length()) {
//            if (path.charAt(tail) == '/'){
//                if (tail == 0){
//                    tail += 1;
//                    continue;
//                }
//                if (path.charAt(tail-1) == '/'){
//                    tail +=1;
//                    head = tail;
//                    continue;
//                }
//                String dir = path.substring(head, tail);
//                strArr.add(dir);
//                strArr.add("/");
//                tail += 1;
//                head = tail;
//            }else{
//                tail += 1;
//            }
//        }
//
//        if (head<tail){
//            String dir = path.substring(head, tail);
//            strArr.add(dir);
//        }
//
//        if (strArr.size()>1 && Objects.equals(strArr.getLast(), "/")) {
//            strArr.removeLast();
//        }
//
//        int back = 0;
//
//        while(strArr.size() > 0){
//            if (Objects.equals(strArr.getLast(), "..")) {
//                back += 1;
//                strArr.removeLast();
//            } else if (Objects.equals(strArr.getLast(), ".")) {
//                strArr.removeLast();
//            } else {
//                if (Objects.equals(strArr.getLast(), "/")){
//                    if (strArr.size() == 1 && res.length() == 0){
//                        res = strArr.getLast() + res;
//                        strArr.removeLast();
//                        continue;
//                    }
//                    if (res.length()>0&& res.charAt(0) != '/'){
//                        res = strArr.getLast() + res;
//                    }
//                    strArr.removeLast();
//                }else{
//                    if (back > 0){
//                        back -= 1;
//                        strArr.removeLast();
//                    } else {
//                        res = strArr.getLast() + res;
//                        strArr.removeLast();
//                    }
//                }
//            }
//        }
//
//        return res;
//    }
}
