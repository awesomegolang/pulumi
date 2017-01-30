// Copyright 2016 Marapongo, Inc. All rights reserved.

import {LocalVariable} from "./definitions";
import {Expression} from "./expressions";
import {Identifier, Node} from "./nodes";

import * as tokens from "../tokens";

export interface Statement extends Node {}

/** Blocks **/

export interface Block extends Statement {
    kind:       BlockKind;
    statements: Statement[];
}
export const blockKind = "Block";
export type  BlockKind = "Block";

/** Local Variables **/

export interface LocalVariableDeclaration extends Statement {
    kind:  LocalVariableDeclarationKind;
    local: LocalVariable;
}
export const localVariableDeclarationKind = "LocalVariableDeclaration";
export type  LocalVariableDeclarationKind = "LocalVariableDeclaration";

/** Try/Catch/Finally **/

export interface TryCatchFinally extends Statement {
    kind:          TryCatchFinallyKind;
    tryBlock:      Block;
    catchBlocks?:  TryCatchBlock[];
    finallyBlock?: Block;
}
export const tryCatchFinallyKind = "TryCatchFinally";
export type  TryCatchFinallyKind = "TryCatchFinally";

export interface TryCatchBlock extends Node {
    kind:       TryCatchBlockKind;
    block:      Block;
    exception?: LocalVariable;
}
export const tryCatchBlockKind = "TryCatchBlock";
export type  TryCatchBlockKind = "TryCatchBlock";

/** Branches **/

// A `break` statement (valid only within loops).
export interface BreakStatement extends Statement {
    kind:   BreakStatementKind;
    label?: Identifier;
}
export const breakStatementKind = "BreakStatement";
export type  BreakStatementKind = "BreakStatement";

// A `continue` statement (valid only within loops).
export interface ContinueStatement extends Statement {
    kind:   ContinueStatementKind;
    label?: Identifier;
}
export const continueStatementKind = "ContinueStatement";
export type  ContinueStatementKind = "ContinueStatement";

// An `if` statement.  To simplify the AST, this is the only conditional statement available.  All higher-level
// conditional constructs such as `switch`, `if / else if / ...`, etc. must be desugared into it.
export interface IfStatement extends Statement {
    kind:       IfStatementKind;
    condition:  Expression; // a `bool` condition expression.
    consequent: Statement; // the statement to execute if `true`.
    alternate?: Statement; // the statement to execute if `false`.
}
export const ifStatementKind = "IfStatement";
export type  IfStatementKind = "IfStatement";

// A labeled statement associates an identifier with a statement for purposes of labeled jumps.
export interface LabeledStatement extends Statement {
    kind:      LabeledStatementKind;
    label:     Identifier;
    statement: Statement;
}
export const labeledStatementKind = "LabeledStatement";
export type  LabeledStatementKind = "LabeledStatement";

// A `return` statement to exit from a function.
export interface ReturnStatement extends Statement {
    kind:        ReturnStatementKind;
    expression?: Expression;
}
export const returnStatementKind = "ReturnStatement";
export type  ReturnStatementKind = "ReturnStatement";

// A `throw` statement to throw an exception object.
export interface ThrowStatement extends Statement {
    kind:       ThrowStatementKind;
    expression: Expression;
}
export const throwStatementKind = "ThrowStatement";
export type  ThrowStatementKind = "ThrowStatement";

// A `while` statement.  To simplify the AST, this is the only looping statement available.  All higher-level
// looping constructs such as `for`, `foreach`, `for in`, `for of`, `do / while`, etc. must be desugared into it.
export interface WhileStatement extends Statement {
    kind: WhileStatementKind;
    test: Expression; // a `bool` statement indicating whether to continue.
    body: Block;      // the body to execute provided the test remains `true`.
}
export const whileStatementKind = "WhileStatement";
export type  WhileStatementKind = "WhileStatement";

/** Miscellaneous **/

// An empty statement.
export interface EmptyStatement extends Statement {
    kind: EmptyStatementKind;
}
export const emptyStatementKind = "EmptyStatement";
export type  EmptyStatementKind = "EmptyStatement";

// Multiple statements in one (unlike a block, this doesn't introduce a new scope).
export interface MultiStatement extends Statement {
    kind:       MultiStatementKind;
    statements: Statement[];
}
export const multiStatementKind = "MultiStatement";
export type  MultiStatementKind = "MultiStatement";

// A statement that performs an expression, but ignores its result.
export interface ExpressionStatement extends Statement {
    kind:       ExpressionStatementKind;
    expression: Expression;
}
export const expressionStatementKind = "ExpressionStatement";
export type  ExpressionStatementKind = "ExpressionStatement";

