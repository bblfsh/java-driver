package bblfsh;

import org.junit.Test;

import java.io.ByteArrayInputStream;
import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.io.InputStream;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.node.ObjectNode;

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
        // TODO: check output
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
        } catch (CloseException ex) {
        }
    }

    @Test
    public void processError() throws DriverException, CloseException {
        final String input = "{\"content\":\"package\"}\n";
        final InputStream in = new ByteArrayInputStream(input.getBytes());
        final RequestReader reader = new RequestReader(in);

        final ByteArrayOutputStream out = new ByteArrayOutputStream();
        final ResponseWriter writer = new ResponseWriter(out);

        final Driver driver = new Driver(reader, writer);
        driver.processOne();
        assertThat(out.toString()).contains("\"status\":\"error\"");
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
        assertThat(result).contains("{\"status\":\"fatal\",\"errors\":[\"");
        assertThat(result).contains("Unrecognized token 'garbage'");
    }

    @Test
    public void processComment() throws DriverException, CloseException {
        final String input = "{\"content\":\"class EOF_Test { public void method() {\\r\\n   /*\\r\\n*/ } }\"}";
        final Driver driver = process(input);
        driver.processOne();
        // TODO: check output
    }

    @Test
    public void processStringLiteral() throws DriverException, IOException {
        // give
        final String input = "{\"content\":\"class String_Test { String s = \\\"b\\\\nc\\\\41\\\"; \\r\\n }\"}";
        final Driver driver = process(input);

        // when
        driver.processOne();

        // then check a new node is present \w normalized value
        String json = driver.writer.out.toString();
        final ObjectNode node = new ObjectMapper().readValue(json, ObjectNode.class);
        JsonNode newNode = node.findPath("unescapedValue");

        assertThat(newNode.isMissingNode()).isFalse();
        assertThat(newNode.asText()).isEqualTo("b\nc!");
    }

    public Driver process(String input) {
        final InputStream in = new ByteArrayInputStream(input.getBytes());
        final RequestReader reader = new RequestReader(in);

        final ByteArrayOutputStream out = new ByteArrayOutputStream();
        final ResponseWriter writer = new ResponseWriter(out);

        return new Driver(reader, writer);
    }

}
