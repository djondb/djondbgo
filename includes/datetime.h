/*
 * =====================================================================================
 *
 *       Filename:  DateTime.h
 *
 *    Description:  This class is used to represent date time data type, it has features
 *                  like comparison methods (>, <, ==, <=, etc) and substraction and 
 *                  addition
 *
 *        Version:  1.0
 *        Created:  07/07/2010 1:34 PM
 *       Revision:  none
 *       Compiler:  gcc
 *
 *         Author:  Juan Pablo Crossley (Cross), crossleyjuan@gmail.com
 *   Organization:  djondb
 *
 * This file is part of the djondb project, for license information please refer to the LICENSE file,
 * the application and libraries are provided as-is and free of use under the terms explained in the file LICENSE
 * Its authors create this application in order to make the world a better place to live, but you should use it on
 * your own risks.
 * 
 * Also, be adviced that, the GPL license force the committers to ensure this application will be free of use, thus
 * if you do any modification you will be required to provide it for free unless you use it for personal use (you may 
 * charge yourself if you want), bare in mind that you will be required to provide a copy of the license terms that ensures
 * this program will be open sourced and all its derivated work will be too.
 * =====================================================================================
 */ 

#ifndef _DATETIME_H
#define	_DATETIME_H

#include "dtime.h"

#include <string>
#include <ctime>

using namespace std;

class DateTime {
public:
	/** Default constructors **/
	/**
	 * Empty constrcutor, all the values will be set to 0
	 **/
    DateTime();
	 /**
	  * This will construct a date with the values passed as arguments
	  **/
    DateTime(int year, int month, int day, int hour = 0, int min = 0, int secs = 0);
	 /**
	  * Uses a time_t to initialize the date
	  **/
    DateTime(time_t time);
	 /**
	  * Creates a date from string representation
	  * yyyy/mm/dd HH:MM:SS
	  * the hour is 24 format
	  *
	  * examples: 2013/01/10 01:35:00
	  * note: all the fields are required
	  **/
    DateTime(string dateTime);
	 /**
	  * Creates a datetime using the "double" representation, which should be like:
	  * 20130113.013500
	  **/
    DateTime(double value);
	 /**
	  * Copy constructor
	  **/
    DateTime(const DateTime& orig);
    virtual ~DateTime();

	 /**
	  * Retuns today's date
	  * @param includeTime if true then the time will be include, false otherwise
	  **/
    static DateTime today(bool includeTime = false);

	 /**
	  * Add days to the current date, the date will not be affected, it will return
	  * the result
	  **/
    DateTime addDays(int days) const;

	 /**
	  * Calculate the number of days between the current instance and the passed argument
	  **/
    int daysTo(const DateTime& dt) const;


	 /**
	  * day of the week, been 0 - monday and 6 sunday
	  **/
    int dayOfTheWeek() const;

	 /**
	  * returns the char representation of the current date
	  * the caller will be responsible to free the memory allocated
	  *
	  * sample returns:  20130115.013500
	  * @returns 20130115.013500
	  **/
    char* toChar() const;

	 /**
	  * Gets the double representation of the date, as follows
	  * YYYYMMDD.hhmmss
	  *
	  * @return 20130115.013525
	  **/
    double toDouble() const;

    void setSecs(int _secs);
    int getSecs() const;
    void setMin(int _min);
    int getMin() const;
    void setHour(int _hour);
    int getHour() const;
    void setDay(int _day);
    int getDay() const;
    void setMonth(int _month);
    int getMonth() const;
    void setYear(int _year);
    int getYear() const;
    DTime dtime() const;

    struct tm* toTimeStruct() const;

    bool operator >(const DateTime& dateTime) const;
    bool operator <(const DateTime& dateTime) const;
    bool operator >=(const DateTime& dateTime) const;
    bool operator <=(const DateTime& dateTime) const;
    long operator -(const DateTime& dateTimeRight) const;
    bool operator ==(const DateTime& dateTime) const;

private:

    int _year;
    int _month;
    int _day;
    int _hour;
    int _min;
    int _secs;

    void loadDate(string dateTime);
};

#endif	/* _DATETIME_H */

