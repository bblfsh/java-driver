import java.lang.annotation.*;

@Target(ElementType.TYPE_USE)
@interface X{}

class A{
  A() {
    java.util.@X Date date;
  }
}

