package goparser

/*import (
	_ "fmt"
	. "github.com/axh432/gogex"
	"github.com/axh432/goparser/optional"
	_ "strings"
)

var owb = optional.Whitespace.Block


var (
	//Primitives
	whitespaceBlock                  = Range(Whitespace, 1, -1)
	optionalWhitespaceBlock          = Range(Whitespace, 0, -1)
	whitespaceNoNewLine              = SetOfCharacters(" \t")
	whitespaceNoNewLineBlock         = Range(whitespaceNoNewLine, 1, -1)
	optionalWhitespaceNoNewLineBlock = Range(SetOfCharacters(" \t"), 0, -1)
	newline                          = SequenceOfCharacters("\n")
	whitespaceAtLeastOneNewLineBlock = Sequence(optionalWhitespaceNoNewLineBlock, newline, optionalWhitespaceBlock)
	underscore                       = SetOfCharacters("_")
	comma                            = SetOfCharacters(",")
	openBracket                      = SetOfCharacters("(")
	closedBracket                    = SetOfCharacters(")")
	openCurlyBrace                   = SetOfCharacters("{")
	closedCurlyBrace                 = SetOfCharacters("}")
	quote                            = SetOfCharacters(`"`)
	dot                              = SetOfCharacters(".")
	word                             = Range(Letter, 1, -1)
	String                           = Label(Sequence(quote, Range(Set(SetOfNotCharacters(`"`), SequenceOfCharacters(`\"`)), 0, -1), quote), "String:string")
	boolValue                        = Label(Set(SequenceOfCharacters("true"), SequenceOfCharacters("false")), "Bool:bool")
	integerValue                     = Label(Range(Number, 1, -1), "Integer:int")
	listOfIntegerValues              = Label(Sequence(integerValue, Range(Sequence(optionalWhitespaceNoNewLineBlock, comma, optionalWhitespaceBlock, integerValue), 0, -1)), "ListOfIntegers:[]int")

	//name
	letterNumberUnderscoreBlock         = Range(Set(Letter, Number, underscore), 1, -1)
	optionalLetterNumberUnderscoreBlock = Range(Set(Letter, Number, underscore), 0, -1)
	letterOrUnderscore                  = Set(Letter, underscore)
	name                                = Sequence(letterOrUnderscore, optionalLetterNumberUnderscoreBlock)

	typeName     = Label(name, "TypeName:string")
	variableName = Label(name, "VariableName:string")
	returnType   = Label(name, "ReturnType:string")
	functionName = Label(name, "FunctionName:string")

	//importKeyword
	importKeyword         = Label(SequenceOfCharacters("import"), "ImportKeyword:string")
	importAccessType      = Label(Range(SetOfCharacters("_."), 0, 1), "ImportAccessType:string")
	importName            = Label(String, "ImportName:string")
	Importd               = Label(Sequence(importAccessType, optionalWhitespaceNoNewLineBlock, importName), "Import")
	importMultiple        = Sequence(Importd, Range(Sequence(whitespaceAtLeastOneNewLineBlock, Importd), 0, -1))
	importBoundedMultiple = Sequence(openBracket, optionalWhitespaceBlock, importMultiple, optionalWhitespaceBlock, closedBracket)
	importBoundedEmpty    = Sequence(openBracket, optionalWhitespaceBlock, closedBracket)
	importBoundedAll      = Set(importBoundedMultiple, importBoundedEmpty)
	imports               = Label(Set(importBoundedAll, Importd), "Imports:[]Import")
	importDeclaration       = Label(Sequence(importKeyword, optionalWhitespaceBlock, imports), "ImportStatement")

	//Function Parameters
	parameter                     = Label(Sequence(variableName, whitespaceNoNewLineBlock, typeName), "Parameter")
	functionParametersList        = Sequence(parameter, Range(Sequence(optionalWhitespaceNoNewLineBlock, comma, optionalWhitespaceBlock, parameter), 0, -1))
	functionParametersBoundedList = Sequence(openBracket, optionalWhitespaceBlock, functionParametersList, optionalWhitespaceNoNewLineBlock, closedBracket)
	functionParametersEmpty       = Sequence(openBracket, optionalWhitespaceBlock, closedBracket)
	functionParameters            = Set(functionParametersBoundedList, functionParametersEmpty)

	//Function Return Parameters
	returnParametersNamed       = functionParameters
	returnParametersSingle      = returnType
	returnParametersList        = Sequence(returnType, Range(Sequence(optionalWhitespaceNoNewLineBlock, comma, optionalWhitespaceBlock, returnType), 0, -1))
	returnParametersBoundedList = Sequence(openBracket, optionalWhitespaceBlock, returnParametersList, optionalWhitespaceNoNewLineBlock, closedBracket)
	returnParameters            = Set(returnParametersSingle, returnParametersBoundedList, returnParametersNamed)
	optionalReturnParameters    = Range(returnParameters, 0, 1)

	//Function Signature
	Func              = Label(SequenceOfCharacters("func"), "FunctionKeyword")
	functionSignature = Label(Sequence(Func, whitespaceNoNewLineBlock, functionName, optionalWhitespaceNoNewLineBlock, functionParameters, optionalWhitespaceNoNewLineBlock, optionalReturnParameters), "FunctionSignature")

	//Var Assign Statement
	Var                   = Label(SequenceOfCharacters("var"), "VarKeyword")
	varAssignmentOperator = SetOfCharacters("=")
	valuePossibilities    = Set(String, boolValue, listOfIntegerValues, variableName, functionCall)
	optionalTypeName      = Range(typeName, 0, 1)
	varNames              = Sequence(variableName, Range(Sequence(optionalWhitespaceNoNewLineBlock, comma, optionalWhitespaceBlock, variableName), 0, -1))
	varAssignStatement    = Sequence(Var, optionalWhitespaceBlock, varNames, whitespaceNoNewLineBlock, optionalTypeName, optionalWhitespaceNoNewLineBlock, varAssignmentOperator, optionalWhitespaceBlock, valuePossibilities)

	//Var Declaration Statement
	varDeclarationStatement = Sequence(Var, whitespaceBlock, varNames, whitespaceNoNewLineBlock, typeName)

	//Var Block
	varBlockAssignStatement = Sequence(varNames, whitespaceNoNewLineBlock, optionalTypeName, optionalWhitespaceNoNewLineBlock, varAssignmentOperator, optionalWhitespaceBlock, valuePossibilities)
	varMultiple             = Sequence(varBlockAssignStatement, Range(Sequence(whitespaceAtLeastOneNewLineBlock, varBlockAssignStatement), 0, -1))
	varBoundedMultiple      = Sequence(openBracket, optionalWhitespaceBlock, varMultiple, optionalWhitespaceBlock, closedBracket)
	varBoundedEmpty         = Sequence(openBracket, optionalWhitespaceBlock, closedBracket)
	varBoundedAll           = Set(varBoundedMultiple, varBoundedEmpty)
	varBlock                = Sequence(Var, optionalWhitespaceBlock, varBoundedAll)

	//Var Statement
	varStatement = Label(Set(varAssignStatement, varDeclarationStatement, varBlock), "VarStatement:struct")

	//Assign Statement
	assignmentOperator = SequenceOfCharacters(":=")
	assignStatement    = Label(Sequence(varNames, optionalWhitespaceNoNewLineBlock, assignmentOperator, optionalWhitespaceBlock, valuePossibilities), "assignStatement")

	//Function Body
	statement              = Label(Set(varStatement, assignStatement, functionCall), "statement")
	statements             = Sequence(statement, Range(Sequence(whitespaceAtLeastOneNewLineBlock, statement), 0, -1))
	statementsBounded      = Sequence(openCurlyBrace, optionalWhitespaceBlock, statements, optionalWhitespaceBlock, closedCurlyBrace)
	statementsBoundedEmpty = Sequence(openCurlyBrace, optionalWhitespaceBlock, closedCurlyBrace)
	functionBody           = Set(statementsBounded, statementsBoundedEmpty)

	//Function
	functionDeclaration = Label(Sequence(functionSignature, optionalWhitespaceNoNewLineBlock, functionBody), "functionDeclaration")

	//Package
	Package            = SequenceOfCharacters("package")
	packageName        = Label(name, "packagename")
	packageDeclaration = Label(Sequence(Package, whitespaceNoNewLineBlock, packageName), "packageDeclaration")

	//Basic Golang
	Golang = Sequence(packageDeclaration, whitespaceAtLeastOneNewLineBlock, importDeclaration, whitespaceAtLeastOneNewLineBlock, functionDeclaration)
	Golang2 = Sequence(packageDeclaration, whitespaceAtLeastOneNewLineBlock, importDeclaration, whitespaceAtLeastOneNewLineBlock, varStatement, functionDeclaration)
)

//Function Call
func functionCall(iter *Iterator) MatchTree {
	functionCallParameter := Label(Set(variableName, String, functionCall), "parameter")
	functionCallParameters := Sequence(functionCallParameter, Range(Sequence(optionalWhitespaceNoNewLineBlock, comma, optionalWhitespaceBlock, functionCallParameter), 0, -1))
	functionCallParametersBounded := Sequence(openBracket, optionalWhitespaceBlock, functionCallParameters, optionalWhitespaceNoNewLineBlock, closedBracket)
	functionCallParametersEmpty := Sequence(openBracket, optionalWhitespaceBlock, closedBracket)
	functionCallParametersAll := Set(functionCallParametersBounded, functionCallParametersEmpty)

	optionalPackageName := Range(Sequence(packageName, optionalWhitespaceNoNewLineBlock, dot, optionalWhitespaceBlock), 0, 1)

	return Sequence(optionalPackageName, functionName, optionalWhitespaceNoNewLineBlock, functionCallParametersAll)(iter)
}*/
