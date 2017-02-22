package ReadFile;

import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
/**
 * A simple example program that reads a text file line by line and display each line.
 */
public class ReadTextFileExample {
    public static void main(String[] args) {
        BufferedReader br = null;
        try {
            br = new BufferedReader(new FileReader("C:\\temp\\testfile.txt"));
            String line;
            while ((line = br.readLine()) != null) {
                System.out.println(line);
            }
        } catch (IOException e) {
            e.printStackTrace();
        } finally {
            try {
                if (br != null) {
                    br.close();
                }
            } catch (IOException ex) {
                ex.printStackTrace();
            }
        }
    }
}