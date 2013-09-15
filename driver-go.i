%module djondb

%include "std_string.i"
%include "std_vector.i"
%include "windows.i"

%{
#include "includes/bsonarrayobj.h"
#include "includes/bson.h"
#include "includes/bsonobj.h"
#include "includes/bsonparser.h"
#include "includes/filterdefs.h"
#include "includes/bsonutil.h"
#include "includes/djondbconnection.h"
#include "includes/djondbcursor.h"
#include "includes/djondbconnectionmanager.h"
#include "includes/djondb_client.h"
%}

namespace std {
   %template(BSONObjVectorPtr) vector<BSONObj*>;
   %template(StringVector) vector<std::string>;
}

%ignore getDJString;

/* Let's just grab the original header file here */
%include "includes/bsonarrayobj.h"
%include "includes/bson.h"
%include "includes/bsonobj.h"
%include "includes/bsonparser.h"
%include "includes/filterdefs.h"
%include "includes/bsonutil.h"
%include "includes/djondbconnection.h"
%include "includes/djondbcursor.h"
%include "includes/djondbconnectionmanager.h"
%include "includes/djondb_client.h"

