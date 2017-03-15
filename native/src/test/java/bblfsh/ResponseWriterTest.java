package bblfsh;

import org.apache.commons.io.IOUtils;
import org.junit.Test;

import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.nio.charset.StandardCharsets;
import java.util.ArrayList;

import static org.fest.assertions.Assertions.assertThat;

public class ResponseWriterTest {

    @Test
    public void error() throws IOException {
        final ByteArrayOutputStream out = new ByteArrayOutputStream();
        final ResponseWriter writer = new ResponseWriter(out);

        Response response;

        response = new Response();
        response.status = "fatal";
        response.errors = new ArrayList<>();
        response.errors.add("error");
        writer.write(response);

        final String result = new String(out.toByteArray());
        final String expected = "{\"status\":\"fatal\",\"errors\":[\"error\"]}\n";
        assertThat(result).isEqualTo(expected);
    }

    @Test
    public void valid() throws IOException {
        final ByteArrayOutputStream out = new ByteArrayOutputStream();
        final ResponseWriter writer = new ResponseWriter(out);
        final EclipseParser parser = new EclipseParser();
        Response response;

        final String source = IOUtils.toString(
                getClass().getResourceAsStream("/helloWorld.java"),
                StandardCharsets.UTF_8);

        response = new Response();
        response.status = "ok";
        response.ast = parser.parse(source);
        writer.write(response);
        //TODO: check output
    }

}
