package Leetcode;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Letter_Combinations_of_a_Phone_Number_17 {
    public List<String> letterCombinations(String digits) {
        if (digits.length() == 0){
            return new ArrayList<>();
        }
        List<String> res = new ArrayList<>();
        Map<Integer, List<String>> phoneMap = new HashMap<>();
        int charInt = 97;
        for (int i = 2; i<=9; i++){
            phoneMap.put(i,new ArrayList<>());
            List<String> list = phoneMap.get(i);
            if (i == 7 || i == 9){
                int times = 4;
                while (times>0){
                    list.add(Character.toString(charInt));
                    charInt += 1;
                    times -= 1;
                }
            }else{
                int times = 3;
                while (times>0){
                    list.add(Character.toString(charInt));
                    charInt += 1;
                    times -= 1;
                }
            }
        }

        getStr(phoneMap,digits,"",res);
        return res;
    }

    public void getStr(Map<Integer, List<String>> map, String numStr, String resStr, List<String> res){
        if (numStr.length() == 0){
             res.add(resStr);
        }else{
            List<String> list = map.get(Integer.parseInt(Character.toString(numStr.charAt(0))));
            if (numStr != null && numStr.length() > 1) {
                numStr = numStr.substring(1);
            } else {
                numStr = "";
            }
            for (String charStr : list){
                String newResStr = resStr + charStr;
                getStr(map,numStr,newResStr,res);
            }
        }

    }
}
