// =====================================================================================
//  Filename:  bsonarrayobj.h
//
//  Description:  This file contains the definition of the class BSONArrayObj which is an array of BSONObj
//
//  Version:  1.0
//  Created:  02/15/2012 09:07:11 AM
//  Revision:  none
//  Compiler:  gcc
//
//  Author:  Juan Pablo Crossley (crossleyjuan@gmail.com),
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

#ifndef BSONARRAYOBJ_H_INCLUDED
#define BSONARRAYOBJ_H_INCLUDED

#include <vector>
#include "defs.h"

class BSONObj;

class BSONArrayObj {
public:
	BSONArrayObj();
	~BSONArrayObj();
	BSONArrayObj(const BSONArrayObj& orig);

	virtual __int32 length() const;
	/**
	 * @brief Adds the BSONObj to the array
	 *
	 * @param obj
	 */
	void add(const BSONObj& obj);
	/**
	 * @brief Adds all the elements of the argument to the current elements
	 *
	 * @param array
	 */
	void addAll(const BSONArrayObj& array);
	virtual BSONObj* get(__int32 index) const;
   virtual char* toChar() const;
	typedef std::vector<BSONObj*>::iterator iterator;
	typedef std::vector<BSONObj*>::const_iterator const_iterator;
	virtual BSONArrayObj* select(const char* select) const;
	
	virtual iterator begin();
	virtual iterator end();

	virtual const_iterator begin() const;
	virtual const_iterator end() const;


private:
	std::vector<BSONObj*> _elements;
};
#endif // BSONARRAYOBJ_H_INCLUDED
