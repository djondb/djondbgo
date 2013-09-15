#ifndef VERSION_H
#define VERSION_H

#include <string>

class Version
{
public:
    Version(std::string version);
    Version(const Version& version);
    ~Version();

    bool operator>(const Version& version) const;
    bool operator>=(const Version& version) const;
    bool operator<(const Version& version) const;
    bool operator<=(const Version& version) const;
    bool operator==(const Version& version) const;
    operator std::string();
private:
    std::string _version;

    int _mayor;
    int _minor;
    long _revision;
};

#endif // VERSION_H
