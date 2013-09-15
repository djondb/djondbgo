#ifndef BSONPARSER_H
#define BSONPARSER_H

#include <string>
#include "bsondefs.h"
#include "defs.h"

class BSONObj;
class BSONArrayObj;

class BSONParser
{
	public:
		/** Default constructor */
		BSONParser();
		/** Default destructor */
		virtual ~BSONParser();

		static BSONObj* parse(const std::string& sbson);
		static BSONArrayObj* parseArray(const std::string& sbson);

	protected:
	private:
		static BSONObj* parseBSON(const char* c, __int32& pos) throw(BSONException);
		static BSONArrayObj* parseArray(const char* chrs, __int32& pos);
};

#endif // BSONPARSER_H
