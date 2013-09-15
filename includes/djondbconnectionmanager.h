#ifndef DJONDB_CONNECTIONMANAGER_H
#define DJONDB_CONNECTIONMANAGER_H

#include <map>
#include <string>
#include "util.h"

#ifdef WINDOWS
   #define LibraryExport   __declspec( dllexport )
#else
   #define LibraryExport
#endif


namespace djondb {
    class DjondbConnection;

    struct ConnectionReference {
        DjondbConnection* _connection;
        int _references;
    };

    class LibraryExport DjondbConnectionManager
    {
        public:
            /** Default constructor */
            DjondbConnectionManager();
            /** Default destructor */
            virtual ~DjondbConnectionManager();

            static DjondbConnection* getConnection(std::string host);
            static DjondbConnection* getConnection(std::string host, int port);

            static void releaseConnection(DjondbConnection* conn);

        protected:
        private:
            static std::map<std::string, struct ConnectionReference> _connections;

			static bool __initialized;
    };
}

#endif // DJONDB_CONNECTIONMANAGER_H
