package tech.sourced.langparsers;

import com.fasterxml.jackson.core.JsonFactory;
import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.core.util.DefaultPrettyPrinter;
import com.fasterxml.jackson.databind.DeserializationFeature;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.SerializationFeature;
import com.fasterxml.jackson.databind.SerializerProvider;
import com.fasterxml.jackson.databind.module.SimpleModule;
import com.fasterxml.jackson.databind.ser.std.StdSerializer;
import org.apache.commons.io.IOUtils;
import org.eclipse.jdt.core.dom.AST;
import org.eclipse.jdt.core.dom.ASTNode;
import org.eclipse.jdt.core.dom.ASTParser;
import org.eclipse.jdt.core.dom.StructuralPropertyDescriptor;

import java.util.List;

import java.io.*;

public class EclipseParser {

    final private ASTParser parser;
    private int nCount = 0;
    private JsonGenerator jG;
    final private ObjectMapper mapper = new ObjectMapper();

    /**
     * Creates a new EclipseParser
     *
     * @param sourceFile String of source file to read
     * @param outJ       JSON parsed out
     * @throws IOException when file can't be opened or errors in reading/writing
     */
    public EclipseParser(String sourceFile, PrintStream outJ, boolean prettyprint) throws IOException {

        File file = new File(sourceFile);
        final BufferedReader reader = new BufferedReader(new FileReader(file));
        char[] source = IOUtils.toCharArray(reader);
        reader.close();
        this.parser = ASTParser.newParser(AST.JLS8);
        parser.setSource(source);
        parser.setKind(ASTParser.K_COMPILATION_UNIT);

        final JsonFactory jsonF = new JsonFactory();
        jG = jsonF.createGenerator(outJ);
        if (prettyprint) {
            jG.setPrettyPrinter(new DefaultPrettyPrinter());
            mapper.enable(SerializationFeature.INDENT_OUTPUT);
        }
        mapper.disable(SerializationFeature.FAIL_ON_EMPTY_BEANS);
        mapper.disable(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES);
        mapper.enable(DeserializationFeature.ACCEPT_EMPTY_STRING_AS_NULL_OBJECT);
        SimpleModule module = new SimpleModule();
        module.addSerializer(ASTNode.class, new NodeSerializer());
        mapper.registerModule(module);
    }

    /**
     * Parse the code and generate the JSON in outJ
     *
     * @return nNodes visited while the serialization process
     * @throws IOException if anything related to I/O or Json generation failed
     */
    public int getAST() throws IOException {

        final ASTNode cu = parser.createAST(null);

        setHeader();
        jG.writeFieldName("CompilationUnit");

        mapper.writeValue(jG, cu);

        setEnd();
        return nCount;
    }

    private void setHeader() throws IOException {
        jG.writeStartObject();
        jG.writeFieldName("driver");
        jG.writeString("Java8:1.0.0");
        jG.writeFieldName("ast");
        jG.writeStartObject();

    }

    private void setEnd() throws IOException {
        jG.writeEndObject();
        jG.writeEndObject();
        jG.flush();
        jG.close();
    }

    /**
     * Customized ASTNode serializer
     */
    private class NodeSerializer extends StdSerializer<ASTNode> {

        private NodeSerializer() {
            this(null);
        }

        private NodeSerializer(Class<ASTNode> t) {
            super(t);
        }

        @Override
        public void serialize(ASTNode node, JsonGenerator jG, SerializerProvider provider) throws IOException {
            List<StructuralPropertyDescriptor> descriptorList = node.structuralPropertiesForType();
            nCount++;
            jG.writeStartObject();

            for (StructuralPropertyDescriptor descriptor : descriptorList) {
                Object child = node.getStructuralProperty(descriptor);
                if (child instanceof List) {
                    serializeChildList((List<ASTNode>) child, descriptor, provider);
                } else if (child instanceof ASTNode) {
                    serializeChild((ASTNode) child, descriptor, provider);
                } else if (child != null) {
                    jG.writeFieldName(descriptor.getId());
                    jG.writeString(child.toString());
                }
            }
            jG.writeEndObject();
        }

        private void serializeChildList(List<ASTNode> children, StructuralPropertyDescriptor descriptor, SerializerProvider provider) throws IOException {
            if (children.size() < 1) {
                return;
            }

            jG.writeFieldName(descriptor.getId());
            jG.writeStartArray();

            for (ASTNode node : children) {
                serialize(node, jG, provider);
            }
            jG.writeEndArray();
        }

        private void serializeChild(ASTNode child, StructuralPropertyDescriptor descriptor, SerializerProvider provider) throws IOException {
            /*Class cClass = child.getClass();
            jG.writeFieldName("Esto es lo nuevo");
            mapper.writeValue(jG,cClass.cast(child));*/


            jG.writeFieldName(descriptor.getId());
            serialize(child, jG, provider);
        }
    }
}