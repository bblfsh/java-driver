package tech.sourced.langparsers;

import com.fasterxml.jackson.core.JsonGenerationException;
import org.msgpack.core.MessagePack;
import org.msgpack.core.MessagePacker;
import org.msgpack.core.MessageUnpacker;
import org.msgpack.core.buffer.InputStreamBufferInput;
import org.msgpack.core.buffer.OutputStreamBufferOutput;

import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.util.ArrayList;


public class DriverResponse {
    private String status = "ok";
    private ArrayList<String> errors = new ArrayList<String>(0);
    private String driver;
    private String language;
    private String languageVersion;
    private String ast = "empty";

    public DriverResponse(String driver, String language, String languageVersion) {
        this.driver = driver;
        this.language = language;
        this.languageVersion = languageVersion;
    }

    public void makeResponse(String source) {
        try {
            ByteArrayOutputStream baos = new ByteArrayOutputStream();
            EclipseParser parser = new EclipseParser(source, baos, false);
            parser.getAST();

            ast = baos.toString();
        } catch (IOException e) {
            if (e.getClass() == JsonGenerationException.class) {
                errors.add("JsonGenerationException");
                errors.add(e.toString());
                status = "fatal";
            } else {
                errors.add("IOExcepction");
                errors.add(e.toString());
                status = "error";
            }
        }
    }

    public String getAst() {
        return ast;
    }

    public void pack(OutputStreamBufferOutput out) throws IOException {
        MessagePacker packer = MessagePack.newDefaultPacker(out);

        packer.packString(status);
        final int eSize = errors.size();
        packer.packArrayHeader(eSize);
        for (String elem : errors) {
            packer.packString(elem);
        }
        packer.packString(driver);
        packer.packString(language);
        packer.packString(languageVersion);
        packer.packString(ast);
    }

    public static DriverResponse unpack(InputStreamBufferInput in) throws IOException {
        MessageUnpacker unpacker = MessagePack.newDefaultUnpacker(in);

        String status = unpacker.unpackString();
        final int eSize = unpacker.unpackArrayHeader();
        ArrayList<String> errors = new ArrayList<String>();
        for (int i = 0; i < eSize; i++) {
            errors.add(unpacker.unpackString());
        }
        String driver = unpacker.unpackString();
        String language = unpacker.unpackString();
        String languageVersion = unpacker.unpackString();
        String ast = unpacker.unpackString();

        DriverResponse response = new DriverResponse(driver, language, languageVersion);
        response.errors = errors;
        response.ast = ast;
        response.status = status;

        return response;
    }

    public boolean equals(DriverResponse o) {
        return this.status.equals(o.status) && this.ast.equals(o.ast) && this.driver.equals(driver) && this.language.equals(o.language) && this.languageVersion.equals(o.languageVersion);
    }
}
