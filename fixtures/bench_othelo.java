/*
 * To change this template, choose Tools | Templates
 * and open the template in the editor.
 */

/**
 *
 * @author
 */
import java.awt.*;
import java.util.*;

public class othelo {

    private square otheloTable[][] = new square[8][8];
    private int blackcircles, whitecircles;
    private char player;
    private int evaluation;
    private Move lastMove;
    private char previous = 'W';
    public char now = 'B';
    public Runtime r;
    public int depth;

    public othelo(int x, int y, int w, int h, int depth) {
        int X = x;
        for (int j = 0; j < 8; j++) {
            for (int i = 0; i < 8; i++) {
                otheloTable[j][i] = new square(x, y, w, h, i, j);
                x = x + w;
            }
            y = y + h;
            x = X;
            r = Runtime.getRuntime();
        }
        lastMove = new Move();
        setRules('B');
        this.depth = depth;
        r.gc();
    }

    public othelo(othelo oth) {

        Runtime r = Runtime.getRuntime();
        for (int j = 0; j < 8; j++) {
            for (int i = 0; i < 8; i++) {
                Vector<Point> v = new Vector<Point>();
                this.otheloTable[j][i] = new square(oth.getSquare(j, i).getX(), oth.getSquare(j, i).getY(), oth.getSquare(j, i).getW(), oth.getSquare(j, i).getH(), i, j);
                this.otheloTable[j][i].setAdd_circle(oth.getSquare(j, i).getAdd_circle());
                this.otheloTable[j][i].setCircle(oth.getSquare(j, i).getCircle());
                this.otheloTable[j][i].setLabel(oth.getSquare(j, i).getLabel());
                for (int k = 0; k < oth.getSquare(j, i).getChangesSize(); k++) {
                    v.add(oth.getSquare(j, i).getChange(k));
                    this.otheloTable[j][i].addChange(v);
                }

            }
        }
        this.lastMove = new Move(oth.getLastMove().getX(), oth.getLastMove().getY(), oth.getLastMove().getValue());
        r.gc();
    }

    public square getSquare(int j, int i) {
        return otheloTable[j][i];
    }

    public void setBlackcircles() {
        int counter = 0;
        for (int j = 0; j < 8; j++) {
            for (int i = 0; i < 8; i++) {
                if (otheloTable[j][i].getCircle() == 'B') {
                    counter++;
                }
            }
        }
        blackcircles = counter;
    }

    public void setWhitecircles() {
        int counter = 0;
        for (int j = 0; j < 8; j++) {
            for (int i = 0; i < 8; i++) {
                if (otheloTable[j][i].getCircle() == 'W') {
                    counter++;
                }
            }
        }
        whitecircles = counter;
    }

    public int getBlackcircles() {
        return this.blackcircles;
    }

    public int getWhitecircles() {
        return this.whitecircles;
    }
//-----------------------------------------------------------------------------------------------------------------------

    public ArrayList<othelo> getChildren(char player) {


        ArrayList<othelo> children = new ArrayList<othelo>();
        for (int y = 0; y < 8; y++) {
            for (int x = 0; x < 8; x++) {
                if (this.otheloTable[y][x].getAdd_circle()) {
                    othelo child = new othelo(this);
                    child.makeMove(y, x, player);
                    child.lastMove = new Move(y, x);
                    if (player == 'B') {
                        child.setRules('W');
                    } else {
                        child.setRules('B');
                    }
                    children.add(child);

                }
            }
        }

        return children;
    }

    public void setRules(char player) {   //PLAYER OR XROMA
        for (int j = 0; j < 8; j++) {
            for (int i = 0; i < 8; i++) {
                if (!(otheloTable[j][i].getCircle() == 'N')) {
                    otheloTable[j][i].setAdd_circle(false);//an yparxei kuklos de mporei na valei allo
                } else {
                    otheloTable[j][i].setAdd_circle(lookVertical(new Point(j, i), player) | lookHorizontal(new Point(j, i), player) | lookDiagonal(new Point(j, i), player));
                }
                //      System.out.print(j + ", " + i + " :" + otheloTable[j][i].getAdd_circle() + " ");
            }
            //    System.out.println("\n");
        }

    }

    private boolean lookHorizontal(Point square1, char player) {

        return (lookRight(square1, player) | lookLeft(square1, player));
    }

    private boolean lookRight(Point square1, char player) {
        int x = (int) square1.getX();
        int y = (int) square1.getY();
        int next = y + 1;       //to epomeno koutaki stin dieuthinsi pou koitame
        Vector<Point> v = new Vector<Point>(); //tha apothikeuoume tis allages poy tha kanei
        boolean add = false;
        if (y == 7) {
            return false;
        }
        if (player == 'B') {
            if ((next <= 7) && (otheloTable[x][next].getCircle() == 'B')) {
                return false;
            }
            if ((next <= 7) && (otheloTable[x][next].getCircle() == 'W')) {
                v.add(new Point(x, next));
                for (int i = (++next); i <= 7; i++) {
                    if (add) {
                        break;
                    }
                    if (otheloTable[x][i].getCircle() == 'W') {
                        v.add(new Point(x, i));
                    } else if (otheloTable[x][i].getCircle() == 'B') {
                        add = true;
                    } else {
                        return false;
                    }
                }
                if (add) {
                    otheloTable[x][y].addChange(v);
                    return true;
                }
                return false;
            } else {
                return false; //diladi an deksia exei B i N
            }
        }
        if (player == 'W') {
            if ((next <= 7) && (otheloTable[next][y].getCircle() == 'W')) {
                return false;
            }
            if ((next <= 7) && (otheloTable[x][next].getCircle() == 'B')) {
                v.add(new Point(x, next));
                for (int i = (++next); i <= 7; i++) {
                    if (add) {
                        break;
                    }
                    if (otheloTable[x][i].getCircle() == 'B') {
                        v.add(new Point(x, i));
                    } else if (otheloTable[x][i].getCircle() == 'W') {
                        add = true;
                    } else {
                        return false;
                    }
                }
                if (add) {
                    otheloTable[x][y].addChange(v);
                    return true;
                }
                return false;
            } else {
                return false; //diladi an deksia exei W i N
            }
        }
        return false;
    }

    private boolean lookLeft(Point square1, char player) {
        int x = (int) square1.getX();
        int y = (int) square1.getY();
        int previous = y - 1;       //to proigoumeno koutaki stin dieuthinsi pou koitame
        Vector<Point> v = new Vector<Point>(); //tha apothikeuoume tis allages poy tha kanei
        boolean add = false;
        if (y == 0) {
            return false;
        }
        if (player == 'B') {
            if ((previous >= 0) && (otheloTable[x][previous].getCircle() == 'B')) {
                return false;
            }
            if ((previous >= 0) && (otheloTable[x][previous].getCircle() == 'W')) {
                v.add(new Point(x, previous));
                for (int i = (--previous); i >= 0; i--) {
                    if (add) {
                        break;
                    }
                    if (otheloTable[x][i].getCircle() == 'W') {
                        v.add(new Point(x, i));
                    } else if (otheloTable[x][i].getCircle() == 'B') {
                        add = true;
                    } else {
                        return false;
                    }
                }
                if (add) {
                    otheloTable[x][y].addChange(v);
                    return true;
                }
                return false;
            } else {
                return false; //diladi an aristera exei B i N
            }
        }
        if (player == 'W') {
            if ((previous >= 0) && (otheloTable[x][previous].getCircle() == 'W')) {
                return false;
            }
            if ((previous >= 0) && (otheloTable[x][previous].getCircle() == 'B')) {
                v.add(new Point(x, previous));
                for (int i = (--previous); i >= 0; i--) {
                    if (add) {
                        break;
                    }
                    if (otheloTable[x][i].getCircle() == 'B') {
                        v.add(new Point(x, i));
                    } else if (otheloTable[x][i].getCircle() == 'W') {
                        add = true;
                    } else {
                        return false;
                    }
                }
                if (add) {
                    otheloTable[x][y].addChange(v);
                    return true;
                }
                return false;
            } else {
                return false; //diladi an aristera exei W i N
            }
        }
        return false;
    }

    private boolean lookVertical(Point square1, char player) {
        return (lookDown(square1, player)) | (lookUp(square1, player));

    }

    private boolean lookDown(Point square1, char player) {
        int x = (int) square1.getX();
        int y = (int) square1.getY();
        int next = x + 1;       //to epomeno koutaki stin dieuthinsi pou koitame
        Vector<Point> v = new Vector<Point>(); //tha apothikeuoume tis allages poy tha kanei
        boolean add = false;
        if (x == 7) {
            return false;
        }
        if (player == 'B') {
            if ((next <= 7) && (otheloTable[next][y].getCircle() == 'B')) {
                return false;
            }
            if ((next <= 7) && (otheloTable[next][y].getCircle() == 'W')) {
                v.add(new Point(next, y));
                for (int i = (++next); i <= 7; i++) {
                    if (add) {
                        break;
                    }
                    if (otheloTable[i][y].getCircle() == 'W') {
                        v.add(new Point(i, y));
                    } else if (otheloTable[i][y].getCircle() == 'B') {
                        add = true;
                    } else {
                        return false;
                    }

                }
                if (add) {
                    otheloTable[x][y].addChange(v);
                    return true;
                }
                return false;
            }
            return false; //diladi an pano exei B i N

        }
        if (player == 'W') {
            if ((next <= 7) && (otheloTable[next][y].getCircle() == 'W')) {
                return false;
            }
            if ((next <= 7) && (otheloTable[next][y].getCircle() == 'B')) {
                v.add(new Point(next, y));
                for (int i = (++next); i <= 7; i++) {
                    if (add) {
                        break;
                    }
                    if (otheloTable[i][y].getCircle() == 'B') {
                        v.add(new Point(i, y));
                    } else if (otheloTable[i][y].getCircle() == 'W') {
                        add = true;
                    } else {
                        return false;
                    }
                }
                if (add) {
                    otheloTable[x][y].addChange(v);
                    return true;
                }
                return false;
            } else {
                return false; //diladi an pano exei W i N
            }

        }
        return false;
    }

    private boolean lookUp(Point square1, char player) {
        int x = (int) square1.getX();
        int y = (int) square1.getY();
        int previous = x - 1;       //to epomeno koutaki stin dieuthinsi pou koitame
        Vector<Point> v = new Vector<Point>(); //tha apothikeuoume tis allages poy tha kanei
        boolean add = false;
        if (x == 0) {
            return false;
        }
        if (player == 'B') {
            if ((previous <= 7) && (otheloTable[previous][y].getCircle() == 'B')) {
                return false;
            }
            if ((previous >= 0) && (otheloTable[previous][y].getCircle() == 'W')) {
                v.add(new Point(previous, y));
                for (int i = (--previous); i >= 0; i--) {
                    if (add) {
                        break;
                    }
                    if (otheloTable[i][y].getCircle() == 'W') {
                        v.add(new Point(i, y));
                    } else if (otheloTable[i][y].getCircle() == 'B') {
                        add = true;
                    } else {
                        return false;
                    }
                }
                if (add) {
                    otheloTable[x][y].addChange(v);
                    return true;
                }
                return false;
            } else {
                return false; //diladi an kato exei B i N
            }
        }
        if (player == 'W') {
            if ((previous <= 7) && (otheloTable[previous][y].getCircle() == 'W')) {
                return false;
            }
            if ((previous >= 0) && (otheloTable[previous][y].getCircle() == 'B')) {
                v.add(new Point(previous, y));
                for (int i = (--previous); i >= 0; i--) {
                    if (add) {
                        break;
                    }
                    if (otheloTable[i][y].getCircle() == 'B') {
                        v.add(new Point(i, y));
                    } else if (otheloTable[i][y].getCircle() == 'W') {
                        add = true;
                    } else {
                        return false;
                    }
                }
                if (add) {
                    otheloTable[x][y].addChange(v);
                    return true;
                }
                return false;
            } else {
                return false; //diladi an kato exei W i N
            }
        }
        return false;
    }

    private boolean lookDiagonal(Point square1, char player) {
        return (lookDiagonal1Up(square1, player) | lookDiagonal1Down(square1, player) | lookDiagonal2Up(square1, player) | lookDiagonal2Down(square1, player));
    }

    private boolean lookDiagonal1Down(Point square1, char player) {
        int x = (int) square1.getX();
        int y = (int) square1.getY();
        int nextX = x + 1;       //to epomeno koutaki stin dieuthinsi pou koitame
        int nextY = y + 1;
        Vector<Point> v = new Vector<Point>(); //tha apothikeuoume tis allages poy tha kanei
        boolean add = false;
        if (x == 7) {
            return false;
        }
        if (y == 7) {
            return false;
        }
        if (player == 'B') {
            if ((nextX <= 7) && (nextY <= 7) && (otheloTable[nextX][nextY].getCircle() == 'B')) {
                return false;
            }
            if ((nextX <= 7) && (nextY <= 7) && (otheloTable[nextX][nextY].getCircle() == 'W')) {
                v.add(new Point(nextX, nextY));
                for (int i = (++nextX); i <= 7; i++) {
                    if (add) {
                        break;
                    }
                    nextY++;
                    if ((nextY <= 7) && (otheloTable[i][nextY].getCircle() == 'W')) {
                        v.add(new Point(i, nextY));
                    } else if ((nextY <= 7) && (otheloTable[i][nextY].getCircle() == 'B')) {
                        add = true;
                    } else {
                        return false;
                    }
                }
                if (add) {
                    otheloTable[x][y].addChange(v);
                    return true;
                }
                return false;
            } else {
                return false; //diladi an pano exei B i N
            }
        }
        if (player == 'W') {
            if ((nextX <= 7) && (nextY <= 7) && (otheloTable[nextX][nextY].getCircle() == 'W')) {
                return false;
            }
            if ((nextX <= 7) && (nextY <= 7) && (otheloTable[nextX][nextY].getCircle() == 'B')) {
                v.add(new Point(nextX, nextY));
                for (int i = (++nextX); i <= 7; i++) {
                    if (add) {
                        break;
                    }
                    nextY++;
                    if ((nextY <= 7) && (otheloTable[i][nextY].getCircle() == 'B')) {
                        v.add(new Point(i, nextY));
                    } else if ((nextY <= 7) && (otheloTable[i][nextY].getCircle() == 'W')) {
                        add = true;
                    } else {
                        return false;
                    }
                }
                if (add) {
                    otheloTable[x][y].addChange(v);
                    return true;
                }
                return false;
            } else {
                return false; //diladi an pano exei B i N
            }
        }
        return false;
    }

    private boolean lookDiagonal1Up(Point square1, char player) {
        int x = (int) square1.getX();
        int y = (int) square1.getY();
        int previousX = x - 1;       //to epomeno koutaki stin dieuthinsi pou koitame
        int previousY = y - 1;
        Vector<Point> v = new Vector<Point>(); //tha apothikeuoume tis allages poy tha kanei
        boolean add = false;
        if (x == 0) {
            return false;
        }
        if (y == 0) {
            return false;
        }
        if (player == 'B') {
            if ((previousX >= 0) && (previousY >= 0) && (otheloTable[previousX][previousY].getCircle() == 'B')) {
                return false;
            }
            if ((previousX >= 0) && (previousY >= 0) && (otheloTable[previousX][previousY].getCircle() == 'W')) {
                v.add(new Point(previousX, previousY));
                for (int i = (--previousX); i >= 0; i--) {
                    if (add) {
                        break;
                    }
                    previousY--;
                    if ((previousY >= 0) && (otheloTable[i][previousY].getCircle() == 'W')) {
                        v.add(new Point(i, previousY));
                    } else if ((previousY >= 0) && (otheloTable[i][previousY].getCircle() == 'B')) {
                        add = true;
                    } else {
                        return false;
                    }
                }
                if (add) {
                    otheloTable[x][y].addChange(v);
                    return true;
                }
                return false;
            } else {
                return false; //diladi an kato exei B i N
            }
        }
        if (player == 'W') {
            if ((previousX >= 0) && (previousY >= 0) && (otheloTable[previousX][previousY].getCircle() == 'W')) {
                return false;
            }
            if ((previousX >= 0) && (previousY >= 0) && (otheloTable[previousX][previousY].getCircle() == 'B')) {
                v.add(new Point(previousX, previousY));
                for (int i = (--previousX); i >= 0; i--) {
                    if (add) {
                        break;
                    }
                    previousY--;
                    if ((previousY >= 0) && (otheloTable[i][previousY].getCircle() == 'B')) {
                        v.add(new Point(i, previousY));
                    } else if ((previousY >= 0) && (otheloTable[i][previousY].getCircle() == 'W')) {
                        add = true;
                    } else {
                        return false;
                    }
                }
                if (add) {
                    otheloTable[x][y].addChange(v);
                    return true;
                }
                return false;
            } else {
                return false; //diladi an kato exei B i N
            }
        }
        return false;
    }

    private boolean lookDiagonal2Down(Point square1, char player) {
        int x = (int) square1.getX();
        int y = (int) square1.getY();
        int nextX = x + 1;       //to epomeno koutaki stin dieuthinsi pou koitame
        int nextY = y - 1;
        Vector<Point> v = new Vector<Point>(); //tha apothikeuoume tis allages poy tha kanei
        boolean add = false;
        if (x == 7) {
            return false;
        }
        if (y == 0) {
            return false;
        }
        if (player == 'B') {
            if ((nextX <= 7) && (nextY >= 0) && (otheloTable[nextX][nextY].getCircle() == 'B')) {
                return false;
            }
            if ((nextX <= 7) && (nextY >= 0) && (otheloTable[nextX][nextY].getCircle() == 'W')) {
                v.add(new Point(nextX, nextY));
                for (int i = (++nextX); i <= 7; i++) {
                    if (add) {
                        break;
                    }
                    nextY--;
                    if ((nextY >= 0) && (otheloTable[i][nextY].getCircle() == 'W')) {
                        v.add(new Point(i, nextY));
                    } else if ((nextY >= 0) && (otheloTable[i][nextY].getCircle() == 'B')) {
                        add = true;
                    } else {
                        return false;
                    }
                }
                if (add) {
                    otheloTable[x][y].addChange(v);
                    return true;
                }
                return false;
            } else {
                return false; //diladi an pano exei B i N
            }
        }
        if (player == 'W') {
            if ((nextX <= 7) && (nextY >= 0) && (otheloTable[nextX][nextY].getCircle() == 'W')) {
                return false;
            }
            if ((nextX <= 7) && (nextY >= 0) && (otheloTable[nextX][nextY].getCircle() == 'B')) {
                v.add(new Point(nextX, nextY));
                for (int i = (++nextX); i <= 7; i++) {
                    if (add) {
                        break;
                    }
                    nextY--;
                    if ((nextY >= 0) && (otheloTable[i][nextY].getCircle() == 'B')) {
                        v.add(new Point(i, nextY));
                    } else if ((nextY >= 0) && (otheloTable[i][nextY].getCircle() == 'W')) {
                        add = true;
                    } else {
                        return false;
                    }
                }
                if (add) {
                    otheloTable[x][y].addChange(v);
                    return true;
                }
                return false;
            } else {
                return false; //diladi an pano exei B i N
            }

        }
        return false;
    }

    private boolean lookDiagonal2Up(Point square1, char player) {
        int x = (int) square1.getX();
        int y = (int) square1.getY();
        int previousX = x - 1;       //to epomeno koutaki stin dieuthinsi pou koitame
        int previousY = y + 1;
        Vector<Point> v = new Vector<Point>(); //tha apothikeuoume tis allages poy tha kanei
        boolean add = false;
        if (x == 0) {
            return false;
        }
        if (y == 7) {
            return false;
        }
        if (player == 'B') {
            if ((previousX >= 0) && (previousY <= 7) && (otheloTable[previousX][previousY].getCircle() == 'B')) {
                return false;
            }
            if ((previousX >= 0) && (previousY <= 7) && (otheloTable[previousX][previousY].getCircle() == 'W')) {
                v.add(new Point(previousX, previousY));
                for (int i = (--previousX); i >= 0; i--) {
                    if (add) {
                        break;
                    }
                    previousY++;
                    if ((previousY <= 7) && (otheloTable[i][previousY].getCircle() == 'W')) {
                        v.add(new Point(i, previousY));
                    } else if ((previousY <= 7) && (otheloTable[i][previousY].getCircle() == 'B')) {
                        add = true;
                    } else {
                        return false;
                    }
                }
                if (add) {
                    otheloTable[x][y].addChange(v);
                    return true;
                }
                return false;
            } else {
                return false; //diladi an kato exei B i N
            }
        }
        if (player == 'W') {
            if ((previousX >= 0) && (previousY <= 7) && (otheloTable[previousX][previousY].getCircle() == 'W')) {
                return false;
            }
            if ((previousX >= 0) && (previousY <= 7) && (otheloTable[previousX][previousY].getCircle() == 'B')) {
                v.add(new Point(previousX, previousY));
                for (int i = (--previousX); i >= 0; i--) {
                    if (add) {
                        break;
                    }
                    previousY++;
                    if ((previousY <= 7) && (otheloTable[i][previousY].getCircle() == 'B')) {
                        v.add(new Point(i, previousY));
                    } else if ((previousY <= 7) && (otheloTable[i][previousY].getCircle() == 'W')) {
                        add = true;
                    } else {
                        return false;
                    }
                }
                if (add) {
                    otheloTable[x][y].addChange(v);
                    return true;
                }
                return false;
            } else {
                return false; //diladi an kato exei B i N
            }
        }
        return false;
    }

    public void makeMove(int x, int y, char player) {
        if (player == 'B') {
            this.otheloTable[x][y].setCircle('B');
            this.lastMove.setX(y);
            this.lastMove.setY(x);
            for (int i = 0; i < this.otheloTable[x][y].getChangesSize(); i++) {
                //  System.out.println("ALLAGI gia B " + (int) this.otheloTable[x][y].getChange(i).getX() + ", " + (int) this.otheloTable[x][y].getChange(i).getY());
                this.otheloTable[(int) this.otheloTable[x][y].getChange(i).getX()][(int) this.otheloTable[x][y].getChange(i).getY()].setCircle('B');
            }
            setRules('W');
        } else if (player == 'W') {
            this.otheloTable[x][y].setCircle('W');
            for (int i = 0; i < otheloTable[x][y].getChangesSize(); i++) {
                //  System.out.println("ALLAGI gia W" + (int) this.otheloTable[x][y].getChange(i).getX() + ", " + (int) this.otheloTable[x][y].getChange(i).getY());
                this.otheloTable[(int) otheloTable[x][y].getChange(i).getX()][(int) otheloTable[x][y].getChange(i).getY()].setCircle('W');
            }

            setRules('B');
        }

    }

    public char getPlayer() {
        return now;
    }

    public void setPlayer() {
        //     player = p;
        //    System.out.println(player + "player sto othelo");
        char help;
        setRules(previous);
        if (previous == now) {
            if (isTerminal()) {
                now = ' ';
            } else {
                help = now;
                now = previous;
                previous = help;
            }
        } else if (isTerminal()) {
            previous = now;
        } else {
            help = now;
            now = previous;
            previous = help;
        }
        player = now;
        if (now == 'W') {
            state state1 = new state(depth);
            Move m = state1.alpha_beta(this);
            setRules('W');
            this.makeMove(m.getX(), m.getY(), 'W');
            state1 = null;
            r.gc();
            setRules('B');
            setPlayer();

        }

    }

    private int getMovesNumber() {
        int counter = 0;
        for (int j = 0; j < 8; j++) {
            for (int i = 0; i < 8; i++) {
                if (otheloTable[j][i].getAdd_circle()) {
                    counter++;
                }
            }
        }
        return counter;

    }

    public boolean isTerminal() {
        int counter = 0;
        for (int j = 0; j < 8; j++) {
            for (int i = 0; i < 8; i++) {
                if ((this.otheloTable[j][i].getCircle() == 'B') || (this.otheloTable[j][i].getCircle() == 'W')) {
                    counter++;
                }
            }
        }
        if (counter == 64) {
            return true;
        } else {
            return false;
        }

    }
    /////////////////////////////////////////////////////
    ////////////////////////////////////////////////////

    public int evaluate() {
        this.setBlackcircles();
        this.setWhitecircles();

        int eval1 = this.getBlackcircles() - this.getWhitecircles();

        int eval2 = 0;
        int eval3 = 0;
        int sth = 0;
        boolean f = true;
        for (int j = 0; j < 8; j++) {
            for (int i = 0; i < 8; i++) {
                if (this.getSquare(j, i).getCircle() == 'S') {
                    if (getPlayer() == 'B') {
                        eval2 = -10;
                    } else {
                        eval2 = 10;
                    }
                } else if (this.getSquare(j, i).getLabel() == 'X') {
                    if (getPlayer() == 'B') {
                        eval2 = 10;
                    } else {
                        eval2 = -10;
                    }
                } else if (this.getSquare(j, i).getLabel() == 'C') {
                    if (getPlayer() == 'B') {
                        eval2 = 7;
                    } else {
                        eval2 = -7;
                    }
                } else if (this.getSquare(j, i).getLabel() == 'B') {
                    if (getPlayer() == 'B') {
                        eval2 = -5;
                    } else {
                        eval2 = 5;
                    }
                } else if (this.getSquare(j, i).getLabel() == 'A') {
                    if (getPlayer() == 'B') {
                        eval2 = 7;
                    } else {
                        eval2 = -7;
                    }
                }
                if ((i > 0)) {
                    if (this.getSquare(j, i - 1).getCircle() == 'B') {
                        eval3 = -10;
                    }
                } else if ((i < 7)) {
                    if (this.getSquare(j, i + 1).getCircle() == 'B') {
                        eval3 = -10;
                    }
                }
                if (f) {
                    if ((j > 0)) {
                        if (this.getSquare(j - 1, sth).getCircle() == 'B') {
                            eval3 = -10;
                        }
                    } else if ((j < 7)) {
                        if (this.getSquare(j + 1, sth).getCircle() == 'B') {
                            eval3 = -10;
                        }
                    }
                    f = false;
                }

                f = true;


            }
        }

        return (10 * eval1 / 100 + 40 * eval2 / 100 + 50 * eval3 / 100);


    }

    public void setEvaluation(int a) {
        this.evaluation = a;
    }

    public int getEvaluation() {
        return this.evaluation;
    }

    public ArrayList<Point> getMoves() {
        ArrayList<Point> a = new ArrayList<Point>();
        for (int j = 0; j < 8; j++) {
            for (int i = 0; i < 8; i++) {
                if (this.otheloTable[j][i].getAdd_circle()) {
                    a.add(new Point(j, i));
                }
            }
        }
        return a;
    }

    public Move getLastMove() {
        return this.lastMove;
    }
}
