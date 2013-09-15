#ifndef DTIME_H
#define DTIME_H

#include <string>

class DTime
{
public:
    DTime();
    DTime(long secs);
    DTime(int hours, int minutes, int secs);
    DTime(std::string time); // HHMMSS
    DTime(const DTime& time); // HHMMSS
    ~DTime();

    long totalSecs() const;
    long totalMinutes() const;

    int hour() const;
    int minutes() const;
    int seconds() const;

    char* toChar() const;

    void add(long secs);
    void add(const DTime time);

    DTime operator +(const DTime& dtime) const;
    void operator ++(int secs);
    DTime operator -(const DTime& dtime) const;
    void operator --(int secs);


private:
    long _secs;
    void time(int& hour, int& min, int& secs) const;
};

#endif // DTIME_H
