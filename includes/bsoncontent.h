#ifndef BSONCONTENT_H_INCLUDED
#define BSONCONTENT_H_INCLUDED

#include "bsondefs.h"
#include "util.h"
#include <string>

class BSONArrayObj;
class BSONObj;

class BSONContent {
	public:
		BSONContent(BSONTYPE);
		virtual ~BSONContent();

		BSONContent(const BSONContent& orig);

		virtual BSONTYPE type() const {
			return _type;
		};

		virtual void setType(BSONTYPE type) {
			_type = type;
		}

		virtual bool operator ==(const BSONContent& content);
		virtual bool operator !=(const BSONContent& content);

		virtual operator __int32() const;
		virtual operator __int64() const;
		virtual operator double() const;
		virtual operator bool() const;
		virtual operator BSONObj*() const;
		virtual operator BSONArrayObj*() const;
		virtual operator djondb::string() const;

		virtual BSONContent* clone() const = 0;
	protected:
		BSONTYPE _type;
};

class BSONContentBoolean: public BSONContent {
	public:
		BSONContentBoolean(bool element);
		virtual ~BSONContentBoolean();

		BSONContentBoolean(const BSONContentBoolean& orig);

		bool operator ==(const BSONContentBoolean& content);
		bool operator !=(const BSONContentBoolean& content);
		operator bool() const;
		virtual BSONContentBoolean* clone() const;
	public:
		bool _element;
};

class BSONContentInt: public BSONContent {
	public:
		BSONContentInt(__int32 element);
		virtual ~BSONContentInt();

		BSONContentInt(const BSONContentInt& orig);

		bool operator ==(const BSONContentInt& content);
		bool operator !=(const BSONContentInt& content);
		operator __int32() const;
		virtual BSONContentInt* clone() const;
	public:
		__int32 _element;
};

class BSONContentLong: public BSONContent {
	public:
		BSONContentLong(__int64 element);
		virtual ~BSONContentLong();

		BSONContentLong(const BSONContentLong& orig);

		bool operator ==(const BSONContentLong& content);
		bool operator !=(const BSONContentLong& content);
		operator __int64() const;
		virtual BSONContentLong* clone() const;
	public:
		__int64 _element;
};

class BSONContentDouble: public BSONContent {
	public:
		BSONContentDouble(double element);
		virtual ~BSONContentDouble();

		BSONContentDouble(const BSONContentDouble& orig);

		bool operator ==(const BSONContentDouble& content);
		bool operator !=(const BSONContentDouble& content);
		operator double() const;
		virtual BSONContentDouble* clone() const;
	public:
		double _element;
};

class BSONContentString: public BSONContent {
	public:
		BSONContentString(char* element, __int32 len);
		virtual ~BSONContentString();

		BSONContentString(const BSONContentString& orig);

		bool operator ==(const BSONContentString& content);
		bool operator !=(const BSONContentString& content);

		operator djondb::string() const;
		virtual BSONContentString* clone() const;
	public:
		djondb::string _element;
};

class BSONContentBSON: public BSONContent {
	public:
		BSONContentBSON(BSONObj* obj);
		virtual ~BSONContentBSON();

		BSONContentBSON(const BSONContentBSON& orig);

		bool operator ==(const BSONContentBSON& content);
		bool operator !=(const BSONContentBSON& content);
		virtual BSONContentBSON* clone() const;
		operator BSONObj*() const;

	public:
		BSONObj* _element;
};

class BSONContentBSONArray: public BSONContent {
	public:
		BSONContentBSONArray(BSONArrayObj* obj);
		virtual ~BSONContentBSONArray();

		BSONContentBSONArray(const BSONContentBSONArray& orig);

		bool operator ==(const BSONContentBSONArray& content);
		bool operator !=(const BSONContentBSONArray& content);
		operator BSONArrayObj*() const;
		virtual BSONContentBSONArray* clone() const;

	public:
		BSONArrayObj* _element;
};

#endif // BSONCONTENT_H_INCLUDED
