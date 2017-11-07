import java.lang.annotation.*;
class X <@Marker1 @Marker3 F extends @Marker1 @Marker2 String> {
	void foo(F f) {
    Y <@Marker2 @Marker3 ? super @Marker1 @Marker2 String> y = new @Marker2 @Marker1 Y<String>();
  }
}
class Y<@Marker3 T> {
	public int bar(T t) {
		return t instanceof @Marker1 @Marker2 String ? t.toString().length() : 0;
	}
}
@Target (ElementType.TYPE_USE)
@interface Marker1 {}
@Target (ElementType.TYPE_USE)
@interface Marker2 {}
@Target (ElementType.TYPE_USE)
@interface Marker3 {}
