let foo: string = "hello";


function tString(num: number): string {
    return String(num)
}
console.log(foo)
console.log(tString(123))

const x: object = { foo: 123 };
const y: object = [1, 2, 3];
const z: object = (n: number) => {
    return n + 1
};

