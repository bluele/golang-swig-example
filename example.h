#ifndef EXAMPLE_H
#define EXAMPLE_H

class Example {
private:
  int val;
public:
  Example();
  ~Example();
  void setVal(int x);
  int getVal();
  static int add(int x, int y);
};

#endif