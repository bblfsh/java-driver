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

    public CompilationUnitSerializer() {
        this(null);
    }

    public CompilationUnitSerializer(Class<CompilationUnit> t) {
        super(t);
    }

    @Override
    public void serialize(CompilationUnit cu, JsonGenerator jG, SerializerProvider provider) throws IOException {
        jG.writeStartObject();
        jG.writeFieldName("CompilationUnit");

        serializeAll(cu, cu, jG, provider);

        jG.writeEndObject();
    }

    private void serializeAll(CompilationUnit cu, ASTNode node, JsonGenerator jG, SerializerProvider provider) throws IOException {
        List<StructuralPropertyDescriptor> descriptorList = node.structuralPropertiesForType();
        jG.writeStartObject();

        final int Ntype = node.getNodeType();
        String ClassName = node.nodeClassForType(Ntype).getName().substring(25);
        jG.writeFieldName("internalClass");
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
                serializePosition(cu, node, jG);
            }
        }

        if (node == cu) {
            List<Comment> cl = cu.getCommentList();
            if (!cl.isEmpty()) {
                jG.writeFieldName("comments");
                jG.writeStartArray();
                for (Comment c: (List<Comment>) cu.getCommentList()) {
                    if (c.getParent() != null) continue;
                    jG.writeStartObject();
                    final int type = c.getNodeType();
                    String name = c.nodeClassForType(type).getName().substring(25);
                    jG.writeFieldName("internalClass");
                    jG.writeString(name);
                    serializePosition(cu, (ASTNode)c, jG);
                    jG.writeEndObject();
                }
                jG.writeEndArray();
            }
        }

        jG.writeEndObject();
    }

    private void serializeChildList(CompilationUnit cu, List<ASTNode> children, JsonGenerator jG, StructuralPropertyDescriptor descriptor, SerializerProvider provider) throws IOException {
        if (children.size() < 1) {
            return;
        }
        jG.writeFieldName(descriptor.getId());
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
        final int startPosition = node.getStartPosition();
        jG.writeFieldName("startPosition");
        jG.writeNumber(startPosition);
        jG.writeFieldName("startLine");
        jG.writeNumber(cu.getLineNumber(startPosition));
        jG.writeFieldName("startColumn");
        jG.writeNumber(cu.getColumnNumber(startPosition) + 1); // 1-based numbering

        final int endPosition = startPosition + node.getLength();
        jG.writeFieldName("endPosition");
        jG.writeNumber(endPosition);
        jG.writeFieldName("endLine");
        jG.writeNumber(cu.getLineNumber(endPosition));
        jG.writeFieldName("endColumn");
        jG.writeNumber(cu.getColumnNumber(endPosition) + 1); // 1-based numbering
    }
}
