package tech.sourced.langparsers;

import org.msgpack.core.MessagePack;
import org.msgpack.core.MessagePacker;
import org.msgpack.core.MessageUnpacker;
import org.msgpack.core.buffer.InputStreamBufferInput;
import org.msgpack.core.buffer.OutputStreamBufferOutput;

import java.io.IOException;

public class DriverRequest {

    private String action;
    private String language;
    private String languageVersion;
    private String content;

    public DriverRequest(String action, String language, String languageVersion, String content) {
        this.content = content;
        this.action = action;
        this.language = language;
        this.languageVersion = languageVersion;
    }

    public void pack(OutputStreamBufferOutput out) throws IOException {
        MessagePacker packer = MessagePack.newDefaultPacker(out);

        packer.packString(action);
        packer.packString(language);
        packer.packString(languageVersion);
        packer.packString(content);
    }

    public String getContent() {
        return content;
    }

    public static DriverRequest unpack(InputStreamBufferInput in) throws IOException {
        MessageUnpacker unpacker = MessagePack.newDefaultUnpacker(in);

        String action = unpacker.unpackString();
        String language = unpacker.unpackString();
        String languageVersion = unpacker.unpackString();
        String content = unpacker.unpackString();

        return new DriverRequest(action, language, languageVersion, content);
    }

    public boolean equals(DriverRequest o) {

        return this.action.equals(o.action) && this.language.equals(o.language) && this.languageVersion.equals(o.languageVersion) && this.content.equals(o.content);
    }
}
