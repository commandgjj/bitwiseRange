Functionality Overview
======================

In OpenFlow, there is no comparsion. Such as, you cannot match ``1000 <= tcp_port
<= 1999'' simply as you do it in other program language. But, OpenFlow supports
bitwise match.

Range matches can be expressed as a collection of bitwise matches. For example,
suppose that the goal is to match TCP source ports 1000 to 1999, inclusive.
The binary representations of 1000 and 1999 are:

01111101000

11111001111

The following series of bitwise matches will match
1000 and 1999 and all the values in between:

01111101xxx

0111111xxxx

10xxxxxxxxx

110xxxxxxxx

1110xxxxxxx

11110xxxxxx

1111100xxxx

which can be written as the following matches:

tcp,tp_src=0x03e8/0xfff8

tcp,tp_src=0x03f0/0xfff0

tcp,tp_src=0x0400/0xfe00

tcp,tp_src=0x0600/0xff00

tcp,tp_src=0x0700/0xff80

tcp,tp_src=0x0780/0xffc0

tcp,tp_src=0x07c0/0xfff0

This function will return the value and corresponding mask. Such as, for above
input [1000, 1999], it will return:
 [(0x03e8, 0xfff8),

 (0x03f0, 0xfff0),
 
 (0x0400, 0xfe00),
 
 (0x0600, 0xff00),
 
 (0x0700, 0xff80),
 
 (0x0780, 0xffc0),
 
 (0x07c0, 0xfff0)]

Of course, it also can compute IP address bitwise ranges.
