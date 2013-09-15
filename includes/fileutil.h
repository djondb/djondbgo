/* 
 * File:   fileutil.h
 * Author: cross
 *
 * Created on July 7, 2010, 3:08 PM
 */

#ifndef _FILEUTIL_H
#define	_FILEUTIL_H

#include <vector>
#include <string>
#include "defs.h"

char* readFile(char* fileName);
int writeFile(const std::string& fileName, const std::string& text, bool append);
int getdir (const char* dir, std::vector<char*> &files, const char* extension);
bool existFile(const char* fileName);
bool existDir(const char* dir);
bool makeDir(const char* dir);
bool checkFileCreation(const char* dir);
bool removeFile(const char* file);
__int64 fileSize(const char* file);
long pageSize();

//! This method combines two paths into one
/*! it'll add the file separator char 
  if not present in the first path, also it will check if one of the paths is NULL to avoid
  wrong concatenation. The caller should call free on the result.
  \param path A null terminated string, it could be NULL
  \param path2 A null terminated string, it could be NULL
*/
char* combinePath(const char* path, const char* path2);

#ifndef WINDOWS
#define FILESEPARATOR "/"
#else
#define FILESEPARATOR (char*)"\\"
#endif

#endif	/* _FILEUTIL_H */

