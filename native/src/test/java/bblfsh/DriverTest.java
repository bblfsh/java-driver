package bblfsh;

import org.junit.Test;

import java.io.ByteArrayInputStream;
import java.io.ByteArrayOutputStream;
import java.io.InputStream;

import static org.fest.assertions.Assertions.*;

public class DriverTest {

    @Test
    public void process() throws DriverException, CloseException {
        final String input = "{\"content\":\"package foo;\"}\n{\"content\":\"package bar;\"}\n";
        final InputStream in = new ByteArrayInputStream(input.getBytes());
        final RequestReader reader = new RequestReader(in);

        final ByteArrayOutputStream out = new ByteArrayOutputStream();
        final ResponseWriter writer = new ResponseWriter(out);

        final Driver driver = new Driver(reader, writer);
        driver.processOne();
        driver.processOne();
        //TODO: check output
    }

    @Test
    public void exitOnCloseIn() throws DriverException, CloseException {
        final String input = "{\"content\":\"package foo;\"}\n{\"content\":\"package bar;\"}\n";
        final InputStream in = new ByteArrayInputStream(input.getBytes());
        final RequestReader reader = new RequestReader(in);

        final ByteArrayOutputStream out = new ByteArrayOutputStream();
        final ResponseWriter writer = new ResponseWriter(out);

        final Driver driver = new Driver(reader, writer);
        try {
            driver.run();
        } catch (CloseException ex) { }
    }

    @Test
    public void processInvalid() throws DriverException, CloseException {
        final String input = "garbage";
        final InputStream in = new ByteArrayInputStream(input.getBytes());
        final RequestReader reader = new RequestReader(in);

        final ByteArrayOutputStream out = new ByteArrayOutputStream();
        final ResponseWriter writer = new ResponseWriter(out);

        final Driver driver = new Driver(reader, writer);
        driver.processOne();

        final String result = new String(out.toByteArray());
        assertThat(result).isEqualTo("{\"status\":\"fatal\",\"errors\":[\"Unrecognized token 'garbage': was expecting ('true', 'false' or 'null')\\n at [Source: garbage; line: 1, column: 15]\"]}\n");
    }
}
