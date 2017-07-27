import java.util.function.Function;

class A {
  public int a(int x) { return x+1; };
}

class B extends A {
  B() { Function<Integer, Integer> f = super::a; }
}
