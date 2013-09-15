// =====================================================================================
//  Filename:  filterdefs.h
// 
//  Description:  
// 
//  Version:  1.0
//  Created:  07/17/2012 08:44:23 AM
//  Revision:  none
//  Compiler:  gcc
// 
//  Author:  YOUR NAME (), 
// 
// License:
// 
// This file is part of the djondb project, for license information please refer to the LICENSE file,
// the application and libraries are provided as-is and free of use under the terms explained in the file LICENSE
// Its authors create this application in order to make the world a better place to live, but you should use it on
// your own risks.
// 
// Also, be adviced that, the GPL license force the committers to ensure this application will be free of use, thus
// if you do any modification you will be required to provide it for free unless you use it for personal use (you may 
// charge yourself if you want), bare in mind that you will be required to provide a copy of the license terms that ensures
// this program will be open sourced and all its derivated work will be too.
// =====================================================================================
#ifndef FILTERDEFS_INCLUDED_H
#define FILTERDEFS_INCLUDED_H
#include "stringfunctions.h"

enum FILTER_OPERATORS {
	FO_NONE,
	FO_TOKEN,
	FO_AND, FO_OR,
	FO_EQUALS,
	FO_NOT_EQUALS,
	FO_PARENTESIS_OPEN,
	FO_PARENTESIS_CLOSE,
	FO_NOT,
	FO_LESSTHAN,
   FO_LESSEQUALTHAN,
	FO_GREATERTHAN,
	FO_GREATEREQUALTHAN,
	FO_EXISTS
};

enum EXPRESSION_TYPE {
	ET_CONSTANT,
	ET_SIMPLE,
	ET_BINARY,
	ET_UNARY
};

FILTER_OPERATORS parseFilterOperator(const char* text);

class ParseException: public std::exception {
	public:
		ParseException(int code, const char* error);
		ParseException(const ParseException& orig);
		virtual const char* what() const throw();
		int errorCode() const;

	private:
		int _errorCode;
		const char* _errorMessage;
};

#endif // FILTERDEFS_INCLUDED_H
