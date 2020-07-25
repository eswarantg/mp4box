# mp4box

## Motivation
* Didn't find any pure Go Implementation of Mp4 Box
* Optimized for Reader as all the boxes are typically not parsed
* Only when specific Property is required the actual parsing takes place

## Class Diagram
![class diagram](out/class/mp4box.png?raw=true "Class Diagram")

## Implementation Details
* Base frame work
    * Reading of Boxes is the primary goal
    * Writing of File with boxes is also provisioned and MUST work
* Completeness of all Boxes
    * Boxes that are required are added...
    * Contributions are welcome to add as required
* Getter methods for property
    * Getter method for required properties are implemented as required
* Setter methods for property
    * No implementation till now...
    * But accounted for in the design


## Unit Test
* Standard Unit Test framework is included
* Refer "test/urls.txt" and include other samples as required
* Specific field value evaluation is still TBD
