package com.dudu.leetcode.p42;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

/**
 * @author zhaolu
 * @version 1.0
 * @datetime 11/22/2019 2:52 PM
 * @email zhao.lu@parcelx.io
 */
public class Main {
    public static int[] stringToIntegerArray(String input) {
        input = input.trim();
        input = input.substring(1, input.length() - 1);
        if (input.length() == 0) {
            return new int[0];
        }

        String[] parts = input.split(",");
        int[] output = new int[parts.length];
        for(int index = 0; index < parts.length; index++) {
            String part = parts[index].trim();
            output[index] = Integer.parseInt(part);
        }
        return output;
    }

    public static void main(String[] args) throws IOException {
        BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
        String line;
        while ((line = in.readLine()) != null) {
            int[] height = stringToIntegerArray("[0,1,0,2,1,0,1,3,2,1,2,1]");

            int ret = new Solution().trap(height);

            String out = String.valueOf(ret);

            System.out.print(out);
        }
    }
}
