package bblfsh;

import java.io.*;

public class Main {

    public static void main(String args[]) {
        final Driver driver = new Driver(System.in, System.out);

        try {
            driver.run();
        } catch (DriverException e) {
            System.exit(1);
        }
    }
}
