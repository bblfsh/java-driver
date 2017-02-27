package tech.sourced.babelfish;

import org.eclipse.jdt.core.dom.AST;
import org.eclipse.jdt.core.dom.ASTParser;
import org.eclipse.jdt.core.dom.CompilationUnit;

import java.io.IOException;

public class EclipseParser {

    final private ASTParser parser;

    /**
     * Creates a new EclipseParser
     * <p>
     * EclipseParser is not threadsafe
     *
     * @throws IOException when file can't be opened or errors in reading/writing
     */
    public EclipseParser() throws IOException {
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
    public CompilationUnit getAST(String source) throws IOException {
        parser.setSource(source.toCharArray());
        return (CompilationUnit) parser.createAST(null);

    }
}
