package bblfsh;

import org.apache.commons.io.IOUtils;
import org.junit.Test;

import java.io.*;


public class Java8ParserTest {


    //Only to see the JSON output
    @Test
    public void responsePack() throws IOException {
        File file = new File("src/test/resources/helloWorld.java");
        final BufferedReader reader = new BufferedReader(new FileReader(file));
        String source = IOUtils.toString(reader);
        EclipseParser parser = new EclipseParser();
        final RequestResponseMapper mapperGen = new RequestResponseMapper(true);
        final ByteArrayOutputStream baos = new ByteArrayOutputStream();
        final RequestResponseMapper.ResponseMapper responseMapper = mapperGen.getResponseMapper(baos);

        DriverResponse response = new DriverResponse("1.0.0");
        response.setMapper(responseMapper);
        response.makeResponse(parser,source);

        response.pack();

        System.out.println();
    }

    @Test
    public void requestPackUnpack() throws IOException {
        File file = new File("src/test/resources/helloWorld.java");
        final BufferedReader reader = new BufferedReader(new FileReader(file));
        String source = IOUtils.toString(reader);
        final RequestResponseMapper mapperGen = new RequestResponseMapper(true);
        final ByteArrayOutputStream baos = new ByteArrayOutputStream();
        final RequestResponseMapper.RequestMapper responseMapper = mapperGen.getRequestMapper(baos);

        DriverRequest request = new DriverRequest("parse-ast",source);
        request.setMapper(responseMapper);
        request.pack();

        DriverRequest request2 = DriverRequest.unpack(new String(baos.toByteArray()));
        Boolean equals = request.equals(request2);
        assert(equals);
    }
}