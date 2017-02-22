package tech.sourced.langparsers;

import com.fasterxml.jackson.core.JsonEncoding;
import com.fasterxml.jackson.core.JsonFactory;
import com.fasterxml.jackson.core.JsonGenerator;

import com.fasterxml.jackson.core.util.DefaultPrettyPrinter;
import com.fasterxml.jackson.databind.DeserializationFeature;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.SerializationFeature;
import com.fasterxml.jackson.databind.SerializerProvider;
import com.fasterxml.jackson.databind.module.SimpleModule;
import com.fasterxml.jackson.databind.ser.std.StdSerializer;

import org.eclipse.jdt.core.dom.*;

import java.util.HashMap;
import java.util.List;

import java.io.*;
import java.util.Map;


public class EclipseParser {

    final private ASTParser parser;
    private int nCount = 0;
    private JsonGenerator jG;
    private ObjectMapper mapper;
    private CompilationUnit cu;

    /**
     * Creates a new EclipseParser
     *
     * @param source String to parse
     * @param outJ   JSON parsed out
     * @throws IOException when file can't be opened or errors in reading/writing
     */
    public EclipseParser(String source, ByteArrayOutputStream outJ, boolean debug) throws IOException {

        this.parser = ASTParser.newParser(AST.JLS8);
        parser.setSource(source.toCharArray());
        parser.setKind(ASTParser.K_COMPILATION_UNIT);
        mapper = new ObjectMapper();

        final JsonFactory jsonF = new JsonFactory();
        jG = jsonF.createGenerator(outJ, JsonEncoding.UTF8);

        if (debug) {
            mapper.enable(SerializationFeature.INDENT_OUTPUT);
            jG.setPrettyPrinter(new DefaultPrettyPrinter());
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

        cu = (CompilationUnit)parser.createAST(null);
        jG.writeStartObject();

        jG.writeFieldName("CompilationUnit");

        mapper.writeValue(jG, cu);

        jG.writeEndObject();

        jG.flush();
        jG.close();
        return nCount;
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
                    jG.writeFieldName("startPosition");
                    final int position = node.getStartPosition();
                    jG.writeString(String.valueOf(position));
                    jG.writeFieldName("line");
                    jG.writeString(String.valueOf(cu.getLineNumber(position)));
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

            final int Ntype = child.getNodeType();
            String ClassName = child.nodeClassForType(Ntype).getName().substring(25);
            jG.writeFieldName(ClassName.concat("/").concat(descriptor.getId()));
            serialize(child, jG, provider);
        }
    }
}
