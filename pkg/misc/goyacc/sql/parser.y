%{
package sql
%}

%union {
    ident string
    idents []string
    statement StmtNode
}

%token SELECT FROM

%token <ident> IDENT

%type <idents> FieldList
%type <statement> SelectStmt

%%
statement:
SelectStmt
{
    yylex.(*lex).result = $1
}

SelectStmt:
SELECT FieldList FROM IDENT
{
    $$ = &SelectStmt{Table: $4, Columns: $2}
}

FieldList:
IDENT
{
    $$ = []string{$1}
}
| FieldList ',' IDENT
{
    $$ = append($1, $3)
}