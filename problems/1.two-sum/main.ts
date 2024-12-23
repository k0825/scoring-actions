const main = (inputs: string): void => {
  const lines = inputs.split(/\n/);

  const firstLine = lines[0].split(/\s/);
  const A = parseInt(firstLine[0]);
  const B = parseInt(firstLine[1]);

  const secondLine = lines[1];
  const S = secondLine;

  console.log(A + B, S);
};

main(require("fs").readFileSync("/dev/stdin", "utf8"));
