package tech.sourced.babelfish;

import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;

public class DriverRequest {

    public String action;
    public String language;
    public String languageVersion;
    public String content;
    private ObjectMapper mapper;
    private JsonGenerator jG;

    public DriverRequest(String action, String language, String languageVersion, String content) {
        this.content = content;
        this.action = action;
        this.language = language;
        this.languageVersion = languageVersion;
    }

    public DriverRequest() {

    }

    public static DriverRequest unpack(String in) throws IOException {

        ObjectMapper mapper = new ObjectMapper();
        return mapper.readValue(in, DriverRequest.class);
    }

    @JsonIgnore
    public void setMapper(RequestResponseMapper.RequestMapper requestMapper) {
        mapper = requestMapper.mapper;
        jG = requestMapper.jG;
    }

    //Not needed only used on testing
    public void pack() throws IOException {
        if (mapper != null) {
            mapper.writeValue(jG, this);
        } else {
            throw new IOException("Mapper not assigned, use setMapper before packing");
        }
    }

    public String getContent() {
        return content;
    }

    public boolean equals(DriverRequest o) {

        return this.action.equals(o.action) && this.language.equals(o.language) && this.languageVersion.equals(o.languageVersion) && this.content.equals(o.content);
    }
}
