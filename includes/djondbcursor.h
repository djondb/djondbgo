// =====================================================================================
// 
//  @file:  djondbcursor.h
// 
//  @brief:  Cursor client side implementation, this contains a reference to the server cursor
//           and collects it's results
// 
//  @version:  1.0
//  @date:     08/30/2013 12:46:58 PM
//  Compiler:  g++
// 
//  @author:  Juan Pablo Crossley (Cross), crossleyjuan@gmail.com
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
#ifndef DJONDBCURSOR_INCLUDED_H
#define DJONDBCURSOR_INCLUDED_H 

#include "util.h"

#ifdef WINDOWS
   #define LibraryExport   __declspec( dllexport )
#else
   #define LibraryExport
#endif

class BSONObj;
class BSONArrayObj;
class NetworkOutputStream;
class NetworkInputStream;
class CommandWriter;

namespace djondb {


	/**
	 * @brief This class represents a cursor on the server side, and it is implemented for the version 0.3.2 and above
	 * if the server is in a previous version this will crash with unexpected error, it's advisable, at least for the
	 * moment to have server and client aligned in the same version, later on this class could be improved to support
	 * unknown versions
	 */
	class LibraryExport DjondbCursor
	{
		enum CursorStatus {
			CS_LOADING, //!< Cursor is still loading
			CS_RECORDS_LOADED //!< All the records are ready, it's now a client side cursor
		};

		public:
		/**
		 * @brief Constructor to be used from the implementation, is not intended to be used by external users (clients of the
		 * driver)
		 *
		 * @param nos
		 * @param nis
		 * @param writer
		 * @param cursorId contains the id of the cursor
		 * @param firstPage contains the first page of the cursor, the cursor will gain control of this variable, it should not freed from the caller
		 */
			DjondbCursor(NetworkOutputStream* nos, NetworkInputStream* nis, CommandWriter* writer, const char* cursorId, BSONArrayObj* firstPage);
			DjondbCursor(const DjondbCursor& orig);
			~DjondbCursor();

			/**
			 * @brief checks if more elements could be retrieved
			 *
			 * @return true if more BSONObj are ready, false otherwise
			 */
			bool next();
			/**
			 * @brief checks if there are elements in the front of the cursor
			 *
			 * @return true of not BOF, false otherwise
			 */
			bool previous();
			/**
			 * @brief Returns the current loaded element, the client should call next() method before calling this, if not
			 * an unexpected behavior could occur
			 *
			 * @return the current element or NULL if the next method was not called
			 */
			BSONObj* const current();

			/**
			 * @brief This will retrieve the length of the rows contained in the cursor, if the cursor is still loading then
			 * all the rows will be retrieved from the server. This method should be used with care, because it will try to
			 * retrieve every row from the server and it may contain several pages
			 *
			 * @return length of the cursor
			 */
			__int32 length();

			/**
			 * @brief Releases the cursor from the server, the client should use this method if the cursor is no longer required
			 */
			void releaseCursor();

			/**
			 * @brief Will place the current row in the desired position
			 *
			 * @param position
			 */
			void seek(__int32 position);

		private:// Private variables
			/**
			 * @brief Contains the returned elements, could be incomplete until the EOF is reached, please check "_status" to know
			 * if this is fully loaded
			 */
			BSONArrayObj* _rows;
			/**
			 * @brief last element retrieved by next or previous methods
			 */
			BSONObj* _current;
			/**
			 * @brief Current position
			 */
			int _position;
			/**
			 * @brief Number of elements currently loaded in the _rows array
			 */
			int _count;
			/**
			 * @brief Status of the cursor, just a flag to know if the cursor is still loading (more data could be server from the
			 * server or all the data was retrieved)
			 */
			CursorStatus _status;
			/**
			 * @brief Cursor ID
			 */
			char* _cursorId;

			NetworkOutputStream*  _outputStream;
			NetworkInputStream*   _inputStream;
			CommandWriter*        _commandWriter;
			
		private:// Private methods
			void retrieveCurrent();
	};
};
#endif /* DJONDBCURSOR_INCLUDED_H */
