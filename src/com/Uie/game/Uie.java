package com.Uie.game;

import java.util.Scanner;

public class Uie {
    public static void main(String []args){
        Scanner scanner = new Scanner(System.in);
        String in = null;
        System.out.println("1 is Welcome(start)   2 is return  ");
        in = scanner.next();
        String welcome = "1";
        String returno = "2";
        String whilej = "3";
        String f = "uie ip";
        if(in.equals(welcome)){
            Scanner scanner1 = new Scanner(System.in);
            in = null;
            System.out.println("Welcome\nin:");
            in = scanner1.next();
            if(in.equals(f)){
                in = null;
                Scanner scanner2 = new Scanner(System.in);
                String gui = null;
                System.out.println("in ip");
                in = scanner2.next();
                gui = in;
                System.out.println("uie has backed into IP:" + in);
            }
            //
    }//jdjjksdee
}
}
