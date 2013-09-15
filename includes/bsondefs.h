#ifndef BSON_DEFS_H_INCLUDED
#define BSON_DEFS_H_INCLUDED

#include <string>

enum BSONTYPE {
    INT_TYPE,
    DOUBLE_TYPE,
    LONG_TYPE,
	 PTRCHAR_TYPE,
    STRING_TYPE,
    BSON_TYPE,
	 BSONARRAY_TYPE,
    UNKNOWN_TYPE,
	 NULL_TYPE,
    LONG64_TYPE,
	 BOOL_TYPE
//    PTR
};

class BSONException: public std::exception {
	public:
		BSONException(const char* error):
  		  std::exception()	{
			_errorMessage = error;
		}

		BSONException(const BSONException& orig) :
			std::exception()	{
			this->_errorMessage = orig._errorMessage;
		}

		virtual const char* what() const throw() {
			return _errorMessage;
		}

	private:
		const char* _errorMessage;
};

#endif //BSON_DEFS_H_INCLUDED
