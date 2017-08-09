import java.lang.annotation.*;
@Target(ElementType.TYPE_USE) @interface X {}
@Target(ElementType.TYPE_USE) @interface Y {}

class A {
	class B {
		class C {
			public C(@X A.@Y B B.this){}
    }
  }
}
