package com.object;

public class TestMax {
    /** 主方法 */
    public static void main(String[] args) {
        double i = 5.2;
        double j = 2.5;
        double k = max(i, j);
        System.out.println( i + " 和 " + j + " 比较，最大值是：" + k);
    }

    /*返回两个整数变量较大的值*/
    public static int max(int num1, int num2) {

        return Math.max(num1, num2);
    }

    public static double max(double num1, double num2) {

        return Math.max(num1, num2);
    }
}
