package bblfsh;

import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;

/**
 * Class for the java driver request.
 */
public class Request {
    public String action;
    public String content;
    private ObjectMapper mapper;
    private JsonGenerator jG;

    /**
     * Creates a new Request
     *
     * @param action  to do with the content
     * @param content content of the petition, the code to parse
     */
    public Request(String action, String content) {
        this.content = content;
        this.action = action;
    }

    public Request() {

    }

    /**
     * Deserialize a Request object from a String given.
     *
     * @param in String to deserialize
     * @return Request object from the string
     * @throws IOException when it's impossible for Jackson to deserialize
     */
    public static Request unpack(String in) throws IOException {
        ObjectMapper mapper = new ObjectMapper();
        return mapper.readValue(in, Request.class);
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
     * Serialize Request in the output given by the requestMapper assigned before.
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

    public boolean equals(Request o) {
        return this.action.equals(o.action) && this.content.equals(o.content);
    }
}
