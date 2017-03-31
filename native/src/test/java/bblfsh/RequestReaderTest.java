package bblfsh;

import org.junit.Test;

import java.io.ByteArrayInputStream;
import java.io.FileInputStream;
import java.io.IOException;
import java.io.InputStream;

import static org.fest.assertions.Assertions.*;

public class RequestReaderTest {

    @Test
    public void twoValid() throws IOException {
        final String input = "{\"content\":\"package foo;\"}\n{\"content\":\"package bar;\"}\n";
        final InputStream in = new ByteArrayInputStream(input.getBytes());
        final RequestReader reader = new RequestReader(in);

        Request request = reader.read();
        Request expected = new Request();
        expected.content = "package foo;";
        assertThat(request.content).isEqualTo(expected.content);
        assertThat(request).isEqualTo(expected);

        request = reader.read();
        expected = new Request();
        expected.content = "package bar;";
        assertThat(request.content).isEqualTo(expected.content);
        assertThat(request).isEqualTo(expected);
    }

    @Test
    public void oneMalformedOnValid() throws IOException {
        final String input = "foo\n{\"content\":\"package foo;\"}\n";
        final InputStream in = new ByteArrayInputStream(input.getBytes());
        final RequestReader reader = new RequestReader(in);

        boolean thrown = false;
        try {
            reader.read();
        } catch (IOException ex) {
            thrown = true;
        }
        assertThat(thrown).isTrue();

        Request request = reader.read();
        Request expected = new Request();
        expected.content = "package foo;";
        assertThat(request.content).isEqualTo(expected.content);
        assertThat(request).isEqualTo(expected);
    }

    @Test
    public void throwOnClosed() throws IOException {
        final InputStream in = new ByteArrayInputStream(new byte[]{}) {
            @Override
            public int read(byte[] var1) throws IOException {
                throw new IOException("closed");
            }
        };
        final RequestReader reader = new RequestReader(in);

        boolean thrown = false;
        try {
            reader.read();
        } catch (IOException ex) {
            thrown = true;
        }
        assertThat(thrown).isTrue();
    }
}
