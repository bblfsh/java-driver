package bblfsh;

import com.fasterxml.jackson.core.JsonEncoding;
import com.fasterxml.jackson.core.JsonFactory;
import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.core.util.DefaultPrettyPrinter;
import com.fasterxml.jackson.databind.DeserializationFeature;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.SerializationFeature;
import com.fasterxml.jackson.databind.module.SimpleModule;
import org.eclipse.jdt.core.dom.CompilationUnit;

import java.io.ByteArrayOutputStream;
import java.io.IOException;

public class RequestResponseMapper {

    final Boolean prettyPrint;

    public RequestResponseMapper(Boolean prettyPrint) {
        this.prettyPrint = prettyPrint;
    }

    public RequestMapper getRequestMapper(ByteArrayOutputStream out) throws IOException {
        return new RequestMapper(prettyPrint, out);
    }

    public ResponseMapper getResponseMapper(ByteArrayOutputStream out) throws IOException {
        return new ResponseMapper(prettyPrint, out);
    }

    public class RequestMapper {

        public final ObjectMapper mapper = new ObjectMapper();
        private final JsonFactory jsonF = new JsonFactory();
        public JsonGenerator jG;

        public RequestMapper(Boolean prettyPrint, ByteArrayOutputStream out) throws IOException {

            jG = jsonF.createGenerator(out, JsonEncoding.UTF8);

            if (prettyPrint) {
                jG.setPrettyPrinter(new DefaultPrettyPrinter());
                mapper.enable(SerializationFeature.INDENT_OUTPUT);
            }
            mapper.disable(SerializationFeature.FAIL_ON_EMPTY_BEANS);
            mapper.disable(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES);
            mapper.enable(DeserializationFeature.ACCEPT_EMPTY_STRING_AS_NULL_OBJECT);
        }
    }

    public class ResponseMapper {

        public final ObjectMapper mapper = new ObjectMapper();
        private final JsonFactory jsonF = new JsonFactory();
        public JsonGenerator jG;

        public ResponseMapper(Boolean prettyPrint, ByteArrayOutputStream out) throws IOException {
            jG = jsonF.createGenerator(out, JsonEncoding.UTF8);

            if (prettyPrint) {
                jG.setPrettyPrinter(new DefaultPrettyPrinter());
                mapper.enable(SerializationFeature.INDENT_OUTPUT);
            }
            mapper.disable(SerializationFeature.FAIL_ON_EMPTY_BEANS);
            mapper.disable(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES);
            mapper.enable(DeserializationFeature.ACCEPT_EMPTY_STRING_AS_NULL_OBJECT);
            SimpleModule module = new SimpleModule();
            module.addSerializer(CompilationUnit.class, new CompilationUnitSerializer());
            mapper.registerModule(module);
        }
    }
}
