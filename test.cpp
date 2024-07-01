#include <boost/statechart/state_machine.hpp>
#include <boost/statechart/simple_state.hpp>
#include <boost/statechart/transition.hpp>
#include <boost/statechart/termination.hpp>

// Forward declarations
struct Open;
struct bad_data;
struct create, resolve;
struct Alerting;
struct select;
struct Active;
struct finish;
struct deselect;
struct escalate;
struct Escalated;
struct escalate_close;
struct defer, wait;
struct Waiting;
struct timeout, finish;

// State machine
struct StateMachine : boost::statechart::state_machine<StateMachine, Open> {};

// States and transitions
struct Open : boost::statechart::simple_state<Open, StateMachine> {
    typedef boost::mpl::list<
        boost::statechart::transition<Complete, bad_data>,
        boost::statechart::transition<Alerting, create, resolve>
    > reactions;
};

struct bad_data : boost::statechart::simple_state<bad_data, StateMachine> {
};

struct create, resolve : boost::statechart::simple_state<create, resolve, StateMachine> {
};

struct Alerting : boost::statechart::simple_state<Alerting, StateMachine> {
    typedef boost::mpl::list<
        boost::statechart::transition<Active, select>
    > reactions;
};

struct select : boost::statechart::simple_state<select, StateMachine> {
};

struct Active : boost::statechart::simple_state<Active, StateMachine> {
    typedef boost::mpl::list<
        boost::statechart::transition<Complete, finish>,
        boost::statechart::transition<Alerting, deselect>,
        boost::statechart::transition<Escalated, escalate>,
        boost::statechart::transition<Waiting, defer, wait>
    > reactions;
};

struct finish : boost::statechart::simple_state<finish, StateMachine> {
};

struct deselect : boost::statechart::simple_state<deselect, StateMachine> {
};

struct escalate : boost::statechart::simple_state<escalate, StateMachine> {
};

struct Escalated : boost::statechart::simple_state<Escalated, StateMachine> {
    typedef boost::mpl::list<
        boost::statechart::transition<Complete, escalate_close>
    > reactions;
};

struct escalate_close : boost::statechart::simple_state<escalate_close, StateMachine> {
};

struct defer, wait : boost::statechart::simple_state<defer, wait, StateMachine> {
};

struct Waiting : boost::statechart::simple_state<Waiting, StateMachine> {
    typedef boost::mpl::list<
        boost::statechart::transition<Alerting, timeout, finish>
    > reactions;
};

struct timeout, finish : boost::statechart::simple_state<timeout, finish, StateMachine> {
};

// Events
struct Complete : boost::statechart::event<Complete> {};
struct Alerting : boost::statechart::event<Alerting> {};
struct Active : boost::statechart::event<Active> {};
struct Escalated : boost::statechart::event<Escalated> {};
struct Waiting : boost::statechart::event<Waiting> {};

