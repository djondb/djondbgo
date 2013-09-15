// =====================================================================================
//  Filename:  settings.h
// 
//  Description: The class settings will read the djondb.conf file 
// 
//  Version:  1.0
//  Created:  03/11/2012 12:39:40 PM
//  Revision:  none
//  Compiler:  gcc
// 
//  Author: Juan Pablo Crossley (crossleyjuan@gmail.com), 
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
#ifndef SETTINGS_H_INCLUDED
#define SETTINGS_H_INCLUDED

#include "defs.h"
#include <string>

const std::string getSetting(std::string);
void setSetting(std::string key, std::string value);

#endif // SETTINGS_H_INCLUDED
