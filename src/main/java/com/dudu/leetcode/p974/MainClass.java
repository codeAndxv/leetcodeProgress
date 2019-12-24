package com.dudu.leetcode.p974;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.io.IOException;
import java.util.Arrays;
import java.util.List;
import java.util.function.Consumer;
import java.util.function.Function;
import java.util.stream.Collectors;

/**
 * @author zhaolu
 * @version 1.0
 * @datetime 12/20/2019 11:50 AM
 * @email zhao.lu@parcelx.io
 */
public class MainClass {
    public static void main(String[] args) throws IOException {
        Solution solution = new Solution();
        File file = new File("src/main/java/com/dudu/leetcode/p974/input.txt");
        BufferedReader fileReader = new BufferedReader(new FileReader(file));
        StringBuilder stringBuilder = new StringBuilder();
        String str;
        while ((str = fileReader.readLine()) != null) {
            stringBuilder.append(str);
        }
        List<Integer> list = Arrays.stream(stringBuilder.toString().split(",")).map(Integer::parseInt).collect(Collectors.toList());
        final int[] A = new int[list.size()];
        for(int index = 0; index<list.size();index++){
            A[index] = list.get(index);
        }
        solution.subarraysDivByK(A, 586);
    }
}
