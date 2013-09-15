#ifndef BPLUSINDEX_H
#define BPLUSINDEX_H

#include "index.h"
#include <list>

template <class K, class V>
class PriorityCache;

#include <boost/shared_ptr.hpp>

//typedef char* INDEXPOINTERTYPE;
typedef const char* INDEXPOINTERTYPE;

#define COMPAREKEYS(k1, k2) \
	(strcmp(k1, k2) == 0);

const int BUCKET_MAX_ELEMENTS = 5; // Should be even (3, 5, 7)

class IndexPage;

class IndexPage {
	public:
		IndexPage();
		~IndexPage();
		int size;

		Index** elements;
		IndexPage** pointers;
		IndexPage* rightSibling;
		IndexPage* leftSibling;

		IndexPage* parentElement;

		void debugElements() const;
		void debug() const;
		bool isLeaf() const;
		bool isFull() const;
		bool _leaf;
		std::list<Index*> find(FilterParser* parser) const;
		int findInsertPosition(Index* index) const;

		void moveElements(int startPoint, int count);
		void movePointers(int startPoint, int count);

};

class BPlusIndex: public IndexAlgorithm
{
	public:
		BPlusIndex(std::set<std::string> keys);
		virtual ~BPlusIndex();

		virtual void add(const BSONObj& elem, djondb::string documentId, long filePos, long indexPos);
		virtual Index* find(BSONObj* const elem);
		virtual void remove(const BSONObj& elem);
		virtual std::list<Index*> find(FilterParser* parser);

		void debug();
		
	protected:
	private:
		IndexPage* _head;

	private:
		IndexPage* findIndexPage(IndexPage* start, djondb::string key) const;
		Index* findIndex(IndexPage* start, djondb::string key) const;
		void insertIndexElement(IndexPage* page, Index* index);
		void dispose(IndexPage* page);
		void createRoot(Index* element, IndexPage* left, IndexPage* right);
		void addElement(IndexPage* page, Index* element, IndexPage* rightPointer);
		void splitAdd(IndexPage* page, Index* index, IndexPage* rightPointer);
		void splitAddLeaf(IndexPage* page, Index* index);
		void splitAddInner(IndexPage* page, Index* index, IndexPage* rightPointer);

		void copyElements(IndexPage* source, IndexPage* destination, int startIndex, int endIndex);
		void removeElements(IndexPage* source, int startIndex, int endIndex);
		void moveElements(IndexPage* source, IndexPage* destination, int startIndex, int endIndex);
		/*  
			 bool insertElement(const Index& elem);
			 BucketElement* findBucketElement(Bucket* start, const Index& idx, bool create);
			 void initializeBucket(Bucket* const element);
			 void initializeBucketElement(BucketElement* const elem);

			 void insertBucketElement(Bucket* bucket, BucketElement* element);
			 void checkBucket(Bucket* const bucket);

			 std::list<Index*> find(FilterParser* parser, Bucket* bucket);
			 std::list<Index*> findElements(FilterParser* parser, Bucket* bucket);
			 */
};

#endif // BPLUSINDEX_H
