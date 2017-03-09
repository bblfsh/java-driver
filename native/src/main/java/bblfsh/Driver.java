package bblfsh;

import com.fasterxml.jackson.databind.JsonMappingException;

import java.io.*;

public class Driver {

    private BufferedReader in;
    private BufferedOutputStream out;
    private EclipseParser parser;

    private Driver() {

    }

    public Driver(final InputStream in, final OutputStream out) {
        this.in = new BufferedReader(new InputStreamReader(in));
        this.out = new BufferedOutputStream(out);
        this.parser = new EclipseParser();
    }

    public void run() throws DriverException {
        while (true) {
            try {
                this.process();
            } catch (IOException ex) {
                //TODO: handle this finer-grained in process()
                throw new DriverException("IOException in process()", ex);
            }
        }
    }

    private void process() throws DriverException, IOException {
        final RequestResponseMapper mapperGen = new RequestResponseMapper(true);
        final ByteArrayOutputStream baos = new ByteArrayOutputStream();
        final RequestResponseMapper.ResponseMapper responseMapper = mapperGen.getResponseMapper(baos);

        while (true) {
            final Response response = new Response("1.0.0");
            response.setMapper(responseMapper);
            try {
                final String inStr = in.readLine();
                Request request;
                try {
                    if (inStr != null) {
                        request = Request.unpack(inStr);
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

    static private void exceptionPrinter(Exception e, String errorString, ByteArrayOutputStream baos, BufferedOutputStream out, Response response) throws IOException {
        response.cu = null;
        response.errors.add(e.getClass().getCanonicalName());
        response.errors.add(errorString + e.getMessage());
        response.status = "fatal";
        response.pack();
        out.write(baos.toByteArray());
        out.flush();
    }

}
