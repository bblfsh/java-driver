class Code {
  <I extends Appendable & Readable> Code(I a) {
    System.out.println((Appendable & Readable) a);
  }
}
