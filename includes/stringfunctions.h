/*
 * File:   stringfunctions.h
 * Author: cross
 *
 * Created on July 8, 2010, 2:54 PM
 */

#ifndef _STRINGFUNCTIONS_H
#define	_STRINGFUNCTIONS_H

#include <string>
#include <vector>
#include "defs.h"

namespace djondb {

	class stringHolder {
		public:
			stringHolder(char* c) {
				_char = c;
				_constChar = NULL;
				_references = 1;
			}

			stringHolder(const char* c) {
				_constChar = c;
				_char = NULL;
				_references = 1;
			}

			void addReference() {
				_references++;
			}

			void removeReference() {
				_references--;
			}

			__int32 references() const {
				return _references;
			}

			~stringHolder() {
				if (_char != NULL) free(_char);
				_char = NULL;
			}

			operator char*() {
				if (_char != NULL) {
					return _char;
				} else {
					return const_cast<char*>(_constChar);
				}
			}

			operator const char*() {
				if (_constChar != NULL) {
					return _constChar;
				} else {
					return const_cast<const char*>(_char);
				}
			}

		private:
			char* _char;
			const char* _constChar;
			__int32 _references;
	};

	/**
	 * This class is just a wrapper for a char*, it will keep the same char* passed and this
	 * will just work as wrapper container
	 * This class should be used if the char* should keep the same pointed element as the original, 
	 * useful when working with Buffers
	 **/
	class string {
		public:
			string();
			string(char* c, __int32 len);
			string(const char* c, __int32 len);
			string(const string& str);
			virtual ~string();

			const char* c_str() const;
			__int32 length() const;

			int compare(const djondb::string s2) const;
			int compare(const char* s2, __int32 len) const;

			operator const char*() const;
			operator char*() const;

			std::string str() const;

			bool operator ==(const djondb::string& str);
			bool operator !=(const djondb::string& str);
			djondb::string& operator=(const djondb::string& rvar);

		private:
			stringHolder* _holder;
			__int32 _len;
	};
};

char* strcpy(char* str, int len);
char* strcpy(const char* str, int len);
char* strcpy(char* str, int offset, int len);
char* strcpy(std::string str);
bool endsWith(char* source, char* check);
bool startsWith(const char* source, const char* check);
std::vector<std::string*>* tokenizer(const std::string source, const char* tokens);
std::string format(const char * fmt, ...);

std::string toString(double a);
std::string toString(double a, int fixedPrecision);
std::string toString(int a);
std::vector<std::string> split(std::string str, std::string token);
long countChar(const char* s, const char c);
std::string concatStrings(const std::string& a, const std::string& b);

std::vector<std::string> splitLines(std::string);
bool compareInsensitive(const char* text1, const char* text2);

char* trim(char* str, int len);

#endif	/* _STRINGFUNCTIONS_H */

