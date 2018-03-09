package bblfsh;

import org.eclipse.jdt.core.dom.Message;

import sun.reflect.annotation.ExceptionProxy;

import java.io.IOException;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.stream.Stream;
import java.util.stream.Collectors;

public class Driver {

    private final RequestReader reader;
    private final ResponseWriter writer;
    private final EclipseParser parser;

    public Driver(final RequestReader reader, final ResponseWriter writer) {
        this.reader = reader;
        this.writer = writer;
        this.parser = new EclipseParser();
    }

    public void run() throws DriverException, CloseException {
        while (true) {
            this.processOne();
        }
    }

    public void processOne() throws DriverException, CloseException {
        Request request;
        try {
            request = this.reader.read();
        } catch (CloseException ex) {
            throw ex;
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
        this.writer.setContent(request.content);
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
        response.errors = Arrays.stream(response.ast.getMessages())
            .map(Message::getMessage)
            .collect(Collectors.toCollection(ArrayList::new));
        response.status = response.errors.isEmpty() ? "ok" : "error";
        return response;
    }
}
