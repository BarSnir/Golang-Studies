var p1 = {
    first: "Bar",
    last: "Snir",
    age: 29
}

var p2 = p1

p2.first = "Meshi"
p2.last = "Mor"

console.log(p1);
console.log(p2);