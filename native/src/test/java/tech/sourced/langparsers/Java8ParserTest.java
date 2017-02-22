package tech.sourced.langparsers;

import org.apache.commons.io.IOUtils;
import org.junit.Test;

import java.io.*;

import org.msgpack.core.buffer.InputStreamBufferInput;
import org.msgpack.core.buffer.OutputStreamBufferOutput;


public class Java8ParserTest {

    //Standard output
    final private PrintStream out = new PrintStream(System.out);

    //dev/null output
    final private PrintStream nullout = new PrintStream(new OutputStream() {
        @Override
        public void write(int i) throws IOException {
        }
    });
    private int nNodes;

    @Test
    public void requestResponse() throws IOException {

        File file = new File("src/test/resources/simpleLambda.java");
        final BufferedReader reader = new BufferedReader(new FileReader(file));
        String source = IOUtils.toString(reader);
        ByteArrayOutputStream baos = new ByteArrayOutputStream();

        //Manual way
        EclipseParser parser = new EclipseParser(source, baos, false);
        nNodes = parser.getAST();
        String ast1 = baos.toString();

        //Petition way
        DriverResponse response = new DriverResponse("1.0.0", "Java", "8");
        response.makeResponse(source);

        String ast2 = response.getAst();

        assert (ast1.equals(ast2));
    }

    @Test
    public void responsePackUnpack() throws IOException {
        String version = "1.0.0";
        String language = "Java";
        String languageVersion = "8";
        DriverResponse response = new DriverResponse(version, language, languageVersion);

        ByteArrayOutputStream baos = new ByteArrayOutputStream();
        OutputStreamBufferOutput out = new OutputStreamBufferOutput(baos);
        response.pack(out);
        out.writeBuffer(8192);
        ByteArrayInputStream bain = new ByteArrayInputStream(baos.toByteArray());
        InputStreamBufferInput in = new InputStreamBufferInput(bain);

        DriverResponse responseUnpacked = DriverResponse.unpack(in);

        boolean equals = responseUnpacked.equals(response);
        assert (equals);
    }

    @Test
    public void requestPackUnpack() throws IOException {
        String action = "parseAST";
        String language = "Java";
        String languageVersion = "8";
        String content = "empty";

        DriverRequest request = new DriverRequest(action, language, languageVersion, content);

        ByteArrayOutputStream baos = new ByteArrayOutputStream();
        OutputStreamBufferOutput out = new OutputStreamBufferOutput(baos);
        request.pack(out);
        out.writeBuffer(8192);
        ByteArrayInputStream bain = new ByteArrayInputStream(baos.toByteArray());
        InputStreamBufferInput in = new InputStreamBufferInput(bain);

        DriverRequest requestUnpacked = DriverRequest.unpack(in);

        boolean equals = requestUnpacked.equals(request);
        assert (equals);
    }
}