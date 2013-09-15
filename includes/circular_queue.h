/*
 * =====================================================================================
 *
 *       Filename:  circular_queue.h
 *
 *    Description:  Provides a circular fixed size queue
 *
 *        Version:  1.0
 *        Created:  06/24/2012 11:20:29 AM
 *       Revision:  none
 *       Compiler:  gcc
 *
 *         Author:  Juan Pablo Crossley (crossleyjuan@gmail.com), 
 *   Organization:  
 *
 * =====================================================================================
 */

template <class A>
class CircularQueue {
	private:
		class Element {
			private:
				A _value;
				Element* _next;
				Element* _previous;
			public:
				Element(A val) {
					_value = val;
					_next = NULL;
					_previous = NULL;
				}

				A value() {
					return _value;
				}

				Element* next() {
					return _next;
				}

				void setNext(Element* n) {
					_next = n;
				}

				Element* previous() {
					return _previous;
				}

				void setPrevious(Element* p) {
					_previous = p;
				}
		};

	public:
		CircularQueue(int size);
		CircularQueue();
		~CircularQueue();

		void push_back(A val);
		A pop_back();
		void push_front(A val);
		A pop_front();
		A next();
		A previous();
		void reset();
		A current();
		int size() const;

	private:
		int _maxSize;
		int _size;
		Element* _lastElement;
		Element* _firstElement;
		Element* _currentElement;
};

template <class A>
CircularQueue<A>::CircularQueue(int size) {
	_lastElement = NULL;
	_firstElement = NULL;
	_maxSize = size;
	_size = 0;
	_currentElement = NULL;
}

template <class A>
CircularQueue<A>::CircularQueue() {
	_lastElement = NULL;
	_firstElement = NULL;
	_maxSize = -1;
	_size = 0;
	_currentElement = NULL;
}

template <class A>
CircularQueue<A>::~CircularQueue() {
	while (size() > 0) {
		pop_front();
	}
}

template <class A>
void CircularQueue<A>::push_back(A val) {
	Element* element = new Element(val);

	if (_lastElement == NULL) {
		_lastElement = element;
	}
	_lastElement->setNext(element);
	element->setPrevious(_lastElement);
	_lastElement = element;

	if (_firstElement == NULL) {
		_firstElement = element;
	}
	_lastElement->setNext(_firstElement);
	_firstElement->setPrevious(_lastElement);

	_size++;
	if ((_maxSize > -1) && (_size > _maxSize)) {
		pop_front();
	}
	_currentElement = _lastElement;
}

template <class A>
A CircularQueue<A>::pop_back() {
	Element* result = _lastElement;
	A val = result->value();
	if (_size == 1) {
		_lastElement = NULL;
		_firstElement = NULL;
		_currentElement = NULL;
	} else if (_size > 1) {
		_lastElement = _lastElement->previous();
		_lastElement->setNext(_firstElement);
		_firstElement->setPrevious(_lastElement);
	}

	if (result != NULL) delete result;
	_currentElement = _lastElement;
	_size--;
	return val;
}


template <class A>
void CircularQueue<A>::push_front(A val) {
	Element* element = new Element(val);

	if (_firstElement == NULL) {
		_firstElement = element;
	}
	element->setNext(_firstElement);
	_firstElement->setPrevious(element);
	_firstElement = element;

	if (_lastElement == NULL) {
		_lastElement = element;
	}
	_lastElement->setNext(element);
	_firstElement->setPrevious(_lastElement);

	_size++;
	if ((_maxSize > -1) && (_size > _maxSize)) {
		pop_back();
	}
	_currentElement = _lastElement;
}

template <class A>
A CircularQueue<A>::pop_front() {
	Element* result = _firstElement;
	A val = result->value();
	if (_size == 1) {
		_lastElement = NULL;
		_firstElement = NULL;
		_currentElement = NULL;
	} else if (_size > 1) {
		_firstElement = _firstElement->next();
		_lastElement->setNext(_firstElement);
		_firstElement->setPrevious(_lastElement);
	}

	if (result != NULL) delete result;
	_currentElement = _lastElement;
	_size--;
	return val;
}

template <class A>
A CircularQueue<A>::next() {
	Element* ret = _currentElement->next();
	A val = ret->value();
	_currentElement = ret;
	return val;
}

template <class A>
A CircularQueue<A>::previous() {
	Element* ret = _currentElement->previous();
	A val = ret->value();
	_currentElement = ret;
	return val;
}

template <class A>
void CircularQueue<A>::reset() {
	_currentElement = _lastElement;
}

template <class A>
A CircularQueue<A>::current() {
	if (_currentElement != NULL) {
		return _currentElement->value();
	}
}

template <class A>
int CircularQueue<A>::size() const {
	return _size;
}
