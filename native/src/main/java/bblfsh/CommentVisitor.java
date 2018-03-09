package bblfsh;

import org.eclipse.jdt.core.dom.ASTNode;
import org.eclipse.jdt.core.dom.ASTVisitor;
import org.eclipse.jdt.core.dom.BlockComment;
import org.eclipse.jdt.core.dom.Comment;
import org.eclipse.jdt.core.dom.CompilationUnit;
import org.eclipse.jdt.core.dom.Javadoc;
import org.eclipse.jdt.core.dom.LineComment;

public class CommentVisitor extends ASTVisitor {

    CompilationUnit compilationUnit;

    private String[] source;
    private String commentText;

    public CommentVisitor(CompilationUnit compilationUnit, String[] source) {

        super();
        this.compilationUnit = compilationUnit;
        this.source = source;
    }

    public String getCommentText() {
        return commentText;
    }

    public boolean visit(LineComment node) {

        int startLineNumber = compilationUnit.getLineNumber(node.getStartPosition()) - 1;
        int startLineColumn = compilationUnit.getColumnNumber(node.getStartPosition());
        commentText = source[startLineNumber].substring(startLineColumn).trim();

        return true;
    }

    public boolean visit(Javadoc node) {
        return visitBlock(node);
    }

    public boolean visit(BlockComment node) {
        return visitBlock(node);
    }

    private boolean visitBlock(Comment node) {

        int startLineNumber = compilationUnit.getLineNumber(node.getStartPosition()) - 1;
        int startLineColumn = compilationUnit.getColumnNumber(node.getStartPosition());

        int endLineNumber = compilationUnit.getLineNumber(node.getStartPosition() + node.getLength()) - 1;
        int endLineColumn = compilationUnit.getColumnNumber(node.getStartPosition() + node.getLength());

        StringBuffer blockComment = new StringBuffer();

        for (int lineCount = startLineNumber ; lineCount <= endLineNumber; lineCount++) {

            int startCol = lineCount == startLineNumber ? startLineColumn : 0;
            int endCol = lineCount == endLineNumber ? endLineColumn : source[lineCount].length();
            String blockCommentLine = source[lineCount].substring(startLineColumn, endCol).trim();

            blockComment.append(blockCommentLine);
            if (lineCount != endLineNumber) {
                blockComment.append("\n");
            }
        }

        commentText = blockComment.toString();

        return true;
    }
}
