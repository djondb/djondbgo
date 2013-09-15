/*
 * =====================================================================================
 *
 *       Filename:  djon_error_codes.h
 *
 *    Description:  A place to put all the error codes
 *
 *        Version:  1.0
 *        Created:  03/07/2013 10:17:01 PM
 *       Revision:  none
 *       Compiler:  gcc
 *
 *         Author:  Juan Pablo Crossley (Cross), crossleyjuan@gmail.com
 *   Organization:  djondb
 *
 * This file is part of the djondb project, for license information please refer to the LICENSE file,
 * the application and libraries are provided as-is and free of use under the terms explained in the file LICENSE
 * Its authors create this application in order to make the world a better place to live, but you should use it on
 * your own risks.
 * 
 * Also, be adviced that, the GPL license force the committers to ensure this application will be free of use, thus
 * if you do any modification you will be required to provide it for free unless you use it for personal use (you may 
 * charge yourself if you want), bare in mind that you will be required to provide a copy of the license terms that ensures
 * this program will be open sourced and all its derivated work will be too.
 * =====================================================================================
 */ 
#ifndef DJON_ERROR_CODES_INCLUDED_H
#define DJON_ERROR_CODES_INCLUDED_H 

// 100 is reserved for database errors
const int D_ERROR_TOO_MANY_RESULTS = 100;
const int D_ERROR_PARSEERROR = 101;

// 600 is reserved for user errors
const int D_ERROR_CONNECTION = 600;
const int D_ERROR_INVALID_STATEMENT = 601;

#endif /* DJON_ERROR_CODES_INCLUDED_H */
