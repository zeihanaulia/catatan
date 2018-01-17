# ADVANCED DESIGN PATTERNS WITH REACT

## Events

### Syntethic events

- Enable technique to make event accessible
- Object wrap original event object
- Has browser independent properties
- Reused - single global handler

### Store event using react

- Hard to store event and reuse later
    - Reason - becomes null, post action performed

- Solution:
    - Persist method - call to make event presistent

### Convention

- Recalls in the series as event attached
- Fire callbacks when occur

### Other Event

- Event bubbling - attach sigle event handler to root element
- Event delegation:
    - On firing event by browser - React calls the handler
    - Used for memory
    - Speed Optimization

### Note

- Aim 
    - Write less boilerplate and avoid duplicating the code

- Solution
    - write single event handler foreach component

- Aim
    - create new event handler for same component
- Solution:
    - add new case to switch