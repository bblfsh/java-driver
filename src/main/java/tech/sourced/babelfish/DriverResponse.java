package tech.sourced.babelfish;

import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.eclipse.jdt.core.dom.CompilationUnit;

import java.io.IOException;
import java.util.ArrayList;

public class DriverResponse {
    public String status = "ok";
    public ArrayList<String> errors = new ArrayList<String>(0);
    final public String driver;
    final public String language;
    final public String languageVersion;
    @JsonProperty("AST")
    public CompilationUnit cu;
    private ObjectMapper mapper;
    private JsonGenerator jG;

    public DriverResponse(String driver, String language, String languageVersion) throws IOException {
        this.driver = driver;
        this.language = language;
        this.languageVersion = languageVersion;
    }

    @JsonIgnore
    public void setMapper(RequestResponseMapper.ResponseMapper responseMapper) {
        mapper = responseMapper.mapper;
        jG = responseMapper.jG;
    }

    public void makeResponse(EclipseParser parser, String source) {
        try {
            cu = parser.getAST(source);
        } catch (IOException e) {
            errors.add("IOException");
            errors.add(e.toString());
            status = "error";
        }
    }

    public void pack() throws IOException {
        if (mapper != null) {
            mapper.writeValue(jG, this);
        } else {
            throw new IOException("Mapper not assigned, use setMapper before packing");
        }
    }

    public boolean equals(DriverResponse o) {
        return this.status.equals(o.status) && this.cu == o.cu && this.driver.equals(driver) && this.language.equals(o.language) && this.languageVersion.equals(o.languageVersion);
    }
}

