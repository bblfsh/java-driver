package bblfsh;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.databind.DeserializationFeature;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.SerializationFeature;
import com.fasterxml.jackson.databind.module.SimpleModule;
import org.eclipse.jdt.core.dom.CompilationUnit;

import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.io.OutputStream;

public class ResponseWriter {

    private final OutputStream out;
    private final ObjectMapper mapper;

    public ResponseWriter(final OutputStream out) {
        this.out = out;
        this.mapper = new ObjectMapper();
        mapper.setSerializationInclusion(JsonInclude.Include.NON_NULL);
        mapper.disable(SerializationFeature.FAIL_ON_EMPTY_BEANS);
        mapper.disable(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES);
        mapper.enable(DeserializationFeature.ACCEPT_EMPTY_STRING_AS_NULL_OBJECT);
        SimpleModule module = new SimpleModule();
        module.addSerializer(CompilationUnit.class, new CompilationUnitSerializer());
        mapper.registerModule(module);
    }

    public void write(final Response response) throws IOException {
        // mapper closes the output stream after write, that's why we use an
        // intermediate output stream
        ByteArrayOutputStream out = new ByteArrayOutputStream();
        mapper.writeValue(out, response);
        out.write('\n');
        this.out.write(out.toByteArray());
    }

}
