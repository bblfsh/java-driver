class Foo {
    int bar(int val) { return val+1; }
    void foo() {
        Interf a = val -> val+1;
        a = (int val) -> { return val+1; };
        a = Foo::bar;
    }
}