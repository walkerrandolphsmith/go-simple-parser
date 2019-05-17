# combinator parser

I want to create a simple string parser that can handle python like logical operators.

Here is the BNF I plan to go by:

```
<expression>::= <term>{<or><term>}
<term>::= <factor>{<and><factor>}
<factor>::= <constant> | <not><factor> | <group_start><expression><group_end>
<constant>::= false | true
<or>::= or
<and>:= and
<not>:= not
<group_start>:= (
<group_end>:= )
```

## Test
```
go test ./...
```