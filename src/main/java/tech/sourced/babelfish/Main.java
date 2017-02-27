package tech.sourced.babelfish;

import org.apache.commons.io.IOUtils;

import java.io.*;

public class Main {

    public static void main(String args[]) {

        BufferedInputStream in = new BufferedInputStream(System.in);
        BufferedOutputStream out = new BufferedOutputStream(System.out);
        //String sample = "{\"action\" : \"getAST\",\"language\" : \"Java\",\"languageVersion\" : \"8\",\"content\" : \"package Hello;\\n\\n pu####~~~blic class Hello {\\n    1public static void main(String\\n }\\n\"}";
        //BufferedInputStream in = new BufferedInputStream(new ByteArrayInputStream(sample.getBytes()));

        while (true) {
            try {
                process(in, out);
            } catch (IOException e) {
                //This should never happen
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
            final DriverResponse response = new DriverResponse("1.0.0", "Java", "8");
            try {
                inStr = IOUtils.toString(in, "UTF-8");
                final DriverRequest request = DriverRequest.unpack(inStr);
                response.setMapper(responseMapper);
                response.makeResponse(parser, request.content);
                response.pack();

                baos.flush();
                out.write(baos.toByteArray());
                out.flush();
            } catch (IOException e) {

                response.setMapper(responseMapper);
                response.cu = null;
                response.errors.add("IOException");
                response.errors.add("A problem occurred while reading " + e.toString());
                response.pack();

                out.write(baos.toByteArray());
                out.flush();
                return;
            }

        }
    }
}
