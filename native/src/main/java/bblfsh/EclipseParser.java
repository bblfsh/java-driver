package bblfsh;

import org.eclipse.jdt.core.JavaCore;
import org.eclipse.jdt.core.dom.AST;
import org.eclipse.jdt.core.dom.ASTParser;
import org.eclipse.jdt.core.dom.CompilationUnit;

import java.io.IOException;

import java.util.Map;

/**
 * Java AST parser based on Eclipse JDT.
 * <p>
 * This class is not thread-safe.
 * </p>
 */
public class EclipseParser {

    private final ASTParser parser;

    /**
     * Creates a new EclipseParser
     */
    public EclipseParser() {
        this.parser = ASTParser.newParser(AST.JLS8);
        parser.setKind(ASTParser.K_COMPILATION_UNIT);
    }

    /**
     * Parses the given source code text.
     *
     * @param source String to parses
     * @return CompilationUnit of the AST
     * @throws IllegalStateException if parser is not configured properly.
     */
    public CompilationUnit parse(final String source) throws IOException {
        parser.setSource(source.toCharArray());
        Map<String, String> options = JavaCore.getOptions();
        JavaCore.setComplianceOptions(JavaCore.VERSION_1_8, options);
        parser.setCompilerOptions(options);
        return (CompilationUnit) parser.createAST(null);
    }
}
