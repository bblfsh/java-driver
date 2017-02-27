package tech.sourced.babelfish;

import org.apache.commons.io.IOUtils;

import java.io.BufferedInputStream;
import java.io.BufferedOutputStream;
import java.io.ByteArrayOutputStream;
import java.io.IOException;

public class Main {

    public static void main(String args[]) {

        BufferedInputStream in = new BufferedInputStream(System.in);
        BufferedOutputStream out = new BufferedOutputStream(System.out);

        while (true) {
            try {
                process(in, out);
            } catch (IOException e) {
                System.out.println("Can't write in the output given " + e.toString());
            }
        }

    }

    static private void process(BufferedInputStream in, BufferedOutputStream out) throws IOException {

        final EclipseParser parser = new EclipseParser();
        final RequestResponseMapper mapperGen = new RequestResponseMapper(true);
        final ByteArrayOutputStream baos = new ByteArrayOutputStream();
        final RequestResponseMapper.ResponseMapper responseMapper = mapperGen.getResponseMapper(baos);

        while (true) {
            String inStr;
            try {
                inStr = IOUtils.toString(in, "UTF-8");
            } catch (IOException e) {
                out.write(("A problem occurred while reading " + e.toString()).getBytes());
                out.flush();
                return;
            }
            DriverRequest request = DriverRequest.unpack(inStr);
            DriverResponse response = new DriverResponse("1.0.0", "Java", "8");
            response.setMapper(responseMapper);
            response.makeResponse(parser, request.getContent());
            response.pack();

            baos.flush();
            out.write(baos.toByteArray());
            out.flush();
        }
    }
}
