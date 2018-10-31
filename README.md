# gIntersect

Only prints those words that are both in [File A] and [File B]


    usage : intersect [File A] [File B]
    Prints out the lines that are both in [File A] and [File B]


#Example

    $ cat FileA.txt
    hello
    $ cat FileB.txt
    hello
    world
    $ gIntersect FileA.txt FileB.txt
    world


#Memory
**gIntersect** reads the contents of **FileA.txt** into memory then compares it with **FileB.txt** . So performance will deteriorate for very large **FileA.txt** file.

