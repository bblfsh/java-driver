package bblfsh;

import org.eclipse.jdt.core.dom.AST;
import org.eclipse.jdt.core.dom.ASTParser;
import org.eclipse.jdt.core.dom.CompilationUnit;

import java.io.IOException;

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
     * Parse the code and generate the JSON in outJ
     *
     * @param source String to parses
     * @return CompilationUnit of the AST
     * @throws IOException if anything related to I/O or Json generation failed
     */
    public CompilationUnit parse(final String source) throws IOException {
        parser.setSource(source.toCharArray());
        return (CompilationUnit) parser.createAST(null);
    }
}
