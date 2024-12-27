#!/bin/sh
alias antlr4='java -Xmx500M -cp "/usr/local/lib/antlr4-4.13.1-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
antlr4 -Dlanguage=Go -visitor -package parser SysYF.g4 -o ../parser
