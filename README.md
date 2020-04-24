# BeerProto

Another beer format, written in protocol buffer.

Written as a pure computer interchange format for language-neutral, platform-neutral, extensible way of serializing structured data for use in communications protocols and data storage.

#### Protocol Buffer
BeerProto follows the structure of BeerJSON and can take advantages of the official code generating plugins for protocol buffers given support for C++, C#, Dart, Go, Java, Python, Ruby, Objective C, PHP and JavaScript as well as third part plugins such as Swift.

You can also use the Proto3 supports for canonical encoding in JSON, making it easier to share data in a human readable way.

#### Why
The exchange of recipes is still very defective as the most supported format is still BeerXML, this however has may defects as it favours human readability over interoperability. this results in software not being able to convert weights, measurements and timings automatically. 

A proposed update to BeerXML was made with the draft BeerXML 2.0 spec which would resolve many of those issues however this became abandoned.
A second replacement attempt has been made with BeerJSON which took a lot of inspiration from the 2.0 spec but has yet to see much adoption.   
  
One of the biggest limiting factors of BeerJSON is JSON Schema, which has poor support for code generating. This makes it slow for developer integration and is likely holding it back being adopted into software, as such BeerProto has prioritised the developer integration over human readability, while building on top of all the work done for BeerJSON.  

#### Notes

Interoperability with BeerJSON can be achieved, but does require the conversion of the enumerators between int32's and strings and handling non nullables properties. To assit with this we provide a command line tool that can convert between BeerProto and BeerJSON.