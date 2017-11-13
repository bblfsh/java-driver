class Code {
  private int x;
  void Code(int x) {
    synchronized(this) { this.x = x; }
  }
}
