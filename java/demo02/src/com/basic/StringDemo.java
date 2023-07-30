package com.basic;

public class StringDemo {
    //    String str = "Runoff";
//    String str2 = new String("Runoff");
//    String 创建的字符串存储在公共池中，而 new 创建的字符串对象在堆上：

//    public static void main(String[] args) {
//        char[] helloArray = {'r', 'u', 'n', 'o', 'o', 'b'};
//        String helloString = new String(helloArray);
//        System.out.println(helloString);
//    }

    public static void main(String[] args) {
        String site = "www.runoob.com";
        int len = site.length();
        System.out.println("菜鸟教程网址长度 : " + len);
    }
}
