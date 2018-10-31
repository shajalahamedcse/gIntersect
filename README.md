


# gIntersect

Only prints those words that are both in [File A] and [File B]


    usage : intersect [File A] [File B]
    Prints out the lines that are both in [File A] and [File B]


# Example

    $ cat FileA.txt
    hello
    $ cat FileB.txt
    hello
    world
    $ gIntersect FileA.txt FileB.txt
    world
    
# Performance

    $ time grep -w FileA.txt FileB.txt
    real    0m0.028s
    user    0m0.003s
    sys     0m0.000s
    
    $ time ./main FileA.txt FileB.txt
    real    0m0.004s
    user    0m0.001s
    sys     0m0.004s
    

# Memory

**gIntersect** reads the contents of **[FileA.txt]** into (hashmap) memory then compares it with **[FileB.txt]** . 
So performance will deteriorate for very large **[FileA.txt]** file.

