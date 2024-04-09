package design_patterns

// This situation arises when different child are using the same functionality but they are not inheriting it from the interface
// for eg if there is a Vehicle class but there is subclasses Motorcycle , Car , Bicycle for eg they are having the code for eg Display which is there only in the Motorcycle and Car
// Now this causes duplicacy of the code
//
