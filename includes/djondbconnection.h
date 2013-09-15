#ifndef DJONDB_CONNECTION_H
#define DJONDB_CONNECTION_H

#include <string>
#include <vector>
#include "bson.h"
#include "util.h"
#include "filterparser.h"

#ifdef WINDOWS
   #define LibraryExport   __declspec( dllexport )
#else
   #define LibraryExport
#endif

class NetworkOutputStream;
class NetworkInputStream;
class CommandWriter;
class TransactionManager;
class Command;

#define SERVER_PORT 1243

namespace djondb {

	class DjondbCursor;

	class LibraryExport DjondbConnection
	{
		public:
			/** Default constructor */
			DjondbConnection(std::string host);
			DjondbConnection(std::string host, int port);
			DjondbConnection(const DjondbConnection& orig);

			/** Default destructor */
			virtual ~DjondbConnection();

			const char* beginTransaction();
			void commitTransaction();
			void rollbackTransaction();

			bool open();
			void close();
			void internalClose();
			bool isOpen() const;

			bool shutdown() const;

			bool insert(const std::string& db, const std::string& ns, const std::string& json);
			bool insert(const std::string& db, const std::string& ns, const BSONObj& obj);
			BSONObj* findByKey(const std::string& db, const std::string& ns, const std::string& select, const std::string& id);
			BSONObj* findByKey(const std::string& db, const std::string& ns, const std::string& id);
			DjondbCursor* find(const std::string& db, const std::string& ns);
			DjondbCursor* find(const std::string& db, const std::string& ns, const BSONObj& options);
			DjondbCursor* find(const std::string& db, const std::string& ns, const std::string& filter);
			DjondbCursor* find(const std::string& db, const std::string& ns, const std::string& filter, const BSONObj& options);
			DjondbCursor* find(const std::string& db, const std::string& ns, const std::string& select, const std::string& filter);
			DjondbCursor* find(const std::string& db, const std::string& ns, const std::string& select, const std::string& filter, const BSONObj& options);
			bool update(const std::string& db, const std::string& ns, const std::string& json);
			bool update(const std::string& db, const std::string& ns, const BSONObj& bson);
			bool remove(const std::string& db, const std::string& ns, const std::string& id, const std::string& revision);
			BSONArrayObj* executeQuery(const std::string& query);
			bool executeUpdate(const std::string& query);

			bool dropNamespace(const std::string& db, const std::string& ns);
			std::vector<std::string>* dbs() const;
			std::vector<std::string>* namespaces(const std::string& db) const;

			std::string host() const;

		protected:
		private:
			void prepareOptions(Command* cmd);
			Command* parseCommand(const std::string& expression);

		private:
			NetworkOutputStream*  _outputStream;
			NetworkInputStream*   _inputStream;
			CommandWriter*        _commandWriter;
			std::string*          _activeTransactionId;

			std::string _host;
			int _port;
			bool _open;
			Logger* _logger;
	};

}

#endif // DJONDB_CONNECTION_H
