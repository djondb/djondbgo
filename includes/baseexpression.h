// =====================================================================================
//  Filename:  baseexpression.h
// 
//  Description:  
// 
//  Version:  1.0
//  Created:  07/17/2012 08:42:41 AM
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
#ifndef BASEEXPRESSION_INCLUDE_H
#define BASEEXPRESSION_INCLUDE_H

#include "filterdefs.h"

class BSONObj;
class ExpressionResult;

class BaseExpression {
	public:
		BaseExpression(EXPRESSION_TYPE type) {
			_type = type;
		}

		virtual ~BaseExpression() {}

		EXPRESSION_TYPE type() {
			return _type;
		}

		virtual ExpressionResult* eval(const BSONObj& bson) = 0;
		virtual BaseExpression* copyExpression() = 0;
	private:
		EXPRESSION_TYPE _type;
};

#endif //BASEEXPRESSION_INCLUDE_H

