package bblfsh;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.SerializerProvider;
import com.fasterxml.jackson.databind.ser.std.StdSerializer;
import org.eclipse.jdt.core.dom.ASTNode;
import org.eclipse.jdt.core.dom.Comment;
import org.eclipse.jdt.core.dom.CompilationUnit;
import org.eclipse.jdt.core.dom.StructuralPropertyDescriptor;

import java.io.IOException;
import java.util.List;

/**
 * Custom Jackson serializer for jdt.core.dom.CompilationUnit
 */
public class CompilationUnitSerializer extends StdSerializer<CompilationUnit> {

    private String[] contentLines;

    public CompilationUnitSerializer() {
        this(null);
    }

    public CompilationUnitSerializer(Class<CompilationUnit> t) {
        super(t);
    }

    public void setContent(String content) {
        this.contentLines = content.split("\\r?\\n");
    }

    @Override
    public void serialize(CompilationUnit cu, JsonGenerator jG, SerializerProvider provider) throws IOException {
        jG.writeStartObject();
        jG.writeFieldName("CompilationUnit");

        serializeAll(cu, cu, jG, provider);

        jG.writeEndObject();
    }

    private void serializeAll(CompilationUnit cu, ASTNode node, JsonGenerator jG,
                              SerializerProvider provider) throws IOException {
        List<StructuralPropertyDescriptor> descriptorList = node.structuralPropertiesForType();
        jG.writeStartObject();

        final int Ntype = node.getNodeType();
        String ClassName = node.nodeClassForType(Ntype).getName().substring(25);
        jG.writeFieldName("@type");
        jG.writeString(ClassName);

        for (StructuralPropertyDescriptor descriptor : descriptorList) {
            Object child = node.getStructuralProperty(descriptor);
            if (child instanceof List) {
                serializeChildList(cu, (List<ASTNode>) child, jG, descriptor, provider);
            } else if (child instanceof ASTNode) {
                serializeChild(cu, (ASTNode) child, jG, descriptor, provider);
            } else if (child != null) {
                jG.writeFieldName(descriptor.getId());
                jG.writeString(child.toString());
            } else {
                jG.writeFieldName(descriptor.getId());
                jG.writeNull();
            }
        }
        serializePosition(cu, node, jG);

        if (node == cu) {
            List<Comment> cl = cu.getCommentList();
            if (!cl.isEmpty()) {
                jG.writeFieldName("comments");
                jG.writeStartArray();

                for (Comment c: (List<Comment>) cu.getCommentList()) {
                    if (c.getParent() != null)
                        continue;

                    CommentVisitor visitor = new CommentVisitor(cu, contentLines);
                    c.accept(visitor);

                    jG.writeStartObject();
                    final int type = c.getNodeType();
                    jG.writeFieldName("@type");
                    String name = c.nodeClassForType(type).getName().substring(25);
                    jG.writeString(name);
                    jG.writeFieldName("text");
                    jG.writeString(visitor.getCommentText());
                    serializePosition(cu, (ASTNode)c, jG);
                    jG.writeEndObject();
                }
                jG.writeEndArray();
            }
        }

        jG.writeEndObject();
    }

    private void serializeChildList(CompilationUnit cu, List<ASTNode> children,
                                    JsonGenerator jG,
                                    StructuralPropertyDescriptor descriptor,
                                    SerializerProvider provider) throws IOException {
        jG.writeFieldName(descriptor.getId());
        if (children.size() < 1) {
            jG.writeNull();
            return;
        }
        jG.writeStartArray();
        for (ASTNode node : children) {
            serializeAll(cu, node, jG, provider);
        }
        jG.writeEndArray();
    }

    private void serializeChild(CompilationUnit cu, ASTNode child, JsonGenerator jG, StructuralPropertyDescriptor descriptor, SerializerProvider provider) throws IOException {
        jG.writeFieldName(descriptor.getId());
        serializeAll(cu, child, jG, provider);
    }

    private void serializePosition(CompilationUnit cu, ASTNode node, JsonGenerator jG) throws IOException {
        jG.writeFieldName("@pos");
        jG.writeStartObject();

        jG.writeFieldName("@type");
        jG.writeString("uast:Positions");

        final int startPosition = node.getStartPosition();
        serializePositionFields(cu, jG, startPosition, "start");

        final int endPosition = startPosition + node.getLength();
        serializePositionFields(cu, jG, endPosition, "end");

        jG.writeEndObject();
    }

    private void serializePositionFields(CompilationUnit cu, JsonGenerator jG, int pos, String name) throws IOException {
        jG.writeFieldName(name);
        jG.writeStartObject();

        jG.writeFieldName("@type");
        jG.writeString("uast:Position");

        jG.writeFieldName("offset");
        jG.writeNumber(pos);
        final int line = cu.getLineNumber(pos);
        if (line >= 0) {
            jG.writeFieldName("line");
            jG.writeNumber(line);
        }
        jG.writeFieldName("col");
        jG.writeNumber(cu.getColumnNumber(pos) + 1); // 1-based numbering

        jG.writeEndObject();
    }
}
