#ifndef INDEXFACTORY_H
#define INDEXFACTORY_H

#include <map>
#include <string>
#include <set>
#include <vector>

class IndexAlgorithm;

using namespace std;

class BSONObj;

typedef std::vector<IndexAlgorithm*> listAlgorithmsType;
typedef listAlgorithmsType* listAlgorithmsTypePtr;
typedef map<std::string, listAlgorithmsTypePtr> listByNSType;
typedef listByNSType* listByNSTypePtr;
typedef map<std::string, listByNSTypePtr> listByDbType;

class IndexFactory
{
    public:
        virtual ~IndexFactory();

        static IndexFactory indexFactory;
        IndexAlgorithm* index(const char* db, const char* ns, const std::set<std::string>& keys);
        IndexAlgorithm* index(const char* db, const char* ns, const std::string& key);
		  bool containsIndex(const char* db, const char* ns, const std::string& key);
		  bool containsIndex(const char* db, const char* ns, const std::set<std::string>& keys);
        void dropIndex(const char* db, const char* ns, const std::string& key);
		  void dropIndex(const char* db, const char* ns, const std::set<std::string>& keys);
		  void dropIndexes(const char* db, const char* ns);
    protected:
    private:
        IndexFactory();

			listAlgorithmsTypePtr findAlgorithms(const char* db, const char* ns);
		  IndexAlgorithm* findIndex(const listAlgorithmsTypePtr& algorithms, const std::set<std::string>& keys);
		  void dropIndex(const listAlgorithmsTypePtr& algorithms, const std::set<std::string>& keys);

    private:

        listByDbType _indexes;
};

#endif // INDEXFACTORY_H
