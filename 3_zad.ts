const b = 2.5;
const bCubed = b ** 3;

const Func = (x: number) => {
    const sum = bCubed + Math.pow(x, 3);
    return (1 + Math.sin(sum) ** 2) / Math.cbrt(sum);
};

console.log("Задача А");
console.log("   x    |     y");
console.log(`  1.28  | ${Func(1.28).toFixed(6)}`);
console.log(`  1.68  | ${Func(1.68).toFixed(6)}`);
console.log(`  2.08  | ${Func(2.08).toFixed(6)}`);
console.log(`  2.48  | ${Func(2.48).toFixed(6)}`);
console.log(`  2.88  | ${Func(2.88).toFixed(6)}`);
console.log(`  3.28  | ${Func(3.28).toFixed(6)}`);

console.log("\nЗадача Б");
console.log("   x    |     y");
console.log(`  1.1   | ${Func(1.1).toFixed(6)}`);
console.log(`  2.4   | ${Func(2.4).toFixed(6)}`);
console.log(`  3.6   | ${Func(3.6).toFixed(6)}`);
console.log(`  1.7   | ${Func(1.7).toFixed(6)}`);
console.log(`  3.9   | ${Func(3.9).toFixed(6)}`);