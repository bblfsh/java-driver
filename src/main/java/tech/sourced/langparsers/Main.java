package tech.sourced.langparsers;

import org.msgpack.core.buffer.InputStreamBufferInput;
import org.msgpack.core.buffer.OutputStreamBufferOutput;

import java.io.IOException;

public class Main {

    public static void main(String args[]) throws IOException {

        InputStreamBufferInput in = new InputStreamBufferInput(System.in);
        OutputStreamBufferOutput out = new OutputStreamBufferOutput(System.out);

        DriverRequest request = DriverRequest.unpack(in);
        DriverResponse response = new DriverResponse("1.0.0", "Java", "8");
        response.makeResponse(request.getContent());

        response.pack(out);
    }
}
