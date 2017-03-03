package tech.sourced.babelfish;

import com.fasterxml.jackson.databind.JsonMappingException;

import java.io.*;

public class Main {

    public static void main(String args[]) {
        final BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
        final BufferedOutputStream out = new BufferedOutputStream(System.out);

        while (true) {
            try {
                process(in, out);
            } catch (IOException e) {
                //This exception only occurs when you can't write in System.out
                System.err.println("Can't write in the output given " + e.toString());
                System.exit(1);
            }
        }

    }

    static private void process(BufferedReader in, BufferedOutputStream out) throws IOException {
        final EclipseParser parser = new EclipseParser();
        final RequestResponseMapper mapperGen = new RequestResponseMapper(true);
        final ByteArrayOutputStream baos = new ByteArrayOutputStream();
        final RequestResponseMapper.ResponseMapper responseMapper = mapperGen.getResponseMapper(baos);

        while (true) {
            final DriverResponse response = new DriverResponse("1.0.0");
            response.setMapper(responseMapper);
            try {
                final String inStr = in.readLine();
                DriverRequest request;
                try {
                    if (inStr != null) {
                        request = DriverRequest.unpack(inStr);
                    } else {
                        exceptionPrinter(new NullPointerException(), "reading string ", baos, out, response);
                        return;
                    }
                } catch (JsonMappingException e) {
                    exceptionPrinter(e, "Error reading the petition: ", baos, out, response);
                    return;
                }
                if (request.content != null && request.action != null) {
                    response.makeResponse(parser, request.content);
                } else {
                    exceptionPrinter(new JsonMappingException(""), "Null request ", baos, out, response);
                    return;
                }
                response.pack();
                out.write(baos.toByteArray());
                baos.flush();
                baos.reset();
                out.flush();
            } catch (JsonMappingException e) {
                exceptionPrinter(e, "Error serializing the AST to JSON: ", baos, out, response);
                return;
            } catch (IOException e) {
                exceptionPrinter(e, "A problem occurred while processing the petition: ", baos, out, response);
                return;
            }
        }
    }

    static private void exceptionPrinter(Exception e, String errorString, ByteArrayOutputStream baos, BufferedOutputStream out, DriverResponse response) throws IOException {
        response.cu = null;
        response.errors.add(e.getClass().getCanonicalName());
        response.errors.add(errorString + e.getMessage());
        response.status = "fatal";
        response.pack();
        out.write(baos.toByteArray());
        out.flush();
    }
}
