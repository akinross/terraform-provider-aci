# Redesign Approach

Phased approach:

1. Changed the generator logic and templates to output the same `resource(_test).go`, `datasource(_test).go`, `examples.tf` and `documentation.md` as we currently have.
2. Make improvements to the generated resource.go and datasource.go code (current tests should pass)
3. Make improvements to testing

## Phase 1: Generator Redesign

* clean-up unused functionality in generator
* simplify templates
* per class definition file
* unmarshall to predeterminged struct for the definition load into class details
    - removes the need for all the looping to search for values in definition file
* store metadata per class in a map with references to child identifier (classname) instead of full child included for each class
    - removes duplication of metadata storage
* model test values to property struct
    - in code generate references for test values when applicable
    - pre-calculate the possible test values, choice from valid values, overwrites etc...
* model migrations to property/class structs
    - allow for multiple versions to be specified for migration ( state upgrade expansion )
    - allow for resource renaming ( state move ) 
* use functions to store all required documentation data in the class and property struct
    - description of resource 
    - additional notes & warnings
    - supported versions
    - dn formats
    - UI location
    - examples
    - description for attributes including references to classes
    - valid values and defaults for attributes
    - list of child resources
    - import examples for cli and block
* improved logging and log file output during generator execution
* increase inline commenting to explain behaviour
* generator documentation
* developer experience improvement
    - limit generation to single class
    - generator performance enhancements

## Phase 2: Resource/Datasource Code Changes

* clean-up unused attributes ( like ctx, diags etc... )
* define and reuse types and models for each class
* fix request to include children only 1 time rsp-subtree-class
* add computed id attribute in children for RN or/and DN of child

## Phase 3: Modifying Tests

* analyse current tests
    - coverage
    - defined tests
* create new tests
