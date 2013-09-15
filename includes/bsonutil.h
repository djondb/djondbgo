/*
 * =====================================================================================
 *
 *       Filename:  bsonutil.h
 *
 *    Description:  This contains utility functions for bson
 *
 *        Version:  1.0
 *        Created:  09/09/2012 07:22:40 PM
 *       Revision:  none
 *       Compiler:  gcc
 *
 *         Author:  Juan Pablo Crossley (crossleyjuan@gmail.com), 
 *   Organization:  
 *
 * =====================================================================================
 */
#ifndef BSONUTIL_INCLUDE_H
#define BSONUTIL_INCLUDE_H

#include <string>
#include <set>
#include "bsondefs.h"

std::set<std::string> bson_splitSelect(const char* select);
char* bson_subselect(const char* select, const char* elment);

void* convert(BSONTYPE fromType, BSONTYPE toType, void* value);

#endif // BSONUTIL_INCLUDE_H

