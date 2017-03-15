package bblfsh;

import java.io.*;
import java.util.ArrayList;

public class Driver {

    private final RequestReader reader;
    private final ResponseWriter writer;
    private final EclipseParser parser;

    public Driver(final RequestReader reader, final ResponseWriter writer) {
        this.reader = reader;
        this.writer = writer;
        this.parser = new EclipseParser();
    }

    public void run() throws DriverException {
        while (true) {
            this.processOne();
        }
    }

    public void processOne() throws DriverException {
        Request request;
        try {
            request = this.reader.read();
        } catch (Exception ex) {
            final Response response = createFatalResponse(ex);
            try {
                this.writer.write(response);
            } catch (IOException ex2) {
                throw new DriverException("exception while writing fatal response", ex2);
            }

            return;
        }

        final Response response = this.processRequest(request);
        try {
            this.writer.write(response);
        } catch (IOException ex) {
            throw new DriverException("exception writing response", ex);
        }
    }

    private Response createFatalResponse(final Exception e) {
        final Response r = new Response();
        r.status = "fatal";
        r.errors = new ArrayList<>();
        r.errors.add(e.getMessage());
        return r;
    }

    private Response processRequest(final Request request) {
        Response response = new Response();
        try {
            response.ast = parser.parse(request.content);
        } catch (IOException e) {
            return createFatalResponse(e);
        }

        response.status = "ok";
        return response;
    }
}
