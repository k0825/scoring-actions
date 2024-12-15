// main.ts
const main = (inputs: string): string => {
  const numbers = inputs.split(" ").map(Number);
  const sum = numbers.reduce((a, b) => a + b, 0);
  return sum.toString();
};

if (require.main === module) {
  const inputs = process.argv[2];
  const output = main(inputs);
  console.log(output);
}
