package bblfsh;

import com.fasterxml.jackson.databind.DeserializationFeature;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.SerializationFeature;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;

public class RequestReader {

    private final BufferedReader reader;
    private final ObjectMapper mapper;

    public RequestReader(final InputStream in) {
        this.reader = new BufferedReader(new InputStreamReader(in));
        this.mapper = new ObjectMapper();
        mapper.disable(SerializationFeature.FAIL_ON_EMPTY_BEANS);
        mapper.disable(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES);
        mapper.enable(DeserializationFeature.ACCEPT_EMPTY_STRING_AS_NULL_OBJECT);
    }

    public Request read() throws IOException {
        final String line = this.reader.readLine();
        if (line == null) {
            throw new CloseException("exception while reading line (null)");
        }
        return mapper.readValue(line, Request.class);
    }

}
