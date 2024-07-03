#pragma once
#include <string>
#include <vector>
#include <map>
#include <memory>

// Forward declaration
class ForwardDeclaredClass;

//0
// Class with different access specifiers
class A {
private:
    int x;
    float y;
protected:
    std::string s;
    std::vector<bool> vb;
public:
    int get_x() const;
    float get_y() const;
    std::string get_s() const;
    std::vector<bool> get_vb() const;
    void set_x(int new_x);
};

//1
// Struct (default public)
struct B {
    int x;
    float y;
    std::string s;
    std::vector<bool> vb;
    int get_x() const;
    float get_y() const;
    std::string get_s() const;
    std::vector<bool> get_vb() const;
};

//2
// Template class
template <typename T>
class TemplateClass {
private:
    std::vector<T> items;
public:
    void add_item(const T& item);
    T get_item(size_t index) const;
};

//3
// Class with inheritance
class DerivedClass : public A, private B {
public:
    void derived_method();
};

//4
// Class with virtual methods and abstract class
class AbstractBase {
public:
    virtual void pure_virtual_method() = 0;
    virtual ~AbstractBase() = default;
};

//5
class ConcreteClass : public AbstractBase {
public:
    void pure_virtual_method() override;
};

//6
// Class with static members and friend function
class StaticExample {
private:
    static int static_var;
public:
    static void static_method();
    friend void friend_function(StaticExample&);
};

//7
// Namespace usage
namespace MyNamespace {
    class NamespacedClass {
    public:
        void namespaced_method();
    };
}

//8
// Class with complex method signatures
class ComplexMethods {
public:
    std::vector<std::string> process(const std::map<int, std::string>& input) const;
    std::shared_ptr<ForwardDeclaredClass> create(int id, const std::string& name);
};

//9
// Enum class
enum class Color {
    Red,
    Green,
    Blue
};

//10
// Class using enum and typedef
class EnumUser {
public:
    typedef std::vector<Color> ColorList;
    ColorList favorite_colors;
    void add_color(Color c);
};