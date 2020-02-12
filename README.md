# Vocabulary:

- Business Case: A high level description of something the system should be doing.
    Ideally it should be almost able to be read by a non-programmer.
    It has no dependencies to anything concrete. It is doubly isolated from details, on one hand from the I/O mechanism 
    (Adapters do that), and on the other hand from the specific way that lower-level operations are done, 
    (Services do that). Abbreviated to Case from now on.
    
- Entry point: Any entry point to the system (a main function), for example a main that starts an http server, 
   or one that is triggered by a cli.
   
- Adapter: A function that sits between an entry point and the Case. It gets whatever input the nature of that
    entry point can provide and adapts it to a Case's input. It then takes the Case's output and adapts it
    to the output that the entry point allows.
    
- Service: A struct that provides an API to do some low-level operations. It is never referenced directly inside a Case,
    but only through an Interface that the Service satisfies. Examples of Services are Persistence, Email, KeyValue.
    
- DI: A DI Container that holds all the Services. It is injected inside the Cases, so they can use the Services.
    The services are declared in terms of their Interface, so the Cases are not aware of the concrete implementation.

```
project/                
├── entry_points/       Entry points for the system.
│   ├── cli/            
│   │   ├── commands/   
│   │   └── main.go     
│   └── http            
│       └── main.go     
├── lib                 Where extensions of core library functionalities live.
└── src                 Here is where the actual system is.
    ├── adapters/       Adapters from entry points I/O to Business Cases I/O. 
    │   ├── cli/        
    │   └── http/       
    ├── core            The Core of the system. Immune to I/0 & infrastructure changes.
    │   ├── cases/      Business cases. High level. Can only do stuff because it is injected a Di with services.
    │   └── entities/   Entities.
    └── di              
        ├── di.go       
        └── services/   Services the Di holds so Business Cases can do low-level stuff.
```

# How to see it work:
- Clone repo
- ```docker-compose up``` to fire up the containers
- ```docker-compose exec go bash``` to log in the project container
- ```go test``` to run the tests
