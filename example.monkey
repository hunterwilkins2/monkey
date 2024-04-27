/* © 2024 Hunter Wilkins 
* Example Monkey lang program
* https://monkeylang.org/
*/

let arr = [1, 2, 3, 4, 5];

let map = fn(arr, f) {
    let iter = fn(arr, accumulated) {
        if (len(arr) == 0) {
            return accumulated;
        }
        iter(rest(arr), push(accumulated, f(first(arr))))
    }
    iter(arr, [])
};

let reduce = fn(arr, initial, f) {
    let iter = fn(arr, result) {
        if (len(arr) == 0) {
            return result;
        }
        iter(rest(arr), f(result, first(arr)))
    }
    iter(arr, initial)
}

let filter = fn(arr, pred) {
    let iter = fn(arr, accumulated) {
        if (len(arr) == 0) {
            return accumulated;
        }
        if (pred(first(arr))) {
            iter(rest(arr), push(accumulated, first(arr)))
        } else {
            iter(rest(arr), accumulated)
        }
    }
    iter(arr, [])
}

let double = fn(x) { x * 2 };

// Returns [2, 4, 6, 8, 10]
puts(map(arr, double));

// Returns 12
puts(reduce(filter(arr, fn(x) { x > 2 }), 0, fn(a, b) { a + b }));

let search = fn(arr, val) {
    let binarySearch = fn(arr, val, start, end) {
        if (start > end) {
            return false;
        }
        if (start == end) {
            return false;
        }

        let cur = start + (end - start) / 2;
        if (arr[cur] == val) {
            return true;
        }

        if (arr[cur] < cur) {
            return binarySearch(arr, val, cur + 1, end);
        } else {
            return binarySearch(arr, val, start, cur - 1);
        }
    };

    binarySearch(arr, val, 0, len(arr) - 1);
}

// Returns true
puts(search(arr, 3));

// Returns false
puts(search(arr, 8));

let people = [
    {"name": "Anna", "age": 24},
    {"name": "Bob", "age": 39},
    {"name": "Jane", "age": 42}
];

let printPerson = fn(person) { 
    puts(person["name"], person["age"], ""); 
};
puts("");

// Prints: Anna 24
printPerson(people[0]);

// Prints: Bob 39
printPerson(people[1]);

// Prints: Jane 42
printPerson(people[2]);

// Prints: 3
puts("Length", len(people))
