package tech.sourced.babelfish;

import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;

/**
 *  Class for the java driver request.
 */
public class DriverRequest {

    public String action;
    public String language;
    public String languageVersion;
    public String content;
    private ObjectMapper mapper;
    private JsonGenerator jG;

    /**
     * Creates a new DriverRequest
     *
     * @param action to do with the content
     * @param language languaje of the content always java
     * @param languageVersion version of the lenguaje
     * @param content content of the petition, the code to parse
     */
    public DriverRequest(String action, String language, String languageVersion, String content) {
        this.content = content;
        this.action = action;
        this.language = language;
        this.languageVersion = languageVersion;
    }

    public DriverRequest() {

    }

    /**
     * Deserialize a DriverRequest object from a String given.
     *
     * @param in String to deserialize
     * @return DriverRequest object from the string
     * @throws IOException when it's impossible to Jackson to deserialize
     */
    public static DriverRequest unpack(String in) throws IOException {

        ObjectMapper mapper = new ObjectMapper();
        return mapper.readValue(in, DriverRequest.class);
    }

    /**
     * Set a previously configured Jackson ObjectMapper to driverRequest.
     *
     * @param requestMapper the mapper to set
     */
    @JsonIgnore
    public void setMapper(RequestResponseMapper.RequestMapper requestMapper) {
        mapper = requestMapper.mapper;
        jG = requestMapper.jG;
    }

    /**
     * Serialize DriverRequest in the output given by the requestMapper assigned before.
     *
     * @throws IOException when the write failed or mapper is not assigned
     */
    public void pack() throws IOException {
        if (mapper != null) {
            mapper.writeValue(jG, this);
        } else {
            throw new IOException("Mapper not assigned, use setMapper before packing");
        }
    }

    public boolean equals(DriverRequest o) {

        return this.action.equals(o.action) && this.language.equals(o.language) && this.languageVersion.equals(o.languageVersion) && this.content.equals(o.content);
    }
}
